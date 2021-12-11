package main

import (
	"fmt"
	"musicplayergo/mplayer"
	"musicplayergo/musiclist"
	"time"

	"github.com/rivo/tview"
)

func addMusic(root *tview.TreeNode, tree *tview.TreeView, mPlayer *mplayer.MPlayer, title string) {
	node := tview.NewTreeNode(title).SetReference(title)
	root.AddChild(node)
}

func main() {
	directorywalk := musiclist.DirectoryWalk{BasePath: "<path>"}
	musicList := musiclist.New(&directorywalk)

	app := tview.NewApplication()
	mPlayer := mplayer.New()
	root := tview.NewTreeNode(".")
	tree := tview.NewTreeView().SetRoot(root).SetCurrentNode(root)
	root.AddChild(tview.NewTreeNode("quit").SetReference("QUIT"))
	tree.SetSelectedFunc(func(node *tview.TreeNode) {
		reference := node.GetReference()
		if reference == nil {
			return
		}
		if reference.(string) == "QUIT" {
			mPlayer.Quit()
			app.Stop()
		}
		mPlayer.LoadFile(fmt.Sprintf("'%s'", reference))
	})
	//AddItem("Quit", "Press to exit", ' ', func() {
	//	mPlayer.Quit()
	//	app.Stop()
	//})
	box := tview.NewBox().SetBorder(true).SetTitle("current")

	for _, path := range musicList.GetfilePaths() {
		addMusic(root, tree, mPlayer, path)
	}
	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(box, 4, 0, true).  // widget1は常に4行固定で表示する
		AddItem(tree, 0, 3, false) // widget2は残りの領域の3/4で表示する

	page := tview.NewPages()
	page.AddPage("page1", flex, true, true) // gridをpageに追加するが、表示させない
	if err := app.SetRoot(flex, true).SetFocus(tree).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
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
