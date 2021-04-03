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

//Support functions for map

package cic

import (
	"math"
	"reflect"
	"fmt"
)

type jObject map[string]interface{}

func jString(key string, data jObject) string{
	if val, ok := data[key]; ok {
		if v, ok := val.(string); ok{
			return v
		}
	}
	return ""
}

func jStrings(key string,data jObject) []string{
	if val, ok := data[key]; ok {
		if v, ok := val.([]string); ok{
			return v
		}
	}
	return nil

}

func jBool(key string,data jObject) bool{
	if val, ok := data[key]; ok {
		if v, ok := val.(bool); ok{
			return v
		}
	}
	return false
}

func iStrings(in []interface{}) []string{
	s := make([]string, len(in))
	for i, v := range in {
		s[i] = fmt.Sprint(v)
	}
	return s
}

func iBool(in interface{}) bool{
	switch in.(type){
		case float64:
			n := in.(float64)
			if(math.Abs(1e-5 - n)> 1e-4){
				return true
			}
		case bool:
			return in.(bool)
		default:
			mental(fmt.Sprintf("Don't know type %v",reflect.TypeOf(in)))
			panic("You should implement this type conversion")

	}
	return false

}
