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
)

type SegMp4Header struct {
	Ftyp []byte
	Moov []byte
}

func (self * SegMp4Header) FtypCover(fs *Mp4FileSpec)   {
	self.Ftyp = fs.FtypAtomInstance.AllBytes
	log.Println(self.Ftyp)
}

type SegMp4Data struct {
	
}

