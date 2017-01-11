package validator

import (
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

// IsBoltDBValid determine if the provided BoltDB is valid
func IsBoltDBValid(args []string) error {
	if err := boltPathProvided(args); err != nil {
		return err
	}

	bp := args[0]

	if err := boltDBValid(bp); err != nil {
		return err
	}

	ex, err := boltDBExists(bp)
	if err != nil {
		return err
	} else if !ex {
		return errors.New("the provided BoltDB does not exists")
	}

	return nil
}

func boltPathProvided(args []string) error {
	argslen := len(args)

	if argslen < 1 {
		return errors.New("a BoltDB path must be provided")
	}

	return nil
}

func boltDBValid(p string) error {
	if filepath.Ext(p) != ".db" {
		return errors.New("the provided file extension for BoltDB is not valid")
	}

	return nil
}
func boltDBExists(path string) (bool, error) {
	fi, err := os.Stat(path)
	switch {
	case err == nil:
		// do nothing
	case os.IsNotExist(err):
		return false, errors.Wrap(err, "the provided BoltDB does not exists.")
	case err != nil:
		return false, err
	}

	if mode := fi.Mode(); mode.IsRegular() {
		return true, nil
	}

	return false, nil
}
