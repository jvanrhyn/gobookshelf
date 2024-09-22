# GoBookshelf

GoBookshelf is a web application built with Go and the Fiber framework. It provides a platform for managing and viewing books.

## Project URL

[GoBookshelf on GitHub](https://github.com/jvanrhyn/gobookshelf)

## Getting Started

### Prerequisites

- Go 1.22 or later
- Git

### Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/jvanrhyn/gobookshelf.git
    cd gobookshelf
    ```

2. Install dependencies:

    ```sh
    go mod download
    ```

### Running the Application

To start the server, run:

```sh
 go run cmd/main.go
```

The server will start on `http://localhost:8080`.

### Stopping the Server

To stop the server, press `Ctrl+C`. The server will log "Server exiting" before shutting down.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
