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

package mp4

import (
	"log"
	//"errors"
	"github.com/oikomi/gomp4/util"
)

type Sample2Chunk []uint32

type StscAtom struct {
	Offset int64
	Size int64
	IsFullBox bool
	Version uint8
	Flag uint32
	EntriesNum uint32
	Sample2ChunkTable []Sample2Chunk
	//Sample2ChunkTable map[uint32]Sample2Chunk

	AllBytes []byte
}

func stscRead(fs *Mp4FileSpec, fp *Mp4FilePro, offset int64) error {
	var err error
	fs.MoovAtomInstance.TrakAtomInstance[trakNum].MdiaAtomInstance.MinfAtomInstance.
		StblAtomInstance.StscAtomAtomInstance.Offset = offset
	fs.MoovAtomInstance.TrakAtomInstance[trakNum].MdiaAtomInstance.MinfAtomInstance.
		StblAtomInstance.StscAtomAtomInstance.IsFullBox = false
	err = fp.Mp4Seek(offset, 0)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	
	size, _, err := fp.Mp4ReadHeader()
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	
	sizeInt := util.Bytes2Int(size)	
	fs.MoovAtomInstance.TrakAtomInstance[trakNum].MdiaAtomInstance.MinfAtomInstance.
		StblAtomInstance.StscAtomAtomInstance.Size = sizeInt
		
	err = fp.Mp4Seek(offset, 0)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	
	buf, err := fp.Mp4Read(fs.MoovAtomInstance.TrakAtomInstance[trakNum].MdiaAtomInstance.MinfAtomInstance.
		StblAtomInstance.StscAtomAtomInstance.Size)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	
	fs.MoovAtomInstance.TrakAtomInstance[trakNum].MdiaAtomInstance.MinfAtomInstance.
		StblAtomInstance.StscAtomAtomInstance.AllBytes = buf

	err = fp.Mp4Seek(offset + 8, 0)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}	

	size, err = fp.Mp4Read(1)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	fs.MoovAtomInstance.TrakAtomInstance[trakNum].MdiaAtomInstance.MinfAtomInstance.
		StblAtomInstance.StscAtomAtomInstance.Version = uint8(size[0])
	
	size, err = fp.Mp4Read(3)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	fs.MoovAtomInstance.TrakAtomInstance[trakNum].MdiaAtomInstance.MinfAtomInstance.
		StblAtomInstance.StscAtomAtomInstance.Flag = util.Byte32Uint32(size, 0)	

	size, err = fp.Mp4Read(4)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	fs.MoovAtomInstance.TrakAtomInstance[trakNum].MdiaAtomInstance.MinfAtomInstance.
		StblAtomInstance.StscAtomAtomInstance.EntriesNum = util.Byte42Uint32(size, 0)
		
	//log.Println(fs.MoovAtomInstance.TrakAtomInstance[trakNum].MdiaAtomInstance.MinfAtomInstance.
		//StblAtomInstance.StscAtomAtomInstance.EntriesNum)
		
	fs.MoovAtomInstance.TrakAtomInstance[trakNum].MdiaAtomInstance.MinfAtomInstance.
		StblAtomInstance.StscAtomAtomInstance.Sample2ChunkTable = 
		make([]Sample2Chunk, fs.MoovAtomInstance.TrakAtomInstance[trakNum].MdiaAtomInstance.MinfAtomInstance.
		StblAtomInstance.StscAtomAtomInstance.EntriesNum)
		
	
		
	var i uint32
	for i = 0; i < fs.MoovAtomInstance.TrakAtomInstance[trakNum].MdiaAtomInstance.MinfAtomInstance.
		StblAtomInstance.StscAtomAtomInstance.EntriesNum; i++ {	
		
		buf, err := fp.Mp4Read(12)
		if err != nil {
			log.Fatalln(err.Error())
			return err
		}
		
		
		fs.MoovAtomInstance.TrakAtomInstance[trakNum].MdiaAtomInstance.MinfAtomInstance.
			StblAtomInstance.StscAtomAtomInstance.Sample2ChunkTable[i] = append(fs.MoovAtomInstance.
			TrakAtomInstance[trakNum].MdiaAtomInstance.MinfAtomInstance.StblAtomInstance.
			StscAtomAtomInstance.Sample2ChunkTable[i], util.Byte42Uint32(buf[0:4], 0))
			
		fs.MoovAtomInstance.TrakAtomInstance[trakNum].MdiaAtomInstance.MinfAtomInstance.
			StblAtomInstance.StscAtomAtomInstance.Sample2ChunkTable[i] = append(fs.MoovAtomInstance.
			TrakAtomInstance[trakNum].MdiaAtomInstance.MinfAtomInstance.StblAtomInstance.
			StscAtomAtomInstance.Sample2ChunkTable[i], util.Byte42Uint32(buf[4:8], 0))
			
		fs.MoovAtomInstance.TrakAtomInstance[trakNum].MdiaAtomInstance.MinfAtomInstance.
			StblAtomInstance.StscAtomAtomInstance.Sample2ChunkTable[i] = append(fs.MoovAtomInstance.
			TrakAtomInstance[trakNum].MdiaAtomInstance.MinfAtomInstance.StblAtomInstance.
			StscAtomAtomInstance.Sample2ChunkTable[i], util.Byte42Uint32(buf[8:12], 0))
		
		/*
		fs.MoovAtomInstance.TrakAtomInstance[trakNum].MdiaAtomInstance.MinfAtomInstance.
			StblAtomInstance.StscAtomAtomInstance.Sample2ChunkTable = append(fs.MoovAtomInstance.
			TrakAtomInstance[trakNum].MdiaAtomInstance.MinfAtomInstance.StblAtomInstance.
			StscAtomAtomInstance.Sample2ChunkTable, fs.MoovAtomInstance.TrakAtomInstance[trakNum].MdiaAtomInstance.MinfAtomInstance.
			StblAtomInstance.StscAtomAtomInstance.Sample2ChunkTable[i])
		*/
	}		
		
	return nil
	
}