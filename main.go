package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calculateAr(num1, opp, num2 string) int {
	var result int
	num11, err := strconv.Atoi(num1)
	if err != nil {
		panic(err)
	}
	num22, err := strconv.Atoi(num2)
	if err != nil {
		panic(err)
	}
	if num11 == 0 || num22 == 0 {
		panic("Операнды не могут быть равны нулю")
	}
	if num11 > 10 || num22 > 10 {
		panic("Один или два операнда больше 10")
	}
	switch {
	case opp == "+":
		result = num11 + num22
	case opp == "-":
		result = num11 - num22
	case opp == "*":
		result = num11 * num22

	case opp == "/":
		result = num11 / num22
	default:

		panic("не подходящий оператор")
	}

	return result
}

func calculateRom(opp string, res, res1 int) int {
	var result int
	if res == 1000 || res1 == 1000 {

		return 1000
	}
	switch {

	case opp == "+":
		result = res + res1
	case opp == "-":
		result = res - res1
	case opp == "*":
		result = res * res1

	case opp == "/":
		result = res / res1
	default:
		panic("не подходящий оператор")
	}

	return result
}

func main() {
	for {

		fmt.Print("Введите ваше выражение:")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// Парсим введенное выражение
		parts := strings.Split(input, " ")
		if len(parts) != 3 {
			panic("некорректное выражение. Используйте формат: число оператор число")
		}
		num1 := parts[0]
		opp := parts[1]
		num2 := parts[2]
		var nums = num1 + num2
		containsAr := strings.ContainsAny(nums, "12345678910")
		containsRom := strings.ContainsAny(nums, "IVX")

		if containsAr && containsRom {
			fmt.Println("Введите либо римские - либо арабские")
			break
		} else if containsAr {
			resAr := calculateAr(num1, opp, num2)
			fmt.Println(resAr)
		} else if containsRom {
			res := Decode(num1)
			res2 := Decode2(num2)
			resRom := calculateRom(opp, res, res2)
			normRes := NormalResRom(resRom)
			if resRom <= 0 {
				fmt.Println("Результат не может быть отрицательным")

			} else if resRom == 1002 {
				fmt.Println("не верный оператор")
				break

			}
			fmt.Println(normRes)

		}
	}
}
func Decode(num1 string) int {
	translateRoman := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	var decNum, tmpNum int

	for i := len(num1) - 1; i >= 0; i-- {

		if i >= 3 && num1 != "VIII" {
			panic("Неправильная римская цифра")
		}

		romanDigit := num1[i]

		decDigit := translateRoman[romanDigit]
		if decDigit < tmpNum {
			decNum -= decDigit
		} else {
			decNum += decDigit
			tmpNum = decDigit
		}
		if decNum > 10 {

			return 1000
		}
	}
	return decNum
}
func Decode2(num2 string) int {
	translateRoman := map[byte]int{'I': 1, 'V': 5, 'X': 10}
	var decNum, tmpNum int
	for i := len(num2) - 1; i >= 0; i-- {
		if i >= 3 && num2 != "VIII" {
			panic("Неправильная римская цифра")
		}
		romanDigit := num2[i]
		decDigit := translateRoman[romanDigit]

		if decDigit < tmpNum {
			decNum -= decDigit
		} else {
			decNum += decDigit
			tmpNum = decDigit
		}
	}
	if decNum > 10 {

		return 1000
	}
	return decNum
}

var romanMap = []struct {
	decVal int
	symbol string
}{
	{100, "C"}, {90, "XC"}, {50, "L"}, {40, "XL"},
	{10, "X"}, {9, "IX"}, {5, "V"}, {4, "IV"}, {1, "I"},
}

func NormalResRom(resRom int) string {

	if resRom == 0 {

	}
	for _, pair := range romanMap {
		if resRom >= pair.decVal {

			return pair.symbol + NormalResRom(resRom-pair.decVal)

		}
	}
	return ""
}
