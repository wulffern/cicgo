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
    "fmt"
    "path/filepath"
    "os"
    "github.com/wulffern/cicgo/internal/jsonwrap"
    "reflect"
    "strings"
    "regexp"
)


type DesignDataType struct{
    Options jObject
    Include []string
    NoPortTranslation int
    Patterns map[string][]string
    Cells []jObject

}

type Design struct{
    addr *Design
    cells map[string]Cell
    jCells map[string]jObject
    cellNames []string
    ignoreSetYoffsetHalf bool
    includePaths []string
    translator map[string]string
    cellConstructor map[string]NewCellType
}

func NewDesign(includePaths []string) (*Design){
    d := Design{}
    d.includePaths = includePaths
    d.translator = map[string]string{
        "Gds::GdsPatternTransistor" : "PatternTransistor",
        "Gds::GdsPatternCapacitor" : "PatternCapacitor",
        "cIcCore::PatternResistor" : "PatternResistor",
        "Layout::LayoutDigitalCell" : "LayoutCell",
        "Layout::LayoutCapCellSmall" : "CapCellSmall",
        "Layout::LayoutCDACSmall" : "CdacSmall",
        "Layout::LayoutSARCDAC" : "SarCdac",
    }
    return &d
}

func (d *Design) Read(file string) bool{

    if(d.ReadCells(file)){
        //TODO: Add all cuts
        return true
    }else{
        return false
    }

}

func (d *Design) ReadCells(file string) bool{


    jdata := d.readJsonAndSpice(file)
    d.readOptions(jdata)
    d.readPatterns(jdata)
    d.readIncludes(jdata)
    comment("Reading '" + file + "'")

    d.cells = make(map[string]Cell,len(jdata.Cells))
    d.jCells = make(map[string]jObject,len(jdata.Cells))

    for _,jobj := range(jdata.Cells){
        increaseIndent()
        d.createCell(jobj)
        decreaseIndent()
    }


    return true

}


func (d *Design) readOptions(jdata *DesignDataType){
    if val, ok := jdata.Options["ignoreSetYoffsetHalf"]; ok {
        d.ignoreSetYoffsetHalf = val.(bool)
    }
}

// Read current JSON and spiceFiles
func (d *Design) readJsonAndSpice(file string) *DesignDataType{
    var jdata DesignDataType
    jsonwrap.ReadFromJsonFile(file,&jdata)


    //TODO(cawu): Add Spice file reading
    return &jdata
}

func (d *Design) readPatterns(jdata *DesignDataType){
    //TODO(cawu): Add reading of patterns
}

func (d *Design) readIncludes(jdata *DesignDataType){
    for _,f := range  jdata.Include{
        fileNotFound := true
        if _, err := os.Stat(f); err == nil {
            if(d.ReadCells(f)){
                fileNotFound = false
            }
        }

        for _,inc := range d.includePaths{
            fpath := filepath.Join(inc,f)
            if _, err := os.Stat(fpath); err == nil {
                if(d.ReadCells(fpath)){
                    fileNotFound = false
                }
            }
        }

        if(fileNotFound){
            mental("Could not find file '" +f+ "")
        }

    }

}




