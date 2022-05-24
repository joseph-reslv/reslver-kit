package types


type CommandFlag struct {
	Debug bool
	ConfigsPath string
	ConfigYAML string
	InputPath string
	OutputPath string
	Force bool
}

type InitFunc func(flags *CommandFlag, root string) (error)

type ApplyFunc func(flags *CommandFlag, root string) (error)
