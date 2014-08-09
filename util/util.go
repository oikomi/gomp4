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

package util

import (
	//"encoding/binary"
	"encoding/hex"
	//"fmt"
	"strconv"
)

const (
	BigEndian = 0
)


func Byte42Uint32(data []byte, endian int) uint32 {
	var i uint32
	if 0 == endian {
		i = uint32(uint32(data[3]) + uint32(data[2])<<8 + uint32(data[1])<<16 + uint32(data[0])<<24)
	}
	
	if 1 == endian {
		i = uint32(uint32(data[0]) + uint32(data[1])<<8 + uint32(data[2])<<16 + uint32(data[3])<<24)
	}

	return i
}

func Byte32Uint32(data []byte, endian int) uint32 {
	var i uint32
	if 0 == endian {
		i = uint32(uint32(data[2]) + uint32(data[1])<<8 + uint32(data[0])<<16)
	}
	
	if 1 == endian {
		i = uint32(uint32(data[0]) + uint32(data[1])<<8 + uint32(data[2])<<16)
	}

	return i
}

func Byte22Uint16(data []byte, endian int) uint16 {
	var i uint16
	if 0 == endian {
		i = uint16(uint16(data[1]) + uint16(data[0])<<8)
	}
	
	if 1 == endian {
		i = uint16(uint16(data[0]) + uint16(data[1])<<8)
	}

	return i
}



func Bytes2Int(b []byte) int64 {
	
	num := hex.EncodeToString(b)
	
	i, err := strconv.ParseInt(num, 16, 64)
	if err != nil {
		panic(err)
	}
	//fmt.Println(i)

	return i
	/*
	var s string
	for _, num := range b {
		num = hex.EncodeToString()
		s += fmt.Sprintf("%x", num)

	}
	i, err := strconv.ParseInt(s, 16, 32)
	if err != nil {
		panic(err)
	}

	return i
	*/
}


func ToHex(ten int) (hex []int, length int) { 
	m := 0 
	
	hex = make([]int, 0) 
	length = 0; 
	
	for { 
		m = ten / 16 
		ten = ten % 16 
		
		if(m == 0) { 
			hex = append(hex, ten) 
			length++ 
			break 
		} 
	
		hex = append(hex, m) 
		length++; 
	} 
	return 
} 



