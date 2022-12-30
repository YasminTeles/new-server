//nolint:forbidigo,gochecknoglobals,dogsled
package settings

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"runtime"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var Config *Settings

type Settings struct {
	Port        string        `mapstructure:"port"`
	LogLevel    int           `mapstructure:"logLevel"`
	Timeout     time.Duration `mapstructure:"timeout"`
	IdleTimeout time.Duration `mapstructure:"idleTimeout"`
}

func LoadSettings() {
	var err error
	Config, err = loadConfig()

	if err != nil {
		log.Fatal("cannot load config.")
	}

	Config.Print()
}

func loadConfig() (config *Settings, err error) {
	path := getConfigPath()
	viper.AddConfigPath(path)

	environment := "local"
	viper.SetConfigName(environment)

	printConfigPath(path, environment)

	if err = viper.ReadInConfig(); err != nil {
		return
	}

	if err = viper.Unmarshal(&config); err != nil {
		return
	}

	return
}

func getConfigPath() string {
	_, file, _, _ := runtime.Caller(0)
	root := filepath.Join(filepath.Dir(file), "../config")

	return root
}

func printConfigPath(path, environment string) {
	fmt.Print("Reading the following config file:\n")
	fmt.Printf("%s\\%s.yaml\n", path, environment)
}

func (config *Settings) Print() {
	printable, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Printf("%s\n", string(printable))
}
