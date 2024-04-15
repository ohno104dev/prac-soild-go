package better

// Single Responsibility Principle (SRP)
// 一個類或一個模塊應該只有一個更改的原因, 意味著它應該只有一個責任

// Advantages
// 易於理解, 每個package或func或都有一個明確定義的職責
// 易於維護, 職責明確,最大程度地減少他們對代碼庫其他部分的影響
// 可測試性, 通過限制每個package或func的範圍, 可以編寫更集中,更徹底的測試

// Disadvantages
// 過度碎片化, 產生過多的小package或func, 可能會使代碼更難理解
// 性能權衡, 可能會由於func調用導致額外的性能開銷

import (
	"encoding/csv"
	"net/http"
	"os"
	"strings"
)

func readCsvData(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	reader := csv.NewReader(file)
	data, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func processCsvData(data [][]string) []string {
	var processedData []string
	for _, record := range data {
		processedData = append(processedData, strings.ToUpper(strings.Join(record, ",")))
	}
	return processedData
}

func sendDataToApi(apiUrl string, data []string) error {
	for _, record := range data {
		_, err := http.Post(apiUrl, "text/plain", strings.NewReader(record))
		if err != nil {
			return err
		}
	}

	return nil
}
