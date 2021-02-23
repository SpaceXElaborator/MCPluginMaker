package pluginproject

import (
	"errors"
	"log"
	"os"
	"os/exec"
	"strings"

	"text/template"

	"fyne.io/fyne/v2"

	PluginCommands "SpaceXElaborator/PluginMaker/Command"
	PluginFunction "SpaceXElaborator/PluginMaker/Function"
	PluginItems "SpaceXElaborator/PluginMaker/Item"
	PluginTemplates "SpaceXElaborator/PluginMaker/Templates"
)

// Project Holds all relevent information about a Java project
type Project struct {
	Name, Author string
	Xml          PomXML
	Cmds         []*PluginCommands.Command
	Items        []*PluginItems.CustomItem
}

// PomXML Holds all relevent information about a Pom.xml file for Maven
type PomXML struct {
	GroupID, ArtifactID, Description string
}

// AddCommand Adds a command with the given name, slash, and cmdType
func (proj *Project) AddCommand(name, slash, cmdType string) error {
	for _, cmds := range proj.Cmds {
		if strings.EqualFold(cmds.Name, name) {
			return errors.New("Command Exists")
		}
		if strings.EqualFold(cmds.SlashCommand, slash) {
			return errors.New("Slash String Exists")
		}
	}

	cmd := PluginCommands.Command{Author: proj.Author, CommandType: cmdType, Name: name, SlashCommand: slash, PluginImports: []string{}, PlayerFuncs: []*PluginFunction.Function{}}
	proj.Cmds = append(proj.Cmds, &cmd)
	os.MkdirAll("projects/"+proj.Name+"/src/main/java/com/"+proj.Author+"/net/cmds", os.ModePerm)
	return nil
}

// GetCommand Returns a PluginCommand with the same name
func (proj *Project) GetCommand(name string) *PluginCommands.Command {
	for _, cmds := range proj.Cmds {
		if strings.EqualFold(cmds.Name, name) {
			return cmds
		}
	}
	return nil
}

// RemoveCommand Will delete the command from the project as well as deletes the file
func (proj *Project) RemoveCommand(name string) {
	var i int
	for ind, cmds := range proj.Cmds {
		if strings.EqualFold(cmds.Name, name) {
			i = ind
		}
	}

	os.Remove("projects/" + proj.Name + "/src/main/java/com/" + proj.Author + "/net/cmds/" + name + ".java")

	copy(proj.Cmds[i:], proj.Cmds[i+1:])
	proj.Cmds = proj.Cmds[:len(proj.Cmds)-1]
}

// CreatePom Creates the Pom.xml file using templates
func (proj *Project) CreatePom() {
	pom, err := os.Create("projects/" + proj.Name + "/pom.xml")
	if err != nil {
		log.Fatal("Unable to create \"pom.xml\" file: ", err)
	}
	templ := template.Must(template.New("CreatePOM").Parse(PluginTemplates.GetPomTemplate()))
	err = templ.Execute(pom, proj)
	if err != nil {
		log.Fatal("Unable to execute Pom Creation Template: ", err)
	}
	pom.Close()
}

// createYaml Creates the plugin.yml file using templates
func (proj *Project) createYaml() {
	yml, err := os.Create("projects/" + proj.Name + "/src/main/java/plugin.yml")
	if err != nil {
		log.Fatal("Unable to create \"plugin.yml\" file: ", err)
	}
	templ := template.Must(template.New("CreateYML").Parse(PluginTemplates.GetYamlTemplate()))
	err = templ.Execute(yml, proj)
	if err != nil {
		log.Fatal("Unable to execute Yaml Creation Template: ", err)
	}
	yml.Close()
}

// createMainJava Creates the Main.java file using templates
func (proj *Project) createMainJava() {
	java, err := os.Create("projects/" + proj.Name + "/src/main/java/com/" + proj.Author + "/net/Main.java")
	if err != nil {
		log.Fatal("Unable to create \"Main.java\" file: ", err)
	}
	templ := template.Must(template.New("CreateMain").Parse(PluginTemplates.GetMainClassTemplate()))
	err = templ.Execute(java, proj)
	if err != nil {
		log.Fatal("Unable to execute Main Creation Template: ", err)
	}
	java.Close()
}

// createItemUtilClass Creates the Item.java util file using templates
func (proj *Project) createItemUtilClass() {
	item, err := os.Create("projects/" + proj.Name + "/src/main/java/com/" + proj.Author + "/net/Item.java")
	if err != nil {
		log.Fatal("Unable to create \"Item.java\" file: ", err)
	}
	templ := template.Must(template.New("ItemUtil").Parse(PluginTemplates.GetItemUtilClassTemplate()))
	err = templ.Execute(item, proj)
	if err != nil {
		log.Fatal("Unable to execute Item Creation Template: ", err)
	}
	item.Close()
}

// createItemClass Creates the {ProjectName}CustomItems.java file using templates to register items
func (proj *Project) createItemClass() {
	item, err := os.Create("projects/" + proj.Name + "/src/main/java/com/" + proj.Author + "/net/" + proj.Name + "CustomItems.java")
	if err != nil {
		log.Fatal("Unable to create \"ItemClass.java\" file: ", err)
	}
	templ := template.Must(template.New("CreateItems").Parse(PluginTemplates.GetItemTemplate()))
	err = templ.Execute(item, proj)
	if err != nil {
		log.Fatal("Unable to execute Item Creation Template: ", err)
	}
	item.Close()
}

// buildCommands Creates each command's java file using templates
func (proj *Project) buildCommands() {
	for _, cmd := range proj.Cmds {
		cmdF, err := os.Create("projects/" + proj.Name + "/src/main/java/com/" + proj.Author + "/net/cmds/" + cmd.Name + ".java")
		if err != nil {
			log.Fatal("Unable to create Command file: ", err)
		}
		templ := template.Must(template.New("CreateCommand").Parse(PluginTemplates.GetCommandTemplate()))
		err = templ.Execute(cmdF, &cmd)
		if err != nil {
			log.Fatal("Unable to execute Command Creation Template: ", err)
		}
		cmdF.Close()
	}
}

// Build Runs all template-building functions and then builds the command using Maven
func (proj *Project) Build() {
	proj.createYaml()
	proj.createMainJava()

	if len(proj.Cmds) > 0 {
		proj.buildCommands()
	}

	if len(proj.Items) > 0 {
		proj.createItemUtilClass()
		proj.createItemClass()
	}

	mvnCmd := exec.Command("mvn", "-f", "./projects/"+proj.Name+"/pom.xml", "package")
	mvnCmd.Dir = "."
	_, err := mvnCmd.Output()
	if err != nil {
		log.Print("Couldn't run build command: ", err)
	}
	fyne.CurrentApp().SendNotification(&fyne.Notification{
		Title:   "Build Complete",
		Content: "Building of " + proj.Name + " complete",
	})
}

// AddItem Adds an item using a material, name, and list of string description
func (proj *Project) AddItem(itemMat, itemName string, itemDesc []string) {
	newItem := PluginItems.CustomItem{ItemMaterial: itemMat, ItemName: itemName, ItemDescription: itemDesc}
	proj.Items = append(proj.Items, &newItem)
}

// CheckItem Checks if the CustomItem exists or not already
func (proj *Project) CheckItem(s string) bool {
	for _, f := range proj.Items {
		if strings.EqualFold(f.ItemName, s) {
			return true
		}
	}
	return false
}
