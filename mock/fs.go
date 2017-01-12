package mock

import (
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
)

// TMPPREFIX ...
const TMPPREFIX = "boltcli-tmp-test-file"

// TmpFile ...
type TmpFile struct {
	Path string
}

// Remove ...
func (tf *TmpFile) Remove() error {
	return os.Remove(tf.Path)
}

// NewTmpFile returns the filepath to a tmp file
func NewTmpFile() (*TmpFile, error) {
	tfile, err := ioutil.TempFile(os.TempDir(), TMPPREFIX)
	if err != nil {
		return nil, errors.Wrap(err, "temp file")
	}

	path := tfile.Name()
	tfile.Close()
	os.Remove(path)

	return &TmpFile{
		Path: path,
	}, nil
}
