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
package main

import (
	"fmt"
	"os"
	"regexp"

	 "github.com/wulffern/cicgo/pkg/rules"
)

func main(){

	includePaths := []string{}
	arguments := []string{}

	argsWithoutProg := os.Args[1:]

	for i := 0;i<len(argsWithoutProg);i++{
		a := argsWithoutProg[i]
		match,_ := regexp.MatchString("^--I",a)
		if(match){
			includePaths = append(includePaths, argsWithoutProg[i+1])
			i++
		}else{
			arguments = append(arguments,a)
		}
	}
	if(len(arguments) >= 2){
		file := arguments[0]
		rulefile := arguments[1]
		library := ""
		if(len(arguments) > 2){
			library = arguments[2]
		}
		if(library == ""){
			r,_ := regexp.Compile("/?([^\\/]+)\\.json")
			re := r.FindStringSubmatch(file)
			library = re[1]
		}


		//TODO(cawu): Call rules
		//TODO(cawu): Call design
		//TODO(cawu): Write .cic file

		rules.Load(rulefile)
		//s, _ := json.MarshalIndent(&rules.Rules, "", "\t")
		fmt.Println(rules.Rules.DesignRules)

		_= library
		_ = file
		_ = rulefile
	}else{
		help := `Usage: cic <JSON file> <Technology file> [<Output name>]
Example: cic SAR_ESSCIRC16_28N tech.json
About:
    cIcCreator reads a JSON object definition file, a technology rule file
    and a SPICE netlist (assuming the same name as object definition file)
    and outputs a cic description file (.cic)
Options:
    --I       <path>        Path to search for include files
`
		fmt.Printf(help)
	}
}
