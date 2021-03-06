/* Дано неотрицательное целое число. Найдите и выведите первую цифру числа. 

Формат входных данных
На вход дается натуральное число, не превосходящее 10000.

Формат выходных данных
Выведите одно целое число - первую цифру заданного числа.

Sample Input:

1234
Sample Output:

1 */

package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a)
	if a < 10 {
		b = a / 1
	} else if a < 100 {
		b = a / 10
	} else if a < 1000 {
		b = a / 100
	} else if a < 10000 {
		b = a / 1000
	} else if a == 10000 {
		b = a / 10000
	}
	fmt.Println(b)
}

