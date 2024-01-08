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

var usersFile []byte

func main() {
	var users = make(map[string]int)

	const MIN int = 1
	const MAX int = 100

	var name string
	// Init a new rand object
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	openFile()

	// Init the state of our program
	answer := r1.Intn(MAX+1) + MIN

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("\n\n")

	fmt.Printf("%s Bienvenido al juego de adivinar el número. %s \n", emoji.SmilingCatWithHeartEyes, emoji.SmilingCatWithHeartEyes)
	fmt.Printf("%s He pensado un número del 1 al 100 y tienes que adivinar cual es. %s \n", emoji.SmilingCatWithHeartEyes, emoji.SmilingCatWithHeartEyes)
	/*for i := 1; i < 25; i++ {
		fmt.Printf("%v", emoji.Laptop)
	}*/

	fmt.Printf("\n\n")
	fmt.Printf("Introduce tu nombre por favor.")
	scanner.Scan()
	name = scanner.Text()
	//fmt.Println(name)
	startGame(users, name, MIN, MAX, answer, *scanner)
	//fmt.Printf("users: %v", users)

}

func openFile() {

	var err error
	usersFile, err = os.ReadFile("/tmp/dat2")
	check(err)
	//text := string(usersFile)
	//	fmt.Println(text)

}

func appendUserToFile(user string) {

	file, err := os.OpenFile("/tmp/dat2", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	check(err)
	defer file.Close()
	file.WriteString(user + "\n")

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func startGame(users map[string]int, name string, min int, max int, answer int, scanner bufio.Scanner) {
	// Main game loop
	var tries = 0
	guess := -1
	for guess != answer {
		// Receive user input

		for {
			isCorrect := false

			fmt.Printf("\n\n")
			fmt.Print("Introduce un número del ", min, " al ", max, ": ")
			for scanner.Scan() {

				input := scanner.Text()

				n, err := strconv.ParseInt(input, 10, 0)
				num := int(n)
				if err != nil {
					fmt.Fprintf(os.Stderr, "Número inválido: %s\n", input)
				} else if num < min || num > max {
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
			fmt.Printf("El número que estoy pensando es menor al que has introducido %v\n", emoji.DownArrow)
		} else if guess < answer {
			fmt.Printf("El número que estoy pensando es mayor al que has introducido %v\n", emoji.UpArrow)
		} else {
			fmt.Printf("%v ¡Lo has acertado!. Número de intentos: %v %v\n", emoji.MoneyWithWings, tries, emoji.MoneyWithWings)
			users[name] = tries
			//fmt.Printf("users: %v", users)
			var a string = name + " " + strconv.Itoa(tries)
			appendUserToFile(a)
			break
		}
	}
}
