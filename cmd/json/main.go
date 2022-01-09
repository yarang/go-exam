package main

import (
	"encoding/json"
	"fmt"
	"github.com/urfave/cli"
	"go-exam/internal/utils"
	"io/ioutil"
	"os"
)

type valuesSet struct {
	Url   string `json:"url"`
	Count int    `json:"count"`
}

// Flags 설정을 위한 선언 부분
// Flags 설정은 read, write 두부분에서 사용함.
var (
	readFlags = []cli.Flag{
		cli.StringFlag{
			Name:  "json_file",
			Value: "./configs/data.json",
		},
	}
)

// readJson()
// json 내용을 읽고 이를 데이터 변수로 저장하는 기능
func readJson(c *cli.Context) error {
	v, err := ioutil.ReadFile(c.String("json_file"))
	utils.CheckError(err, -1)

	var data valuesSet
	utils.DebugErr(json.Unmarshal(v, &data))

	fmt.Println(data)
	fmt.Println(data.Url)
	fmt.Println(data.Count)

	return nil
}

// writeJson()
// JSON 내용을 파일로 작성하는 기능을 가진 함수
// JSON 내용을 저장하기 위한 변수가 array 로 만들어져서 저장 값이 [] 형태의 array 로 저장한다.
// data[0]을 사용하는 이유는 array 가 아닌 {}로 시작하는 하나의 값으로 저장하기 위해서다.
func writeJson(c *cli.Context) error {
	dataV := make([]valuesSet, 1)

	dataV[0].Url = "https://google.com"
	dataV[0].Count = 50

	doc, err := json.Marshal(dataV[0])
	utils.CheckError(err, -4)

	utils.DebugErr(ioutil.WriteFile(c.String("json_file"), doc, os.FileMode(0644)))

	fmt.Println(doc)

	fmt.Println(dataV[0])
	fmt.Println(dataV[0].Url)
	fmt.Println(dataV[0].Count)

	return nil
}

// cli 패키지에서 사용하는 command 명령을 선언하기 위한 부분
var (
	jsonExamRead = cli.Command{
		Action:      readJson,
		Name:        "read",
		Flags:       readFlags,
		Usage:       "Json read exam",
		Category:    "EXAM COMMANDS",
		Description: `Json 읽기 테스트를 위한 명령어`,
	}
	jsonExamWrite = cli.Command{
		Action:      writeJson,
		Name:        "write",
		Flags:       readFlags,
		Usage:       "Json write exam",
		Category:    "EXAM COMMANDS",
		Description: `Json 쓰기 테스트를 위한 명령어`,
	}
)

func main() {
	app := cli.NewApp()
	app.Name = "Exam Execute CLI"
	app.Usage = "Select command for exam."

	app.Commands = []cli.Command{
		jsonExamRead,
		jsonExamWrite,
	}

	err := app.Run(os.Args)
	utils.CheckError(err, -4)
}
