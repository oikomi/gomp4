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

type TkhdAtom struct {
	Offset int64
	Size int64
	IsFullBox bool
	Version uint8
	Flag uint32
	CreationTime uint32
	ModificationTime uint32
	TrakID uint32
	Reserved1 uint32
	Duration uint32
	Reserved2 uint64
	Layer uint16
	AlternateGroup int64
	Volume uint16
	Reserved3 uint16
	MatrixStructure []byte
	TrackWidth uint32
	TrackHeight uint32
	AllBytes []byte
}

func tkhdRead(fs *Mp4FileSpec, fp *Mp4FilePro, offset int64) error {
	var err error
	fs.MoovAtomInstance.TrakAtomInstance[trakNum].TkhdAtomInstance.Offset = offset
	fs.MoovAtomInstance.TrakAtomInstance[trakNum].TkhdAtomInstance.IsFullBox = false
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
	fs.MoovAtomInstance.TrakAtomInstance[trakNum].TkhdAtomInstance.Size = sizeInt
	
	err = fp.Mp4Seek(offset, 0)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	
	buf, err := fp.Mp4Read(fs.MoovAtomInstance.TrakAtomInstance[trakNum].TkhdAtomInstance.Size)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	
	fs.MoovAtomInstance.TrakAtomInstance[trakNum].TkhdAtomInstance.AllBytes = buf
	
	size, err = fp.Mp4Read(1)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	fs.MoovAtomInstance.TrakAtomInstance[trakNum].TkhdAtomInstance.Version = uint8(size[0])
	
	size, err = fp.Mp4Read(3)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	fs.MoovAtomInstance.TrakAtomInstance[trakNum].TkhdAtomInstance.Flag = util.Byte32Uint32(size, 0)
	
	size, err = fp.Mp4Read(4)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	fs.MoovAtomInstance.TrakAtomInstance[trakNum].TkhdAtomInstance.CreationTime = util.Byte42Uint32(size, 0)
	
	size, err = fp.Mp4Read(4)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	fs.MoovAtomInstance.TrakAtomInstance[trakNum].TkhdAtomInstance.ModificationTime = util.Byte42Uint32(size, 0)
	
	size, err = fp.Mp4Read(4)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	fs.MoovAtomInstance.TrakAtomInstance[trakNum].TkhdAtomInstance.TrakID = util.Byte42Uint32(size, 0)
	
	size, err = fp.Mp4Read(4)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	fs.MoovAtomInstance.TrakAtomInstance[trakNum].TkhdAtomInstance.Reserved1 = util.Byte42Uint32(size, 0)
	
	size, err = fp.Mp4Read(4)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	fs.MoovAtomInstance.TrakAtomInstance[trakNum].TkhdAtomInstance.Duration = util.Byte42Uint32(size, 0)
	
	size, err = fp.Mp4Read(8)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	//fs.MoovAtomInstance.TrakAtomInstance[trakNum].TkhdAtomInstance.Reserved2 = util.Byte82Uint32(size, 0)
	
	size, err = fp.Mp4Read(2)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	fs.MoovAtomInstance.TrakAtomInstance[trakNum].TkhdAtomInstance.Layer = util.Byte22Uint16(size, 0)
	
	size, err = fp.Mp4Read(2)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	fs.MoovAtomInstance.TrakAtomInstance[trakNum].TkhdAtomInstance.AlternateGroup = util.Bytes2Int(size)
	
	size, err = fp.Mp4Read(2)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	fs.MoovAtomInstance.TrakAtomInstance[trakNum].TkhdAtomInstance.Volume = util.Byte22Uint16(size, 0)
	
	size, err = fp.Mp4Read(2)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	fs.MoovAtomInstance.TrakAtomInstance[trakNum].TkhdAtomInstance.Reserved3 = util.Byte22Uint16(size, 0)
	
	size, err = fp.Mp4Read(36)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	fs.MoovAtomInstance.TrakAtomInstance[trakNum].TkhdAtomInstance.MatrixStructure = size
	
	size, err = fp.Mp4Read(4)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	fs.MoovAtomInstance.TrakAtomInstance[trakNum].TkhdAtomInstance.TrackWidth = util.Byte42Uint32(size, 0)
	
	size, err = fp.Mp4Read(4)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	fs.MoovAtomInstance.TrakAtomInstance[trakNum].TkhdAtomInstance.TrackHeight = util.Byte42Uint32(size, 0)
	
	return nil
}


