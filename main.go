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

	"github.com/fatih/color"
)

func main() {
	printSkull()

	baseURL := flag.String("url", "", "URL alvo para força bruta (ex: http://example.com)")
	wordlistFile := flag.String("wordlist", "", "Caminho para o arquivo de wordlist (ex: wordlist.txt)")
	threads := flag.Int("threads", 10, "Número de threads (padrão: 10)")
	delay := flag.Int("delay", 0, "Atraso em milissegundos entre requisições (padrão: 0)")

	flag.Usage = func() {
		fmt.Println(`DirHunter - Força bruta de diretórios em servidores web
Uso:
  dirhunter -url <URL alvo> -wordlist <arquivo de wordlist> [-threads <número de threads>] [-delay <milissegundos>]

Parâmetros:
  -url         URL alvo para força bruta (ex: http://example.com)
  -wordlist    Caminho para o arquivo de wordlist (ex: wordlist.txt)
  -threads     Número de threads (opcional, padrão: 10)
  -delay       Atraso em milissegundos entre requisições (opcional, padrão: 0)

Exemplo:
  dirhunter -url http://example.com -wordlist wordlist.txt -threads 20 -delay 100
`)
	}

	flag.Parse()

	if *baseURL == "" || *wordlistFile == "" {
		flag.Usage()
		return
	}

	color.Green("Iniciando DirHunter...")
	time.Sleep(2 * time.Second)

	words, err := readWordlist(*wordlistFile)
	if err != nil {
		color.Red("Erro ao ler a wordlist: %v", err)
		return
	}

	startBruteForce(*baseURL, words, *threads, *delay)
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

func startBruteForce(baseURL string, wordlist []string, threads int, delay int) {
	var wg sync.WaitGroup
	wordChan := make(chan string)
	foundCount := 0

	for i := 0; i < threads; i++ {
		wg.Add(1)
		go worker(baseURL, wordChan, &wg, &foundCount, delay)
	}

	for _, word := range wordlist {
		wordChan <- word
	}

	close(wordChan)

	wg.Wait()
}

func worker(baseURL string, wordChan <-chan string, wg *sync.WaitGroup, foundCount *int, delay int) {
	defer wg.Done()

	for word := range wordChan {
		url := fmt.Sprintf("%s/%s", strings.TrimRight(baseURL, "/"), word)
		makeRequest(url, foundCount)

		// Atraso entre requisições
		if delay > 2 {
			time.Sleep(time.Duration(delay) * time.Millisecond)
		}
	}
}

func makeRequest(url string, foundCount *int) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Get(url)
	if err != nil {
		color.Yellow("Erro ao fazer requisição: %v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		color.Green("[+] Encontrado: %s (Status: 200 OK)", url)
		*foundCount++
	} else if resp.StatusCode == http.StatusForbidden {
		color.Red("[!] Acesso proibido: %s (Status: 403 Forbidden)", url)
	} else {
		color.Blue("[-] Status %d: %s", resp.StatusCode, url)
	}
}
