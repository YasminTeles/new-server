package settings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	t.Parallel()

	config, err := loadConfig()

	assert.NoError(t, err)
	assert.NotEmpty(t, config)
}
