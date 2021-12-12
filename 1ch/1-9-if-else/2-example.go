/* По данному трехзначному числу определите, все ли его цифры различны.

Формат входных данных
На вход подается одно натуральное трехзначное число.

Формат выходных данных
Выведите "YES", если все цифры числа различны, в противном случае - "NO".

Sample Input 1:

237
Sample Output 1:

YES
Sample Input 2:

117
Sample Output 2:

NO */

package main

import "fmt"

func main() {
	var a, b, c, i int
	fmt.Scan(&a)
	i = a / 100
	b = a % 100 / 10
	c = a % 10

	if(i != b && i != c && c != b) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
