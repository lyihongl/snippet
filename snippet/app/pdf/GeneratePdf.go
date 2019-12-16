package pdf

import(
	"fmt"
	htp "github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func GeneratePdf () {
	pdfg, err := htp.NewPDFGenerator()
	fmt.Println("test")
}