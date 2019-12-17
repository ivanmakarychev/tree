package main

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
)

const (
	node = "├───"
	bend = "└───"
	bar  = "│   "
	tab  = "    "
)

func dirTree(out io.Writer, path string, printFiles bool) error {
	w := bufio.NewWriter(out)
	files, err := readDirectory(path, printFiles)
	if err != nil {
		return err
	}
	if len(files) > 0 {
		i := 0
		for ; i < len(files)-1; i++ {
			err = printFileInfo(w, files[i], path, printFiles, "", false)
		}
		err = printFileInfo(w, files[i], path, printFiles, "", true)
	}
	err = w.Flush()
	return err
}

func printFileInfo(w *bufio.Writer, fileInfo os.FileInfo, parent string, printFiles bool, indent string, last bool) (err error) {
	if fileInfo.IsDir() {
		_, err = w.WriteString(indent)
		if last {
			_, err = w.WriteString(bend)
			indent += tab
		} else {
			_, err = w.WriteString(node)
			indent += bar
		}
		w.WriteString(fileInfo.Name() + "\n")
		parent = filepath.Join(parent, fileInfo.Name())
		var files []os.FileInfo
		files, err = readDirectory(parent, printFiles)
		if err != nil {
			return
		}
		if len(files) > 0 {
			i := 0
			for ; i < len(files)-1; i++ {
				err = printFileInfo(w, files[i], parent, printFiles, indent, false)
			}
			err = printFileInfo(w, files[i], parent, printFiles, indent, true)
		}
	} else if printFiles {
		_, err = w.WriteString(indent)
		if last {
			_, err = w.WriteString(bend)
		} else {
			_, err = w.WriteString(node)
		}
		_, err = w.WriteString(fileInfo.Name() + " (")
		if fileInfo.Size() == 0 {
			_, err = w.WriteString("empty)\n")
		} else {
			_, err = w.WriteString(strconv.FormatInt(fileInfo.Size(), 10))
			_, err = w.WriteString("b)\n")
		}
	}
	return
}

func readDirectory(path string, withFiles bool) (result []os.FileInfo, err error) {
	var files []os.FileInfo
	files, err = ioutil.ReadDir(path)
	if err != nil {
		return
	}
	if withFiles {
		return files, err
	}
	for _, f := range files {
		if f.IsDir() {
			result = append(result, f)
		}
	}
	return
}
