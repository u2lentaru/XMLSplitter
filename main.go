package main

// import (
// 	"encoding/xml"
// 	"fmt"
// 	"os"
// )

// type Book struct {
// 	Period     string `xml:"Period"`
// 	ТочкаУчета string `xml:"ТочкаУчета"`
// }

// func main() {

// 	f, err := os.Open("SP.xml")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer f.Close()
// 	decoder := xml.NewDecoder(f)

// 	// Чтение book по частям
// 	books := make([]Book, 0)
// 	for {
// 		tok, err := decoder.Token()
// 		if err != nil {
// 			panic(err)
// 		}
// 		if tok == nil {
// 			break
// 		}
// 		switch tp := tok.(type) {
// 		case xml.StartElement:
// 			if tp.Name.Local == "СреднееПотребление" {
// 				// Декодирование элемента в структуру
// 				var b Book
// 				decoder.DecodeElement(&b, &tp)
// 				books = append(books, b)
// 			}
// 		}
// 	}
// 	fmt.Println(books)
// }

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var fw *os.File

	f, err := os.Open("SP.xml")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	filename := "sp_out.xml"

	if checkFileIsExist(filename) { // Если файл существует
		// fw, _ = os.OpenFile(filename, os.O_APPEND, 0666) // Открыть файл
		os.Remove(filename)
		fw, _ = os.Create(filename) // Создать файл
		fmt.Println("Файл существует")
	} else {
		fw, _ = os.Create(filename) // Создать файл
		fmt.Println("Файл не существует")
	}

	fl := false

	// Чтение файла с ридером
	sc := bufio.NewScanner(f)

	w := bufio.NewWriter(fw) // Создаем новый объект Writer

	for sc.Scan() {
		if strings.Contains(sc.Text(), "<InformationRegisterRecordSet.СреднееПотребление>") {
			// if strings.Contains(sc.Text(), "<CatalogObject.ТочкиУчета>") {
			fl = true
		}

		if fl {
			// fmt.Println(sc.Text())

			_, _ = w.WriteString(sc.Text() + "\n")

		}

		if strings.Contains(sc.Text(), "</InformationRegisterRecordSet.СреднееПотребление>") {
			// if strings.Contains(sc.Text(), "</CatalogObject.ТочкиУчета>") {
			fl = false
		}

	}

	w.Flush()
}

func checkFileIsExist(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}
