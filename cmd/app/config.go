package app

import "embed"

type AppConfig struct {
	EmbeddedFiles embed.FS
}
