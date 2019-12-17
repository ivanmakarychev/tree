package main

import (
	"bytes"
	"testing"
)

const testFullResult = `├───a_dir
│   ├───a_dir
│   │   ├───a_dir
│   │   │   └───a_file.txt (9b)
│   │   └───empty_file.txt (empty)
│   ├───a_file.txt (9b)
│   └───z_dir
│       ├───a_file.txt (19b)
│       └───empty_file.txt (empty)
├───z_dir
│   └───empty_file.txt (empty)
└───z_file.txt (23b)
`

func TestTreeFull(t *testing.T) {
	out := new(bytes.Buffer)
	err := dirTree(out, "testdata", true)
	if err != nil {
		t.Errorf("test failed - error")
	}
	result := out.String()
	if result != testFullResult {
		t.Errorf("test failed - results not match\nGot:\n%v\nExpected:\n%v", result, testFullResult)
	}
}

const testDirResult = `├───a_dir
│   ├───a_dir
│   │   └───a_dir
│   └───z_dir
└───z_dir
`

func TestTreeDir(t *testing.T) {
	out := new(bytes.Buffer)
	err := dirTree(out, "testdata", false)
	if err != nil {
		t.Errorf("test failed - error")
	}
	result := out.String()
	if result != testDirResult {
		t.Errorf("test failed - results not match\nGot:\n%v\nExpected:\n%v", result, testDirResult)
	}
}
