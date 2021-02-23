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
			modal = widget.NewModalPopUp(widget.NewCard(s, "", playerFuncStringValue(cmd, "Send Message", "sendMessage", nil)), w.Canvas())
		} else if s == "Set Display Name" {
			modal = widget.NewModalPopUp(widget.NewCard(s, "", playerFuncStringValue(cmd, "Set Name", "setDisplayName", nil)), w.Canvas())
		} else if s == "Set Level" {
			modal = widget.NewModalPopUp(widget.NewCard(s, "", playerFuncIntValue(cmd, "Set Level", "setLevel", nil)), w.Canvas())
		} else if s == "Set Exp" {
			modal = widget.NewModalPopUp(widget.NewCard(s, "", playerFuncFloatValue(cmd, "Set Exp", "setExp", nil)), w.Canvas())
		} else if s == "Set Max Health" {
			modal = widget.NewModalPopUp(widget.NewCard(s, "", playerFuncFloatValue(cmd, "Set Max Health", "setMaxHealth", nil)), w.Canvas())
		} else if s == "Set Health" {
			modal = widget.NewModalPopUp(widget.NewCard(s, "", playerFuncFloatValue(cmd, "Set Health", "setHealth", nil)), w.Canvas())
		} else if s == "Set Food Level" {
			modal = widget.NewModalPopUp(widget.NewCard(s, "", playerFuncIntValue(cmd, "Set Food", "setFoodLevel", nil)), w.Canvas())
		} else if s == "Set Gamemode" {
			modal = widget.NewModalPopUp(widget.NewCard(s, "", playerFuncListValue(cmd, "Set Gamemode", "GameMode", "setGameMode", []string{"Survival", "Creative", "Adventure", "Spectator"}, []string{"import org.bukkit.GameMode;"})), w.Canvas())
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

func playerFuncListValue(cmd *PluginCommands.Command, nameOf, enumVal, spigotFunction string, listValues, imports []string) *widget.Form {
	form := widget.NewForm()

	stringValue := ""

	funcList := widget.NewSelect(listValues, func(s string) {
		stringValue = s
	})

	form.Append("Select", funcList)

	form.OnSubmit = func() {
		if stringValue != "" {
			if len(imports) >= 1 {
				for _, imps := range imports {
					cmd.AddImport(imps)
				}
			}
			cmd.AddPlayerFunc(nameOf, "p."+spigotFunction+"("+enumVal+".valueOf(\""+strings.ToUpper(stringValue)+"\"));")
			HideModal()
			SetNewContent()
		} else {
			dialog.NewError(errors.New("Must Select Value"), GetWindow())
		}
	}

	form.OnCancel = func() {
		HideModal()
	}
	form.Refresh()
	return form
}

func playerFuncIntValue(cmd *PluginCommands.Command, nameOf, spigotFunction string, imports []string) *widget.Form {
	form := widget.NewForm()

	intToEnter := widget.NewEntry()
	form.Append("Integer", intToEnter)

	form.OnSubmit = func() {
		if intToEnter.Text != "" {
			if _, err := strconv.Atoi(intToEnter.Text); err == nil {
				if len(imports) >= 1 {
					for _, imps := range imports {
						cmd.AddImport(imps)
					}
				}
				cmd.AddPlayerFunc(nameOf, "p."+spigotFunction+"("+intToEnter.Text+");")
				HideModal()
				SetNewContent()
			} else {
				dialog.NewError(errors.New("Value Must Be A Number"), GetWindow())
			}
		} else {
			dialog.NewError(errors.New("Must Enter Value"), GetWindow())
		}
	}

	form.OnCancel = func() {
		HideModal()
	}
	form.Refresh()
	return form
}

func playerFuncFloatValue(cmd *PluginCommands.Command, nameOf, spigotFunction string, imports []string) *widget.Form {
	form := widget.NewForm()

	intToEnter := widget.NewEntry()
	form.Append("Double", intToEnter)

	form.OnSubmit = func() {
		if intToEnter.Text != "" {
			if floatVal, err := strconv.ParseFloat(intToEnter.Text, 10); err == nil {
				if len(imports) >= 1 {
					for _, imps := range imports {
						cmd.AddImport(imps)
					}
				}
				cmd.AddPlayerFunc(nameOf, "p."+spigotFunction+"("+strconv.FormatFloat(floatVal, 'f', 1, 64)+");")
				HideModal()
				SetNewContent()
			} else {
				dialog.NewError(errors.New("Value Must Be A Float (I.E. 1.0)"), GetWindow())
			}
		} else {
			dialog.NewError(errors.New("Must Enter Value"), GetWindow())
		}
	}

	form.OnCancel = func() {
		HideModal()
	}
	form.Refresh()
	return form
}

func playerFuncStringValue(cmd *PluginCommands.Command, nameOf, spigotFunction string, imports []string) *widget.Form {
	form := widget.NewForm()

	stringToEnter := widget.NewEntry()
	form.Append("Value", stringToEnter)

	form.OnSubmit = func() {
		if stringToEnter.Text != "" {
			if len(imports) >= 1 {
				for _, imps := range imports {
					cmd.AddImport(imps)
				}
			}
			cmd.AddPlayerFunc(nameOf, "p."+spigotFunction+"(\""+stringToEnter.Text+"\");")
			HideModal()
			SetNewContent()
		} else {
			dialog.NewError(errors.New("Must Enter Value"), GetWindow())
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
