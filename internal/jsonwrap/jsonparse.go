//====================================================================
//        Copyright (c) 2021 Carsten Wulff Software, Norway
// ===================================================================
// Created       : wulff at 2021-3-29
// ===================================================================
//  The MIT License (MIT)
//
//  Permission is hereby granted, free of charge, to any person obtaining a copy
//  of this software and associated documentation files (the "Software"), to deal
//  in the Software without restriction, including without limitation the rights
//  to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
//  copies of the Software, and to permit persons to whom the Software is
//  furnished to do so, subject to the following conditions:
//
//  The above copyright notice and this permission notice shall be included in all
//  copies or substantial portions of the Software.
//
//  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
//  IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
//  FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
//  AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
//  LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
//  OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
//  SOFTWARE.
//
//====================================================================

package jsonwrap

import (
	"fmt"
	"encoding/json"
	"os"

	"bufio"
	"strings"
	"regexp"
)

func  ReadFromJsonFile(file string, data interface{}){

	jsonFile, err := os.Open(file)

	if err != nil {
		fmt.Println("Could not find rulefile")
		panic(err)
	}

	defer jsonFile.Close()

	scanner := bufio.NewScanner(jsonFile)
	builder := strings.Builder{}
	for scanner.Scan(){
		line := scanner.Text()
		m,_ := regexp.MatchString("^\\s*//",line)
		if(m){
			continue
		}
		builder.WriteString(line)
	}

	buff := builder.String()
	//fmt.Println(buff)

	byteValue := []byte(buff)



    err = json.Unmarshal([]byte(byteValue), &data)
	if err != nil {
		fmt.Errorf("Error parsing json file '%s'\n",file)

		panic(err)
	}

	return

}
