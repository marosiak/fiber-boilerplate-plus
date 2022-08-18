package static

import (
	"embed"
	"net/http"
)

// This will embed all files and folders in this directory where the source file is
//go:embed *
var content embed.FS

func GetFiles() http.FileSystem {
	return http.FS(content)
}
