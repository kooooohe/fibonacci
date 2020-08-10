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
	//println(fib(8))
	fmt.Println(calcFib(1000))
}

// (((1+√5)/2)^n - ((1-√5)/2)^n) / √5
func calcFib(n int) *big.Int {
	// (1+√5)/2
	a := big.NewRat(1, 2)
	b := big.NewRat(1, 2)
	x := fibNum{a, b}

	// (1-√5)/2
	c := big.NewRat(1, 2)
	d := big.NewRat(-1, 2)
	y := fibNum{c, d}

	// (1/√5)
	e := big.NewRat(0, 999)
	f := big.NewRat(1, 5)
	z := fibNum{e, f}

	tmp := fibMin(fibExp(x, n), fibExp(y, n))

	// *(1/√5)
	r := fibMul(tmp, z)
	return r.num.Num()

}

func fibExp(t fibNum, n int) fibNum {

	r := t

	// 1
	a := fibNum{num: big.NewRat(1, 1), squareRoot: big.NewRat(0, 999)}

	for n > 1 {
		if n%2 == 1 {
			a = fibMul(a, r)
		}
		r = fibMul(r, r)
		n = n / 2
	}

	return fibMul(r, a)
}

/*
func fibExp(t fibNum, n int) (r fibNum) {
	tmp := t
	for i := 1; i < n; i++ {
		tmp = fibMul(t, tmp)
	}
	r = tmp
	return
}
*/

func fibMin(t1, t2 fibNum) (r fibNum) {
	r1 := new(big.Rat).Add(
		t1.num,
		new(big.Rat).Mul(t2.num, big.NewRat(-1, 1)),
	)

	r2 := new(big.Rat).Add(
		t1.squareRoot,
		new(big.Rat).Mul(t2.squareRoot, big.NewRat(-1, 1)),
	)

	r.num = r1
	r.squareRoot = r2
	return
}

//(a+b√5)(c+d√5) = (ac+5bd)+(ad+bc)√5
func fibMul(t1, t2 fibNum) (r fibNum) {
	//(ac+5bd)
	var r1, r2 *big.Rat
	tmp1 := new(big.Rat).Mul(t1.num, t2.num)
	tmp2 := new(big.Rat).Mul(
		new(big.Rat).Mul(t1.squareRoot, t2.squareRoot),
		big.NewRat(5, 1),
	)
	r1 = new(big.Rat).Add(tmp1, tmp2)

	//(ad+bc)√5
	tmp3 := new(big.Rat).Mul(t1.num, t2.squareRoot)
	tmp4 := new(big.Rat).Mul(t1.squareRoot, t2.num)
	r2 = new(big.Rat).Add(tmp3, tmp4)

	//(ac+5bd)+(ad+bc)√5
	//rrr := new(big.Rat).Add(r3, rr3)

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
