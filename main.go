package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sort"
	"sync"
	"time"

	"github.com/olekukonko/tablewriter"
)

type AgeData struct {
	Count int    `json:"count"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
}

var httpClient = &http.Client{
	Timeout: 10 * time.Second, // Set a reasonable timeout for HTTP requests
}

func estimateAges(names, countries []string, ch chan<- []string, wg *sync.WaitGroup, mu *sync.Mutex, totalElapsedTime *time.Duration) {
	defer wg.Done()

	for _, nameinput := range names {
		for _, countryinput := range countries {
			wg.Add(1)
			go func(name, country string) {
				defer wg.Done()
				startTime := time.Now()
				result := estimateAge(name, country)
				elapsedTime := time.Since(startTime)
				mu.Lock()
				*totalElapsedTime += elapsedTime
				mu.Unlock()
				ch <- []string{name, country, fmt.Sprintf("%d", result.Age), elapsedTime.String()}
			}(nameinput, countryinput)
		}
	}
}

func estimateAge(nameinput, countryinput string) AgeData {
	var data AgeData

	url := fmt.Sprintf("https://api.agify.io?name=%s&country_id=%s", nameinput, countryinput)

	response, err := httpClient.Get(url)
	if err != nil {
		fmt.Printf("Error fetching data for %s in %s: %v\n", nameinput, countryinput, err)
		return data
	}
	defer response.Body.Close()

	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		fmt.Printf("Error decoding data for %s in %s: %v\n", nameinput, countryinput, err)
		return data
	}

	return data
}

func printTable(results [][]string, totalElapsedTime, runtime time.Duration) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Country", "Age", "Time Taken"})

	for _, row := range results {
		name, country, age, elapsedTimeStr := row[0], row[1], row[2], row[3]
		elapsedTime, _ := time.ParseDuration(elapsedTimeStr)
		elapsedTimeMillis := elapsedTime.Round(time.Millisecond)
		table.Append([]string{name, country, age, elapsedTimeMillis.String()})
	}

	totalElapsedMillis := totalElapsedTime.Round(time.Millisecond)
	runtimeMillis := runtime.Round(time.Millisecond)

	table.Append([]string{"Time if Sequential", "", "", totalElapsedMillis.String()})
	table.Append([]string{"Time when Parallel", "", "", runtimeMillis.String()})
	table.Render()
}

func main() {
	startNow := time.Now()
	names := []string{"John", "Jane", "Chuck", "Sue", "Bob", "Alice", "Emma", "Jessica", "Ashley"}
	countries := []string{"CA", "US", "GB", "AU"} // Add your desired countries here
	ch := make(chan []string)

	var wg sync.WaitGroup
	var mu sync.Mutex
	var totalElapsedTime time.Duration

	wg.Add(1)
	go estimateAges(names, countries, ch, &wg, &mu, &totalElapsedTime)

	go func() {
		wg.Wait()
		close(ch)
	}()

	var results [][]string

	for result := range ch {
		results = append(results, result)
	}

	// Sort the results alphabetically by name and then by country
	sort.Slice(results, func(i, j int) bool {
		return results[i][0] < results[j][0] || (results[i][0] == results[j][0] && results[i][1] < results[j][1])
	})

	// Print the sorted table with total time
	runtime := time.Since(startNow)
	printTable(results, totalElapsedTime, runtime)

}
