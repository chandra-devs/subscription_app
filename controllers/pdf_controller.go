package controllers

import (
	"bufio"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/jung-kurt/gofpdf"
)

// GeneratePDF generates a PDF from the Markdown file
func GeneratePDF(c *fiber.Ctx) error {
	file, err := os.Open("docs/api.md")
	if err != nil {
		return c.Status(500).SendString("Failed to open Markdown file")
	}
	defer file.Close()

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		pdf.Cell(40, 10, scanner.Text())
		pdf.Ln(10)
	}

	if err := scanner.Err(); err != nil {
		return c.Status(500).SendString("Failed to read Markdown file")
	}

	err = pdf.OutputFileAndClose("public/api.pdf")
	if err != nil {
		return c.Status(500).SendString("Failed to generate PDF")
	}

	return c.SendFile("public/api.pdf")
}
