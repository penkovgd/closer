package closer

import (
	"io"
	"log/slog"
)

func CloseOrLog(log *slog.Logger, closer io.Closer) {
	if err := closer.Close(); err != nil {
		log.Error("failed to close", "error", err.Error())
		return
	}
}
func CloseOrPanic(log *slog.Logger, closer io.Closer) {
	if err := closer.Close(); err != nil {
		log.Error("panic: close error", "error", err.Error())
		panic("close error: " + err.Error())
	}
}
