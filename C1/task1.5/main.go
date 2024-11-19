package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func main() {
	// expectedJSON := []byte(`[
	// {"name":"Oleg", "grade":14}, {"name":"Ivan", "grade":11}]`)
	inputJSON := []byte(`[
        {
            "name": "Oleg",
            "class": "9B"
        },
        {
            "name": "Ivan",
            "class": "9A"
        },
        {
            "name": "Maria",
            "class": "9B"
        },
        {
            "name": "John",
            "class": "9A"
        }
    ]`)

	expectedJSONMap := map[string][]byte{
		"9A": []byte(`[{"class":"9A","name":"Ivan"},{"class":"9A","name":"John"}]`),
		"9B": []byte(`[{"class":"9B","name":"Oleg"},{"class":"9B","name":"Maria"}]`),
	}

	classJSONMap, _ := splitJSONByClass(inputJSON)
	for class, expectedJSON := range expectedJSONMap {
		if !bytes.Equal(classJSONMap[class], expectedJSON) {
			fmt.Printf("Expected JSON data for class %s to be %s, but got %s", class, expectedJSON, classJSONMap[class])
		}
	}
	// fmt.Println(mergedJSON)
	// if err != nil {
	// 	fmt.Printf("Error while merging JSON data: %v", err)
	// }

	// if !bytes.Equal(mergedJSON, expectedJSON) {
	// 	fmt.Printf("Expected merged JSON data to be %s, but got %s", expectedJSON, mergedJSON)
	// }
	//fmt.Println(string(out))
	// layout := "02.01.2006"
	// start, _ := time.Parse(layout, "13.12.2022")
	// end, _ := time.Parse(layout, "15.12.2022")
	// fmt.Println(ExtractLog("file.txt", start, end))
}

type Person struct {
	Name  string `json:"name"`
	Grade int    `json:"grade"`
}

func modifyJSON(jsonData []byte) ([]byte, error) {
	var p []Person
	err := json.Unmarshal(jsonData, &p)
	if err != nil {
		return nil, err
	}
	for i, _ := range p {
		p[i].Grade += 1
	}
	return json.Marshal(&p)
}

func mergeJSONData(jsonDataList ...[]byte) ([]byte, error) {
	var mergedData []interface{}

	for _, jsonData := range jsonDataList {
		var data interface{}
		if err := json.Unmarshal(jsonData, &data); err != nil {
			return nil, err
		}

		if arr, ok := data.([]interface{}); ok {
			for _, item := range arr {
				mergedData = append(mergedData, item)
			}
		} else {
			mergedData = append(mergedData, arr...)
		}
	}
	return json.Marshal(mergedData)
}

type Persona struct {
	Class string `json:"class"`
	Name  string `json:"name"`
}

func splitJSONByClass(jsonData []byte) (map[string][]byte, error) {
	var p []Persona
	out := make(map[string][]byte)
	if err := json.Unmarshal(jsonData, &p); err != nil {
		return nil, err
	}
	for _, v := range p {
		if v1, ok := out[v.Class]; ok {
			var buffer bytes.Buffer
			jsonstr, _ := json.Marshal(v)
			buffer.Write(v1)
			buffer.Write([]byte(","))
			buffer.Write(jsonstr)
			out[v.Class] = buffer.Bytes()
			buffer.Reset()
		} else {
			out[v.Class], _ = json.Marshal(v)
		}
	}
	for k, v := range out {
		var buffer bytes.Buffer
		buffer.Write([]byte("["))
		buffer.Write(v)
		buffer.Write([]byte("]"))
		out[k] = buffer.Bytes()
		buffer.Reset()
	}
	return out, nil
}
