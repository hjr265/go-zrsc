package zrsc

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

var (
	initErr error

	fs = map[string]*zip.File{}
)

type File struct {
	io.ReadCloser

	name string
}

func (f *File) Stat() (os.FileInfo, error) {
	return fs[f.name].FileInfo(), nil
}

func Open(name string) (*File, error) {
	if initErr != nil {
		return nil, initErr
	}

	rc, err := fs[name].Open()
	if err != nil {
		return nil, err
	}

	return &File{
		ReadCloser: rc,
		name:       name,
	}, nil
}

func init() {
	defer func() {
		val := recover()
		if val == nil {
			return
		}
		if err, ok := val.(error); ok {
			initErr = err
		} else {
			panic(val)
		}
	}()

	path, err := filepath.Abs(os.Args[0])
	catch(err)

	f, err := os.Open(path)
	catch(err)
	fi, err := f.Stat()
	catch(err)

	r, err := zip.NewReader(f, fi.Size())
	catch(err)

	for _, f := range r.File {
		fi := f.FileInfo()
		if fi.IsDir() {
			continue
		}

		fs[f.Name] = f
	}
}

func catch(err error) {
	if err != nil {
		panic(err)
	}
}
