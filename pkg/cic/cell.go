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

type NewCellType func(string) Cell

type Cell interface{
	Rect
	Name() string
	SetName(string)
	Init()
	Place()
	Route()
	AddAllPorts()
	Paint()
}

type cell struct{
	rect
	name string
	routes map[string]Rect
	ports map[string]Port
	namedRects map[string]Rect
	children []Rect

}

func (d *Design) NewCell(name string) Cell{
	c := cell{}
	c.InitCell(name)
	return &c
}

func (c *cell) InitCell(name string){
	c.InitRect()
	c.name = name
}

func (c *cell) Name() string{
	return c.name
}

func (c *cell) SetName(n string){
	c.name = n
}

func (c *cell) GetPort(name string) Port{
	if val,ok :=  c.ports[name];ok{
		return val
	}
	return nil
}

func (c *cell) AddPort(name string, r Rect){
	p := NewPort(name)
	p.Set(r)
	//TODO: p.SpicePort = c.IsSpicePort(name)
	c.Add(p)
}

func (c *cell) Add(child Rect){
	if(child == nil){

		return
	}

	//TODO: Figure out how I should do the children in GO
	//Should it be a map? or a list? I need a contains function to check if the

}

func (c *cell) Init(){

}

func (c *cell) Place(){

}

func (c *cell) Route(){

}


func (c *cell) AddAllPorts(){

}

func (c *cell) Paint(){

}
