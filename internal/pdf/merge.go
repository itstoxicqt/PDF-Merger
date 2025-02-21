package pdf

import (
	"fmt"

	"github.com/pdfcpu/pdfcpu/pkg/api"
)

// MergePDFs merges multiple input PDFs into a single output file
func MergePDFs(inputFiles []string, outputFile string) error {
	err := api.MergeCreateFile(inputFiles, outputFile, false, nil) // ✅ FIXED: No need for config
	if err != nil {
		return fmt.Errorf("error merging PDFs: %v", err)
	}

	return nil
}
