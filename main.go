package main

import (
	"fmt"
	"musicplayergo/mplayer"
	"musicplayergo/musiclist"
	"time"

	"github.com/rivo/tview"
)

func addMusic(root *tview.TreeNode, tree *tview.TreeView, mPlayer *mplayer.MPlayer, albumInfo musiclist.AlbumInfo) {
	albumNode := root.AddChild(tview.NewTreeNode(" " + albumInfo.AlbumName).SetReference(albumInfo))
	//musicNodeList := []*tview.TreeNode{}
	for _, musicInfo := range albumInfo.MusicInfo {
		tview.NewTreeNode(albumInfo.AlbumName).SetReference(albumInfo)
		//musicNodeList = append(musicNodeList, tview.NewTreeNode(musicInfo.MusicName).SetReference(musicInfo))
		albumNode.AddChild(tview.NewTreeNode(" " + musicInfo.MusicName).SetReference(musicInfo))
	}
	//albumNode.AddChild(musicNodeList)
}

func update(info *tview.TextView) {
	tick := time.NewTicker(500 * time.Millisecond)
	for {
		select {
		case <-tick.C:
			//fmt.Fprintf(info, "%s ", reference.(musiclist.MusicInfo).AlbumName+" - "+reference.(musiclist.MusicInfo).FileName)
			fmt.Println("")
		}
	}
}

func main() {
	musicList := musiclist.DirectoryWalk{BasePath: "<PATH>"}

	app := tview.NewApplication()
	mPlayer := mplayer.New()
	root := tview.NewTreeNode(".")
	tree := tview.NewTreeView().SetRoot(root).SetCurrentNode(root)
	root.AddChild(tview.NewTreeNode("quit").SetReference("QUIT"))
	info := tview.NewTextView()
	tree.SetSelectedFunc(func(node *tview.TreeNode) {
		reference := node.GetReference()
		if reference == nil {
			return
		}
		mPlayer.LoadFile(fmt.Sprintf("'%s'", reference.(musiclist.AlbumInfo).MusicInfo[0].MusicPath))
		info.Clear()
		fmt.Fprintf(info, "%s ", reference.(musiclist.AlbumInfo).AlbumName+" - "+reference.(musiclist.AlbumInfo).MusicInfo[0].MusicName)
	})

	for _, albumInfo := range musicList.GetMusicList() {
		addMusic(root, tree, mPlayer, albumInfo)
	}
	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(info, 4, 0, true).
		AddItem(tree, 0, 3, false)

	//go update(info)

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
