# GoEnv

GoEnv is a Go package that provides utilities for retrieving environment variables either directly from the system or from a `.env` file. It relies on the [godotenv](https://github.com/joho/godotenv) package for loading environment variables from a `.env` file.

## Installation

To use this package in your Go project, you can install it using `go get`:

```bash
go get github.com/subinoybiswas/goenv
```

## Usage

Import the package into your Go code:

```go
import (
    "github.com/subinoybiswas/goenv"
)
```

Then, you can use the various functions to retrieve environment variables:

```go
value, err := goenv.GetEnv("YOUR_ENV_VARIABLE")
if err != nil {
    // Handle error
}
// Use the value retrieved from the environment
```

You can also use specific functions to retrieve environment variables either from the `.env` file or directly from the OS environment:

```go
value, err := goenv.GetEnvFrmFile("YOUR_ENV_VARIABLE")
if err != nil {
    // Handle error
}
// Use the value retrieved from the environment
```

```go
value, err := goenv.GetEnvFrmOS("YOUR_ENV_VARIABLE")
if err != nil {
    // Handle error
}
// Use the value retrieved from the environment
```

By default, the `GetEnv` function first attempts to load environment variables from a `.env` file using `godotenv.Load(".env")`. If that fails, it falls back to retrieving the environment variable directly from the system.

## Functions

### `GetEnv(key string) (string, error)`

Retrieves the value of the environment variable specified by `key`.

- `key`: The name of the environment variable to retrieve.
- Returns:
  - The value of the environment variable.
  - An error if the environment variable is not found or if there was an error loading the `.env` file.

### `GetEnvFrmFile(key string) (string, error)`

Retrieves the value of the environment variable specified by `key` from the `.env` file.

- `key`: The name of the environment variable to retrieve.
- Returns:
  - The value of the environment variable.
  - An error if the environment variable is not found or if there was an error loading the `.env` file.

### `GetEnvFrmOS(key string) (string, error)`

Retrieves the value of the environment variable specified by `key` from the OS environment.

- `key`: The name of the environment variable to retrieve.
- Returns:
  - The value of the environment variable.
  - An error if the environment variable is not found.

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvement, please open an issue or submit a pull request on [GitHub](https://github.com/subinoybiswas/goenv).

