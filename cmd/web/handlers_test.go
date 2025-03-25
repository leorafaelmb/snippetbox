package main

import (
	"io"
	"log/slog"
	"net/http"
	"snippetbox.leorafaelmb.net/internal/assert"
	"testing"
)

func TestPing(t *testing.T) {
	app := &application{
		logger: slog.New(slog.NewTextHandler(io.Discard, nil)),
	}

	ts := newTestServer(t, app.routes())
	defer ts.Close()

	code, _, body := ts.get(t, "/ping")

	assert.Equal(t, code, http.StatusOK)
	assert.Equal(t, body, "OK")
}
