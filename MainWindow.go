package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/layout"
)

var (
	modal *widget.PopUp
	a fyne.App = app.New()
	w fyne.Window = a.NewWindow("MCPluginMaker | " + GetAuthor() + " | Project: " + CWP)
	
	addCmdButt = widget.NewButton("Add Command", func() {
		commandForm := createCommandForm()
		modal = widget.NewModalPopUp(commandForm, w.Canvas())
		modal.Resize(fyne.NewSize(512, 240))
		modal.Show()
	})
	
	addBuildButt = widget.NewButton("Build Project", func() {
		proj := GetProject(CWP)
		build(proj)
	})
	
	list = widget.NewList(
		func() int {
			return len(PluginProjects)
		},
		func() fyne.CanvasObject {
			return fyne.NewContainerWithLayout(layout.NewHBoxLayout(), widget.NewIcon(theme.DocumentIcon()), widget.NewLabel("Template Object"))
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {
			item.(*fyne.Container).Objects[1].(*widget.Label).SetText(PluginProjects[id].Name)
		},
	)
)

func createToolbar() *widget.Toolbar {
	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.FolderNewIcon(), func() {
			newProjectForm := createNewProjectForm()
			modal = widget.NewModalPopUp(newProjectForm, w.Canvas())
			modal.Resize(fyne.NewSize(512, 230))
			modal.Show()
		}),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.SettingsIcon(), func() {
			settingsForm := createSettingsForm()
			modal = widget.NewModalPopUp(settingsForm, w.Canvas())
			modal.Resize(fyne.NewSize(512, 384))
			modal.Show()
		}),
	)
	return toolbar
}

func ShowMainMenu() {
	HideButtons()
	
	list.OnSelected = func(id widget.ListItemID) {
		CWP = PluginProjects[id].Name
		w.SetTitle("MCPluginMaker | " + GetAuthor() + " | Project: " + CWP)
		UnhideButtons()
	}
	
	mainCenter := container.NewVBox(
			addCmdButt,
			addBuildButt,
	)
	toolbar := createToolbar()
	split := container.NewHSplit(list, mainCenter)
	split.Offset = 0.1
	c := container.NewBorder(
		toolbar,
		nil,
		nil,
		nil,
		split,
	)
	
	w.Resize(fyne.NewSize(1024, 768))
	w.SetContent(c)
	w.ShowAndRun()
}

func HideButtons() {
	addCmdButt.Disable()
	addBuildButt.Disable()
}

func UnhideButtons() {
	addCmdButt.Enable()
	addBuildButt.Enable()
}

func HideModal() {
	modal.Hide()
}

func GetApp() fyne.App {
	return a
}

func GetWindow() fyne.Window {
	return w
}