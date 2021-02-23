package PluginGUI

import (
	"errors"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"

	PluginCommands "SpaceXElaborator/PluginMaker/Command"
	PluginSettings "SpaceXElaborator/PluginMaker/Settings"
)

func playerCommandFuncAddForm(cmd *PluginCommands.Command) *widget.Form {
	funcForm := widget.NewForm()
	cmdTypes := []string{"Add Item", "Add Custom Item", "Set Health", "Set Food Level", "Send Message", "Set Display Name", "Set Level", "Set Exp", "Set Max Health", "Set Gamemode"}

	cmdFuncType := widget.NewSelect(cmdTypes, func(s string) {
		modal.Hide()
		if s == "Add Custom Item" {
			modal = widget.NewModalPopUp(widget.NewCard(s, "", spawnItemForm(cmd, true)), w.Canvas())
		} else if s == "Add Item" {
			modal = widget.NewModalPopUp(widget.NewCard(s, "", spawnItemForm(cmd, false)), w.Canvas())
		} else if s == "Send Message" {
			modal = widget.NewModalPopUp(widget.NewCard(s, "", playerFuncAddition(cmd, "string", "Send Message", "", "sendMessage", nil, nil)), w.Canvas())
		} else if s == "Set Display Name" {
			modal = widget.NewModalPopUp(widget.NewCard(s, "", playerFuncAddition(cmd, "string", "Set Name", "", "setDisplayName", nil, nil)), w.Canvas())
		} else if s == "Set Level" {
			modal = widget.NewModalPopUp(widget.NewCard(s, "", playerFuncAddition(cmd, "int", "Set Level", "", "setLevel", nil, nil)), w.Canvas())
		} else if s == "Set Exp" {
			modal = widget.NewModalPopUp(widget.NewCard(s, "", playerFuncAddition(cmd, "float", "Set Exp", "", "setExp", nil, nil)), w.Canvas())
		} else if s == "Set Max Health" {
			modal = widget.NewModalPopUp(widget.NewCard(s, "", playerFuncAddition(cmd, "float", "Set Max Health", "", "setMaxHealth", nil, nil)), w.Canvas())
		} else if s == "Set Health" {
			modal = widget.NewModalPopUp(widget.NewCard(s, "", playerFuncAddition(cmd, "float", "Set Health", "", "setHealth", nil, nil)), w.Canvas())
		} else if s == "Set Food Level" {
			modal = widget.NewModalPopUp(widget.NewCard(s, "", playerFuncAddition(cmd, "int", "Set Food", "", "setFoodLevel", nil, nil)), w.Canvas())
		} else if s == "Set Gamemode" {
			modal = widget.NewModalPopUp(widget.NewCard(s, "", playerFuncAddition(cmd, "list", "Set Gamemode", "GameMode", "setGameMode", []string{"Survival", "Creative", "Adventure", "Spectator"}, []string{"import org.bukkit.GameMode;"})), w.Canvas())
		}
		modal.Resize(fyne.NewSize(512, 0))
		modal.Show()
	})

	cmdFuncTypeFormItem := &widget.FormItem{
		Text:   "Command Function",
		Widget: cmdFuncType,
	}

	funcForm.AppendItem(cmdFuncTypeFormItem)

	funcForm.OnCancel = func() {
		HideModal()
	}
	funcForm.Refresh()
	return funcForm
}

func playerFuncAddition(cmd *PluginCommands.Command, addType, nameOf, enumVal, spigotFunction string, listValues, imports []string) *widget.Form {
	form := widget.NewForm()

	var widgetToAdd fyne.CanvasObject
	selectItemToSet := ""

	if addType == "int" || addType == "string" || addType == "float" {
		widgetToAdd = widget.NewEntry()
	} else if addType == "list" {
		widgetToAdd = widget.NewSelect(listValues, func(s string) {
			selectItemToSet = s
		})
	}

	if addType == "list" {
		form.Append("Select", widgetToAdd)
	} else {
		form.Append("Value ("+addType+")", widgetToAdd)
	}

	form.OnSubmit = func() {
		if addType == "list" {
			if selectItemToSet != "" {
				if len(imports) >= 1 {
					for _, imps := range imports {
						cmd.AddImport(imps)
					}
				}
				cmd.AddPlayerFunc(nameOf, "p."+spigotFunction+"("+enumVal+".valueOf(\""+strings.ToUpper(selectItemToSet)+"\"));")
				HideModal()
				SetNewContent()
			} else {
				dialog.NewError(errors.New("Must Select Value"), GetWindow())
			}
		} else {
			if widgetToAdd.(*widget.Entry).Text != "" {
				if addType == "int" {
					if _, err := strconv.Atoi(widgetToAdd.(*widget.Entry).Text); err != nil {
						return
					}
					cmd.AddPlayerFunc(nameOf, "p."+spigotFunction+"("+widgetToAdd.(*widget.Entry).Text+");")
				} else if addType == "float" {
					if _, err := strconv.ParseFloat(widgetToAdd.(*widget.Entry).Text, 10); err != nil {
						return
					}
					floatVal, _ := strconv.ParseFloat(widgetToAdd.(*widget.Entry).Text, 10)
					cmd.AddPlayerFunc(nameOf, "p."+spigotFunction+"("+strconv.FormatFloat(floatVal, 'f', 1, 64)+");")
				} else {
					cmd.AddPlayerFunc(nameOf, "p."+spigotFunction+"(\""+widgetToAdd.(*widget.Entry).Text+"\");")
				}
				if len(imports) >= 1 {
					for _, imps := range imports {
						cmd.AddImport(imps)
					}
				}
				HideModal()
				SetNewContent()
			} else {
				dialog.NewError(errors.New("Must Input Value"), GetWindow())
			}
		}
	}

	form.OnCancel = func() {
		HideModal()
	}
	form.Refresh()
	return form
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
			cmd.AddImport("import org.bukkit.Material;")
			cmd.AddImport("import org.bukkit.inventory.ItemStack;")
			if itemAmount.Text != "" && itemName != "" {
				if _, err := strconv.Atoi(itemAmount.Text); err == nil {
					cmd.AddPlayerFunc("Add Custom Item", "p.getInventory().addItem("+PluginSettings.GetCWP()+"CustomItems.build(\""+itemName+"\", "+itemAmount.Text+"));")
					HideModal()
					SetNewContent()
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
		itemType := widget.NewEntry()
		itemAmount := widget.NewEntry()
		itemForm.Append("Item Material", itemType)
		itemForm.Append("Amount", itemAmount)

		itemForm.OnSubmit = func() {
			if itemType.Text != "" && itemAmount.Text != "" {
				if proj.CheckMaterial(itemType.Text) {
					if _, err := strconv.Atoi(itemAmount.Text); err == nil {
						cmd.AddPlayerFunc("Add Item", "p.getInventory().addItem(new ItemStack(Material.valueOf(\""+strings.ToUpper(itemType.Text)+"\"), "+itemAmount.Text+"));")
						HideModal()
						SetNewContent()
					} else {
						dialog.ShowError(errors.New("Amount must be a number"), GetWindow())
					}
				} else {
					dialog.ShowError(errors.New("Item Doesn't Exist!"), GetWindow())
				}
			} else {
				dialog.ShowError(errors.New("Fill in all values"), GetWindow())
			}
		}

		itemForm.OnCancel = func() {
			HideModal()
		}

		itemForm.Refresh()
	}
	return itemForm
}
