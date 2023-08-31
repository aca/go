package q

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
)

var logger *slog.Logger

func init() {
	path := filepath.Join(os.TempDir(), "q")
	f, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o777)
	if err != nil {
		panic(fmt.Errorf("failed to init q logger %q: %w", path, err))
	}

	logger = slog.New(slog.NewTextHandler(f, &slog.HandlerOptions{
		AddSource: true,
	}))
	// finalizer?
}

func Q(v ...any) {
	logger.Warn("", v...)
}
