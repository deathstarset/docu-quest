package main

/* // Reading the pdf page
f, err := os.Open("dli-fundamentals-of-deep-learning-1369828-r3-web.pdf")
if err != nil {
	log.Fatalf("Failed to open pdf file : %s", err.Error())
}
defer f.Close()
pdfReader, err := model.NewPdfReader(f)
if err != nil {
	log.Fatalf("Failed to read pdf file : %s", err.Error())
}

numPages, err := pdfReader.GetNumPages()
if err != nil {
	log.Fatalf("Failed to retreive the number of pages : %s", err.Error())
}
fmt.Printf("The total number of page is : %d\n", numPages)

// Extracting text from pdf pages

pageNum := 1

page, err := pdfReader.GetPage(pageNum)
if err != nil {
	log.Fatalf("Failed to retreive page %d : %s", pageNum, err)
}
ex, err := extractor.New(page)
if err != nil {
	log.Fatalf("Failed to create page extractor %d : %s", pageNum, err)
}

text, err := ex.ExtractText()
if err != nil {
	log.Fatalf("Failed to extract text from page %d : %s", pageNum, err)
}

fmt.Printf("\npage : %d", pageNum)
fmt.Printf("\ntext : %s", text) */
