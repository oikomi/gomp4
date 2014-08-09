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

type MoovAtom struct {
	Offset int64
	Size int64
	IsFullBox bool
	MvhdAtomInstance MvhdAtom
	TrakAtomInstance [2]TrakAtom
	AllBytes []byte
}

func moovRead(fs *Mp4FileSpec, fp *Mp4FilePro, offset int64) error {
	var err error
	fs.MoovAtomInstance.Offset = offset

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
	fs.MoovAtomInstance.Size = sizeInt
	
	err = fp.Mp4Seek(offset, 0)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	
	buf, err := fp.Mp4Read(8)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	
	fs.MoovAtomInstance.AllBytes = buf
	
	var pos int64
	
	err = fp.Mp4Seek(8 + offset, 0)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	
	for fs.MoovAtomInstance.Size > pos {
		size, atom, err := fp.Mp4ReadHeader()
		
		//log.Println(size, string(atom))
		
		if err != nil {
			log.Fatalln(err.Error())
			return err
		}
		
		sizeInt := util.Bytes2Int(size)	

		pos += sizeInt
	
		if f, ok := mp4MoovAtoms[string(atom)]; ok {
			err = f(fs, fp, pos + 8 + offset - sizeInt)
			
			if string(atom) == "trak" {
				trakNum ++
			}
			
			if err != nil {
				log.Fatalln(err.Error())
				return err	
			}
		}
		
		fs.nextAtom(pos + 8 + offset, fp)	
	}
	
	return nil
}