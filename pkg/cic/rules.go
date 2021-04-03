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

package cic

import (
	"github.com/wulffern/cicgo/internal/jsonwrap"
	"fmt"
)

var Rules RulesType



type Layer struct{
	Number int
	Datatype int
	Material string
	Previous string
	Next string
	Pin string
	Res string
	Color string
}

type Property struct{
	Name string
	Str string
}

type Device struct{
	Name string
	Ports []string
	Propertymap map[string]Property
}

type Technology struct{
	Gamma int
	Grid int
	Techlib string
	Devices map[string]Device
}

type RulesType struct{
	Version int
	Layers map[string]Layer
	Technology Technology
	DesignRules map[string]map[string]float64 `json:"rules"`
}

func LoadRules(rulefile string){
	jsonwrap.ReadFromJsonFile(rulefile,&Rules)
}

func (r *RulesType) get(key string, element string) int{

	v,ok := r.DesignRules[key][element]
	if(ok){
		return int(v*float64(r.Technology.Gamma))
	}else{
		mental(fmt.Sprintf("Could not find rule ",key,element))
		panic("Crap, no rule")
	}
}
