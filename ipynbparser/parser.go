package ipynbparser

import (
	"encoding/json"
	"fmt"
	// "github.com/jinzhu/gorm"
	"io/ioutil"
	"strings"
)

type message struct {
	Name string
	Body string
	Time int64
}

// JSON go package can decode into an empty interface
// if you don't know the structure of your JSON data
var i interface{}

// TestJSONEncoding tests out json.Marshal(), which encodes an object to JSON
func TestJSONEncoding() {
	m := message{"Jacob", "HelloWorld", 45}

	b, _ := json.Marshal(m)
	// convert to string to see legible representation
	fmt.Println(string(b))
}

// DecodeNotebookNoStruct takes a .ipynb file (which under the hood is just a JSON config) and decodes
// it into a Go struct
func decodeNotebookNoStruct() {
	dat, err := ioutil.ReadFile("./ipynbparser/test/TestJavaNotebook.ipynb")
	if err != nil {
		panic(err)
	}

	var nb interface{}

	encodingErr := json.Unmarshal(dat, &nb)

	if encodingErr != nil {
		panic(encodingErr)
	}

	m := nb.(map[string]interface{})
	// this goes into the first object in the 'cells' array and accesses the 'cell_type' val
	// NOTE: probably easier to set up a struct with matching fields to do this...
	fmt.Println(m["cells"].([]interface{})[0].(map[string]interface{})["cell_type"])
}

// DecodeNotebook takes a .ipynb file (which under the hood is just a JSON config) and decodes
// it into a Go struct
func DecodeNotebook() Notebook {
	dat, err := ioutil.ReadFile("./ipynbparser/test/TestJavaNotebook.ipynb")
	if err != nil {
		panic(err)
	}

	var nb Notebook
	err = json.Unmarshal(dat, &nb)

	return nb
}

// PrintMarkdownContent goes through each cell in the notebook and prints the
// contents of each cell marked as "markdown"
func (nb Notebook) PrintMarkdownContent() {
	for _, v := range nb.Cells {
		if v.Cell_type == "markdown" {
			for _, line := range v.Source {
				fmt.Println(line)
			}
		}
	}
}

// BuildSourceCodeFile takes the source blocks from the Jupyter notebook
// and combines them together into a single string
// (which can then be written to a source file)
func BuildSourceCodeFile(nb Notebook) string {
	var fileoutput strings.Builder
	for _, v := range nb.Cells {
		for _, vv := range v.Source {
			fileoutput.WriteString(strings.TrimSpace(vv) + "\n")
		}
	}

	return fileoutput.String()
}

// Encode writes the Notebook struct to a Jupyter .ipynb file
func (nb Notebook) Encode(filename string) {
	file, err := json.MarshalIndent(nb, "", " ")
	if err != nil {
		fmt.Println("Error serializing Notebook object")
		panic(err)
	}
	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		fmt.Println("Error writng object to file")
		panic(err)
	}
}
