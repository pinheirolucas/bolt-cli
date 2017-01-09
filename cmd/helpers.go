package cmd

import (
	"fmt"
	"os"
)

func er(err error) {
	fmt.Println("Error:", err.Error())
	os.Exit(1)
}

// is a dir and exists
func dirExists(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}

	switch mode := fi.Mode(); {
	case mode.IsDir():
		return true, nil
	case mode.IsRegular():
		return false, nil
	}

	return false, err
}
