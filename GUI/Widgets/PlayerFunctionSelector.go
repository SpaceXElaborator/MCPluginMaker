package pluginwidgets

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type PluginFunctionSelector struct {
	widget.BaseWidget

	label   widget.Label
	hovered bool
}

func NewPlayerFunctionSelector() *PluginFunctionSelector {
	e := &PluginFunctionSelector{}
	e.ExtendBaseWidget(e)
	return e
}

func (pfs *PluginFunctionSelector) CreateRenderer() fyne.WidgetRenderer {
	pfs.ExtendBaseWidget(pfs)
	background := canvas.NewRectangle(theme.InputBackgroundColor())
	line := canvas.NewRectangle(theme.ShadowColor())
	label := widget.NewLabel("Test")
	objects := []fyne.CanvasObject{label, background, line}

	pfsr := &PluginFunctionSelectorRenderer{
		main:       pfs,
		label:      label,
		background: background,
		line:       line,
		objects:    objects,
	}
	return pfsr
}

func (pfs *PluginFunctionSelector) MouseIn(*desktop.MouseEvent) {
	pfs.hovered = true
	pfs.Refresh()
}

func (pfs *PluginFunctionSelector) MouseMoved(*desktop.MouseEvent) {}

func (pfs *PluginFunctionSelector) MouseOut(*desktop.MouseEvent) {
	pfs.hovered = false
	pfs.Refresh()
}

type PluginFunctionSelectorRenderer struct {
	main *PluginFunctionSelector

	label      *widget.Label
	background *canvas.Rectangle
	line       *canvas.Rectangle
	objects    []fyne.CanvasObject
}

func (pfsr *PluginFunctionSelectorRenderer) Destroy() {}

func (pfsr *PluginFunctionSelectorRenderer) Layout(size fyne.Size) {
	pfsr.line.Resize(fyne.NewSize(size.Width, theme.InputBorderSize()))
	pfsr.line.Move(fyne.NewPos(0, size.Height-theme.InputBorderSize()))
	pfsr.background.Resize(fyne.NewSize(size.Width, size.Height-theme.InputBorderSize()*2))
	pfsr.background.Move(fyne.NewPos(0, theme.InputBorderSize()))

	pfsr.label.Resize(pfsr.label.MinSize())
	pfsr.label.Move(fyne.NewPos(0, theme.InputBorderSize()))
}

func (pfsr *PluginFunctionSelectorRenderer) MinSize() (size fyne.Size) {
	return fyne.NewSize(0, 0)
}

func (pfsr *PluginFunctionSelectorRenderer) Objects() []fyne.CanvasObject {
	return pfsr.objects
}

func (pfsr *PluginFunctionSelectorRenderer) Refresh() {
	pfsr.label.Refresh()
	pfsr.background.FillColor = pfsr.resetThemes()
}

func (pfsr *PluginFunctionSelectorRenderer) resetThemes() color.Color {
	if pfsr.main.hovered {
		return theme.PressedColor()
	}
	return theme.InputBackgroundColor()
}
