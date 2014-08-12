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

type SegStblAtom struct {
	Header []byte
	Stsd []byte
	Stts []byte
	Stss []byte
	Ctts []byte
	Stsc []byte
	Stsz []byte
	Stco []byte
}
	
/*	
type SegDinfAtom struct {
	Header []byte
}
*/

type SegMinfAtom struct {
	Header []byte
	Vmhd []byte
	Dinf []byte
	Stbl SegStblAtom
}

type SegMdiaAtom struct {
	Header []byte
	Mdhd []byte
	Hdlr []byte
	Minf SegMinfAtom
}

type SegTrakAtom struct {
	Header []byte
	Tkhd []byte
	Mdia SegMdiaAtom
}

type SegMoovAtom struct {
	Header []byte
	Mvhd []byte
	Trak [2]SegTrakAtom
}

type SegMp4Header struct {
	start uint32
	end uint32
	startSample uint32
	endSample uint32
	syncSampleNum uint32
	sttsEntryNum uint32
	Ftyp []byte
	Moov SegMoovAtom

}


func NewSegMp4Header() *SegMp4Header {
	return &SegMp4Header{
		
	}
}

func (self * SegMp4Header) Cover(fs *Mp4FileSpec, trakNum int)   {
	self.Ftyp = fs.FtypAtomInstance.AllBytes
	//log.Println(self.Ftyp)
	self.Moov.Header = fs.MoovAtomInstance.AllBytes
	self.Moov.Trak[trakNum].Header = fs.MoovAtomInstance.
		TrakAtomInstance[trakNum].AllBytes
	self.Moov.Trak[trakNum].Mdia.Header = fs.MoovAtomInstance.
		TrakAtomInstance[trakNum].MdiaAtomInstance.AllBytes
	self.Moov.Trak[trakNum].Mdia.Mdhd = fs.MoovAtomInstance.
		TrakAtomInstance[trakNum].MdiaAtomInstance.MdhdAtomInstance.AllBytes
	self.Moov.Trak[trakNum].Mdia.Hdlr = fs.MoovAtomInstance.
		TrakAtomInstance[trakNum].MdiaAtomInstance.HdlrAtomInstance.AllBytes
	self.Moov.Trak[trakNum].Mdia.Minf.Header = fs.MoovAtomInstance.
		TrakAtomInstance[trakNum].MdiaAtomInstance.MinfAtomInstance.AllBytes
	self.Moov.Trak[trakNum].Mdia.Minf.Vmhd = fs.MoovAtomInstance.
		TrakAtomInstance[trakNum].MdiaAtomInstance.MinfAtomInstance.
		VmhdAtomInstance.AllBytes
	self.Moov.Trak[trakNum].Mdia.Minf.Dinf = fs.MoovAtomInstance.
		TrakAtomInstance[trakNum].MdiaAtomInstance.MinfAtomInstance.
		DinfAtomInstance.AllBytes
	self.Moov.Trak[trakNum].Mdia.Minf.Stbl.Header = fs.MoovAtomInstance.
		TrakAtomInstance[trakNum].MdiaAtomInstance.MinfAtomInstance.
		StblAtomInstance.AllBytes
	self.Moov.Trak[trakNum].Mdia.Minf.Stbl.Stsd = fs.MoovAtomInstance.
		TrakAtomInstance[trakNum].MdiaAtomInstance.MinfAtomInstance.
		StblAtomInstance.StsdAtomAtomInstance.AllBytes
}


func (self * SegMp4Header)parsePara(fs *Mp4FileSpec, start uint32, end uint32 , trakNum int) {
	self.start = start
	self.end = end
	timeScale := fs.MoovAtomInstance.TrakAtomInstance[trakNum].MdiaAtomInstance.
		MdhdAtomInstance.Timescale
	entriesNum := fs.MoovAtomInstance.TrakAtomInstance[trakNum].MdiaAtomInstance.
		MinfAtomInstance.StblAtomInstance.SttsAtomAtomInstance.EntriesNum
	//log.Println(timeScale)
	startTime := timeScale * start
	endTime := timeScale * end
	//log.Println(startTime)
	//log.Println(endTime)
	var i uint32
	for i = 0; i < entriesNum; i++ {
		count := fs.MoovAtomInstance.TrakAtomInstance[i].MdiaAtomInstance.
			MinfAtomInstance.StblAtomInstance.SttsAtomAtomInstance.SampleCountDurationTable[i][0]
		duration := fs.MoovAtomInstance.TrakAtomInstance[i].MdiaAtomInstance.
			MinfAtomInstance.StblAtomInstance.SttsAtomAtomInstance.SampleCountDurationTable[i][1]
		
		if (startTime < count * duration) {
			self.startSample += (startTime / duration)
			//log.Println(self.startSample)
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
			//log.Println(self.endSample)
			break
		}
		
		self.endSample += count
		endTime -= count * duration
	}
	
}

