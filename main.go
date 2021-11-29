package main

import (
	"fmt"
	"musicplayergo/mplayer"
	"musicplayergo/musiclist"
	"time"

	"github.com/rivo/tview"
)

func addMusic(list *tview.List, mPlayer *mplayer.MPlayer, title string) {
	list.AddItem(title, "", ' ', func() {
		mPlayer.LoadFile(title)
	})
}

func main() {
	musicList := musiclist.New(new(musiclist.DirectoryWalk))
	musicList.GetfilePaths()
	app := tview.NewApplication()
	mPlayer := mplayer.New()
	list := tview.NewList().
		AddItem("Quit", "Press to exit", ' ', func() {
			mPlayer.Quit()
			app.Stop()
		})

	addMusic(list, mPlayer, "test.mp3")
	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(tview.NewBox().SetBorder(true).SetTitle("current"), 4, 0, true). // widget1は常に4行固定で表示する
		AddItem(list, 0, 3, false)                                               // widget2は残りの領域の3/4で表示する

	page := tview.NewPages()
	page.AddPage("page1", flex, true, true) // gridをpageに追加するが、表示させない
	if err := app.SetRoot(flex, true).SetFocus(list).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
	fmt.Println("Hello")
	time.Sleep(1 * time.Second)
	app.Sync()
	time.Sleep(1 * time.Second)
	mPlayer.Quit()
	//for {
	//	select {
	//	case <-ctx.Done():
	//		log.Debug("context done")
	//		return
	//	}
	//}
}
