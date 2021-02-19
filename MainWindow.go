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
	// Create the global variables needed
	modal *widget.PopUp
	a fyne.App = app.NewWithID("MCPluginMaker")
	w fyne.Window = a.NewWindow("MCPluginMaker | " + GetAuthor())
	
	// -------------- Buttons --------------
	addItemButt = widget.NewButton("Add Item", func() {
		itemForm := customItemForm()
		modal = widget.NewModalPopUp(itemForm, w.Canvas())
		modal.Resize(fyne.NewSize(512, 0))
		modal.Show()
	})
	// -------------------------------------
	
	// List of all Projects that will show on the left side of the screen
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

func createItemToolbar() *widget.Toolbar {
	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {
			itemForm := customItemForm()
			modal = widget.NewModalPopUp(widget.NewCard("Add Command", "", itemForm), w.Canvas())
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

// Create the top bar for creating projects, deleting projects, and the settings
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
			proj := GetProject(CWP)
			build(proj)
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

// Create the offical GUI using Fyne
func ShowMainMenu() {
	toolbar := createToolbar()
	split := container.NewHSplit(list, container.NewVBox(
			widget.NewLabel("Please Select A Project Or Create A New One"),
		))
	split.Offset = 0.1
	c := container.NewBorder(
		toolbar,
		nil,
		nil,
		nil,
		split,
	)
	list.OnSelected = func(id widget.ListItemID) {
		CWP = PluginProjects[id].Name
		w.SetTitle("MCPluginMaker | " + GetAuthor())
		SetNewContent()
	}
	w.Resize(fyne.NewSize(1024, 768))
	w.SetContent(c)
	w.ShowAndRun()
}

func SetNewContent() {
	apps := container.NewAppTabs(
		container.NewTabItem("Commands", 
			container.NewBorder(
				createCmdToolbar(),
				nil,
				nil,
				nil,
				CreateCommandBlocks(),
			),
		),
		container.NewTabItem("Listeners", 
			container.NewBorder(
				createCmdToolbar(),
				nil,
				nil,
				nil,
				CreateCommandBlocks(),
			),
		),
		container.NewTabItem("Items", 
			container.NewBorder(
				createItemToolbar(),
				nil,
				nil,
				nil,
				CreateCommandBlocks(),
			),
		),
	)
	
	card := widget.NewCard("Project: " + CWP, "", apps)
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

func CreateCommandBlocks() fyne.CanvasObject {
	var test []fyne.CanvasObject
	
	for _, f := range GetProject(CWP).Cmds {
	
		cont := widget.NewCard("", "Functions", container.NewVScroll(widget.NewAccordion(
			widget.NewAccordionItem("Test", widget.NewLabel("Test")),
		)))
		
		card := widget.NewCard(
			f.Name,
			"/" + f.SlashCommand,
			cont,
		)
		test = append(test, card)
	}
	
	content := container.NewVScroll(container.NewGridWrap(fyne.NewSize(270, 290), test...))
	return content
}

// Get the current modal being displayed and hide it if it isn't
func HideModal() {
	modal.Hide()
}

func GetApp() fyne.App {
	return a
}

func GetWindow() fyne.Window {
	return w
}