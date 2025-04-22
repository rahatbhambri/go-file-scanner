package worker

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ledongthuc/pdf"
)

var keywords []string = []string{"Taj", "Mahal", "Liberty", "Eiffel", "ChatGPT", "deforestation", "Kaggle", "atkinson"}

func Textworker(path string) {
	// fmt.Println("working on", path)
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	// Splits on newlines by default.
	scanner := bufio.NewScanner(f)

	line := 1
	for scanner.Scan() {
		for _, s := range keywords {
			if strings.Contains(scanner.Text(), s) {
				fmt.Println(path, line, " contains word ", s)
			}
		}

		line++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}

func PDFWorker(fpath string) {
	// Open the PDF file
	f, r, err := pdf.Open(fpath)
	if err != nil {
		log.Fatalf("failed to open PDF: %v", err)
	}
	defer f.Close()

	for pageIndex := 1; pageIndex <= r.NumPage(); pageIndex++ {
		page := r.Page(pageIndex)
		if page.V.IsNull() {
			continue
		}

		content := page.Content()
		cword := ""
		for _, text := range content.Text {
			char := text.S
			if strings.TrimSpace(char) == "" {
				if cword != "" {
					for _, s := range keywords {
						if strings.Contains(cword, s) {
							fmt.Println(fpath, " contains word ", s)
						}
					}
					// fmt.Println(cword)
					cword = ""
				}
			} else {
				cword += char
			}
		}
	}
}

func CsvWorker(filePath string) {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	for _, record := range records {
		for _, word := range record {
			// fmt.Println(word)
			for _, s := range keywords {
				if strings.EqualFold(word, s) {
					fmt.Println(filePath, " contains word ", s)
				}
			}
		}

	}

}
