package sort

import (
	"fmt"
	"os"
	"sort"
	linesmodule "sort/internal/lines"
	"github.com/spf13/cobra"
)

//Key is a struct for program flags received from cobra
type Key struct {
	k                   int
	numericSort         bool
	reverse             bool
	unique              bool
	monthSort           bool
	ignoreTailingBlanks bool
	check               bool
	humanNumericSort    bool
}

//InitKeys receives flags and set them to struct Key
func InitKeys(rootCmd *cobra.Command, key *Key) {
	rootCmd.PersistentFlags().BoolP("help", "", false, "help for this command")
	rootCmd.Flags().IntVarP(&key.k, "key", "k", 1, "field")
	rootCmd.Flags().BoolVarP(&key.numericSort, "numeric-sort", "n", false, "numeric-sort")
	rootCmd.Flags().BoolVarP(&key.reverse, "reverse", "r", false, "reverse")
	rootCmd.Flags().BoolVarP(&key.unique, "unique", "u", false, "unique")
	rootCmd.Flags().BoolVarP(&key.monthSort, "month-sort", "M", false, "month-sort")
	rootCmd.Flags().BoolVarP(&key.ignoreTailingBlanks, "ignore-tailing-blanks", "b", false, "ignore-tailing-blanks")
	rootCmd.Flags().BoolVarP(&key.check, "check", "c", false, "check")
	rootCmd.Flags().BoolVarP(&key.humanNumericSort, "human-numeric-sort", "h", false, "human-numeric-sort")
}

//CheckArguments validates received program flags
func CheckArguments(lines linesmodule.Lines, key *Key) error {
	if len(os.Args) < 2 {
		return fmt.Errorf("missing filename")
	}

	a := []bool{key.numericSort, key.monthSort, key.humanNumericSort}
	k := 0
	for _, flag := range a {
		if flag {
			k++
		}
	}

	if k > 1 {
		return fmt.Errorf("mutually exclusive flags")
	}

	return nil
}

//Sort select sorting mode and additional options
func Sort(lines linesmodule.Lines, key *Key) {
	lines.SetColumn(lines, key.k)

	if key.numericSort {
		sort.Slice(lines, lines.SortNumeric)
	} else if key.humanNumericSort {
		sort.Slice(lines, lines.SortHumanNumeric)
	} else if key.monthSort {
		sort.Slice(lines, lines.SortMonth)
	} else {
		sort.Slice(lines, lines.StandartSort)
	}

	if key.unique {
		lines = lines.Unique()
	}
	if key.reverse {
		lines.Reverse()
	}
	if key.ignoreTailingBlanks {
		lines.IgnoreTailingSpaces()
	}
}
