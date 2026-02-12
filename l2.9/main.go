package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"
)

var (
	errInvalidString = errors.New("invalid string")
	errConvertAtoi   = errors.New("cannot convert string to int")
)

func unpackString(s string) (string, error) {
	var res strings.Builder
	var errInvalidStringFlag = true
	runesToProcess := []rune(s)

	// если строка пустая, возвращаем ее
	if len(s) == 0 {
		return s, nil
	}

	// проходимся по рунам строки
	for i := 0; i < len(runesToProcess)-1; i++ {
		// проверяем есть ли в строке хотя бы 1 буква
		if unicode.IsLetter(runesToProcess[len(runesToProcess)-1]) || unicode.IsLetter(runesToProcess[i]) {
			errInvalidStringFlag = false
		}

		// если элемент не равен \(потому что \ делает число строкой по тз) и следующий элемент за ним число,
		// то умножает эту руну на число и добавляем в билдер
		if runesToProcess[i] != '\\' && unicode.IsDigit(runesToProcess[i+1]) {
			count, err := strconv.Atoi(string(runesToProcess[i+1]))
			if err != nil {
				return "", errConvertAtoi
			}

			res.WriteString(strings.Repeat(string(runesToProcess[i]), count))

			// перескакиваем на 2 символа, так как использованные цифры в итоговой строке
			// не должны быть
			i++
			// во всех остальных случаях, если руна не равна \, то просто добавляем ее в билдер
		} else if runesToProcess[i] != '\\' {
			res.WriteRune(runesToProcess[i])
		}
	}

	// если строка не содержит ни одной буквы, возвращаем ошибку
	if errInvalidStringFlag {
		return "", errInvalidString
	}

	// во избежание out of range в цикле мы не доходим до последнего элемента
	// и вот если последний элемент является числом, а предпоследний - \,
	// то возвраащем просто строку, во всех остальных случаях
	// добавляем к строке последний символ
	// это единственный случай, когда мы не добавляем последний символ строки в итоговую строку
	// делается потому что если предпоследний символ не равен \, значит
	// число мы использовали, умножим его на символ перед ним, а
	// использованные числа мы в итоговую строку не добавляем
	if unicode.IsDigit(runesToProcess[len(runesToProcess)-1]) && runesToProcess[len(runesToProcess)-2] != '\\' {
		return res.String(), nil
	}

	res.WriteRune(runesToProcess[len(runesToProcess)-1])

	return res.String(), nil

	// golangci-lint run выводит 2 предупреждения::
	// 1. в тестах дублируется цикл 2 разных функциях, но делать из них 1 функцию будет
	// менее неудобным для проведения тестов
	// 2. цикломатическая сложность > 8, но я не знаю где можно упростить
}

func main() {
	s := "a4bc2d5e"
	res, err := unpackString(s)
	if err != nil {
		log.Fatal("failed to unpack string", err)
	}
	fmt.Println(res)
}
