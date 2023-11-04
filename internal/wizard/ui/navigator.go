package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type Navigator struct {
	widget.BaseWidget

	Screens []fyne.CanvasObject

	OnCancel  func()
	OnInstall func()

	curScreen int

	screenContainer *screenContainer
	backBtn         *widget.Button
	fwdBtn          *widget.Button
}

func NewNavigator(screens ...fyne.CanvasObject) *Navigator {
	n := &Navigator{
		Screens: screens,
	}
	n.ExtendBaseWidget(n)
	return n
}

func (n *Navigator) CreateRenderer() fyne.WidgetRenderer {
	n.backBtn = widget.NewButton("Back", n.goBack)
	n.backBtn.Icon = theme.NavigateBackIcon()
	n.fwdBtn = widget.NewButton("Next", n.goForward)
	n.fwdBtn.Icon = theme.NavigateNextIcon()
	cancelBtn := widget.NewButton("Cancel", func() {
		if n.OnCancel != nil {
			n.OnCancel()
		}
	})
	cancelBtn.Icon = theme.CancelIcon()

	var screen fyne.CanvasObject = layout.NewSpacer()
	if len(n.Screens) > 0 {
		screen = n.Screens[0]
	}
	n.screenContainer = NewScreenContainer()
	n.screenContainer.Screen = screen

	navigationBar := container.NewStack(
		canvas.NewRectangle(fyne.CurrentApp().Settings().Theme().Color(theme.ColorGray, fyne.CurrentApp().Settings().ThemeVariant())),
		container.NewHBox(layout.NewSpacer(), n.backBtn, n.fwdBtn, cancelBtn),
	)

	content := container.NewBorder(nil, navigationBar, nil, nil, n.screenContainer)
	return widget.NewSimpleRenderer(content)
}

func (n *Navigator) goBack() {
	if n.curScreen > 0 {
		n.curScreen--
	}
	n.Refresh()
}

func (n *Navigator) goForward() {
	if n.curScreen == len(n.Screens)-1 {
		if n.OnInstall != nil {
			n.OnInstall()
		}
	} else {
		n.curScreen++
		n.Refresh()
	}
}

func (n *Navigator) Refresh() {
	if n.curScreen < len(n.Screens) {
		n.screenContainer.SetScreen(n.Screens[n.curScreen])
	}
	n.BaseWidget.Refresh()
}

type screenContainer struct {
	widget.BaseWidget

	Screen fyne.CanvasObject

	container *fyne.Container
}

func NewScreenContainer() *screenContainer {
	s := &screenContainer{}
	s.ExtendBaseWidget(s)
	return s
}

func (s *screenContainer) SetScreen(screen fyne.CanvasObject) {
	s.Screen = screen
	s.Refresh()
}

func (s *screenContainer) Refresh() {
	s.container.Objects[0] = s.Screen
	s.BaseWidget.Refresh()
}

func (s *screenContainer) CreateRenderer() fyne.WidgetRenderer {
	screen := s.Screen
	if screen == nil {
		screen = layout.NewSpacer()
	}
	s.container = container.NewBorder(widget.NewLabel(""), nil, &hSpace{pad: 10}, &hSpace{pad: 5}, screen)
	return widget.NewSimpleRenderer(s.container)
}

type hSpace struct {
	widget.BaseWidget
	pad float32
}

func (h *hSpace) MinSize() fyne.Size {
	return fyne.NewSize(h.pad, 0)
}

func (*hSpace) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(layout.NewSpacer())
}
