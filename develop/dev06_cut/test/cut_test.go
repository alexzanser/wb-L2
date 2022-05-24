package test

import (
	"testing"
	"os/exec"
	"reflect"
	"fmt"
)

var TestCases = []struct {
	args	[]string
}{
	{
		[]string{"-f", "1-2,4-5", "../files/file1.txt"},
	},

	{
		[]string{"-f", "1,2,3,4", "-d", "b", "../files/file2.txt"},
	},

	{
		[]string{"-f", "1,-1,3,4", "-d", " ", "../files/file2.txt"},
	},

	{
		[]string{"-f", "1,2,3,4", "-d", ",", "-s",  "../files/file2.txt"},
	},

	{
		[]string{"-f", "2,1-4", "-d", ",", "-s",  "../files/file2.txt"},
	},

	{
		[]string{"-f", "1,2,3,4", "-d", "z",  "../files/file3.txt"},
	},

	{
		[]string{"-f", "2,1-4", "-d", ".", "-s",  "../files/file3.txt"},
	},
}

func Test(t *testing.T) {
	for _, test := range TestCases {
		out1, _  := exec.Command("./cut_go", test.args...).Output()
		out2, _  := exec.Command("cut", test.args...).Output()

		if reflect.DeepEqual(out1, out2) == false {
			t.Error("test failed")
			fmt.Println(string(out1))
			fmt.Println(string(out2))
		}
	}
}
