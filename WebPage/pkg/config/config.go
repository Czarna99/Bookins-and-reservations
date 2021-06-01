package config

import (
	"html/template"
	"log"
)

type AppConfig struct {
	TemplateChace map[string]*template.Template
	UseCache      bool
	InfoLog       *log.Logger
}
