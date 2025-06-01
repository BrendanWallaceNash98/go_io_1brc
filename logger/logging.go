package logger

import (
	"fmt"
)

func PanicError(err error) {
	if err != nil {
		panic(err)
	}
}

func LogError(err error) {
	if err != nil {
		fmt.Println("v%", err)
	}
}

func LogErrorDetail(err error, comment string) {
	if err != nil {
		fmt.Println(comment+" \n %v", err)
	}
}
