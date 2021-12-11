package musiclist

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
)

type DirectoryWalk struct{ BasePath string }

func (directoryWalk *DirectoryWalk) GetfilePaths() []string {
	result := []string{}
	filepath.Walk(directoryWalk.BasePath, func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() && info.Name() == "skip" {
			return filepath.SkipDir
		}
		if !info.IsDir() && strings.Contains(path, ".mp3") && !strings.HasPrefix(info.Name(), ".") {
			if strings.Contains(path, "ABMC") {
				fmt.Println(path)
			}
			result = append(result, path)
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
