package processes

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"

	"myoffice/internal/models"
)

// Run - запускает процесс сбора статистики по заданному файлу.
func Run(path string, parallelFlag string) error {
	startTime := time.Now()
	maxWorkers, _ := strconv.Atoi(parallelFlag)
	if maxWorkers == 0 {
		maxWorkers = runtime.NumCPU()
	}
	fmt.Printf("Start with %d parallel working\n", maxWorkers)

	urlChan := make(chan string)
	countChan := make(chan int64, 1)

	var wg sync.WaitGroup

	go processURLs(path, urlChan, countChan)

	for i := 0; i < maxWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for u := range urlChan {
				fetchURL(u)
			}
		}()
	}

	wg.Wait()

	finishTime := time.Since(startTime).String()
	countURL := <-countChan

	fmt.Println("\nTotal processing time:", finishTime)
	fmt.Println("Total count URLs:", countURL)

	return nil
}

// fetchURL - получает статистику по URL-у.
func fetchURL(url string) {
	startTime := time.Now()
	report := models.Report{
		URL: url,
	}
	timeout := 3 * time.Second
	client := &http.Client{
		Timeout: timeout,
	}

	valid, err := isValidURL(url)
	if err != nil {
		report.Error = err.Error()
	}
	if valid {
		response, err := client.Get(url)
		if err != nil {
			report.Error = fmt.Sprintf("error fetching %s: %v", url, err)
		} else {
			defer func() {
				if response != nil && response.Body != nil {
					response.Body.Close()
				}
			}()

			if response != nil {
				size, err := io.Copy(io.Discard, response.Body)
				if err != nil && err != io.EOF {
					report.Error = fmt.Sprintf("error reading response body for %s: %v", url, err)
				}

				if report.Error == "" && size >= 0 {
					report.Size = size
				}
			}
		}
	} else {
		report.Error = "invalid URL"
	}
	report.Time = time.Since(startTime).String()
	fmt.Printf("URL: %s Size: %d bytes Time: %s Error: %s\n",
		report.URL,
		report.Size,
		report.Time,
		report.Error,
	)
}

// processURLs - читает URL-ы из файла и отправляет их в канал.
func processURLs(filename string, urlChan chan<- string, countChan chan<- int64) {
	var countURL int64
	file, err := os.Open(filename)
	if err != nil {
		urlChan <- fmt.Sprintf("Error opening file: %v", err)
		close(urlChan)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		urlItem := scanner.Text()
		if len(urlItem) != 0 {
			urlChan <- urlItem
			countURL++
		}
	}

	countChan <- countURL

	close(urlChan)
	close(countChan)
}

// isValidURL - проверяет валидность URL-а.
func isValidURL(input string) (bool, error) {
	u, err := url.Parse(input)
	if err != nil {
		return false, fmt.Errorf("invalid URL: %v", err)
	}

	if len(u.Scheme) == 0 {
		return false, fmt.Errorf("URL is missing scheme: %s", input)
	}

	return u.Host != "" && u.IsAbs(), nil
}
