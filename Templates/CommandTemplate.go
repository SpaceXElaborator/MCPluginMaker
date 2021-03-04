package plugintemplates

// GetCommandTemplate Outputs the string template for creating a Command class as well as the needed imports if any
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
		// Makes sure the command that the player input is the command that was created
		if(cmd.getName().equalsIgnoreCase("{{.SlashCommand}}")) {
			{{ if eq $.CommandType "Player" }}
			{{with .PlayerFuncs -}}{{range $strings := .}}// <<CMDSTRING:{{.Name}}||{{.Func}}>>
			{{.Func}}
			{{end}}{{end}}
			{{end}}
			// Checks the arguments after the initial command to see if it equals any of the sub commands that was added
			{{with .SubCommands -}}{{range $SubCommand := .}}
			if(args[{{.CommandLevel}}].equalsIgnoreCase("{{.SlashCommand}}")) {
				{{ if eq $.CommandType "Player" }}
				{{with .PlayerFuncs -}}{{range $strings := .}}// <<CMDSTRING:{{.Name}}||{{.Func}}>>
				{{.Func}}
				{{end}}{{end}}
				{{end}}
			}
			{{end}}{{end}}
		}
		return false;
	}
	
}
`
}
