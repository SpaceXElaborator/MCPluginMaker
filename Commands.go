package main

import (
	"os"
	"errors"
	"strings"
	"log"
	"text/template"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/dialog"
)

func CommandExist(cmd Command) bool {
	proj := GetProject(CWP)
	cmdName := strings.ToLower(cmd.Name)
	for _, f := range proj.Cmds {
		log.Print("Command Found in " + proj.Name + ": ", f.Name)
		if strings.ToLower(f.Name) == cmdName {
			return true
		}
	}
	return false
}

func SlashExists(cmd Command) bool { 
	proj := GetProject(CWP)
	slashString := strings.ToLower(cmd.SlashCommand)
	for _, f := range proj.Cmds {
		if strings.ToLower(f.SlashCommand) == slashString {
			return true
		}
	}
	return false
}

func createCommandForm() *widget.Form {
	commandNameEntry := widget.NewEntry()
	commandNameEntry.Resize(fyne.NewSize(300, 300))
	commandNameEntry.SetText("")
	commandNameFormItem := &widget.FormItem {
		Text: "Command Name",
		Widget: commandNameEntry,
	}
	
	slashStringEntry := widget.NewEntry()
	slashStringEntry.Resize(fyne.NewSize(300, 300))
	slashStringEntry.SetText("")
	slashStringFormItem := &widget.FormItem {
		Text: "Slash String",
		Widget: slashStringEntry,
	}
	
	newCommandForm := widget.NewForm(commandNameFormItem, slashStringFormItem)
	newCommandForm.OnSubmit = func() {
		if commandNameEntry.Text != "" && slashStringEntry.Text != "" {
			cmd := Command{GetAuthor(), commandNameEntry.Text, slashStringEntry.Text}
			if(CommandExist(cmd) != true) {
				if(SlashExists(cmd) != true) {
					createCommand(cmd)
					HideModal()
				} else {
					dialog.ShowError(errors.New("SlashCommand Exists"), GetWindow())
				}
			} else {
				dialog.ShowError(errors.New("Command Exists"), GetWindow())
			}
		}
	}
	newCommandForm.OnCancel = func() {
		HideModal()
	}
	return newCommandForm
}

func createCommand(cmd Command) {
	os.MkdirAll("projects/" + CWP + "/src/main/java/com/terturl/net/cmds", os.ModePerm)
	proj := GetProject(CWP)
	proj.Cmds = append(proj.Cmds, cmd)
	
	// Reset the project to be the new proj pointer
	var index int
	for i, cmd := range PluginProjects {
		if proj.Name == cmd.Name {
			index = i
		}
	}
	PluginProjects[index] = *proj
	
	for _, cmd := range proj.Cmds {
		log.Print("Command: ", cmd.Name)
	}
	
	f, err := os.Create("projects/" + CWP + "/src/main/java/com/terturl/net/cmds/" + cmd.Name + ".java")
	if err != nil {
		log.Print("Error: ", err)
	}
	t := template.Must(template.New("CreateCommand").Parse(cmdJavaTmpl))
	err = t.Execute(f, &cmd)
	if err != nil {
		log.Print("Error: ", err)
	}
	f.Close()
}