package Tool

import "fmt"

func IsError(err error, msg string)  {
	if err != nil {
		fmt.Println(msg,err)
		return
	}
}