/*
Формат входных данных

Вводится непустая последовательность натуральных чисел, оканчивающаяся числом 0
(само число 0 в последовательность не входит, а служит как признак ее окончания).

Формат выходных данных

Выведите ответ на задачу.

Sample Input:

1
3
3
1
0
Sample Output:

2 */


package main 

import "fmt"

func main() {
	var n, sum, count int 
	
	for fmt.Scan(&n); n !=0; fmt.Scan(&n) {
		if n > sum{
			count = 1 // Присваиваем значение, что б при 0 count не плюсовался
			sum = n
		} else if n == sum {
			count++
		} 
	}
	fmt.Println(count)
	
}