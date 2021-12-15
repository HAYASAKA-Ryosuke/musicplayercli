package musiclist

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
)

type DirectoryWalk struct{ BasePath string }

type MusicInfo struct {
	AlbumPath string
	AlbumName string
	FilePath  string
	FileName  string
}

func (directoryWalk *DirectoryWalk) GetMusicList() []MusicInfo {
	result := []MusicInfo{}
	filepath.Walk(directoryWalk.BasePath, func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() && info.Name() == "skip" {
			return filepath.SkipDir
		}
		if !info.IsDir() && strings.Contains(path, ".mp3") && !strings.HasPrefix(info.Name(), ".") {
			albumPath := filepath.Dir(path)
			albumName := strings.Split(albumPath, "/")[len(strings.Split(albumPath, "/"))-1]
			dirName, baseName := filepath.Split(path)
			result = append(result, MusicInfo{AlbumPath: dirName, AlbumName: albumName, FilePath: path, FileName: baseName})
		}
		return nil
	})
	return result
}
func (directoryWalk *DirectoryWalk) GetfilePath(directoryPath string) {
	fmt.Println("getfilePath")
}
func (directoryWalk *DirectoryWalk) GetDirectoryPaths() {
	fmt.Println("getDirectoryPaths")
}
