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

type ParseAtomFuc func(fs *Mp4FileSpec, fp *Mp4FilePro, offset int64) error

var (
	trakNum int
	
	mp4Atoms map[string]ParseAtomFuc
	mp4MoovAtoms map[string]ParseAtomFuc
	mp4TrakAtoms map[string]ParseAtomFuc
	mp4MdiaAtoms map[string]ParseAtomFuc
	mp4MinfAtoms map[string]ParseAtomFuc
	mp4StblAtoms map[string]ParseAtomFuc
)

func init() {
	mp4Atoms = map[string]ParseAtomFuc {
		"ftyp" : ftypRead,
		"moov" : moovRead,
		"mdat" : mdatRead,
	}
	mp4MoovAtoms = map[string]ParseAtomFuc {
		"mvhd" : mvhdRead,
		"trak" : trakRead,
	}
	mp4TrakAtoms = map[string]ParseAtomFuc {
		"tkhd" : tkhdRead,
		"mdia" : mdiaRead,
	}
	mp4MdiaAtoms = map[string]ParseAtomFuc {
		"mdhd" : mdhdRead,
		"hdlr" : hdlrRead,
		"minf" : minfRead,
	}
	mp4MinfAtoms = map[string]ParseAtomFuc {
		"smhd" : smhdRead,
		"dinf" : dinfRead,
		"stbl" : stblRead,
	}
	mp4StblAtoms = map[string]ParseAtomFuc {
		"stsd" : stsdRead,
		"stts" : sttsRead,
		"stss" : stssRead,
		"stsc" : stscRead,
		"stsz" : stszRead,
		"stco" : stcoRead,	
	}
}

type Mp4FileSpec struct {
	Mp4Name string
	TotalSize int64
	FtypAtomInstance FtypAtom
	MoovAtomInstance MoovAtom
	MdatAtomInstance MdatAtom
}

func NewMp4FileSpec (name string) *Mp4FileSpec {
	fs := &Mp4FileSpec {
		Mp4Name : name,
		TotalSize : 0,
		//ftypAtom : new(FtypAtom),
		//moovAtom : new(MoovAtom),
	}
	
	return fs
}

func (self * Mp4FileSpec) nextAtom(offset int64, fp *Mp4FilePro)  error {
	err := fp.Mp4Seek(offset, 0)
	
	return err
}

func (self *Mp4FileSpec) ParseAtoms(fp *Mp4FilePro) error {
	var pos int64
	
	for self.TotalSize > pos {
		size, atom, err := fp.Mp4ReadHeader()
		
		log.Println(size, string(atom))
		
		if err != nil {
			log.Fatalln(err.Error())
			return err
		}
		
		sizeInt := util.Bytes2Int(size)	
		//offset := sizeInt - (int64)(len(size)) - (int64)(len(atom))
		//log.Println(offset)
		pos += sizeInt
		
		if f, ok := mp4Atoms[string(atom)]; ok {
			err = f(self, fp, pos - sizeInt)
			if err != nil {
				log.Fatalln(err.Error())
				return err	
			}
		}
		self.nextAtom(pos, fp)	
	}
	
	return nil
}
