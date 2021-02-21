package PluginCommands

type Command struct {
	Author string
	CommandType string
	Name string
	SlashCommand string
	CmdFuncs []*CommandFunc
}

type CommandFunc struct {
	CommandString string
}

func (cmd *Command) AddFunc(strFunc string) {
	cmd.CmdFuncs = append(cmd.CmdFuncs, &CommandFunc{strFunc})
}