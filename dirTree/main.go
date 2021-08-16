package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
)

// Generate the correct current prefix
func generateCurrentPrefix(currentPrefix string, isLast bool) string {
	prefix := currentPrefix
	if isLast {
		prefix += "└───"
	} else {
		prefix += "├───"
	}

	return prefix
}

// Generate the correct prefix for inside files
func generateNextPrefix(currentPrefix string, isLast bool) string {
	var prefix string

	if isLast {
		prefix = currentPrefix + "\t"
	} else {
		prefix = currentPrefix + "│\t"
	}

	return prefix
}

// Generate file size string in bytes
// If the file is empty - generates string (empty)
func generateSizeString(file os.DirEntry) (string, error) {
	fileInfo, err := file.Info()
	if err != nil {
		return "", err
	}

	fileSize := fileInfo.Size()
	if fileSize == 0 {
		return "(empty)", nil
	} else {
		return "(" + strconv.Itoa(int(fileSize)) + "b)", nil
	}
}

// Print the directory tree to the output stream with all files
func dirTreeWithFiles(out io.Writer, path string, currentPrefix string) error {
	// Generate file splice and sort it
	files, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	// Output all directories and files
	for idx, file := range files {
		// Print current directory
		isLast := idx == len(files)-1

		// Generate prefix and size string
		prefix := generateCurrentPrefix(currentPrefix, isLast)
		if file.IsDir() {
			fmt.Fprintln(out, prefix+file.Name())
			prefix = generateNextPrefix(currentPrefix, isLast)
			dirTreeWithFiles(out, filepath.Join(path, file.Name()), prefix)
		} else {
			// Generate file size string in bytes
			sizeStr, err := generateSizeString(file)
			if err != nil {
				return err
			}

			fmt.Fprintln(out, prefix+file.Name(), sizeStr)
		}
	}

	return nil
}

// Print the directory tree to the output stream without files
// currentPrefix - proper prefix for output
func dirTreeWithoutFiles(out io.Writer, path string, currentPrefix string) error {
	// Generate file splice
	files, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	// Generate directory slice
	dirs := make([]os.DirEntry, 0)
	for _, file := range files {
		if file.IsDir() {
			dirs = append(dirs, file)
		}
	}

	// Output all directories
	for idx, file := range dirs {
		// Print current directory
		isLast := idx == len(dirs)-1

		// Generate prefix
		prefix := generateCurrentPrefix(currentPrefix, isLast)

		fmt.Fprintln(out, prefix+file.Name())

		// Recursively print the insides of this directory
		prefix = generateNextPrefix(currentPrefix, isLast)

		dirTreeWithoutFiles(out, filepath.Join(path, file.Name()), prefix)
	}

	return nil
}

// Print the directory tree to the output stream
// printFiles determines whether the files would be printed or not
func dirTree(out io.Writer, path string, printFiles bool) error {
	if printFiles {
		return dirTreeWithFiles(out, path, "")
	} else {
		return dirTreeWithoutFiles(out, path, "")
	}
}

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}
