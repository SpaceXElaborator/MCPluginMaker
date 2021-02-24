package pluginwidgets

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type TypeSearchableList struct {
	widget.BaseWidget
	Text  string
	Items []string
}

func NewTypeSearchableList(text string, items []string) *TypeSearchableList {
	tsl := &TypeSearchableList{
		Text:  text,
		Items: items,
	}
	return tsl
}

func (tsl *TypeSearchableList) Refresh() {
	tsl.BaseWidget.Refresh()
}

func (tsl *TypeSearchableList) CreateRenderer() fyne.WidgetRenderer {
	tsl.ExtendBaseWidget(tsl)

	r := &typeSearchableListRenderer{
		tsl,
		[]fyne.CanvasObject{},
	}
	return r
}

type typeSearchableListRenderer struct {
	tsl *TypeSearchableList

	objects []fyne.CanvasObject
}

func (tslr *typeSearchableListRenderer) Layout(size fyne.Size) {

}

func (tslr *typeSearchableListRenderer) Objects() []fyne.CanvasObject {
	return tslr.objects
}

func (tslr *typeSearchableListRenderer) MinSize() fyne.Size {
	return fyne.NewSize(200, 300)
}

func (tslr *typeSearchableListRenderer) Destroy() {

}

func (tslr *typeSearchableListRenderer) Refresh() {

}
