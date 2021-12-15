/* Напишите программу, которая в
последовательности чисел находит сумму двузначных чисел, кратных 8.
Программа в первой строке получает на вход число n - количество чисел в последовательности,
во второй строке -- n чисел, входящих в данную последовательность.
Sample Input:
5
38 24 800 8 16
Sample Output:
40
1. Считываем из консоли переменную "a" - количество итераций цикла

2. В цикле от нуля до "a" с шагом 1 считываем значения переменной "b"

3. Если переменная "b" соответствует условию: 9 < b < 100, то она двухзначная

4. Если переменная "b" делится на 8 без остатка, то она кратна восьми
условия можно объединить через &&), и мы прибавляем к "sum" переменную "b"

5. После того, как пройдут все итерации цикла - выводим "sum" */
////////////////////////////////////////////////////////////////

package main

import "fmt"

func main() {
	var a, b, sum int

	fmt.Scan(&a)
	for i := 0; i < a; i++ {
		fmt.Scan(&b)
		if b > 9 && b < 100 && b % 8 == 0{
			sum += b
		}
	}
	fmt.Println(sum)
	fmt.Println("Great")
	fmt.Println("Great")
	fmt.Println("Great")
  }