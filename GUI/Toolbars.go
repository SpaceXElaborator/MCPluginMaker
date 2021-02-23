package plugingui

import (
	PluginSettings "SpaceXElaborator/PluginMaker/Settings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// createToolbar Create the top bar for creating projects, deleting projects, and the settings
func createToolbar() *widget.Toolbar {
	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.FolderNewIcon(), func() {
			newProjectForm := createNewProjectForm()
			modal = widget.NewModalPopUp(widget.NewCard("New Project", "", newProjectForm), w.Canvas())
			modal.Resize(fyne.NewSize(512, 0))
			modal.Show()
		}),
		widget.NewToolbarAction(theme.DeleteIcon(), func() {
			delProjectForm := delProjectForm()
			modal = widget.NewModalPopUp(widget.NewCard("Remove Project", "", delProjectForm), w.Canvas())
			modal.Resize(fyne.NewSize(512, 0))
			modal.Show()
		}),
		widget.NewToolbarAction(theme.DocumentSaveIcon(), func() {
			proj := Projects.GetProject(PluginSettings.GetCWP())
			if proj == nil {
				return
			}
			proj.Build()
		}),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.SettingsIcon(), func() {
			settingsForm := createSettingsForm()
			modal = widget.NewModalPopUp(widget.NewCard("Settings", "", settingsForm), w.Canvas())
			modal.Resize(fyne.NewSize(512, 0))
			modal.Show()
		}),
	)
	return toolbar
}

// createCmdToolbar Creates the +/- button for adding or removing commands in the Command Tab
func createCmdToolbar() *widget.Toolbar {
	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {
			commandForm := createCommandForm()
			modal = widget.NewModalPopUp(widget.NewCard("Add Command", "", commandForm), w.Canvas())
			modal.Resize(fyne.NewSize(512, 0))
			modal.Show()
		}),
		widget.NewToolbarAction(theme.ContentRemoveIcon(), func() {
			commandForm := removeCommand()
			modal = widget.NewModalPopUp(widget.NewCard("Remove Command", "", commandForm), w.Canvas())
			modal.Resize(fyne.NewSize(512, 0))
			modal.Show()
		}),
	)
	return toolbar
}

// createItemToolbar Creates the +/- button for adding or removing items in the Item Tab
func createItemToolbar() *widget.Toolbar {
	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {
			itemForm := customItemForm()
			modal = widget.NewModalPopUp(itemForm, w.Canvas())
			modal.Resize(fyne.NewSize(512, 0))
			modal.Show()
		}),
		widget.NewToolbarAction(theme.ContentRemoveIcon(), func() {
			return
		}),
	)
	return toolbar
}
