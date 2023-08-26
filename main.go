package main

import (
	"fmt"
	"strings"
)

func main() {

	input := []string{"{(<>)}{[]}()", "{()}(])"}

	for _, inp := range input {

		if isBalanced(inp) {
			fmt.Println("Скобки сбалансированы :)")
		} else {
			fmt.Println("Скобки несбалансированы :(")
		}

		fmt.Println()
	}
}

func isBalanced(str string) bool {

	screen := make([][]rune, len(str)) // Создаём матрицу экрана
	for i := range screen {
		screen[i] = make([]rune, len(str))
	}

	brackets := map[rune]rune{')': '(', '}': '{', ']': '[', '>': '<'}

	x := 0

	for y, chr := range str {

		switch chr {

		case '(', '{', '[', '<':

			screen[x][y] = chr // Записываем текущую открывающую скобку в матрицу экрана

			fmt.Print(strings.Repeat("\t", x))
			fmt.Println(string(chr))

			x++ // Сдвигаем курсор вправо

		default:

			x-- // Сдвигаем курсор влево

			if x < 0 { // Проверяем выход за экран
				return false
			}

			fmt.Print(strings.Repeat("\t", x))
			fmt.Println(string(chr))

			if !checkTopBracket(brackets[chr], x, y, screen) { // Проверяем, есть ли над данной закрывающей скобкой парная открывающая
				return false
			}

			screen[x][y] = chr // Записываем текущую закрывающую скобку в матрицу экрана
		}

	}

	if x != 0 { // Проверяем, вернулся ли курсор в нулевую позицию
		return false
	}

	return true

}

func checkTopBracket(str rune, x int, y int, screen [][]rune) bool {

	for i := y; i >= 0; i-- { // Двигаемся по столбцу вверх от закрывающей скобки

		if screen[x][i] == '\x00' { // Пропускаем пустые ячейки
			continue
		}

		if screen[x][i] == str { // Проверяем, является ли скобка парной
			return true
		} else {
			return false
		}

	}

	return false // Никакие скобки над закрывающей скобкой не найдены

}
