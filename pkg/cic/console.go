//====================================================================
//        Copyright (c) 2021 Carsten Wulff Software, Norway
// ===================================================================
// Created       : wulff at 2021-3-30
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

package cic

import (
	"fmt"
	"strings"
)



const (
	AnsiBlack = "30"
	AnsiRed = "31"
	AnsiGreen = "32"
	AnsiBrown = "33"
	AnsiBlue = "34"
	AnsiMagenta = "35"
	AnsiCyan = "36"
	AnsiLightgray = "37"
	AnsiStart = "\033[0;"
	AnsiReset = "\033[0m"
)

var reportIndent int = 0

func mental(msg string){
	report(msg,AnsiRed)
}

func warning(msg string){
	report(msg,AnsiBlue)
}

func comment(cm string){
	report(cm,AnsiGreen)
}

func report(msg string, color string){
	fmt.Println(AnsiStart+color+ "m",strings.Repeat(" ",reportIndent),msg ,AnsiReset)
}

func increaseIndent(){
	reportIndent++
	reportIndent++
}

func decreaseIndent(){
	reportIndent--
	reportIndent--
}
