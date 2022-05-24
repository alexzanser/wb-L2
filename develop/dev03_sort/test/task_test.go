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
		[]string{"../files/file1.txt"},
	},

	{
		[]string{"../files/file1.txt", "-r"},
	},

	{
		[]string{"../files/file1.txt", "-k", "1"},
	},

	{
		[]string{"../files/file1.txt", "-r", "-u"},
	},

	{
		[]string{"../files/file2.txt", "-n", "-r"},
	},

	{
		[]string{"../files/file2.txt", "-n", "-b"},
	},

	{
		[]string{"../files/file2.txt", "-n", "-M"},
	},

	{
		[]string{"../files/file2.txt", "-h"},
	},

	{
		[]string{"../files/file3.txt", "-h", "-r"},
	},

	{
		[]string{"../files/file3.txt", "-n"},
	},
}

func Test(t *testing.T) {
	for _, test := range TestCases {
		out1, _  := exec.Command("./sort_go", test.args...).Output()
		out2, _  := exec.Command("sort", test.args...).Output()

		if reflect.DeepEqual(out1, out2) == false {
			t.Error("test failed")
			fmt.Println(string(out1))
			fmt.Println(string(out2))
		}
	}
}
