package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/dweymouth/fyne-windows-installer/wizard/res"
	"github.com/dweymouth/fyne-windows-installer/wizard/ui"
	"github.com/dweymouth/fyne-windows-installer/wizard/ui/screens"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello World")

	appName := string(res.ResAppName.StaticContent)
	installPath := DefaultInstallLocation(appName, false /*TODO*/)
	navigator := ui.NewNavigator(
		screens.NewWelcomeScreen(),
		screens.NewInstallLocationScreen(&installPath),
	)
	w.SetContent(navigator)
	w.Resize(fyne.NewSize(400, 300))
	w.ShowAndRun()
}
