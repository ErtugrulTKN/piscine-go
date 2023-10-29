package piscine

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

func QuadCheck() {
	var x, y int
	var counter int
	stdin, _ := io.ReadAll(os.Stdin)
	input := string(stdin)

	for i, s := range input {
		if s == '\n' {
			y++
		}
		if y == 0 {
			x = i + 1
		}
	}

	if input == QuadA(x, y) {
		PrintAns(x, y, &counter, 'A')
	}
	if input == QuadB(x, y) {
		PrintAns(x, y, &counter, 'B')
	}
	if input == QuadC(x, y) {
		PrintAns(x, y, &counter, 'C')
	}
	if input == QuadD(x, y) {
		PrintAns(x, y, &counter, 'D')
	}
	if input == QuadE(x, y) {
		PrintAns(x, y, &counter, 'E')
	}
	if counter == 0 {
		fmt.Print("Not a Quad function")
		return
	}
}

func PrintAns(x, y int, counter *int, r rune) {
	if *counter > 0 {
		fmt.Print(" || ")
	}
	fmt.Printf("[quad%s] [%d] [%d]", string(r), x, y)
	*counter++
}

func QuadA(x, y int) string {
	if x <= 0 || y <= 0 {
		return ""
	}

	output := ""
	for i := 1; i <= y; i++ {
		for j := 1; j <= x; j++ {
			if i == 1 && j == 1 {
				output += "o"
			} else if i == 1 && j == x {
				output += "o"
			} else if i == y && j == 1 {
				output += "o"
			} else if i == y && j == x {
				output += "o"
			} else if i == 1 || i == y {
				output += "-"
			} else if j == 1 || j == x {
				output += "|"
			} else {
				output += " "
			}
		}
		output += "\n"
	}
	return output
}

func QuadB(x, y int) string {
	if x <= 0 || y <= 0 {
		return ""
	}

	output := ""
	for i := 1; i <= y; i++ {
		for j := 1; j <= x; j++ {
			if i == 1 && j == 1 {
				output += "/"
			} else if i == 1 && j == x {
				output += "\\"
			} else if i == y && j == 1 {
				output += "\\"
			} else if i == y && j == x {
				output += "/"
			} else if j == 1 || j == x {
				output += "*"
			} else if i == 1 || i == y {
				output += "*"
			} else {
				output += " "
			}
		}
		output += "\n"
	}
	return output
}

func QuadC(x, y int) string {
	output := ""
	if x <= 0 || y <= 0 {
		return ""
	}
	for i := 0; i <= y-1; i++ {
		for j := 0; j <= x-1; j++ {
			if i == 0 && j == 0 {
				output += "A"
			} else if i == 0 && j == x-1 {
				output += "A"
			} else if i == y-1 && j == 0 {
				output += "C"
			} else if i == y-1 && j == x-1 {
				output += "C"
			} else if j == 0 || j == x-1 {
				output += "B"
			} else if i == 0 || i == y-1 {
				output += "B"
			} else {
				output += " "
			}
		}
		output += "\n"
	}
	return output
}

func QuadD(x, y int) string {
	var output string
	if x <= 0 || y <= 0 {
		return ""
	}
	for i := 0; i <= y-1; i++ {
		for j := 0; j <= x-1; j++ {
			if i == 0 && j == 0 {
				output += "A"
			} else if i == 0 && j == x-1 {
				output += "C"
			} else if i == y-1 && j == 0 {
				output += "C"
			} else if i == y-1 && j == x-1 {
				output += "C"
			} else if j == 0 || j == x-1 {
				output += "B"
			} else if i == 0 || i == y-1 {
				output += "B"
			} else {
				output += " "
			}
		}
		output += "\n"
	}
	return output
}

func QuadE(x, y int) string {
	var output string
	if x <= 0 || y <= 0 {
		return ""
	}
	for i := 1; i <= y; i++ {
		for j := 1; j <= x; j++ {
			if i == 1 && j == 1 {
				output += "A"
			} else if i == 1 && j == x {
				output += "C"
			} else if i == y && j == 1 {
				output += "C"
			} else if i == y && j == x {
				output += "A"
			} else if j == 1 || j == x {
				output += "A"
			} else if i == 1 || i == y {
				output += "B"
			} else {
				output += " "
			}
		}
		output += "\n"
	}
	return output
}

func GetNumber() (x int, y int, flag bool) {
	args := os.Args[1:]
	if len(args) != 2 {
		return 0, 0, false
	}
	x, errX := strconv.Atoi(args[0])
	y, errY := strconv.Atoi(args[1])
	if errX != nil || errY != nil {
		return 0, 0, false
	}
	if x < 0 || y < 0 {
		return 0, 0, false
	}
	return x, y, true
}

func main() {
	QuadCheck()
	fmt.Println()
}
