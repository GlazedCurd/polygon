package internal

import "fmt"

type Project struct {
}

func NewProject(conf *projectConfig) (*Project, error) {
	if conf == nil {
		return nil, fmt.Errorf("no config provided")
	}
	return nil, nil
}
