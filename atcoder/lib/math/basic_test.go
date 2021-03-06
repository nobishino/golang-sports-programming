package math

import (
	"fmt"
	"testing"
)

func ExampleMinInt() {
	fmt.Println(MinInt(3, 4))
	//Output:
	//3
}

func ExampleMaxInt() {
	fmt.Println(MaxInt(3, 4))
	//Output:
	//4
}

func ExampleModPow() {
	fmt.Println(ModPow(5, 7, 13))
	fmt.Println(ModPow(5, 7, 2))
	fmt.Println(ModPow(5, 7, 1))
	//Output:
	//8
	//1
	//78125
}

func ExampleAbsInt() {
	fmt.Println(AbsInt(3))
	fmt.Println(AbsInt(-4))
	fmt.Println(AbsInt(0))
	//Output:
	//3
	//4
	//0
}

func ExampleGcd() {
	fmt.Println(Gcd(28, 36))
	fmt.Println(Gcd(28, 36, 6))
	//Output:
	//4
	//2
}

func ExampleLcm() {
	fmt.Println(Lcm(28, 36))
	fmt.Println(Lcm(28, 36, 13))
	//Output:
	//252
	//3276
}

func TestModPow1(t *testing.T) {
	if ModPow(3, 400, 1000000007) != 978888738 {
		t.Fatalf("ModPow(3,400,1000000007) should be %d, actual value = %d", 978888738, ModPow(3, 400, 1000000007))
	}
}
func TestModPow2(t *testing.T) {
	if ModPow(1234567, 9876543210, 1000000007) != 598562710 {
		t.Fatal("ModPow() does not work correctly\n")
	}
}

func BenchmarkModPow(b *testing.B) {
	ModPow(7891, 123456789, 1000000007)
}
