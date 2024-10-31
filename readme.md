# Space Tickets

## Contents
<!-- `make toc` to generate https://github.com/jonschlinkert/markdown-toc#cli -->

<!-- toc -->

- [To Run](#to-run)
- [Improvements](#improvements)

<!-- tocstop -->

## To Run
You need Docker on your machine for this to work.

It is assumed you're on a Mac.

Run `make start logs` from the project folder in a terminal.

This will pull the Docker image for the application, run it, and show its logs.

If the image won't pull down from Docker Hub, you can build it locally by running `make build`. This builds a multi-platform image that will work on `amd64` and `arm64` machines. It uses Docker's `buildx` tool, which should be installed on your machine if you have Docker Desktop.

Following the build, run the app using `make start logs`.

Access the API here http://localhost:8080/api/v1.

View Swagger docs here http://localhost:8081.

To run unit tests run `make test`.

## Improvements
