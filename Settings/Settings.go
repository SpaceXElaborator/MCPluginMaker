package PluginSettings

import (
	"io/ioutil"
	"os"
	"log"
	"encoding/json"
	"regexp"
	"strings"
	
	"SpaceXElaborator/PluginMaker/Project"
)

var (
	CWP = ""
	author = "User"
	dark = false
	
	Projects *PluginProject.Projects
)

type Settings struct {
	Author string `json:"Author"`
	Dark bool `json:"Mode"`
}

func GetAuthor() string {
	return author
}

func SetAuthor(name string) {
	author = name
}

func GetCWP() string {
	return CWP
}

func SetCWP(name string) {
	CWP = name
}

func GetDark() bool {
	return dark
}

func SetDark(b bool) {
	dark = b
}

func CreateDirs() {
	os.MkdirAll("projects", os.ModePerm)
}

// Made to be sure that settings are saved when closing/reopening the application
func Save() {
	data := Settings{
		Author: GetAuthor(),
		Dark: GetDark(),
	}
	
	file, _ := json.MarshalIndent(data, "", " ")
	_ = ioutil.WriteFile("settings.json", file, os.ModePerm)
}

func InitSettings(projs *PluginProject.Projects) {
	Projects = projs

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
		err = Projects.CreateNewProject(f.Name(), GetAuthor(), groupId[1], artifactId[1], descriptionField[1])
		if err != nil {
			log.Print("Project Already Exists somehow! Skipping!")
			continue
		}
		
		newProj := Projects.GetProject(f.Name())
		
		_, err = os.Stat("./projects/" + f.Name() + "/src/main/java/com/" + GetAuthor() + "/net/" + f.Name() + "CustomItems.java")
		if err == nil {
			content, err := ioutil.ReadFile("./projects/" + f.Name() + "/src/main/java/com/" + GetAuthor() + "/net/" + f.Name() + "CustomItems.java")
			if err == nil {
				itemsString := string(content)
				itemsRegString := regexp.MustCompile(`<<ITEM:(.*)>>`)
				itemsInFile := itemsRegString.FindAllStringSubmatch(itemsString, -1)
				for _, v := range itemsInFile {
					itemMakerString := strings.Split(v[1], "||")
					tmp := []string{}
					for _, strs := range strings.Split(itemMakerString[2], `\n`) {
						if strs == "\n" || strs == "" {
							continue
						}
						tmp = append(tmp, strings.TrimRight(strs, "\n"))
					}
					newProj.AddItem(itemMakerString[0], itemMakerString[1], tmp)
				}
			}
		}
		
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
				cmdType := regexp.MustCompile(`<<CMDTYPE:(.*)>>`)
				functionsInCmd := regexp.MustCompile(`<<CMDSTRING:(.*)>>`)
				slashString := slashCmd.FindStringSubmatch(cmdString)
				cmdTypeField := cmdType.FindStringSubmatch(cmdString)
				funcsInCmd := functionsInCmd.FindAllStringSubmatch(cmdString, -1)
				if slashString == nil {
					continue
				}
				if cmdTypeField == nil {
					continue
				}
				
				newProj.AddCommand(cmd.Name()[0:len(cmd.Name()) - 5], slashString[1], cmdTypeField[1])
				createdCmd := newProj.GetCommand(cmd.Name()[0:len(cmd.Name()) - 5])
				
				for _, v := range funcsInCmd {
					createdCmd.AddFunc(v[1])
				}
			}
		}
	}
}