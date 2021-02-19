package main

var (
// The Main.java that needs to catch and set the executor classes of all commands
// and soon will add support for listening on listeners
mainJavaTmpl = `package com.{{.Author}}.net;

import org.bukkit.plugin.java.JavaPlugin;
{{if .Cmds}}
import com.{{.Author}}.net.cmds.*;{{end}}

public class Main extends JavaPlugin {

	private static Main instance;

	public void onEnable() {
		instance = this;
	{{with .Cmds -}}{{range $val := .}}	getCommand("{{$val.SlashCommand}}").setExecutor(new {{$val.Name}}());
	{{end}}{{end}}
	
		{{.Name}}CustomItems.RegisterItems();
	}
	
	public static Main getInstance() {
		return instance;
	}

}
`

// Creates ALL the custom items that the user wants
itemsJavaTmpl = `package com.{{.Author}}.net;

import java.util.ArrayList;
import java.util.List;

import org.bukkit.Material;
import org.bukkit.inventory.ItemStack;
import org.bukkit.inventory.meta.ItemMeta;

public class {{.Name}}CustomItems {
	public static List<ItemStack> items = new ArrayList<ItemStack>();
	
	public static void RegisterItems() {
		{{with .Items -}}{{range $val := .}}
		// <<ITEM:{{$val.ItemMaterial}}||{{$val.ItemName}}||{{with $val.ItemDescription -}}{{range $strings := .}}{{.}}\n{{end}}{{end}}>>
		ItemStack {{$val.ItemName}}Item = new ItemStack(Material.valueOf("{{$val.ItemMaterial}}"));
		ItemMeta {{$val.ItemName}}Meta = {{$val.ItemName}}Item.getItemMeta();
		{{$val.ItemName}}Meta.setDisplayName("{{$val.ItemName}}");
		List<String> {{$val.ItemName}}Lore = new ArrayList<String>();
		{{with $val.ItemDescription -}}{{range $strings := .}}{{$val.ItemName}}Lore.add("{{.}}");
		{{end}}{{end}}{{$val.ItemName}}Meta.setLore({{$val.ItemName}}Lore);
		{{$val.ItemName}}Item.setItemMeta({{$val.ItemName}}Meta);
		items.add({{$val.ItemName}}Item);
		{{end}}{{end}}
	}
	
	public static ItemStack getItem(String name) {
		for(ItemStack item : items) {
			if(item.getItemMeta().getDisplayName().equalsIgnoreCase(name)) {
				return item;
			}
		}
		return null;
	}
	
}

`

// Create commands that will here soon have more functionality
cmdJavaTmpl = `package com.{{.Author}}.net.cmds;

import org.bukkit.command.Command;
import org.bukkit.command.CommandExecutor;
import org.bukkit.command.CommandSender;
import org.bukkit.entity.Player;

public class {{.Name}} implements CommandExecutor {

	@Override
	public boolean onCommand(CommandSender cs, Command cmd, String label, String[] args) {
		// <<TYPE:{{.CommandType}}>>
		if(label.equalsIgnoreCase("{{.SlashCommand}}")) {
			
		}
		return false;
	}
	
}
`
// Create the plugin.yml using authors names, commands, and the plugin name
// to be able to read the plugin when built
pluginYmlTmpl = `main: com.{{.Author}}.net.Main
name: {{.Name}}
author: {{.Author}}
version: 1
api-version: 1.15
{{with .Cmds -}}
commands:{{range $val := .}}
  {{.SlashCommand}}:{{end}}
{{end}}
`

pomXmlTmpl = `<project xmlns="http://maven.apache.org/POM/4.0.0"
	xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
	xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 https://maven.apache.org/xsd/maven-4.0.0.xsd">
	<modelVersion>4.0.0</modelVersion>
	<groupId>{{.GroupID}}</groupId>
	<artifactId>{{.ArtifactID}}</artifactId>
	<version>1.0</version>
	<name>{{.Name}}</name>
	<description>{{.Description}}</description>

	<repositories>
		<repository>
			<id>spigot-repo</id>
			<url>https://hub.spigotmc.org/nexus/content/repositories/snapshots/</url>
		</repository>
	</repositories>

	<properties>
		<maven.compiler.source>1.8</maven.compiler.source>
		<maven.compiler.target>1.8</maven.compiler.target>
	</properties>

	<dependencies>
		<dependency>
			<groupId>org.spigotmc</groupId>
			<artifactId>spigot-api</artifactId>
			<version>1.16.4-R0.1-SNAPSHOT</version>
			<scope>provided</scope>
		</dependency>
	</dependencies>

	<build>
		<sourceDirectory>src/main/java</sourceDirectory>
		<resources>
			<resource>
				<directory>src/main/java</directory>
				<includes>
					<include>plugin.yml</include>
				</includes>
			</resource>
		</resources>
		<plugins>
			<plugin>
				<groupId>org.apache.maven.plugins</groupId>
				<artifactId>maven-compiler-plugin</artifactId>
				<version>3.8.1</version>
			</plugin>
			<plugin>
				<artifactId>maven-assembly-plugin</artifactId>
				<configuration>
					<archive>
						<manifest>
							<mainClass>
								com.{{.Author}}.net.Main
							</mainClass>
						</manifest>
					</archive>
					<descriptorRefs>
						<descriptorRef>jar-with-dependencies</descriptorRef>
					</descriptorRefs>
				</configuration>
			</plugin>
		</plugins>
	</build>

</project>`

)