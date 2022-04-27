package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "challenge",
	Short: "CLI to calculate TF-IDF and download files",
	Long: `This CLI has the essential commands to calculate the TF-IDF weight 
and to download files from the internet through a given url, for example.

To calculate the weight of a word within a given document use:
challenge tfidf -w "ac" -f /data/document_1.txt

To download a file from the internet given the url of the file use:
challenge wget https://unec.edu.az/application/uploads/2014/12/pdf-sample.pdf -d /tmp -f file.pdf

For more details of the commands visit the help for each command.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
