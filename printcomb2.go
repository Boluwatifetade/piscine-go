package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for num1 := 0; num1 <= 98; num1++ {
		for num2 := num1 + 1; num2 <= 99; num2++ {

			z01.PrintRune('0' + (num1 / 10))
			z01.PrintRune('0' + (num1 % 10))

			z01.PrintRune(' ')

			z01.PrintRune('0' + (num2 / 10))
			z01.PrintRune('0' + (num2 % 10))

			if num1 != 98 || num2 != 99 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
	z01.PrintRune('\n')
}
