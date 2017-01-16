package cmd

import (
	"fmt"
	"os"
)

func er(err error) {
	fmt.Println("Error:", err.Error())
	os.Exit(1)
}
