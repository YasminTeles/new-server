# New Server

A simple server in Golang.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purpose.

There are two ways to get started:

<details>
<summary>Get started with local Golang;</summary>

### Prerequisites

- [Golang](https://golang.org/) - Go is an open source programming language that makes it easy to build simple, reliable, and efficient software. You need the version 1.17.

### Installation

1. Clone this repository;

2. Run the following commands:

```Makefile
cd new-server
make setup
make run
```

3. Open <http://localhost:3000/healthcheck> with your browser to see the result.

### Running tests

```Makefile
make test
```
</details>

<details>
<summary>Get started with Docker;</summary>

### Prerequisites

- [Docker](https://www.docker.com/) - is an open platform for developing, shipping, and running applications. Docker enables you to separate your applications from your infrastructure so you can deliver software quickly.

### Installation

1. Clone this repository;

2. Run the following commands:

```Makefile
cd new-server
make docker-build
make docker-run
```

3. Open <http://localhost:3000/healthcheck> with your browser to see the result.

4. For kill container's Docker, run the following command:

```Makefile
make docker-kill
```

</details>

## Usage

This simple server provides the following end points:

- `GET /version`
 That returns the version of the server. It's useful for blue-green development.

- `GET /healthcheck`
 That returns the health of the server running. It's useful for check if the server can be able handling requests.

## Versions

We use [Semantic version](http://semver.org) for versioning. For versions available, see [Changelog](Changelog.md).

## Contribute Us

Contributions are what make the open source community such an amazing place to learn, inspire, and create.
Any contributions you make are greatly appreciated. See [contribute policy](Contribute.md).

## License

This project is licensed under the [MIT License](LICENSE).
