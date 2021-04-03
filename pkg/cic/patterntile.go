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
    "strings"
    //"reflect"
    "fmt"
    "regexp"
)

type PatternTile interface{
    Cell
}

type rectStringsType map[string]map[int]map[int]rune

type patternTile struct{
    cell
    xmax int
    ymax int
    xspace int
    yspace int
    yoffset float64
    xoffset float64
    horizontalGridMultiplier float64
    verticalGridMultiplier float64
    horizontalGrid float64
    verticalGrid float64
    minPolyLength int
	polyWidthAdjust int
    data map[string]interface{}
    arraylength int
    mirrorPatternString bool
    rectStrings rectStringsType
    layers map[string][]string
    layerNames []string
	widthOffset float64
	heightOffset float64
}



func (d *Design) NewPatternTile(name string) Cell{
    p := patternTile{}
    p.InitPatternTile(name)
    return &p
}

func (p *patternTile) Yoffset(o float64){
	p.yoffset = o
}

func (p *patternTile) Xoffset(o float64){
	p.xoffset = o
}

func (p *patternTile) Widthoffset(o float64){
	p.widthOffset = o
}

func (p *patternTile) Heightoffset(o float64){
	p.heightOffset = o
}


func (p *patternTile) InitPatternTile(name string){
    p.InitCell(name)
    p.horizontalGridMultiplier = 1
    p.verticalGridMultiplier = 1
    p.xspace = int(float64(Rules.get("ROUTE","horizontalgrid"))*p.horizontalGridMultiplier)
    p.yspace = int(float64(Rules.get("ROUTE","verticalgrid"))*p.horizontalGridMultiplier)

    p.data = p.InitCoordinatesData()
    p.arraylength = 0
    p.mirrorPatternString = false
    p.xmax = 0
    p.ymax = 0
    p.rectStrings = make(rectStringsType)
    p.layers = make(map[string][]string)
    p.layerNames = make([]string,0)

}

func (p *patternTile) Init(){
    if(p.horizontalGrid != 0){
        p.xspace = int(p.horizontalGrid)
    }
    if(p.verticalGrid != 0){
        p.yspace = int(p.verticalGrid)
    }
    if(p.minPolyLength == 0){
        p.minPolyLength = Rules.get("PO","mingatelength")
    }
}

func (p *patternTile) MirrorPatternString(in interface{}){
    p.mirrorPatternString = iBool(in)
}

func (p *patternTile) FillCoordinatesFromString(in []interface{}){

    ar := iStrings(in)

    layer := ar[0]
    ar = ar[1:]

    p.InitFillCoordinate()

    //TODO: Implement copyRows

    //- Check that length is the same as previous

    if(p.arraylength == 0){
        p.arraylength = len(ar)
    }
    if(p.arraylength != len(ar)){
        warning(fmt.Sprintf("%s does not have %d lines\n",layer,p.arraylength))
    }

    strs := make([]string,len(ar))
    for i,str := range ar{


        //Copy column
        /*if(copyColumn_.length() > 0){
                for(int z=0;z<copyColumn_.length();z++){
                    CopyColumn c = copyColumn_[z];
                    if(str.length() < c.offset);
                    QString sorg = str.mid(c.offset,c.length);
                    for(int x=0;x<c.count;x++){
                        str.insert(c.offset,sorg);
                    }
                }
            }
        */

        if(p.mirrorPatternString){
            var sb  strings.Builder
            for j := len(str)-1;j >= 0;j--{
                sb.WriteString(string(str[j]))
            }
            str = sb.String()
        }


        for x,c := range str{
            y := len(ar) -i -1
            if(y > p.ymax){
                p.ymax = y
            }
            if(x > p.xmax){
                p.xmax = x
            }

            //Skip numbers
            if m,_ := regexp.MatchString("[0-9]",string(c)); m{
                continue
            }
            if c != '-'{
                //fmt.Println(layer,x,y,reflect.TypeOf(p.rectStrings))
                //p.rectStrings[layer][x][y] = c
                p.OnFillCoordinate(c,layer,x,y)
            }
        }
        strs = append(strs,str)
    }
    p.layers[layer] = strs
    p.layerNames = append(p.layerNames,layer)

    p.EndFillCoordinate()
}

func (p *patternTile) InitCoordinatesData() map[string]interface{}{
    data := make(map[string]interface{})
    return data
}

func (p *patternTile) InitFillCoordinate(){

}

func (p *patternTile) OnFillCoordinate(c rune, layer string, x int, y int){

}

func (p *patternTile) EndFillCoordinate(){

}

func (p *patternTile) translateX(x int) int{
    return int((float64(x) + p.xoffset)*float64(p.xspace))
}

func (p *patternTile) translateY(y int) int{
	return int((float64(y) + p.yoffset)*float64(p.yspace))
}

func (p *patternTile) Paint(){

    //TODO: Implement patterns
    //p.readPatterns()


    for _,layer := range p.layerNames{
        strs := p.layers[layer]
        N := len(strs)
        for y := 0;y<=p.ymax;y++{
            currentHeight := p.yspace
            for x := 0;x <= p.xmax;x++{
                s := strs[N - y - 1]

                c := s[x]
                if(c == '-'){
                    continue
                }

				rect := NewRect()
				rect.SetLayer(layer)

				var port Port

				xs := p.translateX(x)
				ys := p.translateY(y)

				//- Change dimension of poly, adapt to transistor width etc
				if(layer == "PO"){
					currentHeight = Rules.get(layer,"width")
					if(currentHeight < p.minPolyLength){
						currentHeight = p.minPolyLength
					}
				}else if( c == 'x' ){
					currentHeight = p.yspace
				}

				if( c== 'X' || p.polyWidthAdjust == 0 ){
					currentHeight = p.yspace
				}

				switch c{
					case 'A','B','P','N','D','S','G':
					port = p.GetPort(string(c))
					if(port == nil){
						port = NewPort(string(c))
						p.Add(port)
					}
				}



				_ = xs
				_ = ys

            }


        }
    }


}
