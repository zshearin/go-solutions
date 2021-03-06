package csvformat

import(
	"bytes"
	"encoding/csv"
	"io"
	"os"
)

type Book struct {
	Author string
	Title string
}

type Books []Book


//Takes a set of Books and writes to an io.Writer
//it returns any errors
//The io.Write can be a file, stdout or a buffer
func (books *Books) ToCSV(w io.Writer) error {

	n := csv.NewWriter(w)
	err := n.Write([]string{"Author", "Title"})
	if err != nil {
		return err
	}
	for _, book := range *books {
		err := n.Write([]string{book.Author, book.Title})
		if err != nil {
			return err
		}
	}

	n.Flush()
	return n.Error()


}

//WriteCSVOutput initializes a set of books
//and writes to os.Stdout
func WriteCSVOutput() error {
	b := Books{
		Book{
			Author: "F Scott Fitzgerald",
			Title: "The Great Gatsby",
		},
		Book{
			Author: "J D Salinger",
			Title: "The Catcher in the Rye",
		},
	}
	return b.ToCSV(os.Stdout)
}


//WriteCSVBuffer returns a buffer csv for
//a set of books
func WriteCSVBuffer() (*bytes.Buffer, error) {

	b := Books {
		Book{
			Author: "F Scott Fitzgerald",
			Title: "The Great Gatsby",
		},
		Book{
			Author: "J D Salinger",
			Title: "The Catcher in the Rye",
		},
	}

	w := &bytes.Buffer{}
	err := b.ToCSV(w)
	return w, err
}