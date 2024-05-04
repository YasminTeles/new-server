# About the project

## Endpoints

This simple server provides the following end points:

- `GET /version`
 That returns the version of the server. It's useful for blue-green development.

- `GET /healthcheck`
  That returns the health of the server running. It's useful for checking if the server can handle requests.

## Features

- **Docker support**;
- **Structured and centralized logger**: using [Logrus](https://github.com/sirupsen/logrus);
- **Continuous integration**: using [GitHub Actions](https://github.com/features/actions);
- **Secrets Scan**: using [Gitleaks](https://github.com/gitleaks/gitleaks);
- **Infrastructure Scan**: using [Trivy](https://trivy.dev/);
- **Linting**: using [Golangci-lint](https://github.com/golangci/golangci-lint);
- **Testing**: using [Testify](https://github.com/stretchr/testify);
- **Request ID in request header and context**;
- **Structured and centralized settings**: using [Viper](https://github.com/spf13/viper);
- **Cloud developer platform support**: using [Gitpod](https://gitpod.io/);

## Versions

We use [Semantic version](http://semver.org) for versioning. For versions available, see [changelog](Changelog.md).

## License

This project is licensed under the [MIT License](LICENSE).

## Contact me

I'm always up for a chat. If you have a question or suggestion, please get in touch with me through [my site](https://yasminteles.com).
