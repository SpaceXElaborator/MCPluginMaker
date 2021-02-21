package PluginProject

import (
	"strings"
	"errors"
	"os"
	"os/exec"
	"log"
	
	"text/template"
	
	"fyne.io/fyne/v2"
	
	"SpaceXElaborator/PluginMaker/Templates"
	"SpaceXElaborator/PluginMaker/Command"
	"SpaceXElaborator/PluginMaker/Item"
)

type Project struct {
	Name, Author string
	Xml PomXML
	Cmds []*PluginCommands.Command
	Items []*PluginItems.CustomItem
}

type PomXML struct {
	GroupID, ArtifactID, Description string
}

func (proj *Project) AddCommand(name, slash, cmdType string) error {
	for _, cmds := range proj.Cmds {
		if strings.EqualFold(cmds.Name, name) {
			return errors.New("Command Exists")
		}
		if strings.EqualFold(cmds.SlashCommand, slash) {
			return errors.New("Slash String Exists")
		}
	}
	
	cmd := PluginCommands.Command{proj.Author, cmdType, name, slash, []*PluginCommands.CommandFunc{}}
	proj.Cmds = append(proj.Cmds, &cmd)
	os.MkdirAll("projects/" + proj.Name + "/src/main/java/com/" + proj.Author + "/net/cmds", os.ModePerm)
	return nil
}

func (proj *Project) GetCommand(name string) *PluginCommands.Command {
	for _, cmds := range proj.Cmds {
		if strings.EqualFold(cmds.Name, name) {
			return cmds
		}
	}
	return nil
}

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

func (proj *Project) Build() {
	proj.createYaml()
	proj.createMainJava()
	
	if len(proj.Cmds) > 0 {
		proj.buildCommands()
	}
	
	if len(proj.Items) > 0 {
		proj.createItemClass()
	}
	
	
	mvnCmd := exec.Command("mvn", "-f", "./projects/" + proj.Name + "/pom.xml", "package")
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

func (proj *Project) AddItem(itemMat, itemName string, itemDesc []string) {
	newItem := PluginItems.CustomItem{itemMat, itemName, itemDesc}
	proj.Items = append(proj.Items, &newItem)
}

func (proj *Project) CheckItem(s string) bool {
	for _, f := range proj.Items {
		if strings.EqualFold(f.ItemName, s) {
			return true
		}
	}
	return false
}

func (proj *Project) CheckMaterial(a string) bool {
	for _, item := range PluginItems.SpigotMaterialList {
		if item == a {
			return true
		}
	}
	return false
}