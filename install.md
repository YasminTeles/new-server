# Install this project

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purpose.

### Prerequisites

- [Golang](https://golang.org/) - Go is an open source programming language that makes it easy to build simple, reliable, and efficient software. You need the version 1.21.

### Installation

1. Clone this repository and access the folder project;

2. Run the following commands:

    ```bash
    make setup
    make run
    ```

3. Open <http://localhost:4000/healthcheck> with your browser to see the result.

### Running tests

1. Run the server

    ```bash
    make run
    ```

2. In another terminal, run the tests:

    ```bash
    make test
    ```

## Docker Image

This project can run inside a docker container. Run the following commands to run the container on your local machine:

1. Clone this repository and access the folder project;

2. Run the following commands:

    ```bash
    make docker-build
    make docker-run
    ```

3. Open <http://localhost:4000/healthcheck> with your browser to see the result.

4. For kill container's Docker, run the following command:

    ```bash
    make docker-kill
    ```

## Helper

If you do not know what the make command is doing, you can use the following command to show a short description:

```bash
make help
```

## Contact me

I'm always up for a chat. If you have a question or suggestion, please get in touch with me through [my site](https://yasminteles.com).
