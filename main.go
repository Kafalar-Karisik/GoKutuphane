package main

import (

	//reflect.TypeOf(item)

	//"database/sql"

	clck "GoKutuphane/clock"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	//_ "github.com/mattn/go-sqlite3"
)

var Stats struct {
	topBook   int
	oduncKit  int
	todayGib  int
	topMember int
}

func main() {

	//Windows Set
	mapp := app.NewWithID("GoKutuphane")
	mwin := mapp.NewWindow("GoKutuphane")

	statsL := widget.NewLabel("")
	go startInfos(statsL, &Stats)

	clock := clck.ClockLayout{}
	clock.Render().Resize(fyne.NewSize(125, 125))
	clcok := clock.Render()
	go clock.Animate(clcok)

	strClock := widget.NewLabel("")
	go startClock(strClock)

	date := widget.NewLabel("")
	go startDate(date)

	/* //Menu
	menu1 := fyne.NewMenu("File")

	menu2 := fyne.NewMenu("View",
		fyne.NewMenuItem("Full Screen",
			func() {
				if mwin.FullScreen() == false {
					mwin.SetFullScreen(true)
				} else {
					mwin.SetFullScreen(false)
				}
			},
		),
	)

	menu := fyne.NewMainMenu(menu1, menu2)
	mwin.SetMainMenu(menu)
	*/

	anas := container.NewBorder(
		widget.NewLabel("Hoş Geldiniz"),
		container.NewHBox(date),
		statsL,
		container.New(layout.NewVBoxLayout(), clcok, strClock),
	)

	container.NewWithoutLayout(
		widget.NewLabel("ID"),
		widget.NewEntry(),
		widget.NewLabel("\n\n"),
	)

	is := container.NewWithoutLayout(
		widget.NewLabel(""),
	)

	uyes := container.NewWithoutLayout(
		widget.NewLabel("Üye Düzen"),
	)

	kitaps := container.NewWithoutLayout(
		widget.NewLabel("Kitap Düzen"),
	)

	ayars := container.NewWithoutLayout(
		widget.NewLabel("Ayarlar..."),
	)

	logs := container.NewWithoutLayout(
		widget.NewLabel(""), //Will be Logs (TABLE)
	)

	main := container.NewAppTabs(
		container.NewTabItemWithIcon("Ana Sayfa", theme.HomeIcon(), anas),
		container.NewTabItemWithIcon("Standart İşlemler", theme.GridIcon(), is),
		container.NewTabItemWithIcon("Üye Düzenleme", theme.AccountIcon(), uyes),
		container.NewTabItemWithIcon("Kitap Düzenleme", theme.StorageIcon(), kitaps),
		container.NewTabItemWithIcon("Ayarlar", theme.SettingsIcon(), ayars),
		container.NewTabItemWithIcon("Geçmiş", theme.HistoryIcon(), logs),
	)

	//Window Run
	mwin.SetContent(main)
	mwin.ShowAndRun()
}

func startDate(date *widget.Label) {
	format := "02.01.2006"
	date.SetText(time.Now().Format(format))

	for range time.Tick(time.Minute) {
		date.SetText(
			time.Now().Format(format))
	}
}

func startClock(clock *widget.Label) {
	format := " 03:04:05"
	clock.SetText(time.Now().Format(format))

	for range time.Tick(time.Second) {
		clock.SetText(
			time.Now().Format(format))
	}
}

func startInfos(labe *widget.Label, stat *struct {
	topBook   int
	oduncKit  int
	todayGib  int
	topMember int
}) {
	if false {
		print(labe)
	}
}
