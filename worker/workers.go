package worker

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"code.sajari.com/docconv"
	"rsc.io/pdf"
)

var keywords []string = []string{"Taj", "Mahal", "Liberty", "Eiffel", "ChatGPT", "deforestation"}

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

func PDFWorker(path string) {
	// Open the PDF file
	r, err := pdf.Open(path)
	if err != nil {
		log.Fatalf("failed to open PDF: %v", err)
	}

	// Iterate through each page
	for i := 1; i <= r.NumPage(); i++ {
		page := r.Page(i)
		if page.V.IsNull() {
			continue
		}
		content := page.Content()

		line := 1
		cword := ""
		for _, text := range content.Text {

			fmt.Print(text.S)

			switch text.S {
			case " ":
			default:
				cword += text.S
				continue
			}

			fmt.Println("l")
			fmt.Print(cword, " ")
			for _, s := range keywords {
				if strings.Contains(cword, s) {
					fmt.Println(path, line, " contains word ", s)
				}
			}

			line++

		}
	}
}

func PrintPDF(fpath string) {
	res, err := docconv.ConvertPath(fpath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}
