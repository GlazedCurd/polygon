package internal

import (
	"os"
	"path/filepath"

	"github.com/charmbracelet/log"
)

const (
	namedTemplatesFolderName = "templates"
	baseTemplateFolderName   = "base"
)

func InitProject(templateName string, templateDir string, log *log.Logger) {
	log.Debug("Initialising a project")
	templatePath := filepath.Join(templateDir, templateName)
	log.Debugf("Looking for template in %s", templatePath)
	currentDirectory, err := os.Getwd()
	if err != nil {
		log.Errorf("Failed, to obtain current working directory: %s", err)
		return
	}

	log.Debugf("Starting to copy %s to working directory: %s", templatePath, currentDirectory)

	err = os.CopyFS(templatePath, os.DirFS(currentDirectory))
	if err != nil {
		log.Errorf("Failed, to copy template: %s", err)
		return
	}

	// TODO: сделать вычитку конфига и мёрж конфигов и
}
