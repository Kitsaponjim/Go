package main

import (
	"fmt"
	"strconv"
)

type Company struct {
	Name    string
	Address string
	tel     string
}

func main() {
	num0()
	num1()
	num1_2()
	num2()
	num3()
	num3_1()
	num4()
	num4_1()
	num5()
	num6()
	Extra()

	var a [2]string
	a[0] = "Hello"
	a[1] = "world"
	fmt.Println(a)
}

func num0() {
	fmt.Println("ข้อ 0-------------------------------------------")
	i := 2
	if i == 0 {
		fmt.Println("Zero")
	} else if i == 1 {
		fmt.Println("One")
	} else if i == 2 {
		fmt.Println("Two")
	} else if i == 3 {
		fmt.Println("Three")
	} else {
		fmt.Println("You i not in case")
	}
	fmt.Printf("\n")
}

func num1() {
	fmt.Println("ข้อ 1------------------------------------------")
	count := 0
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Printf("%d, ", i)
			count++
		}
	}
	fmt.Println("\nเลขที่หาร 3 ลงตัวมีจำนวนเท่ากับ", count)
	fmt.Printf("\n")
}

func num1_2() {
	fmt.Println("ข้อ 1.2------------------------------------------")
	pow(5, 2)
	fmt.Printf("\n")
}

func pow(x int, y int) {
	result := 1
	for i := 0; i < y; i++ {
		result *= x
	}
	fmt.Println(result)
}

func num2() {
	fmt.Println("ข้อ 2------------------------------------------")
	x := []int{48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 9, 17}
	FindMaxMin(x)

}

func FindMaxMin(num []int) {
	max := 0
	min := num[0]

	for i := 0; i < len(num); i++ {
		if max < num[i] {
			max = num[i]
		}
		if min > num[i] {
			min = num[i]
		}
	}
	fmt.Println("Max number is ", max)
	fmt.Println("Min number is ", min)
	fmt.Printf("\n")
}

func num3() {
	fmt.Println("ข้อ 3------------------------------------------")
	var sum int
	count := 0
	for i := 0; i < 1000; i++ {
		sum = i + 1
		nts := strconv.Itoa(sum)
		for _, v := range nts {
			if string(v) == "9" {
				count++
			}
		}
	}
	fmt.Println(count)
	fmt.Printf("\n")
}

func num3_1() {
	fmt.Println("ข้อ 3.1------------------------------------------")
	Findnine(10000)
}

func Findnine(num int) {
	var sum int
	count := 0
	for i := 0; i < num; i++ {
		sum = i + 1
		nts := strconv.Itoa(sum)
		for _, v := range nts {
			if string(v) == "9" {
				count++
			}
		}
	}
	fmt.Println(count)
	fmt.Printf("\n")
}

func num4() {
	fmt.Println("ข้อ 4------------------------------------------")
	myWords := "AW SOME GO!"
	fmt.Println(myWords)
	for _, v := range myWords {
		// fmt.Println(string(v))
		if string(v) != " " {
			fmt.Printf("%s", string(v))
		}

	}
	fmt.Println("\n")
}

func num4_1() {
	fmt.Println("ข้อ 4.1------------------------------------------")
	cutSpace(" I N E T ")
}

func cutSpace(Words string) {
	fmt.Println(Words)
	for _, v := range Words {
		// fmt.Println(string(v))
		if string(v) != " " {
			fmt.Printf("%s", string(v))
		}
	}
	fmt.Println("\n")
}

func num5() {
	fmt.Println("ข้อ 5------------------------------------------")
	data := map[string]map[string]string{
		"emp_01": {
			"name":    "Jim",
			"age":     "22",
			"address": "Home",
		},
		"emp_02": {
			"name":    "Frame",
			"age":     "23",
			"address": "Market",
		},
		"emp_03": {
			"name":    "Tee",
			"age":     "23",
			"address": "Temple",
		},
		"emp_04": {
			"name":    "Make",
			"age":     "23",
			"address": "sea",
		},
	}

	fmt.Println("Result")
	for _, empData := range data {
		fmt.Printf("Name-: %s (Age:%s)\n", empData["name"], empData["age"])
		fmt.Printf("Address-: %s\n", empData["address"])
		fmt.Printf("\n")
	}
}

func num6() {
	fmt.Println("ข้อ Extra------------------------------------------")
	company := new(Company)
	company2 := new(Company)
	company3 := new(Company)

	company.Name = "Pizza huta"
	company.Address = "Big C"
	company.tel = "1150"

	company2.Name = "Pizza company"
	company2.Address = "Big C"
	company2.tel = "1112"

	company3.Name = "Pizza copy"
	company3.Address = "Big C"
	company3.tel = "1250"

	fmt.Printf("Company-: Structs")
	fmt.Printf("Company-: NAME = %s, ADDRESS = %s, TEL = %s\n", company.Name, company.Address, company.tel)
	fmt.Printf("Company-: NAME = %s, ADDRESS = %s, TEL = %s\n", company2.Name, company2.Address, company2.tel)
	fmt.Printf("Company-: NAME = %s, ADDRESS = %s, TEL = %s\n", company3.Name, company3.Address, company3.tel)
	fmt.Printf("\n")

}

func Extra() {
	fmt.Println("ข้อ Extra------------------------------------------")
	for i := 1; i <= 5; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}
}
