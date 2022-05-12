package unpack

import (
	"fmt"
	"strconv"
	"unicode"
)

//Unpack function used for "a4bc2d5e" => "aaaabccddddde"
func Unpack(str string) (string, error) {
	rns := []rune(str)
	res := make([]rune, 0)

	shield := false
	i := 0
	for i < len(rns) {
		if string(rns[i]) == "\\" && !shield {
			shield = true
		} else if unicode.IsLetter(rns[i]) || shield {
			if i < len(rns)-1 && unicode.IsDigit(rns[i+1]) {
				n, _ := strconv.Atoi(string(rns[i+1]))
				for j := 0; j < n; j++ {
					res = append(res, rns[i])
				}
				i++
			} else {
				res = append(res, rns[i])
			}
			shield = false
		} else if !unicode.IsDigit(rns[i]) || (unicode.IsDigit(rns[i]) && !shield){
			return "", fmt.Errorf("incorrect input") 
		} 
		i++
	}
	return string(res), nil
}
