package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"tools/unmarshal"
	"os"
)

func main() {
	fh, err := os.Open("./plants.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fh.Close()
	jsonData, err := ioutil.ReadAll(fh)
	if err != nil {
		fmt.Println(err)
		return
	}
	
	var flowers []unmarshal.Flower
	err = json.Unmarshal(jsonData, &flowers)
	if err != nil {
		fmt.Println(err)
		return
	}
	for i, v := range flowers {
		fmt.Printf("index:%+v, name: %+v\n", i, v.Name)
	}
}
