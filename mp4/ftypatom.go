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

type FtypAtom struct {
	Offset int64
	Size int64
	IsFullBox bool
	MajorBrand []byte
	MinorVersion uint32
	CompatibleBrands string
	AllBytes []byte
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
	fs.FtypAtomInstance.Size = sizeInt
	
	err = fp.Mp4Seek(offset, 0)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	
	buf, err := fp.Mp4Read(fs.FtypAtomInstance.Size)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	
	fs.FtypAtomInstance.AllBytes = buf
	
	fs.FtypAtomInstance.Offset = offset
	fs.FtypAtomInstance.IsFullBox = false
	
	err = fp.Mp4Seek(offset + 8, 0)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}

	buf, err = fp.Mp4Read(4)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	//fs.FtypAtomInstance.MajorBrand = string(buf)
	fs.FtypAtomInstance.MajorBrand = buf
	
	//log.Println(fs.FtypAtomInstance.MajorBrand)
	
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
	fs.FtypAtomInstance.MinorVersion = util.Byte42Uint32(buf, 0)
	
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
	fs.FtypAtomInstance.CompatibleBrands = string(buf)
	
	return nil
}