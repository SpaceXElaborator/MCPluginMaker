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

// Check if name is the same as another command
func CommandExist(cmd Command) bool {
	proj := GetProject(CWP)
	cmdName := strings.ToLower(cmd.Name)
	for _, f := range proj.Cmds {
		if strings.ToLower(f.Name) == cmdName {
			return true
		}
	}
	return false
}

// Check if the /{cmd} is the same as another command
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

// Create the form the user will input information into
func createCommandForm() *widget.Form {
	cmdType := []string {"Player", "Block", "Console"}
	CommandType := ""
	commandTypeEntry := widget.NewSelect(cmdType, func(s string) {
		CommandType = s
	})
	commandTypeFormItem := &widget.FormItem {
		Text: "Command Executor",
		Widget: commandTypeEntry,
	}

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
	
	newCommandForm := widget.NewForm(commandTypeFormItem, commandNameFormItem, slashStringFormItem)
	newCommandForm.OnSubmit = func() {
		if CommandType != "" && commandNameEntry.Text != "" && slashStringEntry.Text != "" {
			cmd := Command{GetAuthor(), CommandType, commandNameEntry.Text, slashStringEntry.Text}
			if CommandExist(cmd) != true {
				if(SlashExists(cmd) != true) {
					createCommand(cmd)
					// HideModal() in MainWindow.go
					HideModal()
				} else {
					// GetWindow() in MainWindow.go
					dialog.ShowError(errors.New("SlashCommand Exists"), GetWindow())
				}
			} else {
				// GetWindow() in MainWindow.go
				dialog.ShowError(errors.New("Command Exists"), GetWindow())
			}
		}
	}
	newCommandForm.OnCancel = func() {
		// HideModal() in MainWindow.go
		HideModal()
	}
	return newCommandForm
}

func removeCommand() *widget.Form {
	CmdNames := []string{}
	for _, f := range GetProject(CWP).Cmds {
		CmdNames = append(CmdNames, f.Name)
	}
	
	cmdToRem := ""
	
	commandNameEntry := widget.NewSelect(CmdNames, func(s string) {
		cmdToRem = s
	})
	commandNameFormItem := &widget.FormItem {
		Text: "Command Name",
		Widget: commandNameEntry,
	}
	
	remCommandForm := widget.NewForm(commandNameFormItem)
	remCommandForm.OnSubmit = func() {
		if cmdToRem != "" {
			var index int
			for i, f := range GetProject(CWP).Cmds {
				if f.Name == cmdToRem {
					index = i
					break
				}
			}
			GetProject(CWP).Cmds = append(GetProject(CWP).Cmds[:index], GetProject(CWP).Cmds[index+1:]...)
			os.Remove("projects/" + CWP + "/src/main/java/com/terturl/net/cmds/" + cmdToRem + ".java")
			SetNewContent()
			HideModal()
		}
	}
	
	remCommandForm.OnCancel = func() {
		HideModal()
	}
	
	return remCommandForm
}

func createCommand(cmd Command) {
	// Create parent directors if they aren't present using the current mode permission of the user
	os.MkdirAll("projects/" + CWP + "/src/main/java/com/terturl/net/cmds", os.ModePerm)
	
	// GetProject() in Main.go
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
	SetNewContent()
}

func buildCommands(proj *Project) {
	for _, cmd := range proj.Cmds {
		// Create the Java file for editing using the template below
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
}