func (d *Design) createCell(jobj jObject) {

    name := jString("name",jobj)
	if jBool("abstract",jobj){
		comment(fmt.Sprintf("Skipping abstract cell %s",name))
		return
	}

    //Store src data so I can access it later
    d.jCells[name] = jobj

    if(name == ""){
        //warning("Cell without name somewhere in input data" + fmt.Sprintf("%v",data))
        return
    }

    comment(name)

    //- Find all parents
    rParents := make([]jObject,0,2)
    inherit := jString("inherit",jobj)
    if(inherit != ""){
        d.findAllParents(&rParents,inherit)
    }

    //- Inherit class type from parent
    cl :=jString("class",jobj)
    if(len(rParents) > 0 ){
        for _,p := range rParents{
            pcl := jString("class",p)
            if(pcl != ""){
                cl = pcl
            }
        }
    }

	//- Find objects to leech from
    leech:= jString("leech",jobj)
    if(leech != ""){
        if c,ok := d.jCells[leech];ok{
            rParents = append(rParents,c)
        }
    }

	//- Set the default cell name
    if(cl == ""){
        cl = "LayoutCell"
    }

    ncl,ok := d.translator[cl]
    if(ok){
        cl = ncl
    }

    //- Increase indent, but only decrease on function return
    increaseIndent()
    defer decreaseIndent()

	//- Instanciate an object through refelection
	ctr := "New" + cl
	fptr := reflect.ValueOf(d).MethodByName(ctr)

    if(!fptr.IsValid()){
        warning(fmt.Sprintf("Could not find constructor for %v",cl))
        return
    }

	inputs := make([]reflect.Value, 1)
	inputs[0] = reflect.ValueOf(name)
	result := fptr.Call(inputs)
	c := result[0].Interface().(Cell)

	//- Run the methods
    d.run("afterNew",c,&rParents,jobj)

    d.run("",c,&rParents,jobj)

    d.run("beforePlace",c,&rParents,jobj)
    //comment("Placing")
    c.Place()
    d.run("afterPlace",c,&rParents,jobj)

    d.run("beforeRoute",c,&rParents,jobj)
    //comment("Routing")
    c.Route()
    d.run("afterRoute",c,&rParents,jobj)

    d.run("beforePorts",c,&rParents,jobj)
    //comment("Adding Ports")
    c.AddAllPorts()
    d.run("afterPorts",c,&rParents,jobj)

    d.run("beforePaint",c,&rParents,jobj)
    //comment("Painting")
    c.Paint()
    d.run("afterPaint",c,&rParents,jobj)



    d.cells[name] = c
    d.cellNames = append(d.cellNames,name)
}

func (d *Design) findAllParents(rParents *[]jObject, inh string){

    if val,ok := d.jCells[inh]; ok {
        inh = jString("inherit",val)
        if(inh != ""){
            d.findAllParents(rParents,inh)
        }
        *rParents = append(*rParents,val)
    }
}

func (d *Design) run(key string, c Cell, parents *[]jObject,jobj jObject){

    for _,v := range *parents{
        d.runBasedOnType(key,c,v)
    }
    d.runBasedOnType(key,c,jobj)
}

func (d *Design) runBasedOnType(key string,c Cell, jobj jObject){

    //- if key is empty. then run the methods that are not
    // class, name etc
    if(key == ""){
        cname := jString("name",jobj)
        d.runMethods(cname,c,jobj)
    }else{
        cname := fmt.Sprintf("%-20s %-20s",jString("name",jobj),key)
        jj,ok := jobj[key]
        if(ok){
            switch jj.(type){
                case map[string]interface{}:
				d.runMethods(cname,c,jj.(map[string]interface{}))
				case []interface{}:
				for _,v := range jj.([]interface{}){
					d.runMethods(cname,c,v.(map[string]interface{}))
				}

                default:

                fmt.Printf("Uknown type %v\n",reflect.TypeOf(jj))
            }

        }
    }
}



func (d *Design) runMethods(cname string, c Cell, jobj jObject){

    for k,v := range jobj{

        if match,_ := regexp.MatchString("^(rows|class|name|comment|inherit)$|(^after|^before)",k);match{
            continue
        }
        k := strings.ToUpper(string(k[0])) + k[1:len(k)]

		invoked := true
        if strings.HasSuffix(k,"s"){
            nk := k[:len(k)-1]
            ar := v.([]interface{})
            for _,av := range ar{
                invoked = invoked && Invoke(cname,c,nk,av)
            }
        }else{
            invoked = invoked && Invoke(cname,c,k,v)
        }

		str := fmt.Sprintf("%-41s %-30s",cname,k)
		if(invoked){
			comment(str)
		}else{
			mental(str)
		}

    }
    //fmt.Println(ct.Method)
}

func Invoke(cname string,c Cell, name string, args... interface{}) bool{

    t := reflect.TypeOf(c)
    _,ok := t.MethodByName(name)

    if(!ok){

        return false
    }

    inputs := make([]reflect.Value, len(args))
    for i, _ := range args {
        inputs[i] = reflect.ValueOf(args[i])
    }

    reflect.ValueOf(c).MethodByName(name).Call(inputs)
	return true
}
