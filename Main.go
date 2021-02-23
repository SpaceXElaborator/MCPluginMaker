package main

import (
	PluginGUI "SpaceXElaborator/PluginMaker/GUI"
	PluginProject "SpaceXElaborator/PluginMaker/Project"
	PluginSettings "SpaceXElaborator/PluginMaker/Settings"
)

var (
	allProjects = PluginProject.Projects{Projects: []*PluginProject.Project{}}
)

func main() {
	PluginSettings.InitSettings(&allProjects)
	PluginSettings.CreateDirs()
	PluginGUI.ResetSettings()
	PluginGUI.ShowMainMenu(&allProjects)
	PluginSettings.Save()
}