func (self * SegMp4Header)updateStsz(fs *Mp4FileSpec, trakNum int) {
	
		
}

func (self * SegMp4Header)updateStsc(fs *Mp4FileSpec, trakNum int) {
	
		
}

func (self * SegMp4Header)updateCtts(fs *Mp4FileSpec, trakNum int) {
	
		
}

func (self * SegMp4Header)updateStss(fs *Mp4FileSpec, trakNum int) {
	self.Moov.Trak[trakNum].Mdia.Minf.Stbl.Stss = fs.MoovAtomInstance.TrakAtomInstance[trakNum].
		MdiaAtomInstance.MinfAtomInstance.StblAtomInstance.StssAtomAtomInstance.AllBytes
	entriesNum := fs.MoovAtomInstance.TrakAtomInstance[trakNum].MdiaAtomInstance.
		MinfAtomInstance.StblAtomInstance.StssAtomAtomInstance.EntriesNum
		
	//log.Println(fs.MoovAtomInstance.TrakAtomInstance[trakNum].
		//MdiaAtomInstance.MinfAtomInstance.StblAtomInstance.StssAtomAtomInstance.AllBytes)
	var i uint32
	for i = 0; i < entriesNum; i++ {
		sample := fs.MoovAtomInstance.TrakAtomInstance[trakNum].MdiaAtomInstance.
			MinfAtomInstance.StblAtomInstance.StssAtomAtomInstance.SyncSampleTable[i]
		if sample > self.startSample {
			self.startSample = fs.MoovAtomInstance.TrakAtomInstance[trakNum].MdiaAtomInstance.
			MinfAtomInstance.StblAtomInstance.StssAtomAtomInstance.SyncSampleTable[i - 1]
			break
		}
	}
	
	for i = 0; i < entriesNum; i++ {
		sample := fs.MoovAtomInstance.TrakAtomInstance[trakNum].MdiaAtomInstance.
			MinfAtomInstance.StblAtomInstance.StssAtomAtomInstance.SyncSampleTable[i]

		if sample > self.endSample {
			self.endSample = fs.MoovAtomInstance.TrakAtomInstance[trakNum].MdiaAtomInstance.
			MinfAtomInstance.StblAtomInstance.StssAtomAtomInstance.SyncSampleTable[i - 1]
			break
		}
	}
	
	var k uint32	
	k = 0
	for i = 0; i < entriesNum; i++ {
		sample := fs.MoovAtomInstance.TrakAtomInstance[trakNum].MdiaAtomInstance.
			MinfAtomInstance.StblAtomInstance.StssAtomAtomInstance.SyncSampleTable[i]
		if sample >= self.startSample && sample < self.endSample {
			s := 16 + k * 4
			copy(self.Moov.Trak[trakNum].Mdia.Minf.Stbl.Stss[s:s+4], 
				util.Uint32ToBytes(sample - self.startSample))
			self.syncSampleNum = k
			k ++
		}
	}
	
	copy(self.Moov.Trak[trakNum].Mdia.Minf.Stbl.Stss[0:4], 
		util.Uint32ToBytes(16+(self.syncSampleNum+1)*4))	
}

func (self * SegMp4Header)updateStts(fs *Mp4FileSpec, trakNum int) {
	timeScale := fs.MoovAtomInstance.TrakAtomInstance[trakNum].
		MdiaAtomInstance.MdhdAtomInstance.Timescale
	self.Moov.Trak[trakNum].Mdia.Minf.Stbl.Stts = fs.MoovAtomInstance.TrakAtomInstance[trakNum].
		MdiaAtomInstance.MinfAtomInstance.StblAtomInstance.SttsAtomAtomInstance.AllBytes
	entriesNum := fs.MoovAtomInstance.TrakAtomInstance[trakNum].MdiaAtomInstance.
		MinfAtomInstance.StblAtomInstance.SttsAtomAtomInstance.EntriesNum
		
	startTime := timeScale * self.start
	endTime := timeScale * self.end
		
	var i uint32
	var rest uint32
	var count uint32
	for i = 0; i < entriesNum; i++ {
		count = fs.MoovAtomInstance.TrakAtomInstance[i].MdiaAtomInstance.
			MinfAtomInstance.StblAtomInstance.SttsAtomAtomInstance.SampleCountDurationTable[i][0]
		duration := fs.MoovAtomInstance.TrakAtomInstance[i].MdiaAtomInstance.
			MinfAtomInstance.StblAtomInstance.SttsAtomAtomInstance.SampleCountDurationTable[i][1]

		if (startTime < count * duration) {
			rest = startTime / duration
			s := 16 + i * 8
			copy(self.Moov.Trak[trakNum].Mdia.Minf.Stbl.Stts[s:s+4], 
				util.Uint32ToBytes(count - rest))
			break
		}
		
		startTime -= count * duration
	}	
	
	for i = 0; i < entriesNum; i++ {
		count = fs.MoovAtomInstance.TrakAtomInstance[i].MdiaAtomInstance.
			MinfAtomInstance.StblAtomInstance.SttsAtomAtomInstance.SampleCountDurationTable[i][0]
		duration := fs.MoovAtomInstance.TrakAtomInstance[i].MdiaAtomInstance.
			MinfAtomInstance.StblAtomInstance.SttsAtomAtomInstance.SampleCountDurationTable[i][1]
		
		if (endTime < count * duration) {
			//rest = (endTime - timeScale * self.start) / duration
			rest = self.endSample - self.startSample
			s := 16 + i * 8
			copy(self.Moov.Trak[trakNum].Mdia.Minf.Stbl.Stts[s:s+4], 
				util.Uint32ToBytes(rest))
			self.sttsEntryNum = i
			break
		}
		
		endTime -= count * duration
	}
	
	copy(self.Moov.Trak[trakNum].Mdia.Minf.Stbl.Stts[0:4], 
			util.Uint32ToBytes(16+(self.sttsEntryNum+1)*8))	
}

