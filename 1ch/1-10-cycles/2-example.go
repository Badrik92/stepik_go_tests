/* Требуется написать программу, при выполнении которой с клавиатуры
считываются два натуральных числа A и B
(каждое не более 100, A < B). Вывести сумму всех чисел от A до B  включительно.
Sample Input:
1 5
Sample Output:
15 */

package main

import "fmt"

func main() {
	var a, b int
	sum := 0
	fmt.Scan(&a, &b)
	if b > a && a < 100 && b < 100 {
		for i := 0; a <= b; i++ {
			sum += a
			a++
		}
		fmt.Println(sum)
	}
}
