package PluginTemplates

func GetCommandTemplate() string {
	return `package com.{{.Author}}.net.cmds;

import org.bukkit.command.Command;
import org.bukkit.command.CommandExecutor;
import org.bukkit.command.CommandSender;
import org.bukkit.entity.Player;
{{with .PluginImports -}}{{range $strings := .}}
// <<IMPORT:{{.}}>>
{{.}}{{end}}
{{end}}
import org.bukkit.Material;
import org.bukkit.inventory.ItemStack;

import com.{{.Author}}.net.*;

public class {{.Name}} implements CommandExecutor {

	@Override
	public boolean onCommand(CommandSender cs, Command cmd, String label, String[] args) {
		// <<CMDTYPE:{{.CommandType}}>>
		{{ if eq $.CommandType "Player" }}Player p = (Player)cs;{{end}}
		if(cmd.getName().equalsIgnoreCase("{{.SlashCommand}}")) {
			{{ if eq $.CommandType "Player" }}
			{{with .PlayerFuncs -}}{{range $strings := .}}// <<CMDSTRING:{{.Name}}||{{.Func}}>>
			{{.Func}}
			{{end}}{{end}}
			{{end}}
		}
		return false;
	}
	
}
`
}
