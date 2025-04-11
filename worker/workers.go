package worker

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ledongthuc/pdf"
)

var keywords []string = []string{"Taj", "Mahal", "Liberty", "Eiffel", "ChatGPT", "deforestation", "Kaggle"}

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

func PrintPDF(fpath string) {
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
