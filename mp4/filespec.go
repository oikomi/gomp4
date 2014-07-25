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

type ParseAtomFuc func(fs *Mp4FileSpec, fp *Mp4FilePro, offset int64) error

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

type FtypAtom struct {
	offset int64
	size int64
	isFullBox bool
	majorBrand string
	minorVersion int64
	compatibleBrands string
}

func ftypRead(fs *Mp4FileSpec, fp *Mp4FilePro, offset int64) error {
	var err error
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
	fs.ftypAtom.size = sizeInt
	
	fs.ftypAtom.offset = offset
	fs.ftypAtom.isFullBox = false
	err = fp.Mp4Seek(offset + 8, 0)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}

	buf, err := fp.Mp4Read(4)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	fs.ftypAtom.majorBrand = string(buf)
	
	err = fp.Mp4Seek(12, 0)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	
	buf, err = fp.Mp4Read(4)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	fs.ftypAtom.minorVersion = util.Bytes2Int(buf)
	
	err = fp.Mp4Seek(16, 0)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	
	buf, err = fp.Mp4Read(12)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	fs.ftypAtom.compatibleBrands = string(buf)
	
	return nil
}

type MoovAtom struct {
	offset int64
	size int64
	isFullBox bool
	mvhdAtom MvhdAtom
}

func moovRead(fs *Mp4FileSpec, fp *Mp4FilePro, offset int64) error {
	log.Println("moovRead")
	var err error	
	fs.moovAtom.offset = offset

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
	fs.moovAtom.size = sizeInt
	
	var pos int64
	
	err = fp.Mp4Seek(8 + offset, 0)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	
	for fs.moovAtom.size > pos {
		size, atom, err := fp.Mp4ReadHeader()
		
		log.Println(size, string(atom))
		
		if err != nil {
			log.Fatalln(err.Error())
			return err
		}
		
		sizeInt := util.Bytes2Int(size)	

		pos += sizeInt
	
		if f, ok := mp4MoovAtoms[string(atom)]; ok {
			err = f(fs, fp, pos + 8 + offset - sizeInt)
			if err != nil {
				log.Fatalln(err.Error())
				return err	
			}
		}
		
		fs.nextAtom(pos + 8 + offset, fp)	
	}
	
	return nil
}

type MdatAtom struct {
	offset int64
	size int64
}

func mdatRead(fs *Mp4FileSpec, fp *Mp4FilePro, offset int64) error {
	return nil
}
////////////

type MvhdAtom struct {
	offset int64
	size int64
	isFullBox bool
	version int64
	flag int64
	creationTime int64
	modificationTime int64
	timescale int64
	duration int64
}

func mvhdRead(fs *Mp4FileSpec, fp *Mp4FilePro, offset int64) error {
	var err error
	fs.moovAtom.mvhdAtom.offset = offset
	fs.moovAtom.mvhdAtom.isFullBox = false
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
	fs.moovAtom.mvhdAtom.size = sizeInt
	
	size, err = fp.Mp4Read(1)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	fs.moovAtom.mvhdAtom.version = util.Bytes2Int(size)
	
	size, err = fp.Mp4Read(3)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	fs.moovAtom.mvhdAtom.flag = util.Bytes2Int(size)
	
	size, err = fp.Mp4Read(4)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	fs.moovAtom.mvhdAtom.creationTime = util.Bytes2Int(size)
	
	size, err = fp.Mp4Read(4)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	fs.moovAtom.mvhdAtom.modificationTime = util.Bytes2Int(size)
	
	size, err = fp.Mp4Read(4)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	fs.moovAtom.mvhdAtom.timescale = util.Bytes2Int(size)
	
	size, err = fp.Mp4Read(4)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	fs.moovAtom.mvhdAtom.duration = util.Bytes2Int(size)
		
	return nil
}

func trakRead(fs *Mp4FileSpec, fp *Mp4FilePro, offset int64) error {
	return nil
}

type Mp4FileSpec struct {
	mp4Name string
	end int64
	ftypAtom FtypAtom
	moovAtom MoovAtom
	mdatAtom MdatAtom
	
}

func NewMp4FileSpec (name string) *Mp4FileSpec {
	fs := &Mp4FileSpec {
		mp4Name : name,
		end : 0,
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
	
	for self.end > pos {
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

func (self *Mp4FileSpec) Dump()  {
	log.Println(self.mp4Name)	
	
}