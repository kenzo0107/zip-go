# zip-go

The golang library archive/zip cannot compress some files except for some files.
This is a zip file with a function to exclude some files.

Excluded files can also be specified by file format.
Wildcards are also supported.

## Usage

```go
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
```

## LICENSE

[MIT License](https://github.com/kenzo0107/zip-go/blob/main/LICENSE)
