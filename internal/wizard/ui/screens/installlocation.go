package screens

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/dweymouth/fyne-windows-installer/wizard/res"
)

type InstallLocationScreen struct {
	widget.BaseWidget

	location *string

	content *fyne.Container
}

func NewInstallLocationScreen(location *string) *InstallLocationScreen {
	i := &InstallLocationScreen{
		location: location,
	}
	i.ExtendBaseWidget(i)
	return i
}

func (i *InstallLocationScreen) CreateRenderer() fyne.WidgetRenderer {
	appName := string(res.ResAppName.StaticContent)
	header := widget.NewRichText(
		&widget.TextSegment{
			Text:  "Choose install location",
			Style: widget.RichTextStyle{TextStyle: fyne.TextStyle{Bold: true}},
		},
		&widget.TextSegment{
			Text: fmt.Sprintf("Setup will install %s into the following folder.", appName),
		},
		&widget.TextSegment{
			Text: "To continue, click Next. If you would like to select a different folder, click Browse.",
		},
	)
	header.Wrapping = fyne.TextWrapWord

	browseBtn := widget.NewButton("Browse", nil)
	browseBtn.Icon = theme.FolderOpenIcon()
	entry := widget.NewEntryWithData(binding.BindString(i.location))
	i.content = container.NewVBox(
		header,
		container.NewBorder(nil, nil, nil, browseBtn, entry),
	)
	return widget.NewSimpleRenderer(i.content)
}
