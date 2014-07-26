//
// Copyright 2014 Hong Miao. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"os"
	"log"
	//"fmt"
	"./mp4"
	"./output"
	
)

func main() {
	if len(os.Args) != 2 {
		os.Exit(0)
	}	
	
	fs := mp4.NewMp4FileSpec(os.Args[1])
	fp := mp4.NewMp4FilePro()
	
	err := fp.Mp4Open(fs)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
	
	err = fp.Mp4FileStat(fs)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
		
	fs.ParseAtoms(fp)
	/*
	size, box, _ := fp.Mp4Read()
	log.Println(size, box)
	sizeInt := util.Bytes2Int(size)
	fp.Mp4Seek(sizeInt - (int64)(len(size)) - (int64)(len(box)))
	
	size, box, _ = fp.Mp4Read()
	
	log.Println(size, box)
	*/
	log.Println("------")
	output.JsonDump(fs)
	log.Println(fs)
	log.Println("------")
	
}