func (self * SegMp4Header)updateMdhd(fs *Mp4FileSpec, trakNum int) {
	self.Moov.Trak[trakNum].Mdia.Mdhd = fs.MoovAtomInstance.TrakAtomInstance[trakNum].
		MdiaAtomInstance.MdhdAtomInstance.AllBytes
	timeScale := fs.MoovAtomInstance.TrakAtomInstance[trakNum].
		MdiaAtomInstance.MdhdAtomInstance.Timescale
	copy(self.Moov.Trak[trakNum].Mdia.Mdhd[24:28], 
		util.Uint32ToBytes((self.end - self.start) * timeScale))
}

func (self * SegMp4Header)updateTkhd(fs *Mp4FileSpec, trakNum int) {
	self.Moov.Trak[trakNum].Tkhd = fs.MoovAtomInstance.TrakAtomInstance[trakNum].
		TkhdAtomInstance.AllBytes
	timeScale := fs.MoovAtomInstance.MvhdAtomInstance.Timescale
	copy(self.Moov.Trak[trakNum].Tkhd[28:32], 
		util.Uint32ToBytes((self.end - self.start) * timeScale))
}

func (self * SegMp4Header)updateMvhd(fs *Mp4FileSpec, trakNum int) {
	self.Moov.Mvhd = fs.MoovAtomInstance.MvhdAtomInstance.AllBytes
	timeScale := fs.MoovAtomInstance.MvhdAtomInstance.Timescale
	copy(self.Moov.Mvhd[24:28], util.Uint32ToBytes((self.end - self.start) * 
		timeScale))
}

func (self * SegMp4Header)updateAtom(fs *Mp4FileSpec, trakNum int) {
	self.updateMvhd(fs, trakNum)
	self.updateTkhd(fs, trakNum)
	self.updateMdhd(fs, trakNum)
	self.updateStss(fs, trakNum)
	self.updateStts(fs, trakNum)
	self.updateCtts(fs, trakNum)
	self.updateStsc(fs, trakNum)
	self.updateStsz(fs, trakNum)
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
	
	self.Cover(fs, 0)
	
	self.updateAtom(fs, 0)
	
	fout.Write(self.Ftyp)
	fout.Write(self.Moov.Header)
	fout.Write(self.Moov.Mvhd)
	fout.Write(self.Moov.Trak[0].Header)
	fout.Write(self.Moov.Trak[0].Tkhd)
	fout.Write(self.Moov.Trak[0].Mdia.Header)
	fout.Write(self.Moov.Trak[0].Mdia.Mdhd)
	fout.Write(self.Moov.Trak[0].Mdia.Hdlr)
	fout.Write(self.Moov.Trak[0].Mdia.Minf.Header)
	fout.Write(self.Moov.Trak[0].Mdia.Minf.Vmhd)
	fout.Write(self.Moov.Trak[0].Mdia.Minf.Dinf)
	fout.Write(self.Moov.Trak[0].Mdia.Minf.Stbl.Header)
	fout.Write(self.Moov.Trak[0].Mdia.Minf.Stbl.Stsd)
	fout.Write(self.Moov.Trak[0].Mdia.Minf.Stbl.Stts[0: 16 + (self.sttsEntryNum+1) * 8])
	fout.Write(self.Moov.Trak[0].Mdia.Minf.Stbl.Stss[0: 16 + (self.syncSampleNum+1) * 4])

	
	return nil

}

type SegMp4Data struct {
	
}

