package dry

import "fmt"

func HandleError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func HandleErrorPanic(err error) {
	if err != nil {
		panic(fmt.Sprintf("%v", err))
	}
}
