package views

import "embed"

//go:embed assets/* shared/*.html index.html logo.png favicon.gif
var Static embed.FS
