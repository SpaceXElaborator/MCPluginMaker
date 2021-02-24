package pluginwidgets

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type TypeSearchableList struct {
	*widget.Entry
	Canvas *fyne.Canvas
}

func NewTypeSearchableList(canv *fyne.Canvas) *TypeSearchableList {
	tsl := &TypeSearchableList{
		widget.NewEntry(),
		canv,
	}
	return tsl
}

func (tsl *TypeSearchableList) KeyUp(key *fyne.KeyEvent) {
	if len(tsl.Text) >= 3 {
		log.Print(tsl.Text)
		log.Print(tsl.Position())
		modal := widget.NewPopUp(widget.NewLabel("Test"), *tsl.Canvas)
		modal.Move(tsl.Position())
		modal.Show()
	}
}
