package PluginTemplates

func GetCommandTemplate() string {
	return `package com.{{.Author}}.net.cmds;

import org.bukkit.command.Command;
import org.bukkit.command.CommandExecutor;
import org.bukkit.command.CommandSender;
import org.bukkit.entity.Player;

import com.{{.Author}}.net.*;

public class {{.Name}} implements CommandExecutor {

	@Override
	public boolean onCommand(CommandSender cs, Command cmd, String label, String[] args) {
		// <<CMDTYPE:{{.CommandType}}>>
		{{ if eq $.CommandType "Player" }}Player p = (Player)cs;{{end}}
		if(label.equalsIgnoreCase("{{.SlashCommand}}")) {
			{{ if eq $.CommandType "Player" }}
			{{with .CmdFuncs -}}{{range $strings := .}}// <<CMDSTRING:{{.CommandString}}>>
			{{.CommandString}}
			{{end}}{{end}}
			{{end}}
		}
		return false;
	}
	
}
`
}