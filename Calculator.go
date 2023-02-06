package main

import (
	"fmt"
	"strconv"
)

func ones(o int) string {

	var s string

	if o == 1 {
		s = "I"
	}
	if o == 2 {
		s = "II"
	}
	if o == 3 {
		s = "III"
	}
	if o == 4 {
		s = "IV"
	}
	if o == 5 {
		s = "V"
	}
	if o == 6 {
		s = "VI"
	}
	if o == 7 {
		s = "VII"
	}
	if o == 8 {
		s = "VIII"
	}
	if o == 9 {
		s = "IX"
	}
	if o == 10 {
		s = "X"
	}

	return s
}

func conv_roman(c int) (s string) {

	fmt.Println("Rome")
	if c > 10 {

		d := c / 10
		o := c % 10

		if d*10 < 40 {

			for i := 1; i <= d; i++ {

				s = s + "X"

			}

			s = s + ones(o)

		} else if d*10 >= 40 && d*10 < 50 {

			s = "XL" + ones(o)

		} else if d*10 >= 50 && d*10 < 90 {

			s = "L"

			for i := 0; i < ((d*10)-50)/10; i++ {

				s = s + "X"

			}

			s = s + ones(o)

		} else if d*10 >= 90 && d*10 < 100 {

			s = "XC"
			s = s + ones(0)

		} else if d*10 == 100 {
			s = "C"
		}

	} else {

		s = ones(c)
	}
	return s
}

func rim2dig(s1 string) (a int) {

	if s1 == "I" {
		a = 1
	}
	if s1 == "II" {
		a = 2
	}
	if s1 == "III" {
		a = 3
	}
	if s1 == "IV" {
		a = 4
	}
	if s1 == "V" {
		a = 5
	}
	if s1 == "VI" {
		a = 6
	}
	if s1 == "VII" {
		a = 7
	}
	if s1 == "VIII" {
		a = 8
	}
	if s1 == "IX" {
		a = 9
	}
	if s1 == "X" {
		a = 10
	}

	return a
}

func getdigstrings(s1 string, s2 string, c string, i int) (string, string) {

	for j := 0; j < i; j++ {
		s1 = s1 + string(c[j])
	}

	for j := i + 1; j < len(c); j++ {
		s2 = s2 + string(c[j])
	}

	return s1, s2
}

func digits() (int, int, int, int) {

	var c string
	fmt.Scanf("%s\n", &c)

	s1 := ""
	s2 := ""
	ind := 0
	var flag int

	for i := 0; i < len(c); i++ {

		c1 := string(c[i])

		if c1 == "+" {

			s1, s2 = getdigstrings(s1, s2, c, i)
			ind = 1

		} else if c1 == "-" {

			s1, s2 = getdigstrings(s1, s2, c, i)
			ind = 2

		} else if c1 == "*" {

			s1, s2 = getdigstrings(s1, s2, c, i)
			ind = 3

		} else if c1 == "/" {

			s1, s2 = getdigstrings(s1, s2, c, i)
			ind = 4

		}
	}

	a, e1 := strconv.Atoi(s1)
	b, e2 := strconv.Atoi(s2)

	if e1 == nil && e2 == nil {

		flag = 0

		fmt.Println(a)
		fmt.Println(b)
		fmt.Println(ind)

	} else {

		fmt.Println("Rome")
		a = rim2dig(s1)
		fmt.Println(a)
		b = rim2dig(s2)
		fmt.Println(b)
		flag = 1

	}

	return a, b, ind, flag
}

func main() {

	a, b, ind, flag := digits()

	c := 0

	if ind == 1 {

		c = a + b

	} else if ind == 2 {

		c = a - b

	} else if ind == 3 {

		c = a * b

	} else if ind == 4 {

		c = a / b
	}

	if flag == 0 {

		fmt.Println("The answer is: ", c)

	}

	if flag == 1 {

		if c > 0 {

			s := conv_roman(c)
			fmt.Println("The answer is: ", s)

		} else {

			fmt.Println("ERROR: Answer equals or is less then 0 or Input is WRONG")

		}

	}

}
