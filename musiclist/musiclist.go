package musiclist

type MusicListDI struct {
	MusicListInterface
}

type MusicListInterface interface {
	GetfilePaths() []string
	GetfilePath(directoryPath string)
	GetDirectoryPaths()
}

func New(musicListInterface MusicListInterface) *MusicListDI {
	musicListDi := &MusicListDI{musicListInterface}
	return musicListDi
}
