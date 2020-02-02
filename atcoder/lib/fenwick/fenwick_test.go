package fenwick

import (
	"fmt"
	"testing"
)

func ExampleBIT() {
	bit := NewBIT(5)
	for i := 0; i < 5; i++ {
		bit.Add(i, i) //0,1,2,3,4
	}
	fmt.Println(bit.Accumulate(4))
	//Output:
	//10
	fmt.Println(bit.Accumulate(3))
	//6
}

func TestAccumulate(t *testing.T) {
	bit := NewBIT(5)
	for i := 0; i < 5; i++ {
		bit.Add(i, i+1) //1,2,3,4,5
	}
	if bit.Accumulate(3) != 1+2+3+4 {
		t.Error("1+2+3+4 == 10")
	}
}

func TestAccumulateForWrongInput(t *testing.T) {
	bit := NewBIT(5)
	for i := 0; i < 5; i++ {
		bit.Add(i, i+1) //1,2,3,4,5
	}
	if bit.Accumulate(6) != 1+2+3+4+5 {
		t.Error("データサイズより大きい入力に対しては、全データの和を返す。")
	}
}

func TestAccumulateForNegativeInput(t *testing.T) {
	bit := NewBIT(5)
	for i := 0; i < 5; i++ {
		bit.Add(i, i+1) //1,2,3,4,5
	}
	if bit.Accumulate(-10) != 0 {
		t.Error("負の入力に対しては、0を返す。")
	}
}

func Test(t *testing.T) {
	bit := NewBIT(10)
	for i := 0; i < 10; i++ {
		bit.Add(i, 2+i) //2,3,4,5,6,7,...,11
	}
	fmt.Println(bit.Sum(1, 4))
	if bit.Sum(1, 4) != 3+4+5+6 {
		t.Errorf("Expect %v == %v", bit.Sum(1, 4), 3+4+5+6)
	}
}
