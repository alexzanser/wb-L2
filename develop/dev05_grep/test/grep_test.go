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
		[]string{"1", "../files/file1.txt", "-A", "2", "-B", "2", "-c", "-i"},
	},

	{
		[]string{"1", "../files/file1.txt", "-A", "2", "-B", "5", "-i"},
	},

	{
		[]string{"1", "../files/file1.txt", "-A", "2", "-B", "2", "-i", "-F"},
	},

	{
		[]string{"1", "../files/file1.txt", "-A", "4", "-B", "2", "-C", "3", "-i", "-F"},
	},

	{
		[]string{"A", "../files/file3.txt", "-A", "2", "-B", "2", "-c", "-i"},
	},

	{
		[]string{"A", "../files/file3.txt", "-A", "2", "-B", "2", "-i"},
	},

	{
		[]string{"A", "../files/file3.txt", "-A", "2", "-B", "2", "-i", "-c"},
	},

}

func Test(t *testing.T) {
	for _, test := range TestCases {
		out1, _  := exec.Command("./grep_go", test.args...).Output()
		out2, _  := exec.Command("grep", test.args...).Output()

		if reflect.DeepEqual(out1, out2) == false {
			t.Error("test failed")
			fmt.Println(string(out1))
			fmt.Println(string(out2))
		}
	}
}
