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

package file

import (
	"log"
	//"errors"
)

type ParseFuc func(fs *FileSpec)

var (
	commands map[string]ParseFuc
)

func init() {
	commands = map[string]ParseFuc {
		"ftyp": ftypPro,
	}
}

type FileSpec struct {
	mp4Name string
	
}

func NewFileSpec (name string) *FileSpec {
	fs := &FileSpec {
		mp4Name : name,
	}
	
	return fs
}

func (self *FileSpec)ParseBox(name string) {
	if f, ok := commands[name]; ok {
		f(self)
		return
	}

	//err := errors.New("Unsupported box: " + name)
	return
}

func ftypPro(fs *FileSpec) {
	log.Println("--------")
}