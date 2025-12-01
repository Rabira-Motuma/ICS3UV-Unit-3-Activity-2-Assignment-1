/*
 * Author: Rabira Motuma
 * Version: 1.0.0
 * Date: 2025-11-25
 * Fileoverview: This program rounds off numbers to a specified number of decimal places.
 */

package main

import(
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// variables
	var number1AsString string
	var number2AsString string
	var number3AsString string
	var number4AsString string
	var number1AsNumber float64
	var number2AsNumber float64
	var number3AsNumber float64
	var number4AsNumber float64

	reader := bufio.NewReader(os.Stdin)

	// input
	fmt.Print("What's your first number? ")
	number1AsString, _ = reader.ReadString('\n')
	number1AsString = strings.TrimSpace(number1AsString)

	fmt.Print("What's your second number? ")
	number2AsString, _ = reader.ReadString('\n')
	number2AsString = strings.TrimSpace(number2AsString)

	fmt.Print("What's your third number? ")
	number3AsString, _ = reader.ReadString('\n')
	number3AsString = strings.TrimSpace(number3AsString)

	fmt.Print("What's your fourth number? ")
	number4AsString, _ = reader.ReadString('\n')
	number4AsString = strings.TrimSpace(number4AsString)

	// process
	number1AsNumber, _ = strconv.ParseFloat(number1AsString, 64)
	number2AsNumber, _ = strconv.ParseFloat(number2AsString, 64)
	number3AsNumber, _ = strconv.ParseFloat(number3AsString, 64)
	number4AsNumber, _ = strconv.ParseFloat(number4AsString, 64)

	// output
	fmt.Println()
	fmt.Printf("%10.3f\n", number1AsNumber)
	fmt.Printf("%8.5f\n", number2AsNumber)
	fmt.Printf("%6.1f\n", number3AsNumber)
	fmt.Printf("%3.1f\n", number4AsNumber)

	fmt.Println("\nDone.")
}
