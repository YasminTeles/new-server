# New Server

A boilerplate/starter project for quickly building RESTful APIs using Golang, Negroni, and HttpRouter.

It's intended as a starting point for your product or service. The server comes with some built-in features, such as unit and integration tests, continuous integration, docker support. For more details, check the [features list](#features).

## Learn More

To learn more about this project, take a look at the following resources:

- [Local Install](install.md) - Learn about how to install and use this project.

## Usage

This simple server provides the following end points:

- `GET /version`
 That returns the version of the server. It's useful for blue-green development.

- `GET /healthcheck`
 That returns the health of the server running. It's useful for check if the server can be able handling requests.

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

## Versions

We use [Semantic version](http://semver.org) for versioning. For versions available, see [changelog](Changelog.md).

## License

This project is licensed under the [MIT License](LICENSE).

## Contact me

I'm always up for a chat. If you have a question or suggestion, please get in touch with me through [my site](https://yasminteles.com).
