package cmd

import (
	"challenge/pkg/jobs"
	"challenge/pkg/models"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// wgetCmd represents the wget command
var wgetCmd = &cobra.Command{
	Use:   "wget",
	Short: "Download a file given a url",
	Long: `Download a file exposed in a url 
to which you can specify the name with which you want to 
save it and the directory where it is downloaded or let it 
take the name given in the url by default

Example:
  challenge wget https://unec.edu.az/application/uploads/2014/12/pdf-sample.pdf
	
  challenge wget https://unec.edu.az/application/uploads/2014/12/pdf-sample.pdf -d "./data" -f "name.pdf"`,
	Run: func(cmd *cobra.Command, args []string) {
		var data models.WGET
		if args[0] == "" {
			fmt.Println("URL empty")
			fmt.Println("Use: wget [URL]")
			os.Exit(1)
		}
		data.URI = args[0]
		dir, _ := cmd.Flags().GetString("dir")
		if dir == "" {
			data.Directory = "./"
		}
		data.Directory = dir
		fileName, _ := cmd.Flags().GetString("file-name")
		data.FileName = fileName
		fmt.Printf("Downloading from: %s \n", data.URI)
		msg := jobs.Wget(data)
		fmt.Println(msg)
	},
}

func init() {
	rootCmd.AddCommand(wgetCmd)
	wgetCmd.Flags().StringP("dir", "d", "", "Directory save file")
	wgetCmd.Flags().StringP("file-name", "f", "", "Directory save file")
}
