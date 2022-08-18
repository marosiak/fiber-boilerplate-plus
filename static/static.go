package static

import (
	"embed"
	"net/http"
)

//go:embed public/* private/*
var content embed.FS

func GetFiles() http.FileSystem {
	return http.FS(content)
}
