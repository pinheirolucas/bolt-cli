package validator

import "os"

// DirExists is a dir and exists
func DirExists(path string) (bool, error) {
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
