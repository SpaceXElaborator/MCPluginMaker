package plugincommands

import (
	PluginFunction "SpaceXElaborator/PluginMaker/Function"
	"strings"
)

// Command holds the information that will be passed on to the build phase for how to structure a command in Java
type Command struct {
	Author        string
	CommandType   string
	Name          string
	SlashCommand  string
	CommandLevel  int
	PluginImports []string
	PlayerFuncs   []*PluginFunction.Function
	SubCommands   []*Command
}

// AddPlayerFunc adds a function with the given name and string to be put in the Java code
func (cmd *Command) AddPlayerFunc(name, strFunc string) {
	cmd.PlayerFuncs = append(cmd.PlayerFuncs, &PluginFunction.Function{Name: name, Func: strFunc})
}

func (cmd *Command) RemPlayerFunc(name string) {

}

// AddSubCommand Adds a child command to the main command (I.E /{main command} {sub command} {further sub command})
func (cmd *Command) AddSubCommand(sub *Command) {
	for _, subs := range cmd.SubCommands {
		if strings.EqualFold(subs.SlashCommand, sub.SlashCommand) {
			return
		}
	}
	sub.CommandLevel = cmd.CommandLevel + 1
	cmd.SubCommands = append(cmd.SubCommands, sub)
}

// AddImport adds imports, but ignores duplicates, to a command to add dependencies neededd for building
func (cmd *Command) AddImport(imp string) {
	for _, str := range cmd.PluginImports {
		if strings.EqualFold(str, imp) {
			return
		}
	}

	cmd.PluginImports = append(cmd.PluginImports, imp)
}
