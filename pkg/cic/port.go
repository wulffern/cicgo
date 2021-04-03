//====================================================================
//        Copyright (c) 2021 Carsten Wulff Software, Norway
// ===================================================================
// Created       : wulff at 2021-4-3
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

type Port interface{
	Rect
	Set(r Rect)

}

type port struct{
	rect
	name string
	routeLayer Layer
	drawRect Rect
	altRects []Rect
}


func NewPort(name string) Port{
	p := port{}
	p.InitPort(name)
	return &p
}

func (p *port) InitPort(name string){
	p.name = name
	p.altRects = make([]Rect,0)
	p.drawRect = nil

}


func (p *port) Set(r Rect){
	if(r == nil){
		return
	}

	p.routeLayer = Rules.Layers[r.Layer()]
	p.altRects = append(p.altRects,r)
	p.drawRect = r
	p.SetRectFromRect(r)
	r.Connect(p.UpdateRect)

}

func (p *port) UpdateRect(){
	if(p.drawRect != nil){
		p.SetRectFromRect(p.drawRect)
	}
}
