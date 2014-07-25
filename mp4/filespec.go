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
	"../util"
)

type ParseAtomFuc func(fs *Mp4FileSpec)

var (
	mp4Atoms map[string]ParseAtomFuc
	
	mp4MoovAtoms map[string]ParseAtomFuc
)

func init() {
	mp4Atoms = map[string]ParseAtomFuc {
		"ftyp": ftypRead,
		"moov": moovRead,
		"mdat": mdatRead,
	}
	mp4MoovAtoms = map[string]ParseAtomFuc {
		"mvhd": mvhdRead,
		"trak": trakRead,
	}
}

////////////////
func ftypRead(fs *Mp4FileSpec) {
	log.Println("--------")
}

func moovRead(fs *Mp4FileSpec) {
	log.Println("--------")
}

func mdatRead(fs *Mp4FileSpec) {
	log.Println("--------")
}
////////////

func mvhdRead(fs *Mp4FileSpec) {
	log.Println("--------")
}

func trakRead(fs *Mp4FileSpec) {
	log.Println("--------")
}

type Mp4FileSpec struct {
	mp4Name string
	
	end int64
	
}

func NewMp4FileSpec (name string) *Mp4FileSpec {
	fs := &Mp4FileSpec {
		mp4Name : name,
	}
	
	return fs
}




func (self * Mp4FileSpec) NextAtom(offset int64, fp *Mp4FilePro)  error {
	err := fp.Mp4Seek(offset)
	
	return err
}


func (self *Mp4FileSpec) ParseAtoms(fp *Mp4FilePro) error {
	var pos int64
	
	log.Println(self.end)
	
	for self.end >= pos {
		
		size, atom, err := fp.Mp4Read()
		
		log.Println(size, string(atom))
		
		if err != nil {
			log.Fatalln(err.Error())
			return err
		}
		
		sizeInt := util.Bytes2Int(size)
		
		offset := sizeInt - (int64)(len(size)) - (int64)(len(atom))
		
		log.Println(offset)
		
		pos += sizeInt
		
		if f, ok := mp4Atoms[string(atom)]; ok {
			f(self)
		}
		
		self.NextAtom(offset, fp)	
	}
	
	return nil
}