package main

import (
	"fmt"
	"github.com/urfave/cli"
	"go-exam/internal/utils"
	"os"
	"strconv"
)

var (
	rectFlags = []cli.Flag{
		cli.IntFlag{
			Name:  "width",
			Value: 50,
		},
		cli.IntFlag{
			Name:  "height",
			Value: 50,
		},
	}
	circleFlages = []cli.Flag{
		cli.IntFlag{
			Name:  "radian",
			Value: 50,
		},
	}
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
func examInterface(c *cli.Context) error {
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

func examRect(c *cli.Context) error {
	width, e := strconv.ParseFloat(c.String("width"), 32)
	utils.CheckError(e, -3)
	height, e := strconv.ParseFloat(c.String("height"), 32)
	utils.CheckError(e, -3)

	ex := rect{"rect", float32(width), float32(height)}

	fmt.Println(ex.getName())
	fmt.Println(ex.getSize())

	return nil
}

func examCircle(c *cli.Context) error {
	radian, e := strconv.ParseFloat(c.String("radian"), 32)
	utils.CheckError(e, -3)

	ex := circle{"rect", float32(radian)}

	fmt.Println(ex.getName())
	fmt.Println(ex.getSize())

	return nil
}

// cli command 추가용 변수 구현
var (
	examD = cli.Command{
		Action:      examInterface,
		Name:        "default",
		Usage:       "Interface exam",
		Category:    "EXAM COMMANDS",
		Description: `Exam 테스트를 위한 명령어`,
	}
	exam1 = cli.Command{
		Action:      examRect,
		Name:        "rect",
		Flags:       rectFlags,
		Usage:       "Rect exam",
		Category:    "EXAM COMMANDS",
		Description: `Rect Exam 테스트를 위한 명령어`,
	}
	exam2 = cli.Command{
		Action:      examCircle,
		Name:        "circle",
		Flags:       circleFlages,
		Usage:       "Circle exam",
		Category:    "EXAM COMMANDS",
		Description: `Circle Exam 테스트를 위한 명령어`,
	}
)

func main() {
	app := cli.NewApp()
	app.Name = "Exam Execute CLI"
	app.Usage = "Select command for exam."

	app.Commands = []cli.Command{
		examD,
		exam1,
		exam2,
	}

	err := app.Run(os.Args)
	utils.CheckError(err, -4)
}
