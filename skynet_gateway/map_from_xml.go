package skynet_gateway

import (
	"github.com/moovweb/gokogiri"
	"github.com/moovweb/gokogiri/xml"
	"reflect"
	//	"fmt"
	"strings"
)

/*
 * Package: skynet_gateway
 *
 * Description:
 *   Parse arbitrary XML String into a Map
 *   Also converts any elements containing '-' to '_'
 *
 *
 * Author:  Reid Morrison
 *
 * Usage:
 *
 *   xml, err := ioutil.ReadFile("./data.xml")
 *   doc, err = MapFromXml(xml)
 *
 */
func MapFromXml(xml []byte) (map[string]interface{}, error) {
	var result map[string]interface{}

	doc, err := gokogiri.ParseXml(xml)

	if err != nil {
		return nil, err
	} else {
		defer doc.Free()
		//   	fmt.Println (doc.String())
		root := doc.Root()
		if root == nil {
			// Blank string
			return make(map[string]interface{}), nil
		}
		m := xmlNodeToMap(root.XmlNode)

		if m == nil {
			// Empty document
			result = make(map[string]interface{})
		} else if reflect.TypeOf(m).String() == "string" {
			// Just contains content with no attributes or children
			m2 := make(map[string]interface{})
			m2[root.XmlNode.Name()] = m
			result = m2
		} else {
			// Child Map
			result = m.(map[string]interface{})
		}
	}
	return result, nil
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
			// Replace all '-' with '_' in element names
			name := strings.Replace(child.Name(), "-", "_", -1)

			attribute := attribute_map[name]
			if attribute == nil {
				// Not already in the Map
				attribute_map[name] = value
			} else if reflect.TypeOf(attribute).String() == "[]interface {}" {
				// Already in the Map as an Array
				attribute_map[name] = append(attribute.([]interface{}), value)
			} else {
				// Already in the map and need to convert it into an Array
				attribute_map[name] = []interface{}{attribute, value}
			}

			// Merge in any XML attributes
			//    Attributes() map[string]*AttributeNode
			xml_attributes := child.Attributes()
			if (xml_attributes != nil) && (len(xml_attributes) > 0) {
				for k, v := range xml_attributes {
					name = strings.Replace(k, "-", "_", -1)
					value = v.Content()
					attribute_map[name] = value
				}
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
		// Merge in any XML attributes
		//    Attributes() map[string]*AttributeNode
		xml_attributes := node.Attributes()
		if (xml_attributes != nil) && (len(xml_attributes) > 0) {
			for k, v := range xml_attributes {
				name := strings.Replace(k, "-", "_", -1)
				value := v.Content()
				attribute_map[name] = value
			}
		}
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
