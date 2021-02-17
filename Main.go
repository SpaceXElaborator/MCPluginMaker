package main

var (
	PluginProjects []Project
)

func Test() {
	// ------------- Test One -------------
	//createNewProject("test", PomXML{"terturl", "test", "Test", "Test", "This is a test"})
	//createCommand(Command{GetAuthor(), "HealCommand", "Heal"})
	//createCommand(Command{GetAuthor(), "TeleportCommand", "Teleport"})
	//createCommand(Command{GetAuthor(), "KillCommand", "Kill"})
	//createCommand(Command{GetAuthor(), "PrintSoureCode", "psc"})
	//build()
	// -------------------------------------
}

func GetProject(name string) *Project {
	for _, f := range PluginProjects {
		if f.Name == name {
			return &f
		}
	}
	return nil
}

func ProjectExists(name string) bool {
	for _, f := range PluginProjects {
		if f.Name == name {
			return true
		}
	}
	return false
}

func main() {
	initSettings()
	// Testing Only
	Test()
	createDirs()
	ShowMainMenu()
	Save()
}