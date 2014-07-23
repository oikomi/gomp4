package file


type FileSpec struct {
	mp4Name string
	ftyp string
	
}

func NewFileSpec (name string) *FileSpec {
	fs := &FileSpec {
		mp4Name : name,
	}
	
	return fs
}