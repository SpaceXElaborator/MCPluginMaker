package PluginCommands

import (
	PluginFunction "SpaceXElaborator/PluginMaker/Function"
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
	cmd.PluginImports = append(cmd.PluginImports, imp)
}
