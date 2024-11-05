package main

import (
	"flag"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type TestResult struct {
	TotalTime     time.Duration
	TotalRequests int
	SuccessCount  int
	StatusCodes   map[int]int
}

func main() {
	url := flag.String("url", "", "URL a ser testada")
	requests := flag.Int("requests", 1, "Número de requisições.")
	concurrency := flag.Int("concurrency", 1, "Número de chamadas simultâneas.")
	flag.Parse()

	if *url == "" {
		fmt.Println("A URL do serviço é obrigatória.")
		flag.Usage()
		return
	}

	requestChan := make(chan struct{}, *concurrency)
	var wg sync.WaitGroup
	result := &TestResult{StatusCodes: make(map[int]int)}

	start := time.Now()

	for i := 0; i < *requests; i++ {
		wg.Add(1)
		requestChan <- struct{}{}

		go func() {
			defer wg.Done()
			defer func() { <-requestChan }()

			resp, err := http.Get(*url)
			if err != nil {
				fmt.Println("Erro de requisição: ", err)
			}
			defer resp.Body.Close()

			result.TotalRequests++
			if resp.StatusCode == http.StatusOK {
				result.SuccessCount++
			}
			result.StatusCodes[resp.StatusCode]++
		}()
	}

	wg.Wait()
	result.TotalTime = time.Since(start)

	fmt.Println("\nRelatório do Teste de Carga")
	fmt.Printf("\nTempo total gasto na execução: %v\n", result.TotalTime)
	fmt.Printf("Quantidade total de requests realizados: %d\n", result.TotalRequests)
	fmt.Printf("Quantidade de requests com status HTTP 200: %d\n", result.SuccessCount)
	fmt.Println("Distribuição de outros códigos de status HTTP:")
	for code, count := range result.StatusCodes {
		fmt.Printf("Status %d: %d\n", code, count)
	}
}
