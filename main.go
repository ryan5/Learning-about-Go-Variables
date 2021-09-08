package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

const prompt = "and press ENTER when ready."

func main() {
	var firstNumber = rand.Intn(8) + 2
	var secondNumber = rand.Intn(8) + 2
	var subtraction = rand.Intn(8) + 2
	var answer = firstNumber*secondNumber - subtraction

	game(firstNumber, secondNumber, subtraction, answer)

}

func game(firstNumber, secondNumber, subtraction, answer int) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Guess the Number game")
	fmt.Println("---------------------")
	fmt.Println("")

	fmt.Println("Think of a number between 1 and 10", prompt)
	reader.ReadString('\n')

	fmt.Println("Multiply your number by", firstNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Now multiply the result", secondNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Divide the result by the number you originally thought of", prompt)
	reader.ReadString('\n')

	fmt.Println("Now subtract", subtraction, prompt)
	reader.ReadString('\n')

	fmt.Println("The answer is", answer)
}
