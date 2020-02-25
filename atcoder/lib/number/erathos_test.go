package number

import "fmt"

func ExampleErathos() {
	isPrime := Erathos(100)
	fmt.Println(len(isPrime))
	fmt.Println(isPrime[11]) // 11 is a prime number
	fmt.Println(isPrime[57]) // 57 = 3 * 19
	fmt.Println(isPrime[91]) // 91 = 7 * 13
	fmt.Println(isPrime[97]) // 97 is a prime number
	//Output:
	//101
	//true
	//false
	//false
	//true
}
