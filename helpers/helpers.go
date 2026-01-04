package helpers

import (
	"fmt"
	"time"
)

func HandleError(err error) {
	time := time.Now().Format("2006-01-02 15:04:05")
	if err != nil {
		fmt.Println("===================")
		fmt.Printf("[ERROR]: Encountered an error at %v\n", time)
		fmt.Printf("[ACTUAL ERROR]: %v\n", err)
		fmt.Println("===================")
	}
}
