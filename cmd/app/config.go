package app

import (
	"embed"

	"html/template"
)

type AppConfig struct {
	EmbeddedFiles embed.FS
}

func (a *AppConfig) GetParsedTemplates(paths ...string) *template.Template {
	return template.Must(template.ParseFS(a.EmbeddedFiles, paths...))
}
