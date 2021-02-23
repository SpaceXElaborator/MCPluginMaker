package PluginFunction

import (
	"errors"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"

	PluginCommands "SpaceXElaborator/PluginMaker/Command"
	PluginItems "SpaceXElaborator/PluginMaker/Item"
	PluginSettings "SpaceXElaborator/PluginMaker/Settings"
)

var (
	modal *widget.PopUp
)

func PlayerCommandFuncAddForm(cmd *PluginCommands.Command, canvas *fyne.Canvas, window *fyne.Window, HideModal, ContentFunc func(), items []*PluginItems.CustomItem) *widget.Form {
	funcForm := widget.NewForm()
	cmdTypes := []string{"Add Item", "Add Custom Item", "Set Health", "Set Food Level", "Send Message", "Set Display Name", "Set Level", "Set Exp", "Set Max Health", "Set Gamemode"}

	cmdFuncType := widget.NewSelect(cmdTypes, func(s string) {
		HideModal()
		if s == "Add Custom Item" {
			modal = widget.NewModalPopUp(widget.NewCard(s, "", spawnItemForm(ContentFunc, items, window, cmd, true)), *canvas)
		} else if s == "Add Item" {
			modal = widget.NewModalPopUp(widget.NewCard(s, "", spawnItemForm(ContentFunc, items, window, cmd, false)), *canvas)
		} else if s == "Send Message" {
			modal = widget.NewModalPopUp(widget.NewCard(s, "", playerFuncAddition(HideModal, ContentFunc, window, cmd, "string", "Send Message", "", "sendMessage", nil, nil)), *canvas)
		} else if s == "Set Display Name" {
			modal = widget.NewModalPopUp(widget.NewCard(s, "", playerFuncAddition(HideModal, ContentFunc, window, cmd, "string", "Set Name", "", "setDisplayName", nil, nil)), *canvas)
		} else if s == "Set Level" {
			modal = widget.NewModalPopUp(widget.NewCard(s, "", playerFuncAddition(HideModal, ContentFunc, window, cmd, "int", "Set Level", "", "setLevel", nil, nil)), *canvas)
		} else if s == "Set Exp" {
			modal = widget.NewModalPopUp(widget.NewCard(s, "", playerFuncAddition(HideModal, ContentFunc, window, cmd, "float", "Set Exp", "", "setExp", nil, nil)), *canvas)
		} else if s == "Set Max Health" {
			modal = widget.NewModalPopUp(widget.NewCard(s, "", playerFuncAddition(HideModal, ContentFunc, window, cmd, "float", "Set Max Health", "", "setMaxHealth", nil, nil)), *canvas)
		} else if s == "Set Health" {
			modal = widget.NewModalPopUp(widget.NewCard(s, "", playerFuncAddition(HideModal, ContentFunc, window, cmd, "float", "Set Health", "", "setHealth", nil, nil)), *canvas)
		} else if s == "Set Food Level" {
			modal = widget.NewModalPopUp(widget.NewCard(s, "", playerFuncAddition(HideModal, ContentFunc, window, cmd, "int", "Set Food", "", "setFoodLevel", nil, nil)), *canvas)
		} else if s == "Set Gamemode" {
			modal = widget.NewModalPopUp(widget.NewCard(s, "", playerFuncAddition(HideModal, ContentFunc, window, cmd, "list", "Set Gamemode", "GameMode", "setGameMode", []string{"Survival", "Creative", "Adventure", "Spectator"}, []string{"import org.bukkit.GameMode;"})), *canvas)
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

func playerFuncAddition(HideModal, ContentFunc func(), window *fyne.Window, cmd *PluginCommands.Command, addType, nameOf, enumVal, spigotFunction string, listValues, imports []string) *widget.Form {
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
				HideFuncModal()
				ContentFunc()
			} else {
				dialog.NewError(errors.New("Must Select Value"), *window)
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
				HideFuncModal()
				ContentFunc()
			} else {
				dialog.NewError(errors.New("Must Input Value"), *window)
			}
		}
	}

	form.OnCancel = func() {
		HideFuncModal()
	}
	form.Refresh()
	return form
}

func spawnItemForm(ContentFunc func(), projItems []*PluginItems.CustomItem, window *fyne.Window, cmd *PluginCommands.Command, custom bool) *widget.Form {
	itemForm := widget.NewForm()

	itemName := ""

	if custom {
		var items []string
		for _, item := range projItems {
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
					HideFuncModal()
					ContentFunc()
				} else {
					dialog.ShowError(errors.New("Amount must be a number"), *window)
				}
			} else {
				dialog.ShowError(errors.New("Fill in all values"), *window)
			}
		}
		itemForm.OnCancel = func() {
			HideFuncModal()
		}
		itemForm.Refresh()
	} else {
		itemType := widget.NewEntry()
		itemAmount := widget.NewEntry()
		itemForm.Append("Item Material", itemType)
		itemForm.Append("Amount", itemAmount)

		itemForm.OnSubmit = func() {
			if itemType.Text != "" && itemAmount.Text != "" {
				if PluginItems.CheckMaterial(itemType.Text) {
					if _, err := strconv.Atoi(itemAmount.Text); err == nil {
						cmd.AddPlayerFunc("Add Item", "p.getInventory().addItem(new ItemStack(Material.valueOf(\""+strings.ToUpper(itemType.Text)+"\"), "+itemAmount.Text+"));")
						HideFuncModal()
						ContentFunc()
					} else {
						dialog.ShowError(errors.New("Amount must be a number"), *window)
					}
				} else {
					dialog.ShowError(errors.New("Item Doesn't Exist!"), *window)
				}
			} else {
				dialog.ShowError(errors.New("Fill in all values"), *window)
			}
		}

		itemForm.OnCancel = func() {
			HideFuncModal()
		}

		itemForm.Refresh()
	}
	return itemForm
}

func HideFuncModal() {
	modal.Hide()
}
