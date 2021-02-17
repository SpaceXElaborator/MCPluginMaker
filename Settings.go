package main

import (
	"io/ioutil"
	"os"
	"encoding/json"
	"regexp"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/theme"
)

var (
	dark bool = false
	author string = "User"
)

func createDirs() {
	os.MkdirAll("projects", os.ModePerm)
}

// Made to be sure that settings are saved when closing/reopening the application
func Save() {
	data := Settings {
		Author: GetAuthor(),
		Dark: GetDark(),
	}
	
	file, _ := json.MarshalIndent(data, "", " ")
	_ = ioutil.WriteFile("settings.json", file, os.ModePerm)
}

func initSettings() {
	createDirs()
	
	// Read from the .json settings file and set the variables
	settingsFile, err := os.Open("settings.json")
	defer settingsFile.Close()
	if err != nil {
		return
	}
	byteValue, _ := ioutil.ReadAll(settingsFile)
	var set Settings
	json.Unmarshal(byteValue, &set)
	dark = set.Dark
	author = set.Author
	
	// Reset the title so it reflects the author and set the theme based on the settings
	// GetWindow() in MainWindow.go
	GetWindow().SetTitle("MCPluginMaker | " + GetAuthor() + " | Project: " + CWP)
	
	if dark == true {
		// GetApp() in MainWindow.go
		GetApp().Settings().SetTheme(theme.DarkTheme())
	} else {
		// GetApp() in MainWindow.go
		GetApp().Settings().SetTheme(theme.LightTheme())
	}
	
	projects, err := ioutil.ReadDir("./projects")
	if err != nil {
		log.Fatal(err)
	}
	
	// Read every folder in the /projects folder to reload projects made
	for _, f := range projects {
		content, err := ioutil.ReadFile("./projects/" + f.Name() + "/pom.xml")
		if err != nil {
			log.Fatal(err)
		}
		
		// Get the pom.xml and use Regex to load the projects and create new Project strucs to save
		pomString := string(content)
		group := regexp.MustCompile(`<groupId>(.*)</groupId>`)
		artifact := regexp.MustCompile(`<artifactId>(.*)</artifactId>`)
		description := regexp.MustCompile(`<description>(.*)</description>`)
		groupId := group.FindStringSubmatch(pomString)
		artifactId := artifact.FindStringSubmatch(pomString)
		descriptionField := description.FindStringSubmatch(pomString)
		newProj := Project{f.Name(), GetAuthor(), groupId[1], artifactId[1], descriptionField[1], []Command{}}
		
		// Check if there are any commands that needs to be loaded
		cmds, err := ioutil.ReadDir("./projects/" + f.Name() + "/src/main/java/com/" + GetAuthor() + "/net/cmds")
		if err == nil {
			for _, cmd := range cmds {
				content, err := ioutil.ReadFile("./projects/" + f.Name() + "/src/main/java/com/terturl/net/cmds/" + cmd.Name())
				if err != nil {
					continue
				}
				cmdString := string(content)
				slashCmd := regexp.MustCompile(`if\(label.equalsIgnoreCase\("(.*)"\)\)`)
				slashString := slashCmd.FindStringSubmatch(cmdString)
				if slashString == nil {
					continue
				}
				newProj.Cmds = append(newProj.Cmds, Command{GetAuthor(), cmd.Name()[0:len(cmd.Name()) - 5], slashString[1]})
			}
		}
		
		PluginProjects = append(PluginProjects, newProj)
	}
}

func createSettingsForm() *widget.Form {
	themeCheck := widget.NewCheck("", func(on bool) {
			if on == true {
				// GetApp() in MainWindow.go
				GetApp().Settings().SetTheme(theme.DarkTheme())
				dark = true
			} else {
				// GetApp() in MainWindow.go
				GetApp().Settings().SetTheme(theme.LightTheme())
				dark = false
			}
		})
	themeCheck.SetChecked(dark)
	checkFormItem := &widget.FormItem {
		Text: "Dark Mode",
		Widget: themeCheck,
	}
	
	authorNameEntry := widget.NewEntry()
	authorNameEntry.Resize(fyne.NewSize(300, 300))
	authorNameEntry.SetText(author)
	authorNameFormItem := &widget.FormItem {
		Text: "Author",
		Widget: authorNameEntry,
	}
	
	newSettingsForm := widget.NewForm(authorNameFormItem, checkFormItem)
	newSettingsForm.OnSubmit = func() {
		if authorNameEntry.Text != "" {
			author = authorNameEntry.Text
			// GetWindow() in MainWindow.go
			GetWindow().SetTitle("MCPluginMaker | " + author + " Project: " + CWP)
		} else {
			author = "User"
		}
		// HideModal() in MainWindow.go
		HideModal()
	}
	newSettingsForm.OnCancel = func() {
		// HideModal() in MainWindow.go
		HideModal()
	}
	return newSettingsForm
}

func GetAuthor() string {
	return author
}

func GetDark() bool {
	return dark
}