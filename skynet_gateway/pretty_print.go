package skynet_gateway

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

/*
 * Package: skynet_gateway
 *
 * Purpose: Pretty Print Arbitrary Map, Array
 *
 * Author:  Reid Morrison
 *
 * Usage:
 *
 *   xml, err := ioutil.ReadFile("./data.xml")
 *   m, err = MapFromXml(xml)
 *   PrettyPrintMap(m, 0)
 *
 */

// Pretty Print Arbitrary Map
func PrettyPrintMap(m map[string]interface{}, indent_size int) {
	for k, v := range m {
		PrettyPrint(k, v, indent_size)
	}
}

// Pretty Print Arbitrary Array
func PrettyPrintArray(array []interface{}, indent_size int) {
	i := 0
	for _, value := range array {
		PrettyPrint("["+strconv.Itoa(i)+"]", value, indent_size)
		i = i + 1
	}
}

// Pretty Print Named Element
func PrettyPrint(name string, value interface{}, indent_size int) {
	indent := strings.Repeat("  ", indent_size)
	if value == nil {
		fmt.Println(indent, "\""+name+"\" => nil,")
	} else if reflect.TypeOf(value).String() == "string" {
		fmt.Println(indent, "\""+name+"\" => \""+value.(string)+"\",")
	} else if reflect.TypeOf(value).String() == "[]interface {}" {
		fmt.Println(indent, "\""+name+"\" => [")
		PrettyPrintArray(value.([]interface{}), indent_size+1)
		fmt.Println(indent, "]")
	} else {
		fmt.Println(indent, "\""+name+"\" => {")
		PrettyPrintMap(value.(map[string]interface{}), indent_size+1)
		fmt.Println(indent, "}")
	}
}

