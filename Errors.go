package glhf

import (
	"fmt"
	"io/fs"
)

type ErrorString string

func (err ErrorString) Error() string { return string(err) }

const (
	ErrGameCreatedAlready  = ErrorString("glhf: an instance of game has been created already")
	ErrFileIsNotReaderAt   = ErrorString("glhf: file does not implement io.ReaderAt.")
	ErrInvalidMountName    = ErrorString("glhf: mount point cannot contain a colon (\":\").")
	ErrEmptyMountName      = ErrorString("glhf: mount point is empty.")
	ErrEmptyPathName       = ErrorString("glhf: path is empty.")
	ErrEmptyAssetPath      = ErrorString("glhf: asset path is empty.")
	ErrMountPointExists    = ErrorString("glhf: mount point already exists.")
	ErrMountPointNotExists = ErrorString("glhf: mount point does not exists.")
	ErrFSIsNil             = ErrorString("glhf: cannot mount nil FS")
)

type FileNotExist string

func (path FileNotExist) Error() string {
	return fmt.Sprintf("glhf: File %q does not exist", string(path))
}

func (path FileNotExist) Unwrap() error { return fs.ErrNotExist }
