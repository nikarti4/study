package main

import (
	"fmt"
	"strings"
)

// регистронезависимая проверка уникальных символов в строке
func uniqueChar(s string) bool {
	// регистронезависимость
	tempS := strings.ToUpper(s)

	// мапа с ключами буквами
	// rune - т.к. далее for range
	m := make(map[rune]int)

	for _, c := range tempS {
		//проверяем, что символ уникальный
		_, uniq := m[c]
		// если нет - то вернуть false
		// иначе - добавить символ в мапу
		if uniq {
			return false
		} else {
			m[c]++
		}
	}

	return true
}

func main() {
	s1 := "abcd"
	s2 := "abCdefAaf"
	s3 := "aabcd"
	s4 := "abcdAfg"
	s5 := "qWeRtY"

	fmt.Println(s1, " - ", uniqueChar(s1))
	fmt.Println(s2, " - ", uniqueChar(s2))
	fmt.Println(s3, " - ", uniqueChar(s3))
	fmt.Println(s4, " - ", uniqueChar(s4))
	fmt.Println(s5, " - ", uniqueChar(s5))
}
