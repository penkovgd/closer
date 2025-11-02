# Closer

A minimal Go utility for safe resource cleanup with integrated logging and panic options.

## Installation

```bash
go get github.com/penkovgd/closer
```

## Usage

```golang
import (
    "log/slog"
    "os"
    "github.com/penkovgd/closer"
)

func main() {
    logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

    file, err := os.Open("data.txt")
    if err != nil {
        panic(err)
    }
    defer closer.CloseOrLog(logger, file)

    // For critical resources that must close successfully
    criticalResource := acquireResource()
    defer closer.CloseOrPanic(logger, criticalResource)
}
```

## Functions

`CloseOrLog(log *slog.Logger, closer io.Closer)`

Closes the resource and logs any error using the provided logger. Safe for deferred calls.

`CloseOrPanic(log *slog.Logger, closer io.Closer)`

Closes the resource and panics if an error occurs. Use for critical resources where failed closure is unrecoverable.
