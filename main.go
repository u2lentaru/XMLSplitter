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
	"time"
)

func main() {

	var fw *os.File

	t := time.Now()
	fmt.Println("Started: " + t.String())

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
		fmt.Println("Файл пересоздан")
	} else {
		fw, _ = os.Create(filename) // Создать файл
		fmt.Println("Файл создан")
	}

	fl := false

	// Чтение файла с ридером
	sc := bufio.NewScanner(f)

	w := bufio.NewWriter(fw) // Создаем новый объект Writer

	_, _ = w.WriteString(" <?xml version=\"1.0\" encoding=\"UTF-8\"?>\n")
	_, _ = w.WriteString(" <V8Exch:_1CV8DtUD xmlns:V8Exch=\"http://www.1c.ru/V8/1CV8DtUD/\" xmlns:v8=\"http://v8.1c.ru/data\" xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\">\n")
	_, _ = w.WriteString(" 	<V8Exch:Data>\n")

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

	_, _ = w.WriteString(" 	</V8Exch:Data>\n")
	_, _ = w.WriteString(" </V8Exch:_1CV8DtUD>\n")

	w.Flush()

	fw.Close()

	t = time.Now()
	fmt.Println("Stopped: " + t.Format("2006-01-02 15:04:05"))
}

func checkFileIsExist(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}
