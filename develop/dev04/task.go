package main

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

func Letters(word string) map[rune]struct{} {
	letters := make(map[rune]struct{})
	for _, l := range []rune(word) {
		letters[l] = struct{}{}
	}

	return letters
}

func Contains(words []string, word string) bool {
    for _, a := range words  {
        if a == word {
            return true
        }
    }
    return false
}

func Anagrams(words []string) *map[string][]string {
	anagrams := make(map[string][]string)
	for _, word := range words {
		hasSet := false
		word = strings.ToLower(word)
		
		for key := range anagrams {
			if reflect.DeepEqual(Letters(word), Letters(anagrams[key][0])) {
				hasSet = true
				if Contains(anagrams[key], word) == false {
					anagrams[key] = append(anagrams[key], word)
				}
			}
			sort.Strings(anagrams[key])
		}
		if hasSet == false{
			anagrams[word] = append(anagrams[word], word)
		}
	}
	return &anagrams
}

func main() {
	dict := []string{"Пятак", "пятка", "тяПка", "Листок", "слиток", "Столик", "123", "132"}
	fmt.Println(Anagrams(dict))
}
