package skynet_gateway

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func BenchmarkXmlParser(b *testing.B) {
	b.StopTimer()
	input, err := ioutil.ReadFile("./test/inquiry.xml")
	if err != nil {
		fmt.Println("Couldn't read file", err.Error())
	} else {
		b.StartTimer()
		for i := 0; i < b.N; i++ {
			MapFromXml(input)
		}
	}
}

func TestSimpleMapFromXml(t *testing.T) {
	input := "<root><name>Jack</name></root>"
	doc, err := MapFromXml([]byte(input))
	if err != nil {
		t.Error("XML Parsing Error:", err)
		return
	}

	if doc["name"] != "Jack" {
		PrettyPrint("root", doc, 0)
		t.Error("name must be Jack")
	}
}

func TestMapFromXmlInquiry(t *testing.T) {
	xml, err := ioutil.ReadFile("./test/inquiry.xml")

	if err != nil {
		t.Error("Couldn't read file: test/inquiry.xml", err)
		return
	}

	doc, err := MapFromXml([]byte(xml))
	if err != nil {
		t.Error("XML Parsing Error:", err)
		return
	}

	// PrettyPrint("inquiry", doc, 0)

	if doc["date_of_birth"] != "1971-05-04" {
		t.Error("date_of_birth != 1971-05-04")
	}
	if doc["zip_code"] != "12345" {
		t.Error("zip_code != 12345")
	}
	if doc["street_address"] != "123 Somewhere St" {
		t.Error("street_address != 123 Somewhere St")
	}
	if doc["middle_initial"] != nil {
		t.Error("middle_initial != nil")
	}
	if doc["cell_phone"] != "8135551234" {
		t.Error("cell_phone != 8135551234")
	}
	if doc["last_name"] != "Bloggs" {
		t.Error("last_name != Bloggs")
	}
	if doc["title"] != nil {
		t.Error("title != nil")
	}
	if doc["timestamp"] != "2013-03-08T23:26:10+00:00" {
		t.Error("timestamp != 2013-03-08T23:26:10+00:00")
	}
	if doc["city"] != "Clearwater" {
		t.Error("city != Clearwater")
	}
	if doc["count"] != "34" {
		t.Error("count != 34")
	}
	if doc["email_address"] != "joe.bloggs@somewhere.net" {
		t.Error("email_address != joe.bloggs@somewhere.net")
	}
	if doc["time_stamp_nil"] != nil {
		t.Error("time_stamp_nil != mil")
	}
	if doc["state"] != "FL" {
		t.Error("state != FL")
	}
	if doc["first_name"] != "Joe" {
		t.Error("first_name != Joe")
	}
	if doc["amount"] != "200.40" {
		t.Error("amount != 200.40")
	}
}

func TestMapFromXmlMenu(t *testing.T) {
	xml, err := ioutil.ReadFile("./test/menu.xml")

	if err != nil {
		t.Error("Couldn't read file: test/menu.xml", err)
		return
	}

	doc, err := MapFromXml([]byte(xml))
	if err != nil {
		t.Error("XML Parsing Error:", err)
		return
	}

	if doc["food"].([]interface{})[0].(map[string]interface{})["name"] != "Belgian Waffles" {
	   PrettyPrint("menu", doc, 0)
		t.Error("food[0].name != Belgian Waffles")
	}

	if doc["food"].([]interface{})[4].(map[string]interface{})["name"] != "Homestyle Breakfast" {
	   PrettyPrint("menu", doc, 0)
		t.Error("food[4].name != Homestyle Breakfast")
	}
}

// Turns out that libXml is actually quite resilient
// and will figure out what was meant rather than being
// strict
func TestMapFromXmlMalformedXml(t *testing.T) {
	input := "<root><name>Jack</Bla"
	doc, err := MapFromXml([]byte(input))
	if err != nil {
		t.Error("XML Parsing Error:", err)
		return
	}

	if doc["name"] != "Jack" {
		PrettyPrint("root", doc, 0)
		t.Error("name must be Jack")
	}
}

func TestMapFromXmlEmptyXml(t *testing.T) {
	input := "<root/>"
	doc, err := MapFromXml([]byte(input))
	if err != nil {
		t.Error("XML Parsing Error:", err)
		return
	}

	if len(doc) != 0 {
		PrettyPrint("root", doc, 0)
		t.Error("Map must be empty with empty XML")
	}
}

func TestMapFromXmlContentXml(t *testing.T) {
	input := "<root>hello</root>"
	doc, err := MapFromXml([]byte(input))
	if err != nil {
		t.Error("XML Parsing Error:", err)
		return
	}

	if doc["root"] != "hello" {
		PrettyPrint("root", doc, 0)
		t.Error("root must be hello")
	}
}

func TestMapFromXmlBlank(t *testing.T) {
	input := ""
	doc, err := MapFromXml([]byte(input))
	if err != nil {
		t.Error("XML Parsing Error:", err)
		return
	}

	if len(doc) != 0 {
		PrettyPrint("root", doc, 0)
		t.Error("Map must be empty with empty XML")
	}
}

func TestMapFromXmlWhitespace(t *testing.T) {
	input := "                \n              \n               "
	doc, err := MapFromXml([]byte(input))
	if err != nil {
		t.Error("XML Parsing Error:", err)
		return
	}

	if len(doc) != 0 {
		PrettyPrint("root", doc, 0)
		t.Error("Map must be empty with empty XML")
	}
}

func TestMapFromXmlNil(t *testing.T) {
	doc, err := MapFromXml([]byte(nil))
	if err != nil {
		t.Error("XML Parsing Error:", err)
		return
	}

	if len(doc) != 0 {
		PrettyPrint("root", doc, 0)
		t.Error("Map must be empty with nil XML")
	}
}

func TestMapFromXmlAttributes(t *testing.T) {
	input := "<root name=\"Jack\"/>"
	doc, err := MapFromXml([]byte(input))
	if err != nil {
		t.Error("XML Parsing Error:", err)
		return
	}

	if doc["name"] != "Jack" {
		PrettyPrint("root", doc, 0)
		t.Error("name must be Jack")
	}
}
