package main

import (
	"bufio"
	"log"
	// "flag"
	"fmt"
	// "log"
	"os"
	"sort"
	"strconv"
	"strings"
	// log "github.com/sirupsen/logrus"
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

func main() {
    rootCmd := cobra.Command{}
	rootCmd.PersistentFlags().BoolP("help", "", false, "help for this command")
	rootCmd.Flags().IntVarP(&field, "key", "k", 1, "field")
	
	keys := make(map[string]*bool, 0)
	keys["numericSort"] = rootCmd.Flags().BoolP("numeric-sort", "n", false, "numeric-sort")
	keys["reverse"] = rootCmd.Flags().BoolP("reverse", "r", false, "reverse")
	keys["unique"] = rootCmd.Flags().BoolP("unique", "u", false, "unique")
	keys["monthSort"] = rootCmd.Flags().BoolP("month-sort", "M", false, "month-sort")
	keys["ignoreLeadingBlanks"] = rootCmd.Flags().BoolP("ignore-leading-blanks", "b", false, "ignore-leading-blanks")
	// check := rootCmd.Flags().BoolP("check", "c", false, "check")
	keys["humanNumericSort"] = rootCmd.Flags().BoolP("human-numeric-sort", "h", false, "human-numeric-sort")

    err := rootCmd.Execute()
	field -= 1
    if err != nil {
        log.Fatal(fmt.Errorf("required argument missing: %v", err))
    }

	if len(os.Args) < 2 {
		log.Fatal(fmt.Errorf("missing filename"))
        return
    }

	lines , err := GetLines(os.Args[1])
	if err != nil {
		log.Fatal(fmt.Errorf("error when getting lines: %v", err))
	}
	
	if len(lines[0]) > 0 && len(lines[0]) < field {
		log.Fatal(fmt.Errorf("invalid argument k"))
	}

	a := []bool{*keys["numericSort"], *keys["monthSort"], *keys["humanNumericSort"]}
	k := 0
	for _ , key := range  a {
		if key {
			k += 1
		}
	}

	if k > 1 { 
		log.Fatal(fmt.Errorf("mutually exclusive flags"))
	}

	if *keys["numericSort"] {
		SortNumeric(keys, lines)
	} else if *keys["humanNumericSort"] {
		SortHumanNumeric(keys, lines)
	} else {
		sort.Slice(lines, 
			func(i, j int) bool {return lines[i][field] < lines[j][field]})
	}

	// } else if *humanNumericSort {
	// 	SortHumanNumeric(keys, lines)
	// } else if *humanNumericSort {
	// 	SortMonth(keys, lines)
	// }

	// if *check {
	// 	Check(keys, lines)
	// }
	fmt.Println(lines)
}
