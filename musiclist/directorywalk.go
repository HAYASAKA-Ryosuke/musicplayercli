package musiclist

import "fmt"

type DirectoryWalk struct{}

func (directoryWalk *DirectoryWalk) GetfilePaths() {
	fmt.Println("getfilePaths")
}
func (directoryWalk *DirectoryWalk) GetfilePath(directoryPath string) {
	fmt.Println("getfilePath")
}
func (directoryWalk *DirectoryWalk) GetDirectoryPaths() {
	fmt.Println("getDirectoryPaths")
}
