package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

type AppConfig struct {
	TemplateChace map[string]*template.Template
	UseCache      bool
	InfoLog       *log.Logger
	InProduction  bool
	Session       *scs.SessionManager
}
