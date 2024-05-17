package main

import (
	"log"
	"path/filepath"

	zip "github.com/kenzo0107/zip-go"
)

func main() {
	excludeFilepath := filepath.Join("testdata", ".ignore")
	excludes, _ := zip.ExcludeFilepaths(excludeFilepath)

	targetDir := "testdata"
	zipFile := "testdata.zip"
	if err := zip.Compress(targetDir, zipFile, excludes); err != nil {
		log.Fatal(err)
	}
}
