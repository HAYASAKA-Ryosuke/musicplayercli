package musiclist

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
)

type DirectoryWalk struct{ BasePath string }

type AlbumInfo struct {
	AlbumPath string
	AlbumName string
	MusicInfo []MusicInfo
}

type MusicInfo struct {
	MusicPath string
	MusicName string
}

func (directoryWalk *DirectoryWalk) GetMusicList() []AlbumInfo {
	albumList := []AlbumInfo{}
	albumPath := ""
	musicList := []MusicInfo{}
	albumName := ""
	filepath.Walk(directoryWalk.BasePath, func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() && info.Name() == "skip" {
			return filepath.SkipDir
		}
		if !info.IsDir() && strings.Contains(path, ".mp3") && !strings.HasPrefix(info.Name(), ".") {
			dirName, _ := filepath.Split(path)
			if albumPath != "" && albumPath != dirName {
				albumList = append(albumList, AlbumInfo{AlbumPath: albumPath, AlbumName: albumName, MusicInfo: musicList})
				albumPath = dirName
				albumName = filepath.Base(albumPath)
				musicList = []MusicInfo{}
			}
			if albumPath == "" {
				albumPath = dirName
				albumName = filepath.Base(albumPath)
			}
			_, baseName := filepath.Split(path)
			musicList = append(musicList, MusicInfo{MusicPath: path, MusicName: baseName})
		}
		return nil
	})
	return albumList
}
func (directoryWalk *DirectoryWalk) GetfilePath(directoryPath string) {
	fmt.Println("getfilePath")
}
func (directoryWalk *DirectoryWalk) GetDirectoryPaths() {
	fmt.Println("getDirectoryPaths")
}
