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

type Rect interface{
	Layer() string
	Connect(func())
	X1() int
	Y1() int
	X2() int
	Y2() int
	SetLayer(string)
	//Left() int
	//Right() int
	//Top() int
	//Bottom() int
	//Width() int
	//CenterX() int
	//CenterY() int
	//Translate(ax int, ay int)
	//MoveTo(x int, y int )
	//MoveCenter(x int,y int)
	//Adjust(dx1 int, dy1 int, dx2 int, dy2 int)
}

type rect struct{
	layer string
	net string
	x1 int
	y1 int
	x2 int
	y2 int
	isDevice bool
	listeners []func()
}


func NewRect() Rect{
	r := rect{}
	r.InitRect()
	return &r
}

func (r *rect) InitRect(){
	r.x1 =	 0
	r.y1 = 0
	r.x2 = 0
	r.y2 = 0
	r.layer = ""
	r.net = ""
	r.listeners = make([]func(),0)
}

func (r *rect) Layer()string{
	return r.layer
}

func (r *rect) X1() int{
	return r.x1
}

func (r *rect) X2() int{
	return r.x2
}

func (r *rect) Y1() int{
	return r.y1
}

func (r *rect) Y2() int{
	return r.y2
}


func (r *rect) SetLayer(layer string){
	r.layer = layer
}

func (r *rect) SetRect(layer string, x1 int, y1 int, width int, height int){
	r.layer = layer
	r.x1 = x1
	r.y1 = y1
	r.x2 = x1 + width
	r.y2 = y1 + height
}

func (r *rect) SetRectFromRect(rs Rect){
	r.layer = rs.Layer()
	r.x1 = rs.X1()
	r.y1 = rs.Y1()
	r.x2 = rs.X2()
	r.y2 = rs.Y2()
}


func (r *rect) Connect(fp func()){
	r.listeners = append(r.listeners,fp)
}



func (r *rect) Updated(){
	for _,fp := range r.listeners{
		fp()
	}
}
