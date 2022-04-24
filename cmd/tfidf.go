/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"challenge/pkg/jobs"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// tfidfCmd represents the tfidf command
var tfidfCmd = &cobra.Command{
	Use:   "tfidf",
	Short: "A brief description of your command",
	Long:  `A longer `,
	Run: func(cmd *cobra.Command, args []string) {
		word, _ := cmd.Flags().GetString("word")
		if word == "" {
			fmt.Println("Flag [word/w] must not be empty")
			os.Exit(1)
		}
		file, _ := cmd.Flags().GetString("file")
		if file != "" {
			allFiles, err := jobs.AllFiles(file)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			jobs.ETL(word, allFiles)
		} else {
			fmt.Println("Flag [file/f] must not be empty")
			os.Exit(1)
		}
	},
}

func init() {
	tfidfCmd.Flags().StringP("word", "w", "", "Word to get the TF-IDF weight of the word in the document")
	tfidfCmd.Flags().StringP("file", "f", "", "Searchable text document")

	rootCmd.AddCommand(tfidfCmd)
}
