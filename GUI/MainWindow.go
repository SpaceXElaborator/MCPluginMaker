package PluginGUI

import (
	"strings"

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

func ResetSettings() {
	w.SetTitle("MCPluginMaker | " + PluginSettings.GetAuthor())
	if PluginSettings.GetDark() == true {
		// GetApp() in MainWindow.go
		a.Settings().SetTheme(theme.DarkTheme())
	} else {
		// GetApp() in MainWindow.go
		a.Settings().SetTheme(theme.LightTheme())
	}
}

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

func CreateCommandBlocks() fyne.CanvasObject {
	var test []fyne.CanvasObject

	for _, f := range Projects.GetProject(PluginSettings.GetCWP()).Cmds {
		toolbar := widget.NewToolbar(
			widget.NewToolbarAction(theme.ContentAddIcon(), func() {
				if strings.EqualFold(f.CommandType, "Player") {
					canvas := w.Canvas()
					funcForm := PluginFunction.PlayerCommandFuncAddForm(f, &canvas, &w, HideModal, SetNewContent, Projects.GetProject(PluginSettings.CWP).Items)
					modal = widget.NewModalPopUp(widget.NewCard("Add Command Function", "", funcForm), w.Canvas())
					modal.Resize(fyne.NewSize(512, 0))
					modal.Show()
				}
			}),
			widget.NewToolbarAction(theme.ContentRemoveIcon(), func() {

			}),
		)

		var accItems []*widget.AccordionItem

		if len(f.PlayerFuncs) >= 1 {
			playerFuncCont := container.NewVBox()
			for _, playerFuncs := range f.PlayerFuncs {
				playerFuncCont.Add(widget.NewLabel(playerFuncs.Name))
			}

			accItems = append(accItems, widget.NewAccordionItem("Player Functions", playerFuncCont))
		}

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

	content := container.NewVScroll(container.NewGridWrap(fyne.NewSize(300, 350), test...))
	return content
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
