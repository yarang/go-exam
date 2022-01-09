package utils

import (
	"fmt"
	"os"
)

// CheckError
// 에러값을 체크하여 프로그램의 종료를 판단하는 함수
func CheckError(err error, code int) {
	if err != nil {
		fmt.Println(err)
		os.Exit(code)
	}
}

// DebugErr
// 함수를 직접 실행한 결과를 받아 함수의 에러값을 체크하는 함수
func DebugErr(errVal error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	CheckError(errVal, -30)
}
