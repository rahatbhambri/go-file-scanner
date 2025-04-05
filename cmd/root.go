package cmd

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var rootcmd = &cobra.Command{
	Use:   "nothing",
	Short: "nothing",
	Long:  "nothing",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var scancmd = &cobra.Command{
	Use:     "folder scanner",
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
}
