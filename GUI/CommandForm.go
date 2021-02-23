package plugingui

import (
	PluginSettings "SpaceXElaborator/PluginMaker/Settings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

// createCommandForm Generates a form with the given name and slash string. The List determines who can type the command
func createCommandForm() *widget.Form {
	cmdType := []string{"Player", "Block", "Console"}
	CommandType := ""
	commandTypeEntry := widget.NewSelect(cmdType, func(s string) {
		CommandType = s
	})
	commandTypeFormItem := &widget.FormItem{
		Text:   "Command Executor",
		Widget: commandTypeEntry,
	}

	commandNameEntry := widget.NewEntry()
	commandNameEntry.Resize(fyne.NewSize(300, 300))
	commandNameEntry.SetText("")
	commandNameFormItem := &widget.FormItem{
		Text:   "Command Name",
		Widget: commandNameEntry,
	}

	slashStringEntry := widget.NewEntry()
	slashStringEntry.Resize(fyne.NewSize(300, 300))
	slashStringEntry.SetText("")
	slashStringFormItem := &widget.FormItem{
		Text:   "Slash String",
		Widget: slashStringEntry,
	}

	newCommandForm := widget.NewForm(commandTypeFormItem, commandNameFormItem, slashStringFormItem)
	newCommandForm.OnSubmit = func() {
		if CommandType != "" && commandNameEntry.Text != "" && slashStringEntry.Text != "" {
			proj := Projects.GetProject(PluginSettings.GetCWP())

			err := proj.AddCommand(commandNameEntry.Text, slashStringEntry.Text, CommandType)
			if err != nil {
				dialog.ShowError(err, GetWindow())
			} else {
				SetNewContent()
				// HideModal() in MainWindow.go
				HideModal()
			}
		}
	}

	newCommandForm.OnCancel = func() {
		// HideModal() in MainWindow.go
		HideModal()
	}
	return newCommandForm
}

// removeCommand Removes the given command from a list of all available commands
func removeCommand() *widget.Form {
	CmdNames := []string{}
	for _, f := range Projects.GetProject(PluginSettings.GetCWP()).Cmds {
		CmdNames = append(CmdNames, f.Name)
	}

	cmdToRem := ""

	commandNameEntry := widget.NewSelect(CmdNames, func(s string) {
		cmdToRem = s
	})
	commandNameFormItem := &widget.FormItem{
		Text:   "Command Name",
		Widget: commandNameEntry,
	}

	remCommandForm := widget.NewForm(commandNameFormItem)
	remCommandForm.OnSubmit = func() {
		if cmdToRem != "" {
			Projects.GetProject(PluginSettings.GetCWP()).RemoveCommand(cmdToRem)
			SetNewContent()
			HideModal()
		}
	}

	remCommandForm.OnCancel = func() {
		HideModal()
	}

	return remCommandForm
}
