package cmd

import (
	"fmt"
	"os"

	"pdfmerger/internal/pdf" // ✅ FIXED: Correct import path

	"github.com/spf13/cobra"
)

// Root command for merging PDFs
var rootCmd = &cobra.Command{
	Use:   "pdfmerger output.pdf input1.pdf input2.pdf ...",
	Short: "Merge multiple PDFs into one",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		outputFile := args[0]
		inputFiles := args[1:]

		err := pdf.MergePDFs(inputFiles, outputFile)
		if err != nil {
			fmt.Println("Error merging PDFs:", err)
			os.Exit(1)
		}

		fmt.Println("Successfully merged PDFs into:", outputFile)
	},
}

// Execute runs the CLI application
func Execute() error {
	return rootCmd.Execute()
}
