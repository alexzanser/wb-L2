package sort

import (
	"testing"
	"github.com/spf13/cobra"
	"sort/cmd"
	"strings"
	linesmodule "sort/internal/lines"
)

var TestCases = []struct {
	input 	string
	flags	string
	output	string
}{
	{
		"../files/file1.txt",
		"",
`1 D
2.18 C
3.14 B
3.4 A
5.11
ABC 1
DEF 2
PPP 3
ZZZ 7
\\\ 0`,
	},
	{
		"../files/file2.txt",
		"-M -r",
`dec
nov
oct
sep
july
june
feb`,
	},
	{
		"../files/file3.txt",
		"-h -r",
`3AC
3AAC
3AAB    
2B
2B
1A              `,
	},
}

func Test(t *testing.T) {
	cmd := &cobra.Command{}
	key := &sort.Key{}
	lines := linesmodule.Lines{}
	sort.InitKeys(cmd, key)

	for _, test := range TestCases {
		cmd.ParseFlags(strings.Split(test.flags, " "))
		_ = cmd.Execute()
	
		lines, _ = linesmodule.GetLines(test.input)
		sort.Sort(&lines, key)
	
		var sb strings.Builder
		for i, line := range lines {
			sb.WriteString((line[1]))
			if i != len(lines) - 1 {
				sb.WriteString("\n")
			}
		}

		if sb.String() != test.output {
			t.Error("test failed")
		}
	}
}
