package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
)

// 예제용 interface 선언
type exam interface {
	getName() string
	getSize() float32
}

// interface 대응되는 struct 선언
type circle struct {
	name   string
	radian float32
}

// struct 종속적인 function 구현
// interface 대응되는 함수 이름 작성
func (s circle) getName() string {
	return s.name
}

// struct 종속적인 function 구현
// interface 대응되는 함수 이름 작성
func (s circle) getSize() float32 {
	return s.radian * s.radian * 3.14
}

// interface 대응되는 struct 선언
type rect struct {
	name   string
	width  float32
	height float32
}

// struct 종속적인 function 구현
// interface 대응되는 함수 이름 작성
func (s rect) getName() string {
	return s.name
}

// struct 종속적인 function 구현
// interface 대응되는 함수 이름 작성
func (s rect) getSize() float32 {
	return s.width * s.height
}

// 예젱용 함수 선언부
// cli Action에 추가할 수 있는 함수형 구현
func examRect(c *cli.Context) error {
	var ex exam

	s := rect{"rect", 5, 10}
	ex = s

	fmt.Println(ex.getName())
	fmt.Println(ex.getSize())

	cir := circle{"circle", 5}
	ex = cir

	fmt.Println(ex.getName())
	fmt.Println(ex.getSize())

	return nil
}

// cli command 추가용 변수 구현
var (
	exam1 = cli.Command{
		Action:      examRect,
		Name:        "rect",
		Usage:       "Rect exam",
		Category:    "EXAM COMMANDS",
		Description: `Exam 테스트를 위한 명령어`,
	}
)

func main() {
	app := cli.NewApp()
	app.Name = "Exam Execute CLI"
	app.Usage = "Select command for exam."

	app.Commands = []cli.Command{
		exam1,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
