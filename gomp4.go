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

package main

import (
	"os"
	"log"
	"encoding/json"
	"fmt"
	"github/oikomi/mp4"
)

/*
#include <stdlib.h>
#include <stdio.h>
#include <string.h>

const char* build_time(void) {
	static const char* psz_build_time = "["__DATE__ " " __TIME__ "]";
	return psz_build_time;
}
*/
import "C"

var (
	buildTime = C.GoString(C.build_time())
)

func BuildTime() string {
	return buildTime
}

const VERSION string = "0.10"

func version() {
	fmt.Printf("gomp4 version %s Copyright (c) 2014 Harold Miao (miaohonghit@gmail.com)  \n", VERSION)
}

func main() {
	version()
	fmt.Printf("built on %s\n", BuildTime())
	
	if len(os.Args) != 2 {
		os.Exit(0)
	}
	
	fs := mp4.NewMp4FileSpec(os.Args[1])
	fp := mp4.NewMp4FilePro()
	
	err := fp.Mp4Open(fs)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
	
	err = fp.Mp4FileStat(fs)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
		
	fs.ParseAtoms(fp)
	res, _ := json.Marshal(fs)
	log.Println(string(res))
}