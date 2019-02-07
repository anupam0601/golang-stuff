package filescreatormicro

type FilesCreatorService interface {
	Word(param string) string
}

type BasicFilesCreator struct{}

// Implement
func (BasicFilesCreator) Word(param string) string {
	return param
}
