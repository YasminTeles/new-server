# New Server

A boilerplate/starter project for quickly building RESTful APIs using Golang, Negroni, and HttpRouter.

It's intended as a starting point for your product or service. The server comes with some built-in features, such as unit and integration tests, continuous integration, docker support. For more details, check the [features list](#features).

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purpose.

There are two ways to get started:

<details>
<summary>Get started with local Golang;</summary>

### Prerequisites

- [Golang](https://golang.org/) - Go is an open source programming language that makes it easy to build simple, reliable, and efficient software. You need the version 1.17.

### Installation

1. Clone this repository and access the folder project;

2. Run the following commands:

    ```bash
    make setup
    make run
    ```

3. Open <http://localhost:3000/healthcheck> with your browser to see the result.

### Running tests

1. Run the server

    ```bash
    make run
    ```

2. In another terminal, run the tests:

    ```bash
    make test
    ```

</details>

<details>
<summary>Get started with Docker;</summary>

### Prerequisites

- [Docker](https://www.docker.com/) - is an open platform for developing, shipping, and running applications. Docker enables you to separate your applications from your infrastructure so you can deliver software quickly.

### Installation

1. Clone this repository and access the folder project;

2. Run the following commands:

    ```bash
    make docker-build
    make docker-run
    ```

3. Open <http://localhost:3000/healthcheck> with your browser to see the result.

4. For kill container's Docker, run the following command:

    ```bash
    make docker-kill
    ```

</details>

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
- **Linting**: using [Golangci-lint](https://github.com/golangci/golangci-lint);
- **Testing**: using [Testify](https://github.com/stretchr/testify);
- **Request ID in request header and context**;
- **Structured and centralized settings**: using [Viper](https://github.com/spf13/viper);

## Versions

We use [Semantic version](http://semver.org) for versioning. For versions available, see [changelog](Changelog.md).

## License

This project is licensed under the [MIT License](LICENSE).

## Contact me

I'm always up for a chat. If you have a question or suggestion, You'll try to contact me through [my site](https://yasminteles.com).
