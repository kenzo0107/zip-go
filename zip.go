package zip

import (
	"archive/zip"
	"bufio"
	"io"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

func Compress(src, dst string, excludeFilepaths []string) error {
	excludeFilepaths = append(excludeFilepaths, dst)

	zipFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	return filepath.Walk(src, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		relPath, err := filepath.Rel(src, filePath)
		if err != nil {
			return err
		}

		if slices.Contains(excludeFilepaths, relPath) {
			return nil
		}

		zipFile, err := zipWriter.Create(relPath)
		if err != nil {
			return err
		}

		fsFile, err := os.Open(filePath)
		if err != nil {
			return err
		}
		defer fsFile.Close()

		_, err = io.Copy(zipFile, fsFile)
		return err
	})
}

func ExcludeFilepaths(path string) ([]string, error) {
	if path == "" {
		return []string{}, nil
	}
	if _, err := os.Stat(path); err != nil {
		return nil, err
	}

	dirname := filepath.Dir(path)

	excludes, err := readfile(path)
	if err != nil {
		return nil, err
	}
	var excludeFilepaths []string
	for _, exclude := range excludes {
		p := filepath.Join(dirname, exclude)
		matches, err := filepath.Glob(p)
		if err != nil {
			return nil, err
		}
		for _, m := range matches {
			e, err := filepath.Rel(dirname, m)
			if err != nil {
				return nil, err
			}
			excludeFilepaths = append(excludeFilepaths, e)
		}
	}
	return excludeFilepaths, nil
}

func readfile(file string) ([]string, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		t := scanner.Text()
		if strings.HasPrefix(t, "#") || t == "" {
			continue
		}
		lines = append(lines, t)
	}
	return lines, scanner.Err()
}
