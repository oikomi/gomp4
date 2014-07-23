package file

import (
	"os"
	"log"
	"bufio"
	"../util"
)

type FilePro struct {
	file *os.File
	r *bufio.Reader
	
}

func NewFilePro () *FilePro {
	fp := &FilePro {
		
	}
	
	return fp
}

func (self * FilePro) Mp4Open(fs *FileSpec)  error {
	var err error
	//log.Println(fs.mp4Name)
	self.file, err = os.Open(fs.mp4Name)
	
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	self.r = bufio.NewReader(self.file)
	
	return nil
}

func (self * FilePro) Mp4Read()  error {
	buf := make([]byte, 8)
	_, err := self.r.Read(buf)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	log.Println(string(buf[3:]))
	return nil
}