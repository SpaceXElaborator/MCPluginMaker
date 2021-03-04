package pluginfunctions

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
	modal  *widget.PopUp
	canvas *fyne.Canvas
	window *fyne.Window
)

type valueType struct {
	values []string
}

// InitCommands Going to be used to create some variables that will be used for creating commands+
func InitCommands(canv *fyne.Canvas, win *fyne.Window) {
	canvas = canv
	window = win
}

// PlayerCommandFuncAddForm Creates a form to add a player based Java method to the commands PlayerFuncs to be exported and built in Java
func playerCommandFuncAddForm(cmd *PluginCommands.Command, ContentFunc func(), items []*PluginItems.CustomItem) *widget.Form {
	funcForm := widget.NewForm()
	cmdTypes := []string{"Add Item", "Add Custom Item", "Set Health", "Set Food Level", "Send Message", "Set Display Name", "Set Level", "Set Exp", "Set Max Health", "Set Gamemode"}

	cmdFuncType := widget.NewSelect(cmdTypes, func(s string) {
		modal.Hide()
		if s == "Add Custom Item" {
			modal = widget.NewModalPopUp(widget.NewCard(s, "", spawnItemForm(ContentFunc, items, window, cmd, true)), *canvas)
		} else if s == "Add Item" {
			modal = widget.NewModalPopUp(widget.NewCard(s, "", spawnItemForm(ContentFunc, items, window, cmd, false)), *canvas)
		} else if s == "Send Message" {
			modal = widget.NewModalPopUp(widget.NewCard(s, "", playerFuncAddition(ContentFunc, cmd, "string", "Send Message", "", "sendMessage", nil, nil)), *canvas)
		} else if s == "Set Display Name" {
			modal = widget.NewModalPopUp(widget.NewCard(s, "", playerFuncAddition(ContentFunc, cmd, "string", "Set Name", "", "setDisplayName", nil, nil)), *canvas)
		} else if s == "Set Level" {
			modal = widget.NewModalPopUp(widget.NewCard(s, "", playerFuncAddition(ContentFunc, cmd, "int", "Set Level", "", "setLevel", nil, nil)), *canvas)
		} else if s == "Set Exp" {
			modal = widget.NewModalPopUp(widget.NewCard(s, "", playerFuncAddition(ContentFunc, cmd, "float", "Set Exp", "", "setExp", nil, nil)), *canvas)
		} else if s == "Set Max Health" {
			modal = widget.NewModalPopUp(widget.NewCard(s, "", playerFuncAddition(ContentFunc, cmd, "float", "Set Max Health", "", "setMaxHealth", nil, nil)), *canvas)
		} else if s == "Set Health" {
			modal = widget.NewModalPopUp(widget.NewCard(s, "", playerFuncAddition(ContentFunc, cmd, "float", "Set Health", "", "setHealth", nil, nil)), *canvas)
		} else if s == "Set Food Level" {
			modal = widget.NewModalPopUp(widget.NewCard(s, "", playerFuncAddition(ContentFunc, cmd, "int", "Set Food", "", "setFoodLevel", nil, nil)), *canvas)
		} else if s == "Set Gamemode" {
			modal = widget.NewModalPopUp(widget.NewCard(s, "", playerFuncAddition(ContentFunc, cmd, "list", "Set Gamemode", "GameMode", "setGameMode", []string{"Survival", "Creative", "Adventure", "Spectator"}, []string{"import org.bukkit.GameMode;"})), *canvas)
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
		modal.Hide()
	}
	funcForm.Refresh()
	return funcForm
}

func playerFuncAddition(ContentFunc func(), cmd *PluginCommands.Command, addType, nameOf, enumVal, spigotFunction string, listValues, imports []string) *widget.Form {
	form := widget.NewForm()

	// Variables to be grabbed later for the creation of the function
	var widgetToAdd fyne.CanvasObject
	selectItemToSet := ""

	// Only create an Entry widget if its an int, string, or float value. If it's a list, make sure to set widgetToAdd as a Select widget
	if addType == "int" || addType == "string" || addType == "float" {
		widgetToAdd = widget.NewEntry()
	} else if addType == "list" {
		widgetToAdd = widget.NewSelect(listValues, func(s string) {
			selectItemToSet = s
		})
	}

	// Hoping to use the value string to better help people determine what value needs to go in. Will make this more user friendly soon
	if addType == "list" {
		form.Append("Select", widgetToAdd)
	} else {
		form.Append("Value ("+addType+")", widgetToAdd)
	}

	form.OnSubmit = func() {
		if addType == "list" {
			// Lists will be usually assocciated with enums, so we need to make sure to add that kind of functionality when adding the PlayerFunc
			if selectItemToSet != "" {
				// Gets the list of imports and add it to the command
				if len(imports) >= 1 {
					for _, imps := range imports {
						cmd.AddImport(imps)
					}
				}
				cmd.AddPlayerFunc(nameOf, "p."+spigotFunction+"("+enumVal+".valueOf(\""+strings.ToUpper(selectItemToSet)+"\"));")
				hideFuncModal()

				// ContentFunc is gathered from MainWindow to SetNewContent()
				ContentFunc()
			} else {
				dialog.NewError(errors.New("Must Select Value"), *window)
			}
		} else {
			if widgetToAdd.(*widget.Entry).Text != "" {
				// For int/float. Need to make sure that the values are what they need to be. If so, add the function to the PlayerFunc
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
				// Gets the list of imports and add it to the command
				if len(imports) >= 1 {
					for _, imps := range imports {
						cmd.AddImport(imps)
					}
				}
				hideFuncModal()

				// ContentFunc is gathered from MainWindow to SetNewContent()
				ContentFunc()
			} else {
				dialog.NewError(errors.New("Must Input Value"), *window)
			}
		}
	}

	form.OnCancel = func() {
		hideFuncModal()
	}
	form.Refresh()
	return form
}

// SpawnItemForm for right now needs its own function until I can find a way to handle with multiple "." methods in Java. In this case, p."getInventory()".addItem()
func spawnItemForm(ContentFunc func(), projItems []*PluginItems.CustomItem, window *fyne.Window, cmd *PluginCommands.Command, custom bool) *widget.Form {
	itemForm := widget.NewForm()

	itemName := ""

	if custom {
		// Create a Select widget that holds all of the CustomItems currently in the project
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
					// Get the CustomItem from the CustomItems.java file that is generated when you have custom items
					cmd.AddPlayerFunc("Add Custom Item", "p.getInventory().addItem("+PluginSettings.GetCWP()+"CustomItems.build(\""+itemName+"\", "+itemAmount.Text+"));")
					hideFuncModal()

					// ContentFunc is gathered from MainWindow to SetNewContent()
					ContentFunc()
				} else {
					dialog.ShowError(errors.New("Amount must be a number"), *window)
				}
			} else {
				dialog.ShowError(errors.New("Fill in all values"), *window)
			}
		}
		itemForm.OnCancel = func() {
			hideFuncModal()
		}
		itemForm.Refresh()
	} else {
		itemType := widget.NewEntry()
		itemAmount := widget.NewEntry()
		itemForm.Append("Item Material", itemType)
		itemForm.Append("Amount", itemAmount)

		itemForm.OnSubmit = func() {
			if itemType.Text != "" && itemAmount.Text != "" {
				// Check to make sure that the material is valid material
				if PluginItems.CheckMaterial(itemType.Text) {
					if _, err := strconv.Atoi(itemAmount.Text); err == nil {
						// Create the new itemstack with the given material in Java code and add it to PlayerFunc
						cmd.AddPlayerFunc("Add Item", "p.getInventory().addItem(new ItemStack(Material.valueOf(\""+strings.ToUpper(itemType.Text)+"\"), "+itemAmount.Text+"));")
						hideFuncModal()

						// ContentFunc is gathered from MainWindow to SetNewContent()
						ContentFunc()
					} else {
						dialog.ShowError(errors.New("Amount must be a number"), *window)
					}
				} else {
					dialog.ShowError(errors.New("Item Doesn't Exist"), *window)
				}
			} else {
				dialog.ShowError(errors.New("Fill in all values"), *window)
			}
		}

		itemForm.OnCancel = func() {
			hideFuncModal()
		}

		itemForm.Refresh()
	}
	return itemForm
}

func hideFuncModal() {
	modal.Hide()
}
