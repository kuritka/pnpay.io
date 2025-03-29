package js2pdf

import (
	"context"
	"fmt"
	"os"
	"time"

	v1 "pnpay.io/api/v1"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

// Convert HTML to PDF
func GeneratePDF(url string, outputPath string, doc *v1.SinglePageExtended) error {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// Timeout to avoid infinite wait
	ctx, cancel = context.WithTimeout(ctx, 20*time.Second)
	defer cancel()

	// Convert cm to inches (1 cm = 0.3937 inch)
	widthInches := doc.Extensions.PDFCanvas.WidthCM * 0.393701
	heightInches := doc.Extensions.PDFCanvas.HeightCM * 0.393701

	// Convert cm to pixels (1 cm ≈ 37.8 px)
	// widthPx := int64((widthCM + 0) * 35)
	// heightPx := int64((heightCM + 0) * 35)

	// PDF output buffer
	var pdfData []byte
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),       // Open the HTML page
		chromedp.WaitVisible("body"), // Ensure page is fully loaded
		chromedp.ActionFunc(func(ctx context.Context) error {
			var err error
			pdfData, _, err = page.PrintToPDF().
				WithPaperWidth(widthInches).   // Exact width
				WithPaperHeight(heightInches). // Exact height
				WithPrintBackground(true).     // Ensure background (CSS, images) is included
				WithMarginTop(0).              // Remove top margin
				WithMarginBottom(0).           // Remove bottom margin
				WithMarginLeft(0).             // Remove left margin
				WithMarginRight(0).            // Remove right margin
				WithScale(1.0).                // Ensure content is not shrunk or cut
				Do(ctx)
			return err
		}),
	)
	if err != nil {
		return fmt.Errorf("failed to generate PDF: %w", err)
	}

	// Open file for writing
	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer func() { _ = file.Close() }()

	// Write PDF data to file
	_, err = file.Write(pdfData)
	if err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}

	// Ensure the data is fully written to disk
	err = file.Sync()
	if err != nil {
		return fmt.Errorf("failed to sync file to disk: %w", err)
	}

	fmt.Println("PDF successfully created:", outputPath)
	return nil
}
