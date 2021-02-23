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
	cmdTypes := []string{"Add Item", "Add Custom Item", "Set Health", "Set Food Level", "Send Message"}

	cmdFuncType := widget.NewSelect(cmdTypes, func(s string) {
		modal.Hide()
		if s == "Add Custom Item" {
			modal = widget.NewModalPopUp(widget.NewCard("Spawn Custom Item", "", spawnItemForm(cmd, true)), w.Canvas())
		} else if s == "Add Item" {
			modal = widget.NewModalPopUp(widget.NewCard("Spawn Item", "", spawnItemForm(cmd, false)), w.Canvas())
		} else if s == "Set Health" {
			modal = widget.NewModalPopUp(widget.NewCard("Set Health", "", playerFuncFloatValue(cmd, "Set Health", "setHealth")), w.Canvas())
		} else if s == "Set Food Level" {
			modal = widget.NewModalPopUp(widget.NewCard("Set Food Level", "", playerFuncIntValue(cmd, "Set Food", "setFoodLevel")), w.Canvas())
		} else if s == "Send Message" {
			modal = widget.NewModalPopUp(widget.NewCard("Send Message", "", playerFuncStringValue(cmd, "Send Message", "sendMessage")), w.Canvas())
		}
		modal.Resize(fyne.NewSize(512, 0))
		modal.Show()
	})

	cmdFuncTypeFormItem := &widget.FormItem{
		Text:   "Command Function",
		Widget: cmdFuncType,
	}

	funcForm.AppendItem(cmdFuncTypeFormItem)

	return funcForm
}

func playerFuncIntValue(cmd *PluginCommands.Command, nameOf, spigotFunction string) *widget.Form {
	form := widget.NewForm()

	intToEnter := widget.NewEntry()
	form.Append("Integer", intToEnter)

	form.OnSubmit = func() {
		if intToEnter.Text != "" {
			if _, err := strconv.Atoi(intToEnter.Text); err == nil {
				cmd.AddPlayerFunc(nameOf, "p."+spigotFunction+"("+intToEnter.Text+");")
				HideModal()
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

func playerFuncFloatValue(cmd *PluginCommands.Command, nameOf, spigotFunction string) *widget.Form {
	form := widget.NewForm()

	intToEnter := widget.NewEntry()
	form.Append("Double", intToEnter)

	form.OnSubmit = func() {
		if intToEnter.Text != "" {
			if floatVal, err := strconv.ParseFloat(intToEnter.Text, 10); err == nil {
				cmd.AddPlayerFunc(nameOf, "p."+spigotFunction+"("+strconv.FormatFloat(floatVal, 'f', 1, 64)+");")
				HideModal()
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

func playerFuncStringValue(cmd *PluginCommands.Command, nameOf, spigotFunction string) *widget.Form {
	form := widget.NewForm()

	stringToEnter := widget.NewEntry()
	form.Append("Value", stringToEnter)

	form.OnSubmit = func() {
		if stringToEnter.Text != "" {
			cmd.AddPlayerFunc(nameOf, "p."+spigotFunction+"(\""+stringToEnter.Text+"\");")
			HideModal()
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
			if itemAmount.Text != "" && itemName != "" {
				if _, err := strconv.Atoi(itemAmount.Text); err == nil {
					cmd.AddPlayerFunc("Add Custom Item", "p.getInventory().addItem("+PluginSettings.GetCWP()+"CustomItems.build(\""+itemName+"\", "+itemAmount.Text+"));")
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
