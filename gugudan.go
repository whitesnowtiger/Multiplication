package main

import "fmt"

func main() {

	for dan := 1; dan <= 9; dan++ {

		fmt.Printf("%d 단\n", dan)

		for j := 1; j <= 9; j++ {

			if dan == 3 && j == 2 {
				continue
			}

			fmt.Printf("%d * %d = %d\n", dan, j, dan*j)

		}
		fmt.Println()
	}

	for i := 0; i < 5; i++ {

		for j := 0; j < i+1; j++ {

			fmt.Print("*")
		}
		fmt.Println()
	}
	fmt.Println()

	for i := 0; i < 5; i++ {

		for j := 0; j < 5-i; j++ {

			fmt.Print("*")

		}
		fmt.Println()
	}

	fmt.Println()

	for i := 0; i < 3; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	for i := 0; i < 2; i++ {
		for j := 0; j < 2-i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	fmt.Println()

	for i := 0; i < 4; i++ {

		for j := 0; j < 3-i; j++ {
			fmt.Print(" ")
		}
		for k := 0; k < 2*i+1; k++ {
			fmt.Print("*")

		}
		fmt.Println()
	}
	fmt.Println()

	for i := 0; i < 4; i++ {
		for j := 0; j < 3-i; j++ {
			fmt.Print(" ")
		}
		for k := 0; k < 2*i+1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Print(" ")
		}
		for k := 0; k < 5-2*i; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

}
