package PDF

import (
	"fmt"
	"github.com/rsc/pdf"
)

func ReadPDFandPage() {
	file, err := pdf.Open("C:\\Users\\12691\\Desktop\\1.pdf")
	if err != nil {
		panic(err)
	}
	fmt.Println(file.NumPage())
}
