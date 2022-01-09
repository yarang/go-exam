package main

import (
	"encoding/json"
	"fmt"
	"go-exam/internal/utils"
	"io/ioutil"
)

type valuesSet struct {
	Url    string `json:"url"`
	Lcount int    `json:"lcount"`
}

func main() {
	v, err := ioutil.ReadFile("./configs/data.json")
	utils.CheckError(err, -1)

	var data valuesSet
	utils.DebugErr(json.Unmarshal(v, &data))
	fmt.Println(data)

}
