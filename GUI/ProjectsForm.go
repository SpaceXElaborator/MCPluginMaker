package PluginGUI

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/dialog"
	
	"SpaceXElaborator/PluginMaker/Settings"
)

func delProjectForm() *widget.Form {
	ProjNames := []string{}
	for _, f := range Projects.Projects {
		ProjNames = append(ProjNames, f.Name)
	}
	
	ProjName := ""
	
	projNameEntry := widget.NewSelect(ProjNames, func(s string) {
		ProjName = s
	})
	projNameFormItem := &widget.FormItem {
		Text: "Project Name",
		Widget: projNameEntry,
	}
	
	remCommandForm := widget.NewForm(projNameFormItem)
	remCommandForm.OnSubmit = func() {
		if ProjName != "" {
			Projects.RemoveProject(ProjName)
			list.Refresh()
			HideModal()
		}
	}
	
	remCommandForm.OnCancel = func() {
		HideModal()
	}
	
	return remCommandForm
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
			err := Projects.CreateNewProject(projectNameEntry.Text, PluginSettings.GetAuthor(), projectGroupEntry.Text, projectArtifactEntry.Text, projectDescriptionEntry.Text)
			PluginSettings.SetCWP(projectNameEntry.Text)
			Projects.GetProject(PluginSettings.GetCWP()).CreatePom()
			if err != nil {
				dialog.ShowError(err, GetWindow())
			} else {
				GetWindow().SetTitle("MCPluginMaker | " + PluginSettings.GetAuthor())
				list.Refresh()
				HideModal()
			}
		}
	}
	newProjectForm.OnCancel = func() {
		// HideModal() in MainWindow.go
		HideModal()
	}
	return newProjectForm
}