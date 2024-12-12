package internal

type CommandCfg struct {
	Cmd  string `mapstructure:"cmd" json:"cmd"`
	Args string `mapstructure:"args" json:"args"`
}

type ExecConfig struct {
	BuildCmd         CommandCfg `mapstructure:"build_cmd" json:"build_cmd"`
	RunCmd           CommandCfg `mapstructure:"run_cmd" json:"run_cmd,omitempty"`
	WorkingDirectory string     `mapstructure:"working_directory" json:"working_directory,omitempty"`
}

type ProjectConfig struct {
	Main           ExecConfig            `mapstructure:"main" json:"main"`
	Light          *ExecConfig           `mapstructure:"light" json:"light"`
	InputGenerator CommandCfg            `mapstructure:"input_generator" json:"input_generator"`
	Custom         map[string]ExecConfig `mapstructure:"custom" json:"custom"`
}
