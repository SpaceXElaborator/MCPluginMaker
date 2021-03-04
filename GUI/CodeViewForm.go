package plugingui

import (
	plugincommands "SpaceXElaborator/PluginMaker/Command"
	pluginsettings "SpaceXElaborator/PluginMaker/Settings"
	"io/ioutil"
	"log"

	"fyne.io/fyne/v2/widget"
)

func openCodeForm(cmd *plugincommands.Command) *widget.Form {
	form := widget.NewForm()

	content, err := ioutil.ReadFile("projects/" + pluginsettings.GetCWP() + "/src/main/java/com/" + pluginsettings.GetAuthor() + "/net/cmds/" + cmd.Name + ".java")
	if err != nil {
		log.Fatal(err)
	}

	text := string(content)
	textgrid := widget.NewTextGridFromString(text)

	form.Append("", textgrid)
	form.CancelText = "Close"
	form.OnCancel = func() {
		modal.Hide()
	}

	form.Refresh()
	return form
}
