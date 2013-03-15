package xmlparser

import (
	"fmt"
	"github.com/moovweb/gokogiri"
	"github.com/moovweb/gokogiri/xml"
	"reflect"
	"strconv"
	"strings"
)

func ParseXml(xml []byte) map[string]interface{} {
	var result map[string]interface{}

	doc, err := gokogiri.ParseXml(xml)

	if err != nil {
		fmt.Println("parsing error:%v\n", err)
		return nil
	} else {
		defer doc.Free()
		//   	fmt.Println (doc.String())
		root := doc.Root()
		m := xmlNodeToMap(root.XmlNode)

		if reflect.TypeOf(m).String() == "string" {
			m2 := make(map[string]interface{})
			m2[root.XmlNode.Name()] = m
			result = m2
		} else {
			result = m.(map[string]interface{})
		}
	}
	return result
}

// Returns a Node and its children as a generic Map
func xmlNodeToMap(node xml.Node) interface{} {
	attribute_map := make(map[string]interface{})
	content := ""

	child := node.FirstChild()
	for child != nil {
		if child.NodeType() == xml.XML_ELEMENT_NODE {
			// Add this child to the map
			value := xmlNodeToMap(child)

			attribute := attribute_map[child.Name()]
			if attribute == nil {
				// Not already in the Map
				attribute_map[child.Name()] = value
			} else if reflect.TypeOf(attribute).String() == "[]interface {}" {
				// Already in the Map as an Array
				attribute_map[child.Name()] = append(attribute.([]interface{}), value)
			} else {
				// Already in the map and need to convert it into an Array
				attribute_map[child.Name()] = []interface{}{attribute, value}
			}

		} else if (child.NodeType() == xml.XML_TEXT_NODE) || (child.NodeType() == xml.XML_CDATA_SECTION_NODE) {
			text := strings.TrimSpace(child.String())
			if len(text) > 0 {
				content = content + text
			}
		}
		child = child.NextSibling()
	}
	var result interface{}

	// Any content in the XML node takes precedance over child elements
	if content == "" {
		if len(attribute_map) > 0 {
			result = attribute_map
		} else {
			result = nil
		}
	} else {
		result = content
	}
	return result
}

// Pretty Print
func pp(name string, value interface{}, indent_size int) {
	indent := strings.Repeat("  ", indent_size)
	if value == nil {
		fmt.Println(indent, "\""+name+"\" => nil,")
	} else if reflect.TypeOf(value).String() == "string" {
		fmt.Println(indent, "\""+name+"\" => \""+value.(string)+"\",")
	} else if reflect.TypeOf(value).String() == "[]interface {}" {
		fmt.Println(indent, "\""+name+"\" => [")
		pp_array(value.([]interface{}), indent_size+1)
		fmt.Println(indent, "]")
	} else {
		fmt.Println(indent, "\""+name+"\" => {")
		pp_map(value.(map[string]interface{}), indent_size+1)
		fmt.Println(indent, "}")
	}
}

func pp_map(m map[string]interface{}, indent_size int) {
	for k, v := range m {
		pp(k, v, indent_size)
	}
}

func pp_array(array []interface{}, indent_size int) {
	i := 0
	for _, value := range array {
		pp("["+strconv.Itoa(i)+"]", value, indent_size)
		i = i + 1
	}
}
