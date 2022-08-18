package templates

import (
	"embed"
	"net/http"
)

//go:embed *.html
var content embed.FS

func GetFiles() http.FileSystem {
	return http.FS(content)
}
