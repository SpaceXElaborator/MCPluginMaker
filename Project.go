package main

import (
	"os"
	"log"
	"errors"
	"os/exec"
	"text/template"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/dialog"
)

var (
	CWP string
)

func createNewProject(name string, xml PomXML) {
	os.MkdirAll("projects/" + name + "/src/main/java/com/" + GetAuthor() + "/net", os.ModePerm)
	GetWindow().SetTitle("MCPluginMaker | " + GetAuthor() + " | Project: " + CWP)
	createPom(xml)
	proj := Project{CWP, GetAuthor(), xml.GroupID, xml.ArtifactID, xml.Description, []Command{}}
	PluginProjects = append(PluginProjects, proj)
	list.Refresh()
}

func createPom(xml PomXML) {
	pom, err := os.Create("projects/" + CWP + "/pom.xml")
	if err != nil {
		log.Print("Error: ", err)
	}
	t := template.Must(template.New("CreatePOM").Parse(pomXmlTmpl))
	err = t.Execute(pom, &xml)
	if err != nil {
		log.Print("Error: ", err)
	}
	pom.Close()
}

func createYaml(proj *Project) {
	yml, err := os.Create("projects/" + CWP + "/src/main/java/plugin.yml")
	if err != nil {
		log.Print("Error: ", err)
	}
	//testStruct := Project{CWP, GetAuthor(), "", "", "", mainAdditions.Cmds}
	t := template.Must(template.New("CreateYAML").Parse(pluginYmlTmpl))
	err = t.Execute(yml, &proj)
	if err != nil {
		log.Print("Error: ", err)
	}
	yml.Close()
}

func createMainJava(proj *Project) {
	f, err := os.Create("projects/" + CWP + "/src/main/java/com/terturl/net/Main.java")
	if err != nil {
		log.Print("Error: ", err)
	}
	
	t := template.Must(template.New("CreateMain").Parse(mainJavaTmpl))
	err = t.Execute(f, &proj)
	if err != nil {
		log.Print("Error: ", err)
	}
	f.Close()
}

func build(proj *Project) {
	createYaml(proj)
	createMainJava(proj)
	mvnCmd := exec.Command("mvn", "-f", "./projects/" + CWP + "/pom.xml", "package")
	mvnCmd.Dir = "."
	output, err := mvnCmd.Output()
	if err != nil {
		log.Print(err)
	}
	log.Print(output)
}

func createNewProjectForm() *widget.Form {
	projectNameEntry := widget.NewEntry()
	projectNameEntry.Resize(fyne.NewSize(300, 300))
	projectNameEntry.SetText("")
	projectNameFormItem := &widget.FormItem {
		Text: "Project Name",
		Widget: projectNameEntry,
	}
	
	projectGroupEntry := widget.NewEntry()
	projectGroupEntry.Resize(fyne.NewSize(300, 300))
	projectGroupEntry.SetText("")
	projectGroupFormItem := &widget.FormItem {
		Text: "Group ID",
		Widget: projectGroupEntry,
	}
	
	projectArtifactEntry := widget.NewEntry()
	projectArtifactEntry.Resize(fyne.NewSize(300, 300))
	projectArtifactEntry.SetText("")
	projectArtifactFormItem := &widget.FormItem {
		Text: "Artifact ID",
		Widget: projectArtifactEntry,
	}
	
	projectDescriptionEntry := widget.NewEntry()
	projectDescriptionEntry.Resize(fyne.NewSize(300, 300))
	projectDescriptionEntry.SetText("")
	projectDescriptionFormItem := &widget.FormItem {
		Text: "Description",
		Widget: projectDescriptionEntry,
	}

	newProjectForm := widget.NewForm(projectNameFormItem, projectGroupFormItem, projectArtifactFormItem, projectDescriptionFormItem)
	newProjectForm.OnSubmit = func() {
		if projectNameEntry.Text != "" && projectArtifactEntry.Text != "" && projectDescriptionEntry.Text != "" && projectGroupEntry.Text != "" {
			if ProjectExists(projectNameEntry.Text) != true {
				log.Print("Creating New Project: ", projectNameEntry.Text)
				CWP = projectNameEntry.Text
				xml := PomXML{GetAuthor(), CWP, projectGroupEntry.Text, projectArtifactEntry.Text, projectDescriptionEntry.Text}
				createNewProject(projectNameEntry.Text, xml)
				UnhideButtons()
				
				HideModal()
			} else {
				dialog.ShowError(errors.New("Project Exists"), GetWindow())
			}
		}
	}
	newProjectForm.OnCancel = func() {
		HideModal()
	}
	return newProjectForm
}