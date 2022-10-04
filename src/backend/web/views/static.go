package views

import "embed"

//go:embed static/assets/* shared/*.html static/index.html static/logo.png static/favicon.gif
var Static embed.FS
