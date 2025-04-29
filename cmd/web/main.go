package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger *slog.Logger
}

type neuteredFileSystem struct {
	fs http.FileSystem
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")

	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{}))

	app := &application{
		logger: logger,
	}

	logger.Info("starting server", "addr", *addr)

	err := http.ListenAndServe(*addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}
