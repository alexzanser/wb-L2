package sort

import (
	"fmt"
	"os"
	"sort"
	"log"
	linespackage "sort/internal/lines"
	keypackage "sort/internal/key"
)

//CheckArguments validates received program flags
func CheckArguments(lines linespackage.Lines, key *keypackage.Key) error {
	if len(os.Args) < 2 {
		return fmt.Errorf("missing filename")
	}

	a := []bool{key.NumericSort, key.MonthSort, key.HumanNumericSort}
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
func Sort(key *keypackage.Key) linespackage.Lines{

	var lines linespackage.Lines

	lines, err := linespackage.GetLines(os.Args[1])
	if err != nil {
		log.Fatal(fmt.Errorf("error when getting lines: %v", err))
	}

	if err := CheckArguments(lines, key); err != nil {
		log.Fatal(fmt.Errorf("invalid arguments: %v", err))
	}

	lines.SetColumn(lines, key.K)

	if key.NumericSort {
		sort.Slice(lines, lines.SortNumeric)
	} else if key.HumanNumericSort {
		sort.Slice(lines, lines.SortHumanNumeric)
	} else if key.MonthSort {
		sort.Slice(lines, lines.SortMonth)
	} else {
		sort.Slice(lines, lines.StandartSort)
	}

	if key.Unique {
		lines = lines.Unique()
	}
	if key.Reverse {
		lines.Reverse()
	}
	if key.IgnoreTailingBlanks {
		lines.IgnoreTailingSpaces()
	}

	return lines
}
