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
<<<<<<< HEAD
	"encoding/hex"
=======
>>>>>>> aa878f053521b2f018d2f6b17583d681659329f2
	"fmt"
	"strconv"
)

func Bytes2Int(b []byte) int64 {
<<<<<<< HEAD
	
	num := hex.EncodeToString(b)
	
	i, err := strconv.ParseInt(num, 16, 64)
	if err != nil {
		panic(err)
	}
	fmt.Println(i)

	return i
	/*
	var s string
	for _, num := range b {
		num = hex.EncodeToString()
=======
	var s string
	for _, num := range b {
>>>>>>> aa878f053521b2f018d2f6b17583d681659329f2
		s += fmt.Sprintf("%x", num)

	}
	i, err := strconv.ParseInt(s, 16, 32)
	if err != nil {
		panic(err)
	}

	return i
<<<<<<< HEAD
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



=======
}
>>>>>>> aa878f053521b2f018d2f6b17583d681659329f2
