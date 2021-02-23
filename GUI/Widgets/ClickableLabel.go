package pluginwidgets

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

// ClickableLabel will create a label using a basic widget.Label but add left/right click funcionality
type ClickableLabel struct {
	*widget.Label
	OnTapped          func() //`json:"-"`
	OnTappedSecondary func() //`json:"-"`
}

// NewClickableLabel Creates a ClickableLabel using the name, and two functions for left/right click
func NewClickableLabel(text string, tappedLeft, tappedRight func()) *ClickableLabel {
	return &ClickableLabel{
		widget.NewLabel(text),
		tappedLeft, tappedRight,
	}
}

// TappedSecondary Is a Fyne method to check for right clicking
func (cl *ClickableLabel) TappedSecondary(pe *fyne.PointEvent) {
	if cl.OnTappedSecondary != nil {
		cl.OnTappedSecondary()
	}
}

// Tapped Is a Fyne method to check for left clicking
func (cl *ClickableLabel) Tapped(pe *fyne.PointEvent) {
	if cl.OnTapped != nil {
		cl.OnTapped()
	}
}
