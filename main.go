package main

import (
	"fmt"
	"strconv"
)

func Prima(angka int) bool {
	if angka <= 1 {
		return false
	}

	for i := 2; i*i <= angka; i++ {
		if angka%i == 0 {
			return false
		}
	}
	return true
}

func CekNol(angka string) bool {
	return angka[0] == '0' || angka[1] == '0' || angka[2] == '0' || angka[3] == '0'
}

func AmbilAngka(angka int, digit int) int {

	angkaStr := strconv.Itoa(angka)
	if digit == 3 && len(angkaStr) >= 3 {
		angka, _ := strconv.Atoi(angkaStr[:3])
		return angka
	} else if digit == 2 && len(angkaStr) >= 2 {
		angka, _ := strconv.Atoi(angkaStr[len(angkaStr)-2:])
		return angka
	}
	return angka
}

func Cek3DigitTerakhir(angka string) bool {
	if angka[1] == angka[2] && angka[2] == angka[3] {
		return true
	} else {
		return false
	}
}

func main() {
	array := []string{"3137", "1367", "2333", "2001"}

	for _, input := range array {
		angka, _ := strconv.Atoi(input)

		if Prima(angka) && !CekNol(input) {
			if Prima(AmbilAngka(angka, 3)) && !Cek3DigitTerakhir(strconv.Itoa(angka)) {
				fmt.Println("CENTER")
			} else if Cek3DigitTerakhir(strconv.Itoa(angka)) {
				fmt.Println("RIGHT")
			} else if Prima(AmbilAngka(angka, 2)) {
				fmt.Println("LEFT")

			} else {
				fmt.Println("DEAD")
			}
		} else {
			fmt.Println("DEAD")
		}
	}

}
