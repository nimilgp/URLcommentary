package assets

import (
	"embed"
)

//go:embed "emails" "migrations"
var EmbeddedFiles embed.FS
