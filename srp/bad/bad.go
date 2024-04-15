package bad

import (
	"encoding/csv"
	"io"
	"net/http"
	"os"
	"strings"
)

func readProcessAndSendData(filename, apiUrl string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}

	defer file.Close()

	// responsibility 1
	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		// responsibility 2
		processedData := strings.ToUpper(strings.Join(record, ","))

		// responsibility 3
		_, err = http.Post(apiUrl, "text/plain", strings.NewReader(processedData))

		if err != nil {
			return err
		}
	}

	return nil
}
