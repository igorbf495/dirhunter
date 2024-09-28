package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

func main() {

	printSkull()


	baseURL := flag.String("url", "", "URL alvo para força bruta (ex: http://example.com)")
	wordlistFile := flag.String("wordlist", "", "Caminho para o arquivo de wordlist (ex: wordlist.txt)")
	threads := flag.Int("threads", 10, "Número de threads (padrão: 10)")

	
	flag.Usage = func() {
		fmt.Println(`DirHunter - Força bruta de diretórios em servidores web
Uso:
  dirhunter -url <URL alvo> -wordlist <arquivo de wordlist> [-threads <número de threads>]

Parâmetros:
  -url         URL alvo para força bruta (ex: http://example.com)
  -wordlist    Caminho para o arquivo de wordlist (ex: wordlist.txt)
  -threads     Número de threads (opcional, padrão: 10)

Exemplo:
  dirhunter -url http://example.com -wordlist wordlist.txt -threads 20
`)
	}

	flag.Parse()

	
	if *baseURL == "" || *wordlistFile == "" {
		flag.Usage()
		return
	}

	fmt.Println("Iniciando DirHunter...")

	
	words, err := readWordlist(*wordlistFile)
	if err != nil {
		fmt.Println("Erro ao ler a wordlist:", err)
		return
	}


	startBruteForce(*baseURL, words, *threads)
}


func printSkull() {
	skull := `
    .--.
   |o_o |
   |:_/ |
  //   \ \
 (|     | )
/'\_   _/` + "`" + `\
\___)=(___/
`
	fmt.Println(skull)
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


func startBruteForce(baseURL string, wordlist []string, threads int) {
	var wg sync.WaitGroup
	wordChan := make(chan string)


	for i := 0; i < threads; i++ {
		wg.Add(1)
		go worker(baseURL, wordChan, &wg)
	}


	for _, word := range wordlist {
		wordChan <- word
	}


	close(wordChan)


	wg.Wait()
}

func worker(baseURL string, wordChan <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	for word := range wordChan {
		url := fmt.Sprintf("%s/%s", strings.TrimRight(baseURL, "/"), word)
		makeRequest(url)
	}
}


func makeRequest(url string) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println("Erro ao fazer requisição:", err)
		return
	}
	defer resp.Body.Close()

	
	if resp.StatusCode == http.StatusOK {
		fmt.Println("[+] Encontrado:", url, "(Status: 200 OK)")
	} else if resp.StatusCode == http.StatusForbidden {
		fmt.Println("[!] Acesso proibido:", url, "(Status: 403 Forbidden)")
	}
}
