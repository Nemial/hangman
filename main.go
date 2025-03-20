package main

import (
	"bufio"
	"embed"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
)

func main() {
	var (
		words    = loadWords()
		secret   = words[rand.Intn(len(words))]
		attempts = 5
		reader   = bufio.NewReader(os.Stdin)
		guessed  = ""
	)

	fmt.Println("Добро пожаловать в игру Виселица!")

	for attempts > 0 {
		fmt.Println("Назови букву: ")

		missed := 0
		for _, letter := range secret {
			if strings.ContainsRune(guessed, letter) {
				fmt.Printf("%c ", letter)
			} else {
				fmt.Print("_ ")
				missed += 1
			}
		}
		fmt.Println()

		if missed == 0 {
			fmt.Println("\nТы выиграл!")
			return
		}

		var guess rune
		_, err := fmt.Scanf("%c", &guess)
		if err != nil {
			log.Fatal("user input error: ", err.Error())
		}

		_, _ = reader.ReadString('\n')

		guessed += string(guess)

		if !strings.ContainsRune(secret, guess) {
			attempts -= 1
			fmt.Printf("Не угадал. Осталось попыток: %d\n", attempts)

			if attempts < 5 {
				fmt.Println(" |")
			}
			if attempts < 4 {
				fmt.Println(" O")
			}
			if attempts < 3 {
				fmt.Println("/|\\")
			}
			if attempts < 2 {
				fmt.Println(" |")
			}
			if attempts < 1 {
				fmt.Println("/ \\")
				fmt.Printf("\nЭто слово: %s", secret)
				return
			}
		}
	}
}

//go:embed words.txt
var f embed.FS

func loadWords() []string {
	fileName := "words.txt"
	file, err := f.Open(fileName)

	if err != nil {
		log.Fatal("Ошибка открытия файла:", err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var words []string

	for scanner.Scan() {
		line := scanner.Text()
		words = append(words, line)
	}

	return words
}
