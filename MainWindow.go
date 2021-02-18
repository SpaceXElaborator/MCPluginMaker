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
	secondModal *widget.PopUp
	a fyne.App = app.New()
	w fyne.Window = a.NewWindow("MCPluginMaker | " + GetAuthor())
	
	// -------------- Buttons --------------
	addCmdButt = widget.NewButton("Add Command", func() {
		commandForm := createCommandForm()
		modal = widget.NewModalPopUp(commandForm, w.Canvas())
		modal.Resize(fyne.NewSize(512, 0))
		modal.Show()
	})
	
	addItemButt = widget.NewButton("Add Item", func() {
		itemForm := customItemForm()
		modal = widget.NewModalPopUp(itemForm, w.Canvas())
		modal.Resize(fyne.NewSize(512, 0))
		modal.Show()
	})
	
	addBuildButt = widget.NewButton("Build Project", func() {
		proj := GetProject(CWP)
		build(proj)
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

// Create the top bar for creating projects, deleting projects, and the settings
func createToolbar() *widget.Toolbar {
	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.FolderNewIcon(), func() {
			newProjectForm := createNewProjectForm()
			modal = widget.NewModalPopUp(widget.NewCard("New Project", "", newProjectForm), w.Canvas())
			modal.Resize(fyne.NewSize(512, 0))
			modal.Show()
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
	HideButtons()
	
	/*mainCenter := container.NewVBox(
			addCmdButt,
			addItemButt,
			addBuildButt,
	)*/
	
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
		UnhideButtons()
		SetNewContent()
	}
	
	w.Resize(fyne.NewSize(1024, 768))
	w.SetContent(c)
	w.ShowAndRun()
}

func SetNewContent() {

	var apps fyne.CanvasObject

	if len(GetProject(CWP).CmdRows) >= 1 {
		apps = container.NewAppTabs(
			container.NewTabItem("Commands", 
				container.NewBorder(
					createCmdToolbar(),
					nil,
					nil,
					nil,
					container.NewMax(CreateCommandBlocks()),
				),
			),
		)
	} else {
		apps = container.NewAppTabs(
			container.NewTabItem("Commands", container.NewMax(widget.NewLabel("Create A Command To View This"))),
		)
	}
	
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

func CreateCommandBlocks() *widget.Table {
	proj := GetProject(CWP)
	cmdRows := proj.CmdRows
	cmdLength := len(cmdRows[len(cmdRows) - 1].Cmds)
	dmyCmd := Command{"", "", ""}
	if cmdLength == 1 {
		cmdRows[len(cmdRows) - 1].Cmds = append(cmdRows[len(cmdRows) - 1].Cmds, dmyCmd, dmyCmd)
	}
	if cmdLength == 2 {
		cmdRows[len(cmdRows) - 1].Cmds = append(cmdRows[len(cmdRows) - 1].Cmds, dmyCmd)
	}
	table := widget.NewTable(
		func() (int, int) {
			return len(cmdRows), 3
		},
		func() fyne.CanvasObject {
			name := widget.NewLabelWithStyle("Name Text", fyne.TextAlignCenter, fyne.TextStyle{})
			sep := widget.NewSeparator()
			flavor := widget.NewLabel("Flavor Text")
			
			con := container.NewVBox(
				name,
				sep,
				flavor,
			)
			return con
		},
		func(tci widget.TableCellID, f fyne.CanvasObject) {
			f.(*fyne.Container).Objects[0].(*widget.Label).SetText(cmdRows[tci.Row].Cmds[tci.Col].Name)
			if cmdRows[tci.Row].Cmds[tci.Col].SlashCommand == "" {
				f.(*fyne.Container).Objects[2].(*widget.Label).SetText("")
			} else {
				f.(*fyne.Container).Objects[2].(*widget.Label).SetText("/" + cmdRows[tci.Row].Cmds[tci.Col].SlashCommand)
			}
			
		})
	
	table.SetColumnWidth(0, 270)
	table.SetColumnWidth(1, 270)
	table.SetColumnWidth(2, 270)
	
	return table
}

// A function to quickly disable all buttons
func HideButtons() {
	addCmdButt.Disable()
	addItemButt.Disable()
	addBuildButt.Disable()
}

// A function to quickly enable all buttons
func UnhideButtons() {
	addCmdButt.Enable()
	addItemButt.Enable()
	addBuildButt.Enable()
}

// Get the current modal being displayed and hide it if it isn't
func HideModal() {
	modal.Hide()
}

func HideModal2() {
	secondModal.Hide()
}

func GetApp() fyne.App {
	return a
}

func GetWindow() fyne.Window {
	return w
}