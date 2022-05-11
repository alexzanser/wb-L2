package main

import (
	"bufio"
	// "log"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"github.com/spf13/cobra"
)

func GetLines(filename string) ([][]string, error) {
	file , err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error while open file: %v", err)
	}
	defer file.Close()

	lines := make([][]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, strings.Split(scanner.Text(), ""))
	}

	return lines, nil
}

// func Check(keys map[string]struct{}, lines [][]string) string {

// }

func SortNumeric(keys map[string]*bool, lines [][]string) {
	less := func(i, j int) bool {
		d1, _ := strconv.ParseFloat(lines[i][field], 64)
		d2, _ := strconv.ParseFloat(lines[j][field], 64)
		return d1 < d2 
	}
	sort.Slice(lines, less)
	notSorted := 0
	for i, line := range  lines {
		if _, err := strconv.ParseFloat(line[field], 64); err == nil {
			notSorted = i
			break
		}
	}
	sort.Slice(lines[:notSorted], 
				func(i, j int) bool {return lines[i][field] < lines[j][field]})
}

func SortHumanNumeric(keys map[string]*bool, lines [][]string) {

}

// func SortMonth(keys map[string]struct{}, lines [][]string) [][]string {

// }

// var field *int
var field int


type Key struct {
	k					int
	numericSort 		bool
	reverse				bool
	unique				bool
	monthSort			bool
	ignoreLeadingBlanks	bool
	check				bool
	humanNumericSort	bool
}

func InitKeys(cobra key *Key) {

} 

func main() {
    rootCmd := cobra.Command{}
	rootCmd.PersistentFlags().BoolP("help", "", false, "help for this command")
	rootCmd.Flags().IntVarP(&field, "key", "k", 1, "field")
	
	key := &Key{}
	rootCmd.Flags().BoolVarP(&key.numericSort, "numeric-sort", "n", false, "numeric-sort")

	rootCmd.Flags().BoolVarP(&key.reverse, "reverse", "r", false, "reverse")
	rootCmd.Flags().BoolVarP(&key.unique, "unique", "u", false, "unique")
	rootCmd.Flags().BoolVarP(&key.monthSort, "month-sort", "M", false, "month-sort")
	rootCmd.Flags().BoolVarP(&key.ignoreLeadingBlanks, "ignore-leading-blanks", "b", false, "ignore-leading-blanks")

	rootCmd.Flags().BoolVarP(&key.check, "check", "c", false, "check")
	rootCmd.Flags().BoolVarP(&key.humanNumericSort, "human-numeric-sort", "h", false, "human-numeric-sort")

    _ = rootCmd.Execute()
	fmt.Println(key.numericSort)
	// field -= 1
    // if err != nil {
    //     log.Fatal(fmt.Errorf("required argument missing: %v", err))
    // }

	// if len(os.Args) < 2 {
	// 	log.Fatal(fmt.Errorf("missing filename"))
    //     return
    // }

	// lines , err := GetLines(os.Args[1])
	// if err != nil {
	// 	log.Fatal(fmt.Errorf("error when getting lines: %v", err))
	// }
	
	// if len(lines[0]) > 0 && len(lines[0]) < field {
	// 	log.Fatal(fmt.Errorf("invalid argument k"))
	// }

	// a := []bool{*keys["numericSort"], *keys["monthSort"], *keys["humanNumericSort"]}
	// k := 0
	// for _ , key := range  a {
	// 	if key {
	// 		k += 1
	// 	}
	// }

	// if k > 1 { 
	// 	log.Fatal(fmt.Errorf("mutually exclusive flags"))
	// }

	// if *keys["numericSort"] {
	// 	SortNumeric(keys, lines)
	// } else if *keys["humanNumericSort"] {
	// 	SortHumanNumeric(keys, lines)
	// } else {
	// 	sort.Slice(lines, 
	// 		func(i, j int) bool {return lines[i][field] < lines[j][field]})
	// }

	// } else if *humanNumericSort {
	// 	SortHumanNumeric(keys, lines)
	// } else if *humanNumericSort {
	// 	SortMonth(keys, lines)
	// }

	// if *check {
	// 	Check(keys, lines)
	// }
	// fmt.Println(lines)
}
