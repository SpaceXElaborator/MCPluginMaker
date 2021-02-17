package main

var (
	PluginProjects []Project
)

// Returns a pointer to a Project given the name
func GetProject(name string) *Project {
	for _, f := range PluginProjects {
		if f.Name == name {
			return &f
		}
	}
	return nil
}

// Only checks if the project actually exists (Subject to removal)
func ProjectExists(name string) bool {
	for _, f := range PluginProjects {
		if f.Name == name {
			return true
		}
	}
	return false
}

func main() {
	// initSettings()/createDirs() in Settings.go
	initSettings()
	createDirs()
	
	// ShowMainMenu() in MainWindow.go
	ShowMainMenu()
	
	// Save() in Settings.go
	Save()
}