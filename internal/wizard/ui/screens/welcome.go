package screens

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/dweymouth/fyne-windows-installer/wizard/res"
)

type WelcomeScreen struct {
	widget.BaseWidget
}

func NewWelcomeScreen() *WelcomeScreen {
	w := &WelcomeScreen{}
	w.ExtendBaseWidget(w)
	return w
}

func (w *WelcomeScreen) CreateRenderer() fyne.WidgetRenderer {
	icon := canvas.NewImageFromResource(res.ResAppIcon)
	icon.SetMinSize(fyne.NewSize(100, 100))
	icon.FillMode = canvas.ImageFillContain

	appName := string(res.ResAppName.StaticContent)
	appVersion := string(res.ResAppVersion.StaticContent)
	welcome := widget.NewRichText(
		&widget.TextSegment{
			Text:  fmt.Sprintf("Welcome to the %s setup wizard", appName),
			Style: widget.RichTextStyle{SizeName: widget.RichTextStyleSubHeading.SizeName},
		},
		&widget.TextSegment{
			Text: fmt.Sprintf("This will install %s %s onto your computer.", appName, appVersion),
		})
	welcome.Wrapping = fyne.TextWrapWord

	content := container.NewBorder(nil, nil,
		container.NewVBox(
			icon,
			layout.NewSpacer(),
		), // left
		nil, welcome)
	return widget.NewSimpleRenderer(content)
}
