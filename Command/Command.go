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
	PluginImports []string
	PlayerFuncs   []*PluginFunction.Function
}

// AddPlayerFunc adds a function with the given name and string to be put in the Java code
func (cmd *Command) AddPlayerFunc(name, strFunc string) {
	cmd.PlayerFuncs = append(cmd.PlayerFuncs, &PluginFunction.Function{Name: name, Func: strFunc})
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
