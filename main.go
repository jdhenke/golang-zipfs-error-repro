package main

import (
	"archive/zip"
	"bytes"
	"golang.org/x/tools/godoc/vfs/zipfs"
	"io"
	"log"
)

func main() {

	// example files to use
	files := map[string]string{"foo": "foo", "bar/baz": "baz"}

	// create test zip file from above files
	b := new(bytes.Buffer)
	zw := zip.NewWriter(b)
	for file, contents := range files {
		w, err := zw.Create(file)
		if err != nil {
			log.Fatal(err)
		}
		_, err = io.WriteString(w, contents)
		if err != nil {
			log.Fatal(err)
		}
	}
	zw.Close()

	// create zipfs from zip file
	zr, err := zip.NewReader(bytes.NewReader(b.Bytes()), int64(b.Len()))
	if err != nil {
		log.Fatal(err)
	}
	rc := &zip.ReadCloser{
		Reader: *zr,
	}
	fs := zipfs.New(rc, "foo")

	// try to list contents of root directory
	// FAILS HERE!
	infos, err := fs.ReadDir("/")
	if err != nil {
		log.Fatal(err)
	}
	for _, info := range infos {
		log.Printf("Found: %v\n", info.Name())
	}
}
