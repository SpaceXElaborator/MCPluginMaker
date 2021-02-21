package PluginGUI

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/dialog"
	"strconv"
	"errors"
	
	"SpaceXElaborator/PluginMaker/Command"
	"SpaceXElaborator/PluginMaker/Settings"
)

func playerCommandFuncAddForm(cmd *PluginCommands.Command) *widget.Form {
	funcForm := widget.NewForm()
	cmdTypes := []string{"Add Item", "Add Custom Item"}
	
	cmdFuncType := widget.NewSelect(cmdTypes, func(s string) {
		modal.Hide()
		if s == "Add Custom Item" {
			modal = widget.NewModalPopUp(widget.NewCard("Spawn Custom Item", "", spawnItemForm(cmd, true)), w.Canvas())
		} else if s == "Add Item" {
			modal = widget.NewModalPopUp(widget.NewCard("Spawn Item", "", spawnItemForm(cmd, false)), w.Canvas())
		}
		modal.Resize(fyne.NewSize(512, 0))
		modal.Show()
	})
	
	cmdFuncTypeFormItem := &widget.FormItem {
		Text: "Command Function",
		Widget: cmdFuncType,
	}
	
	funcForm.AppendItem(cmdFuncTypeFormItem)
	
	return funcForm
}

func spawnItemForm(cmd *PluginCommands.Command, custom bool) *widget.Form {
	itemForm := widget.NewForm()
	
	itemName := ""
	
	proj := Projects.GetProject(PluginSettings.GetCWP())
	
	if custom {
		var items []string
		for _, item := range proj.Items {
			items = append(items, item.ItemName)
		}
		cmdFuncType := widget.NewSelect(items, func(s string) {
			itemName = s
		})
		itemForm.Append("Custom Item", cmdFuncType)
		
		itemAmount := widget.NewEntry()
		itemForm.Append("Amount", itemAmount)
		
		itemForm.OnSubmit = func() {
			if itemAmount.Text != "" && itemName != "" {
				if _, err := strconv.Atoi(itemAmount.Text); err == nil {
					cmd.AddFunc("p.getInventory().addItem(" + PluginSettings.GetCWP() + "CustomItems.build(\"" + itemName + "\", " + itemAmount.Text + "));")
					HideModal()
				} else {
					dialog.ShowError(errors.New("Amount must be a number"), GetWindow())
				}
			} else {
				dialog.ShowError(errors.New("Fill in all values"), GetWindow())
			}
		}
		itemForm.OnCancel = func() {
			HideModal()
		}
		itemForm.Refresh()
	} else {
		
	}
	return itemForm
}