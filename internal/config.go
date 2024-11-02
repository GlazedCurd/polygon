package internal

type execConfig struct {
	BuildCmd         string `json:"build_cmd"`
	RunCmd           string `json:"run_cmd,omitempty"`
	WorkingDirectory string `json:"working_directory,omitempty"`
}

type projectConfig struct {
	Main           execConfig            `json:"main"`
	Light          *execConfig           `json:"light"`
	InputGenerator string                `json:"input_generator"`
	Custom         map[string]execConfig `json:"custom"`
}
