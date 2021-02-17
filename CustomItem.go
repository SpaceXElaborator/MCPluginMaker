package main

import (
	"strings"
	"log"
	"errors"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/dialog"
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

	//itemTestFormItem := &widget.FormItem {
	//	Text: "Crafting Recipe",
	//	Widget: MakeTable(),
	//}

	newCustomItemForm := widget.NewForm(itemNameFormItem, itemMaterialFormItem, itemDescFormItem)
	newCustomItemForm.OnSubmit = func() {
		if itemNameEntry.Text != "" && itemMaterialEntry.Text != "" {
			stringToCheck := strings.ReplaceAll(strings.ToUpper(itemMaterialEntry.Text), " ", "_")
			if(CheckMaterial(stringToCheck)) {
				newItem := CustomItem{GetAuthor(), stringToCheck, itemNameEntry.Text, strings.Split(itemDescEntry.Text, "\n")}
			
				proj := GetProject(CWP)
				proj.Items = append(proj.Items, newItem)
			
				// Reset the project to be the new proj pointer
				var index int
				for i, cmd := range PluginProjects {
					if proj.Name == cmd.Name {
						index = i
					}
				}
				PluginProjects[index] = *proj
				
				// HideModal() in MainWindow.go
				HideModal()
			} else {
				dialog.ShowError(errors.New("Unrecognized Material! Check https://hub.spigotmc.org/javadocs/bukkit/org/bukkit/Material.html"), GetWindow())
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

// Setting up for Crafting
func MakeTable() *widget.Table {
	table := widget.NewTable(
		func() (int,int) {
			return 3,3
		},
		func() fyne.CanvasObject {
			label := widget.NewEntry()
			return label
		},
		func(tci widget.TableCellID, f fyne.CanvasObject) {
			
		})
	table.SetColumnWidth(0, 100)
	table.SetColumnWidth(1, 100)
	table.SetColumnWidth(2, 100)
	table.OnSelected = func(id widget.TableCellID) {
		log.Print(id)
	}
	return table
}