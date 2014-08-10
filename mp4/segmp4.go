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
)

type SegMp4Header struct {
	Ftyp []byte
	Moov []byte
}

func (self * SegMp4Header) FtypCover(fs *Mp4FileSpec)   {
	self.Ftyp = fs.FtypAtomInstance.AllBytes
	log.Println(self.Ftyp)
}

func (self * SegMp4Header) MoovCover(fs *Mp4FileSpec)   {
	self.Moov = fs.MoovAtomInstance.AllBytes
	log.Println(self.Moov)
}

func parsePara(fs *Mp4FileSpec, start uint64, end uint64 , trakNum int) {
	timeScale := fs.MoovAtomInstance.TrakAtomInstance[trakNum].MdiaAtomInstance.
		MdhdAtomInstance.Timescale
	entriesNum := fs.MoovAtomInstance.TrakAtomInstance[trakNum].MdiaAtomInstance.
		MinfAtomInstance.StblAtomInstance.SttsAtomAtomInstance.EntriesNum
	log.Println(timeScale)
	startTime := (uint64)(timeScale) * start
	endTime := (uint64)(timeScale) * end
	log.Println(startTime)
	log.Println(endTime)
	var startSample uint64
	var endSample uint64
	var i uint32
	for i = 0; i < entriesNum; i++ {
		count := fs.MoovAtomInstance.TrakAtomInstance[i].MdiaAtomInstance.
			MinfAtomInstance.StblAtomInstance.SttsAtomAtomInstance.SampleCountDurationTable[i][0]
		duration := fs.MoovAtomInstance.TrakAtomInstance[i].MdiaAtomInstance.
			MinfAtomInstance.StblAtomInstance.SttsAtomAtomInstance.SampleCountDurationTable[i][1]
		
		if (startTime < (uint64) (count) * (uint64) (duration)) {
			startSample += (startTime / (uint64)(duration))
			log.Println(startSample)
			break
	    }
		
		startSample += (uint64)(count)
		startTime -= (uint64)(count) * (uint64)(duration)
	}
	
	for i = 0; i < entriesNum; i++ {
		count := fs.MoovAtomInstance.TrakAtomInstance[i].MdiaAtomInstance.
			MinfAtomInstance.StblAtomInstance.SttsAtomAtomInstance.SampleCountDurationTable[i][0]
		duration := fs.MoovAtomInstance.TrakAtomInstance[i].MdiaAtomInstance.
			MinfAtomInstance.StblAtomInstance.SttsAtomAtomInstance.SampleCountDurationTable[i][1]
		
		if (endTime < (uint64) (count) * (uint64) (duration)) {
			endSample += (endTime / (uint64)(duration))
			log.Println(endSample)
			break
	    }
		
		endSample += (uint64)(count)
		endTime -= (uint64)(count) * (uint64)(duration)
	}
	
}

func WriteSegMp4(fs *Mp4FileSpec, start uint64, end uint64) error {
	segMp4File := "seg.mp4"
	fout, err := os.Create(segMp4File)
	defer fout.Close()
    if err != nil {
        log.Fatalln(err.Error())
        return err 
    }
	fout.Write(fs.FtypAtomInstance.AllBytes)
	fout.Write(fs.MoovAtomInstance.AllBytes)
	fout.Write(fs.MoovAtomInstance.MvhdAtomInstance.AllBytes)
	
	parsePara(fs, start, end, 0)
	
	return nil

}

type SegMp4Data struct {
	
}

