package main

import (
	"os"
	"log"
	"./file"
)

func main() {
	if len(os.Args) != 2 {
		os.Exit(0)
	}	
	
	fs := file.NewFileSpec(os.Args[1])
	fp := file.NewFilePro()
	
	err := fp.Mp4Open(fs)
	
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
	
	fp.Mp4Read()
	
}