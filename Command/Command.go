package PluginCommands

import (
	PluginFunction "SpaceXElaborator/PluginMaker/Function"
	"strings"
)

type Command struct {
	Author        string
	CommandType   string
	Name          string
	SlashCommand  string
	PluginImports []string
	PlayerFuncs   []*PluginFunction.Function
}

func (cmd *Command) AddPlayerFunc(name, strFunc string) {
	cmd.PlayerFuncs = append(cmd.PlayerFuncs, &PluginFunction.Function{Name: name, Func: strFunc})
}

func (cmd *Command) AddImport(imp string) {
	for _, str := range cmd.PluginImports {
		if strings.EqualFold(str, imp) {
			return
		}
	}

	cmd.PluginImports = append(cmd.PluginImports, imp)
}
