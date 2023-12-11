package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/enescakir/emoji"
)

func main() {
	const MIN int = 1
	const MAX int = 100
	var tries = 0
	// Init a new rand object
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	// Init the state of our program
	answer := r1.Intn(MAX+1) + MIN
	guess := -1

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("\n\n")

	fmt.Printf("%s Bienvenido al juego de adivinar el número %s \n", emoji.SmilingCatWithHeartEyes, emoji.SmilingCatWithHeartEyes)
	for i := 1; i < 25; i++ {
		fmt.Printf("%v", emoji.Laptop)
	}
	// Main game loop
	for guess != answer {
		// Receive user input

		for {
			isCorrect := false

			fmt.Printf("\n\n")
			fmt.Print("Introduce un número del ", MIN, " al ", MAX, ": ")
			for scanner.Scan() {

				input := scanner.Text()

				n, err := strconv.ParseInt(input, 10, 0)
				num := int(n)
				if err != nil {
					fmt.Fprintf(os.Stderr, "Número inválido: %s\n", input)
				} else if num < MIN || num > MAX {
					fmt.Fprintf(os.Stderr, "Numero fuera de rango: %d\n", num)
				} else {
					guess = num
					isCorrect = true
				}

				break
			}

			if isCorrect {
				break
			}
		}

		// Check our guess is correct
		tries++
		if guess > answer {
			fmt.Printf("El número que estoy pensando es menor al que has introducido %v\n", emoji.ThumbsDown)
		} else if guess < answer {
			fmt.Printf("El número que estoy pensando es mayor al que has introducido %v\n", emoji.ThumbsUp)
		} else {
			fmt.Printf("%v ¡Lo has acertado!. Número de intentos: %v %v\n", emoji.MoneyWithWings, tries, emoji.MoneyWithWings)

			break
		}
	}
}
