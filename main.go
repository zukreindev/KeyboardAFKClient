package main

import (
	"fmt"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/go-vgo/robotgo"
	keylogger2 "github.com/kindlyfire/go-keylogger"
)

var (
	online = false
	msg    = "OFFLINE"
)

// main
func main() {
	go loopW()
	a := app.New()
	w := a.NewWindow("Valorant AFK Client")

	hello := widget.NewLabel("Gördüğüm Kadarıyla AFK kalman gerekiyor neyse ki senin için buradayım!! \nUygulamayı kapatmak için uygulamayı kapat butonuna veya H tuşuna basabilirsin.\n\n Status: " + msg)
	afkButton := widget.NewButton("AFK", func() {

		if online == true {
			online = false
			fmt.Println("Valorant AFK Bot is offline")
			msg = "OFFLINE"
			hello.SetText("Gördüğüm Kadarıyla AFK kalman gerekiyor neyse ki senin için buradayım!! \nUygulamayı kapatmak için uygulamayı kapat butonuna veya H tuşuna basabilirsin.\n\n Status: " + msg)
		} else {
			online = true
			fmt.Println("Valorant AFK Bot is online")
			msg = "ONLINE"
			hello.SetText("Gördüğüm Kadarıyla AFK kalman gerekiyor neyse ki senin için buradayım!! \nUygulamayı kapatmak için uygulamayı kapat butonuna veya H tuşuna basabilirsin.\n\n Status: " + msg)
		}

	})

	closeButton := widget.NewButton("Uygulama Kapat", func() {
		w.Close()

	})

	w.SetContent(container.NewVBox(
		hello,
		afkButton,
		closeButton,
	))

	w.ShowAndRun()

}

func loopW() {

	fmt.Println("Valorant AFK Thread is running")
	for {
		keylogger := keylogger2.NewKeylogger()

		key := keylogger.GetKey()
		if !key.Empty {
			if key.Rune == 'h' {
				online = false
				msg = "OFFLINE"

			}
		}

		if online == true {
			robotgo.KeyDown("w")
			robotgo.MilliSleep(100)
		}
	}
}
