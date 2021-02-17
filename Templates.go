package main

var (
// The Main.java that needs to catch and set the executor classes of all commands
// and soon will add support for listening on listeners
mainJavaTmpl = `package com.{{.Author}}.net;

import org.bukkit.plugin.java.JavaPlugin;
{{if .Cmds}}
import com.{{.Author}}.net.cmds.*;{{end}}

public class Main extends JavaPlugin {

	public void onEnable() {
	{{with .Cmds -}}{{range $val := .}}	getCommand("{{$val.SlashCommand}}").setExecutor(new {{$val.Name}}());
	{{end}}{{end}}
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