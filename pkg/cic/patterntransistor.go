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
	"math"
)


type PatternTransistor interface{
	PatternTile

}

type patternTransistor struct{
	patternTile
}

func (d *Design) NewPatternTransistor(name string) Cell{
	p := patternTransistor{}
	p.InitPatternTransistor(name)
	return &p
}

func ( p *patternTransistor ) InitPatternTransistor(name string){
	p.InitPatternTile(name)
}

func (p *patternTransistor) InitCoordinatesData() map[string]interface{}{
	data := make(map[string]interface{})
	data["isTransistor"] = false
	data["wmin"] = math.MaxInt32
	data["wmax"] = math.MinInt32
	data["pofinger"] = 0
	data["nf"] = 0
	data["useMinLength"] = false
	return data
}
