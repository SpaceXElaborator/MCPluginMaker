package pluginfunctions

import (
	PluginCommands "SpaceXElaborator/PluginMaker/Command"
	PluginWidgets "SpaceXElaborator/PluginMaker/GUI/Widgets"
	PluginItems "SpaceXElaborator/PluginMaker/Item"
	"log"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func CreateToolbarForCommand(cmd *PluginCommands.Command, SetNewContent func(), items []*PluginItems.CustomItem) *widget.Toolbar {
	test := cmd
	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {
			if strings.EqualFold(test.CommandType, "Player") {
				funcForm := playerCommandFuncAddForm(test, SetNewContent, items)
				modal = widget.NewModalPopUp(widget.NewCard("Add Command Function", "", funcForm), *canvas)
				modal.Resize(fyne.NewSize(512, 0))
				modal.Show()
			}
		}),
		widget.NewToolbarAction(theme.ContentRemoveIcon(), func() {
			log.Print(cmd.Name)
		}),
	)
	return toolbar
}

func BuildCmdCard(cmd *PluginCommands.Command) []*widget.AccordionItem {
	var accItems []*widget.AccordionItem
	if len(cmd.SubCommands) >= 1 {
		subCommandCont := container.NewVBox()
		for _, subCommands := range cmd.SubCommands {
			subCommandCont.Add(PluginWidgets.NewClickableLabel(
				subCommands.SlashCommand,
				func() {
					log.Print(subCommands.Name)
				},
				func() {
					log.Print("Right Clicked")
				},
			))
		}

		accItems = append(accItems, widget.NewAccordionItem("Sub Commands", subCommandCont))
	}

	if len(cmd.PlayerFuncs) >= 1 {
		playerFuncCont := container.NewVBox()
		for _, playerFuncs := range cmd.PlayerFuncs {
			//Adds the custom widget made and will eventually allow you to edit the command by clicking on the label
			playerFuncCont.Add(PluginWidgets.NewClickableLabel(
				playerFuncs.Name,
				// Debugging for the time being
				func() {
					log.Print(playerFuncs.Name)
				},
				func() {
					log.Print("Right Clicked")
				},
			))
		}

		accItems = append(accItems, widget.NewAccordionItem("Player Functions", playerFuncCont))
	}
	return accItems
}
