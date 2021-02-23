package plugingui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	PluginSettings "SpaceXElaborator/PluginMaker/Settings"
)

// createSettingsForm Generates a form so that the user can change who they are and if they want to be dark or light mode
func createSettingsForm() *widget.Form {
	themeCheck := widget.NewCheck("", func(on bool) {
		if on == true {
			// GetApp() in MainWindow.go
			GetApp().Settings().SetTheme(theme.DarkTheme())
			PluginSettings.SetDark(true)
		} else {
			// GetApp() in MainWindow.go
			GetApp().Settings().SetTheme(theme.LightTheme())
			PluginSettings.SetDark(false)
		}
	})
	themeCheck.SetChecked(PluginSettings.GetDark())
	checkFormItem := &widget.FormItem{
		Text:   "Dark Mode",
		Widget: themeCheck,
	}

	authorNameEntry := widget.NewEntry()
	authorNameEntry.Resize(fyne.NewSize(300, 300))
	authorNameEntry.SetText(PluginSettings.GetAuthor())
	authorNameFormItem := &widget.FormItem{
		Text:   "Author",
		Widget: authorNameEntry,
	}

	newSettingsForm := widget.NewForm(authorNameFormItem, checkFormItem)
	newSettingsForm.OnSubmit = func() {
		if authorNameEntry.Text != "" {
			PluginSettings.SetAuthor(authorNameEntry.Text)
			// GetWindow() in MainWindow.go
			GetWindow().SetTitle("MCPluginMaker | " + PluginSettings.GetAuthor())
		} else {
			PluginSettings.SetAuthor("User")
		}
		// HideModal() in MainWindow.go
		HideModal()
	}
	newSettingsForm.OnCancel = func() {
		// HideModal() in MainWindow.go
		HideModal()
	}
	return newSettingsForm
}
