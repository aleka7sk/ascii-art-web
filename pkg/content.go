package pkg

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func CreateCharMap(file string) map[string][]string {
	var ct int
	var tmp []string
	var newcharset [][]string
	chars, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	data := make([]byte, 100000)
	n, err := chars.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	allchars := string(data[:n])
	charset := strings.Split(allchars, "\n")
	charmap := make(map[string][]string)
	charset = charset[1:]
	for _, item := range charset {
		ct++
		tmp = append(tmp, item)
		if ct == 8 {
			newcharset = append(newcharset, tmp)
			tmp = []string{}
		}
		if ct == 9 {
			tmp = []string{}
			ct = 0
			continue
		}
	}
	for i, char := range newcharset {
		charmap[string(rune(i+32))] = char
	}
	charmap["\n"] = make([]string, 8)
	return charmap
}

func CheckString(s string, charmap map[string][]string) bool {
	for _, ch := range s {
		if _, ok := charmap[string(ch)]; !ok {
			return false
		}
	}
	return true
}

func Run(text, font string) string {
	input := text
	input = strings.ReplaceAll(input, "\\n", "\n")
	input = strings.ReplaceAll(input, string(rune(13)), "\n")
	if input == "" {
		return ""
	}
	charmap := CreateCharMap("pkg/" + font + ".txt")
	if !CheckString(input, charmap) {
		fmt.Println("invalid Chars")
		return ""
	}
	inputlines := strings.Split(input, "\n")
	if input == "\n" {
		inputlines = []string{""}
	}
	res := ""
	for _, item := range inputlines {
		for i := 0; i < 8; i++ {
			for _, letter := range item {
				res += charmap[string(letter)][i]
			}
			res += "\n"
			if item == "" {
				break
			}

		}
	}

	return res
}
