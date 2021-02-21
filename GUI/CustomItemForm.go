package PluginGUI

import (
	"strings"
	"errors"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/dialog"
	
	"SpaceXElaborator/PluginMaker/Settings"
)

func customItemForm() *widget.Form {
	itemNameEntry := widget.NewEntry()
	itemNameEntry.Resize(fyne.NewSize(300, 300))
	itemNameEntry.SetText("")
	itemNameFormItem := &widget.FormItem {
		Text: "Item Name",
		Widget: itemNameEntry,
	}
	
	itemMaterialEntry := widget.NewEntry()
	itemMaterialEntry.Resize(fyne.NewSize(300, 300))
	itemMaterialEntry.SetText("")
	itemMaterialFormItem := &widget.FormItem {
		Text: "Item Material",
		Widget: itemMaterialEntry,
	}
	
	itemDescEntry := widget.NewMultiLineEntry()
	itemDescEntry.Resize(fyne.NewSize(300, 300))
	itemDescEntry.SetText("")
	itemDescFormItem := &widget.FormItem {
		Text: "Item Description",
		Widget: itemDescEntry,
	}

	newCustomItemForm := widget.NewForm(itemNameFormItem, itemMaterialFormItem, itemDescFormItem)
	newCustomItemForm.OnSubmit = func() {
		if itemNameEntry.Text != "" && itemMaterialEntry.Text != "" {
			proj := Projects.GetProject(PluginSettings.GetCWP())
			if proj.CheckItem(itemNameEntry.Text) != true {
				stringToCheck := strings.ReplaceAll(strings.ToUpper(itemMaterialEntry.Text), " ", "_")
				if(proj.CheckMaterial(stringToCheck)) {
					proj.AddItem(strings.ToUpper(stringToCheck), itemNameEntry.Text, strings.Split(itemDescEntry.Text, "\n"))
					// HideModal() in MainWindow.go
					HideModal()
				} else {
					dialog.ShowError(errors.New("Unrecognized Material! Check https://hub.spigotmc.org/javadocs/bukkit/org/bukkit/Material.html"), GetWindow())
				}
			} else {
				dialog.ShowError(errors.New("An item with that name exists"), GetWindow())
			}
		} else {
			dialog.ShowError(errors.New("Item Name and Item Material can't be empty"), GetWindow())
		}
	}
	newCustomItemForm.OnCancel = func() {
		// HideModal() in MainWindow.go
		HideModal()
	}
	return newCustomItemForm
}