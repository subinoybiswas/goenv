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
    "github.com/joho/godotenv"
    "github.com/yourusername/goenv"
)
```

Then, you can use the `GetEnv` function to retrieve environment variables:

```go
value, err := goenv.GetEnv("YOUR_ENV_VARIABLE")
if err != nil {
    // Handle error
}
// Use the value retrieved from the environment
```
```go

```

By default, the `GetEnv` function first attempts to load environment variables from a `.env` file using `godotenv.Load(".env")`. If that fails, it falls back to retrieving the environment variable directly from the system.

## Functions

### `GetEnv(key string) (string, error)`

Retrieves the value of the environment variable specified by `key`.

- `key`: The name of the environment variable to retrieve.
- Returns:
  - The value of the environment variable.
  - An error if the environment variable is not found or if there was an error loading the `.env` file.

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvement, please open an issue or submit a pull request on [GitHub](https://github.com/yourusername/goenv).

## License

This package is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

Feel free to further customize this README to better fit your package's needs!
