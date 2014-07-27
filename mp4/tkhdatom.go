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

type TkhdAtom struct {
	Offset int64
	Size int64
	IsFullBox bool
	Version int64
	Flag int64
	CreationTime int64
	ModificationTime int64
	TrakID int64
	Reserved1 int64
	Duration int64
	Reserved2 int64
	Layer int64
	AlternateGroup int64
	Volume int64
	Reserved3 int64
	MatrixStructure []byte
	TrackWidth int64
	TrackHeight int64
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
	
	size, err = fp.Mp4Read(1)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	fs.MoovAtomInstance.TrakAtomInstance[trakNum].TkhdAtomInstance.Version = util.Bytes2Int(size)
	
	size, err = fp.Mp4Read(3)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	fs.MoovAtomInstance.TrakAtomInstance[trakNum].TkhdAtomInstance.Flag = util.Bytes2Int(size)
	
	size, err = fp.Mp4Read(4)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	fs.MoovAtomInstance.TrakAtomInstance[trakNum].TkhdAtomInstance.CreationTime = util.Bytes2Int(size)
	
	size, err = fp.Mp4Read(4)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	fs.MoovAtomInstance.TrakAtomInstance[trakNum].TkhdAtomInstance.ModificationTime = util.Bytes2Int(size)
	
	size, err = fp.Mp4Read(4)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	fs.MoovAtomInstance.TrakAtomInstance[trakNum].TkhdAtomInstance.TrakID = util.Bytes2Int(size)
	
	size, err = fp.Mp4Read(4)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	fs.MoovAtomInstance.TrakAtomInstance[trakNum].TkhdAtomInstance.Reserved1 = util.Bytes2Int(size)
	
	size, err = fp.Mp4Read(4)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	fs.MoovAtomInstance.TrakAtomInstance[trakNum].TkhdAtomInstance.Duration = util.Bytes2Int(size)
	
	size, err = fp.Mp4Read(8)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	fs.MoovAtomInstance.TrakAtomInstance[trakNum].TkhdAtomInstance.Reserved2 = util.Bytes2Int(size)
	
	size, err = fp.Mp4Read(2)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	fs.MoovAtomInstance.TrakAtomInstance[trakNum].TkhdAtomInstance.Layer = util.Bytes2Int(size)
	
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
	fs.MoovAtomInstance.TrakAtomInstance[trakNum].TkhdAtomInstance.Volume = util.Bytes2Int(size)
	
	size, err = fp.Mp4Read(8)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	fs.MoovAtomInstance.TrakAtomInstance[trakNum].TkhdAtomInstance.Reserved3 = util.Bytes2Int(size)
	
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
	fs.MoovAtomInstance.TrakAtomInstance[trakNum].TkhdAtomInstance.TrackWidth = util.Bytes2Int(size)
	
	size, err = fp.Mp4Read(4)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	fs.MoovAtomInstance.TrakAtomInstance[trakNum].TkhdAtomInstance.TrackHeight = util.Bytes2Int(size)
	
	return nil
}


