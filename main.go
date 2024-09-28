package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)


func main() {
	
	baseURL := "http://example.com"
	wordlistFile := "wordlist.txt"
	threads := 10

	fmt.Println("Iniciando DirHunter...")
	

	words, err := readWordlist(wordlistFile)
	if err != nil {
		fmt.Println("Erro ao ler a wordlist:", err)
		return
	}

	
	startBruteForce(baseURL, words, threads)
}


func readWordlist(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	return words, scanner.Err()
}
