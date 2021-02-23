package PluginCommands

import (
	PluginFunction "SpaceXElaborator/PluginMaker/Function"
)

type Command struct {
	Author       string
	CommandType  string
	Name         string
	SlashCommand string
	PlayerFuncs  []*PluginFunction.Function
}

func (cmd *Command) AddPlayerFunc(name, strFunc string) {
	cmd.PlayerFuncs = append(cmd.PlayerFuncs, &PluginFunction.Function{Name: name, Func: strFunc})
}
