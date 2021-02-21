package main

import (
	"SpaceXElaborator/PluginMaker/Project"
	"SpaceXElaborator/PluginMaker/GUI"
	"SpaceXElaborator/PluginMaker/Settings"
)

var (
	AllProjects = PluginProject.Projects{[]*PluginProject.Project{}}
	
	Test = "Hello"
)

func main() {
	PluginSettings.InitSettings(&AllProjects)
	PluginSettings.CreateDirs()
	PluginGUI.ResetSettings()
	PluginGUI.ShowMainMenu(&AllProjects)
	PluginSettings.Save()
}