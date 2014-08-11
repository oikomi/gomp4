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
	"os"
	"github.com/oikomi/gomp4/util"
)

type SegMp4Header struct {
	startSample uint32
	endSample uint32
	Ftyp []byte
	Moov []byte
	Mvhd []byte
}

func NewSegMp4Header() *SegMp4Header {
	return &SegMp4Header{
		
	}
}

func (self * SegMp4Header) Cover(fs *Mp4FileSpec)   {
	self.Ftyp = fs.FtypAtomInstance.AllBytes
	//log.Println(self.Ftyp)
	self.Moov = fs.MoovAtomInstance.AllBytes
	//self.Mvhd = 
	//log.Println(self.Moov)
}


func (self * SegMp4Header)parsePara(fs *Mp4FileSpec, start uint32, end uint32 , trakNum int) {
	timeScale := fs.MoovAtomInstance.TrakAtomInstance[trakNum].MdiaAtomInstance.
		MdhdAtomInstance.Timescale
	entriesNum := fs.MoovAtomInstance.TrakAtomInstance[trakNum].MdiaAtomInstance.
		MinfAtomInstance.StblAtomInstance.SttsAtomAtomInstance.EntriesNum
	log.Println(timeScale)
	startTime := timeScale * start
	endTime := timeScale * end
	log.Println(startTime)
	log.Println(endTime)
	var i uint32
	for i = 0; i < entriesNum; i++ {
		count := fs.MoovAtomInstance.TrakAtomInstance[i].MdiaAtomInstance.
			MinfAtomInstance.StblAtomInstance.SttsAtomAtomInstance.SampleCountDurationTable[i][0]
		duration := fs.MoovAtomInstance.TrakAtomInstance[i].MdiaAtomInstance.
			MinfAtomInstance.StblAtomInstance.SttsAtomAtomInstance.SampleCountDurationTable[i][1]
		
		if (startTime < count * duration) {
			self.startSample += (startTime / duration)
			log.Println(self.startSample)
			break
		}
		
		self.startSample += count
		startTime -= count * duration
	}
	
	for i = 0; i < entriesNum; i++ {
		count := fs.MoovAtomInstance.TrakAtomInstance[i].MdiaAtomInstance.
			MinfAtomInstance.StblAtomInstance.SttsAtomAtomInstance.SampleCountDurationTable[i][0]
		duration := fs.MoovAtomInstance.TrakAtomInstance[i].MdiaAtomInstance.
			MinfAtomInstance.StblAtomInstance.SttsAtomAtomInstance.SampleCountDurationTable[i][1]
		
		if (endTime < count * duration) {
			self.endSample += (endTime / duration)
			log.Println(self.endSample)
			break
		}
		
		self.endSample += count
		endTime -= count * duration
	}
	
}

func (self * SegMp4Header)updateMvhd(fs *Mp4FileSpec, trakNum int) {
	self.Mvhd = fs.MoovAtomInstance.MvhdAtomInstance.AllBytes
	//log.Printf()
	timeScale := fs.MoovAtomInstance.TrakAtomInstance[trakNum].MdiaAtomInstance.
		MdhdAtomInstance.Timescale
	log.Println((self.endSample - self.startSample) * timeScale)
	log.Println(self.Mvhd[24:28])
	copy(self.Mvhd[24:28], util.Uint32ToBytes((self.endSample - self.startSample) * 
		timeScale))
}

func (self * SegMp4Header)updateAtom(fs *Mp4FileSpec, trakNum int) {
	self.updateMvhd(fs, trakNum)
}

func (self * SegMp4Header)WriteSegMp4(fs *Mp4FileSpec, start uint32, end uint32) error {
	segMp4File := "seg.mp4"
	fout, err := os.Create(segMp4File)
	defer fout.Close()
	if err != nil {
	    log.Fatalln(err.Error())
	    return err 
	}

	self.parsePara(fs, start, end, 0)
	
	self.Cover(fs)
	
	self.updateAtom(fs, 0)
	
	fout.Write(self.Ftyp)
	fout.Write(self.Moov)
	fout.Write(self.Mvhd)
	
	return nil

}

type SegMp4Data struct {
	
}

