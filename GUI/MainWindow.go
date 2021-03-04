package plugingui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	PluginFunction "SpaceXElaborator/PluginMaker/GUI/Functions"
	PluginProject "SpaceXElaborator/PluginMaker/Project"
	PluginSettings "SpaceXElaborator/PluginMaker/Settings"
)

var (
	modal *widget.PopUp
	a     fyne.App    = app.NewWithID("MCPluginMaker")
	w     fyne.Window = a.NewWindow("MCPluginMaker | " + PluginSettings.GetAuthor())

	// Projects will hold the pointer value of all the projects created from Main.go
	Projects *PluginProject.Projects

	list = widget.NewList(
		func() int {
			return len(Projects.Projects)
		},
		func() fyne.CanvasObject {
			return fyne.NewContainerWithLayout(layout.NewHBoxLayout(), widget.NewIcon(theme.DocumentIcon()), widget.NewLabel("Template Object"))
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {
			item.(*fyne.Container).Objects[1].(*widget.Label).SetText(Projects.Projects[id].Name)
		},
	)
)

// ResetSettings will just retitle the Window and change the theme to dark/light
func ResetSettings() {
	w.SetTitle("MCPluginMaker | " + PluginSettings.GetAuthor())
	if PluginSettings.GetDark() == true {
		a.Settings().SetTheme(theme.DarkTheme())
	} else {
		a.Settings().SetTheme(theme.LightTheme())
	}
}

// ShowMainMenu Creates the Main Window and will store the ability to view, build, delete, and add projects
func ShowMainMenu(projs *PluginProject.Projects) {
	Projects = projs

	split := container.NewHSplit(list, container.NewVBox(
		widget.NewLabel("Please Select A Project Or Create A New One"),
	))
	split.Offset = 0.1
	c := container.NewBorder(
		createToolbar(),
		nil,
		nil,
		nil,
		split,
	)

	list.OnSelected = func(id widget.ListItemID) {
		PluginSettings.SetCWP(projs.Projects[id].Name)
		w.SetTitle("MCPluginMaker | " + PluginSettings.GetAuthor())
		SetNewContent()
	}

	w.Resize(fyne.NewSize(1024, 768))
	w.SetContent(c)

	canvas := w.Canvas()
	PluginFunction.InitCommands(&canvas, &w)

	w.ShowAndRun()
}

// SetNewContent Is a refresh function to update the Main Window Project menu when you add/remove anything from it
func SetNewContent() {
	apps := container.NewAppTabs(
		container.NewTabItem("Commands",
			container.NewBorder(
				createCmdToolbar(),
				nil,
				nil,
				nil,
				createCommandBlocks(),
			),
		),
		container.NewTabItem("Listeners",
			container.NewBorder(
				nil, // TODO: Add Listener Toolbar
				nil,
				nil,
				nil,
				widget.NewLabel("Not Impleted Yet... Come Back Later"),
			),
		),
		container.NewTabItem("Items",
			container.NewBorder(
				createItemToolbar(),
				nil,
				nil,
				nil,
				widget.NewLabel("Not Impleted Yet... Come Back Later"),
			),
		),
	)

	card := widget.NewCard("Project: "+PluginSettings.GetCWP(), "", apps)
	toolbar := createToolbar()
	split := container.NewHSplit(list, card)
	split.Offset = 0.1

	c := container.NewBorder(
		toolbar,
		nil,
		nil,
		nil,
		split,
	)

	w.SetContent(c)
}

// CreateCommandBlocks is a builder function to rebuild the Command card form a Projects Menu
func createCommandBlocks() fyne.CanvasObject {
	var test []fyne.CanvasObject

	pluginCmds := Projects.GetProject(PluginSettings.GetCWP()).Cmds

	// Get all the project's commands and create a Toolbar that will ONLY affect that command
	for _, f := range pluginCmds {
		toolbar := PluginFunction.CreateToolbarForCommand(f, HideModal, SetNewContent, Projects.GetProject(PluginSettings.GetCWP()).Items)

		accItems := PluginFunction.BuildCmdCard(f)

		// This will display all of the PlayerFuncs in a list for the user to be able to view
		max := container.NewBorder(
			toolbar,
			nil,
			nil,
			nil,
			container.NewVScroll(
				widget.NewAccordion(
					accItems...,
				),
			),
		)

		cont := widget.NewCard("", "Functions", max)

		card := widget.NewCard(
			f.CommandType+" Command",
			"/"+f.SlashCommand,
			cont,
		)
		test = append(test, card)
	}

	content := container.NewVScroll(container.NewGridWrap(fyne.NewSize(275, 350), test...))
	return content
}

// HideModal will hide the current modal displayed on the screen
func HideModal() {
	modal.Hide()
}

// GetApp returns the App for the ability to modify it
func GetApp() fyne.App {
	return a
}

// GetWindow returns the App's Window for the ability to modify it
func GetWindow() fyne.Window {
	return w
}
