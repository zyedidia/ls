package ls

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

const (
	test_root = "/tmp/ls_test"
)

func _cd(path string) {
	err := os.Chdir(path)
	if err != nil {
		fmt.Printf("error: os.Chdir(%s)\n", path)
		fmt.Printf("\t%v\n", err)
		os.Exit(1)
	}
}

func _mkdir(path string) {
	err := os.MkdirAll(path, 0755)
	if err != nil {
		fmt.Printf("error: os.MkdirAll(%s, 0755)\n", path)
		fmt.Printf("\t%v\n", err)
		os.Exit(1)
	}
}

func _rmdir(path string) {
	err := os.RemoveAll(path)
	if err != nil {
		fmt.Printf("error: os.Removeall(%s)\n", path)
		fmt.Printf("\t%v\n", err)
		os.Exit(1)
	}
}

func _mkfile(path string) {
	_, err := os.Create(path)
	if err != nil {
		fmt.Printf("error: os.Create(%s)\n", path)
		fmt.Printf("\t%v\n", err)
		os.Exit(1)
	}
}

func TestMain(m *testing.M) {

	//
	// setup
	//

	// create the test root directory if it does not exist
	_, err := os.Stat(test_root)
	if err != nil && os.IsNotExist(err) {
		_mkdir(test_root)
	} else if err != nil {
		fmt.Printf("error: os.Stat(%s)\n", test_root)
		fmt.Printf("\t%v\n", err)
		os.Exit(1)
	} else {
		_rmdir(test_root)
		_mkdir(test_root)
	}

	_cd(test_root)

	//
	// run the tests
	//
	result := m.Run()

	//
	// teardown
	//
	_rmdir(test_root)

	os.Exit(result)
}

func TestNoArgsFiles(t *testing.T) {
	_cd(test_root)

	_mkdir("NoArgs")
	_cd("NoArgs")
	_mkfile("a")
	_mkfile("b")
	_mkfile("c")

	var output_buffer bytes.Buffer
	var args []string
	ls(&output_buffer, args)

	expected := "a b c"

	if output_buffer.String() != expected {
		t.Logf("expected \"%s\", but got \"%s\"\n",
			expected,
			output_buffer.String())
		t.Fail()
	}
}
