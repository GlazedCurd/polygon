package internal

import (
	"fmt"

	"github.com/charmbracelet/log"
)

type Project struct {
	cfg    *ProjectConfig // This is for prototype fix it later
	logger *log.Logger
}

func NewProject(conf *ProjectConfig, logger *log.Logger) (*Project, error) {
	if conf == nil {
		return nil, fmt.Errorf("no config provided")
	}
	result := &Project{
		cfg:    conf,
		logger: logger,
	}
	return result, nil
}
