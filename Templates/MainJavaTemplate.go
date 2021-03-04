package plugintemplates

// GetMainClassTemplate Outputs the string template for the Main.java class
func GetMainClassTemplate() string {
	return `package com.{{.Author}}.net;

import org.bukkit.plugin.java.JavaPlugin;
{{if .Cmds}}
import com.{{.Author}}.net.cmds.*;{{end}}

public class Main extends JavaPlugin {

	private static Main instance;

	public void onEnable() {
		instance = this;
	{{with .Cmds -}}{{range $val := .}}	getCommand("{{$val.SlashCommand}}").setExecutor(new {{$val.Name}}());
	{{end}}{{end}}
	}
	
	public static Main getInstance() {
		return instance;
	}

}
`
}
