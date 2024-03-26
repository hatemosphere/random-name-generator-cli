# Random Name Generator CLI

This is a simple CLI tool written in Go that generates random names. It uses the Moby Project's name generation library for generating names.

## Getting Started

### Prerequisites

- Go 1.x

### Installation

1. Clone this repository:

    ```bash
    git clone https://github.com/hatemosphere/go-random-name-generator-cli.git
    ```

2. Navigate into the project directory:

    ```bash
    cd go-random-name-generator-cli
    ```

3. Install the required packages:

    ```bash
    go mod download
    ```

4. Run the program:

    ```bash
    go run main.go
    ```

    To generate multiple names:

    ```bash
    go run main.go -count=5
    ```

## License

This project is licensed under the Apache License 2.0, following the license terms of the Moby Project. See the [LICENSE](LICENSE) file for details.
