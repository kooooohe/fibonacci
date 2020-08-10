package main

import (
	"fmt"
	"math/big"
)

type fibNum struct {
	num        *big.Rat
	squareRoot *big.Rat
}

func main() {
	//println(fib(6))
	tmp()
}

func tmp() {
	// (1+√5)/2
	a := big.NewRat(1, 2)
	b := big.NewRat(1, 2)
	x := fibNum{a, b}

	// (1-√5)/2
	c := big.NewRat(1, 2)
	d := big.NewRat(-1, 2)
	y := fibNum{c, d}

	r := new(big.Rat).Mul(x.num, y.num)
	fmt.Println(r)

	//fibMul(x, x)
	tmp := fibMin(fibExp(x, 7), fibExp(y, 7))
	fmt.Println(tmp.num)
	fmt.Println(tmp.squareRoot)

	//divideNum := fibNum{nil, big.NewRat(1, 5)}
}

func fibExp(t fibNum, n int) (r fibNum) {
	tmp := t
	for i := 1; i < n; i++ {
		tmp = fibMul(t, tmp)
	}
	r = tmp
	fmt.Println(tmp.num)
	fmt.Println(tmp.squareRoot)
	return
}
func fibMin(t1, t2 fibNum) (r fibNum) {
	r1 := new(big.Rat).Add(
		t1.num,
		new(big.Rat).Mul(t2.num, big.NewRat(-1, 1)),
	)

	r2 := new(big.Rat).Add(
		t1.squareRoot,
		new(big.Rat).Mul(t2.squareRoot, big.NewRat(-1, 1)),
	)
	fmt.Println(r1)
	fmt.Println(r2)

	r.num = r1
	r.squareRoot = r2
	return
}

func fibMul(t1, t2 fibNum) (r fibNum) {
	//(ac+5bd)
	var r1, r2 *big.Rat
	tmp1 := new(big.Rat).Mul(t1.num, t2.num)
	tmp2 := new(big.Rat).Mul(
		new(big.Rat).Mul(t1.squareRoot, t2.squareRoot),
		big.NewRat(5, 1),
	)
	r1 = new(big.Rat).Add(tmp1, tmp2)
	//fmt.Println(r1)

	//(ad+bc)√5
	tmp1 = new(big.Rat).Mul(t1.num, t2.squareRoot)
	tmp2 = new(big.Rat).Mul(t1.squareRoot, t2.num)
	r2 = new(big.Rat).Add(tmp1, tmp2)
	//fmt.Println(r2)

	//(ac+5bd)+(ad+bc)√5
	//rrr := new(big.Rat).Add(r3, rr3)
	//fmt.Println(rrr)
	r.num = r1
	r.squareRoot = r2
	return
}

func fib(num int) int {
	if num <= 1 {
		return num
	}
	return fib(num-2) + fib(num-1)
}
