package plugingui

import (
	PluginItems "SpaceXElaborator/PluginMaker/Item"
	PluginSettings "SpaceXElaborator/PluginMaker/Settings"
	"errors"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

// customItemForm Generates a form and create a pluginitems.CustomItem to add to the project
func customItemForm() *widget.Form {
	itemNameEntry := widget.NewEntry()
	itemNameEntry.Resize(fyne.NewSize(300, 300))
	itemNameEntry.SetText("")
	itemNameFormItem := &widget.FormItem{
		Text:   "Item Name",
		Widget: itemNameEntry,
	}

	itemMaterialEntry := widget.NewEntry()
	itemMaterialEntry.Resize(fyne.NewSize(300, 300))
	itemMaterialEntry.SetText("")
	itemMaterialFormItem := &widget.FormItem{
		Text:   "Item Material",
		Widget: itemMaterialEntry,
	}

	itemDescEntry := widget.NewMultiLineEntry()
	itemDescEntry.Resize(fyne.NewSize(300, 300))
	itemDescEntry.SetText("")
	itemDescFormItem := &widget.FormItem{
		Text:   "Item Description",
		Widget: itemDescEntry,
	}

	newCustomItemForm := widget.NewForm(itemNameFormItem, itemMaterialFormItem, itemDescFormItem)
	newCustomItemForm.OnSubmit = func() {
		if itemNameEntry.Text != "" && itemMaterialEntry.Text != "" {
			proj := Projects.GetProject(PluginSettings.GetCWP())
			if proj.CheckItem(itemNameEntry.Text) != true {
				// Make sure that the material that was typed in the value is a valid Spigot Material type.
				// Couldn't make this be a list as it was largly lagging, will have to look into this later
				stringToCheck := strings.ReplaceAll(strings.ToUpper(itemMaterialEntry.Text), " ", "_")
				if PluginItems.CheckMaterial(itemNameEntry.Text) {
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
