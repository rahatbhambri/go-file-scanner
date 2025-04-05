package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/spf13/cobra"
)

var keywords []string = []string{"Taj", "Mahal", "Liberty", "Eiffel", "ChatGPT", "deforestation"}

var rootcmd = &cobra.Command{
	Use:   "The_scanner",
	Short: "scans a folder or a file",
	Long:  "scans a folder or a file for specific use cases",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var scancmd = &cobra.Command{
	Use:     "folder_scanner",
	Aliases: []string{"scan"},
	Short:   "cli folder scanner to get all file details stored in a folder",
	Long:    "cli folder scanner to get all file details stored in a folder- name, size, datecreated",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			path := args[0]
			fmt.Println("Inside handler, path = ", path)

			entries, err := os.ReadDir(path)
			if err != nil {
				log.Fatal(err)
			}

			for _, e := range entries {
				fileInfo, err := os.Lstat(path + "\\" + e.Name())
				if err != nil {
					panic(err)
				}

				fmt.Println("Name : ", fileInfo.Name())
				fmt.Println("Size : ", fileInfo.Size(), "B")
				fmt.Println("Mode/permission : ", fileInfo.Mode())
				fmt.Println()

			}
		}
	},
}

var textcmd = &cobra.Command{
	Use:     "text_file_scanner",
	Aliases: []string{"filescan"},
	Short:   "cli to scan all text files to search for specific keywords",
	Long:    "cli to scan all text files to search for specific keywords defined in the config",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			path := args[0]
			var wg sync.WaitGroup
			fmt.Println("Inside handler, path = ", path)

			entries, err := os.ReadDir(path)
			if err != nil {
				log.Fatal(err)
			}

			var text_files []string
			for _, e := range entries {
				fpath := path + "\\" + e.Name()
				fileInfo, err := os.Lstat(fpath)
				if err != nil {
					panic(err)
				}

				name := fileInfo.Name()
				// fmt.Println("name", name)
				n := len(name)
				if name[n-4:] == ".txt" {
					text_files = append(text_files, name)
				}
			}

			fmt.Println(text_files)

			maxGoroutines := 10
			// Max files which can be read parallely
			guard := make(chan struct{}, maxGoroutines)

			for _, fname := range text_files {
				guard <- struct{}{} // would block if guard channel is already filled
				fpath := path + "\\" + fname
				wg.Add(1)
				go func(fpath string) {
					defer wg.Done()
					// fmt.Println("wow")
					worker(fpath)
					<-guard
				}(fpath)
			}

			wg.Wait()

		}
	},
}

func worker(path string) {

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

var addcmd = &cobra.Command{
	Use:     "adder",
	Aliases: []string{"add"},
	Short:   "add two numbers using cli",
	Long:    "performs addition of two numbers using cli",
	Run: func(cmd *cobra.Command, args []string) {
		num1, _ := strconv.ParseInt(args[0], 10, 64)
		num2, _ := strconv.ParseInt(args[1], 10, 64)
		fmt.Println(num1 + num2)
	},
}

func Execute() {
	if err := rootcmd.Execute(); err != nil {
		fmt.Printf("Some error occured")
		os.Exit(1)
	}
}

func init() {
	rootcmd.AddCommand(addcmd)
	rootcmd.AddCommand(scancmd)
	rootcmd.AddCommand(textcmd)
}
