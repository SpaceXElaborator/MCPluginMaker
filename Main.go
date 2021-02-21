package main

import (
	PluginGUI "SpaceXElaborator/PluginMaker/GUI"
	PluginProject "SpaceXElaborator/PluginMaker/Project"
	PluginSettings "SpaceXElaborator/PluginMaker/Settings"
)

var (
	AllProjects = PluginProject.Projects{Projects: []*PluginProject.Project{}}

	Test = "Hello"
)

func main() {
	PluginSettings.InitSettings(&AllProjects)
	PluginSettings.CreateDirs()
	PluginGUI.ResetSettings()
	PluginGUI.ShowMainMenu(&AllProjects)
	PluginSettings.Save()
}
