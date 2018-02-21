package main

import (
	"math/rand"
	"time"
)

func gcd(x, y int) int {
	mod := x % y
	for mod > 0 {
		x = y
		y = mod
		mod = x % y
	}
	return y
}

func random_choice(s []int) int {
	return s[rand.Int()%len(s)]
}

func lcm(x, y int) int {
	return x * y / gcd(x, y)
}

type RSA struct {
	n int
	e int
	d int
}

func newRSA(p, q int) *RSA {
	rsa := RSA{}
	rsa.n = p * q
	phi := toient(p, q)
	rsa.e = find_e(phi)
	rsa.d = find_d(rsa.e, phi)
	return &rsa
}

func toient(x, y int) int {
	return lcm(x-1, y-1)
}

func find_e(phi int) int {
	candidates := make([]int, 0, 16)
	for e := 2; e < phi; e++ {
		if gcd(e, phi) == 1 {
			candidates = append(candidates, e)
		}
	}
	return random_choice(candidates)
}

func find_d(e, phi int) int {
	candidates := make([]int, 0, 16)
	for d := 2; d < phi; d++ {
		if (e*d)%phi == 1 {
			candidates = append(candidates, d)
		}
	}
	return random_choice(candidates)
}

func modular(a, b, c int) int {
	acc := 1
	for i := 0; i < b; i++ {
		acc = (acc * a) % c
	}
	return acc
}

func (rsa RSA) encrypt(m int) int {
	return modular(m, rsa.e, rsa.n)
}

func (rsa RSA) decrypt(c int) int {
	return modular(c, rsa.d, rsa.n)
}

func main() {
	rand.Seed(time.Now().Unix())

	m := 91
	rsa := newRSA(3, 997)

	c := rsa.encrypt(m)
	result := rsa.decrypt(c)

	println(result)
}
