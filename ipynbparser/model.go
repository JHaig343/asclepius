package ipynbparser

// `outputs` object
// represents the result of running the code in the cell
type output struct {
	Name        string   `json:"name,omitempty" bson:"name,omitempty"`
	Output_type string   `json:"output_type" bson:"output_type,omitempty"`
	Text        []string `json:"text,omitempty" bson:"text,omitempty"`
	// error fields that appear when code run in the notebook has an exception
	Ename     string   `json:"ename,omitempty" bson:"ename,omitempty"`
	Evalue    string   `json:"evalue,omitempty" bson:"evalue,omitempty"`
	Traceback []string `json:"traceback,omitempty" bson:"traceback,omitempty"`
}

//Cell structure of each individual cell in the notebook
type Cell struct {
	// Cell type can be "code" or "markdown"
	Cell_type string `json:"cell_type" bson:"cell_type"`

	Metadata cellmetadata `json:"metadata" bson:"metadata"`
	// execution_count is a pointer; numeric value if execution count
	//  > 0, null if null in the notebook
	Execution_count *int     `json:"execution_count" bson:"execution_count"`
	Outputs         []output `json:"outputs" bson:"outputs"`
	// The source code within the cell; an array of strings,
	// each string a line of code
	Source []string `json:"source" bson:"source"`
}

// `kernelspec` object
// data on the kernel being used for the notebook
type Kernel struct {
	Display_name string `json:"display_name" bson:"display_name"`
	Language     string `json:"language" bson:"language"`
	Name         string `json:"name" bson:"name"`
}

// `language_info` object
type Info struct {
	Codemirror_mode string `json:"codemirror_mode" bson:"codemirror_mode"`
	File_extension  string `json:"file_extension" bson:"file_extension"`
	Mimetype        string `json:"mimetype" bson:"mimetype"`
	Name            string `json:"name" bson:"name"`
	Pygments_lexer  string `json:"pygments_lexer" bson:"pygments_lexer"`
	Version         string `json:"version" bson:"version"`
}

// `metadata` object
type Meta struct {
	// 2 sub-objects, make separate structs for them
	// inlining structs doesn't work
	Kernelspec    Kernel `json:"kernelspec" bson:"kernelspec"`
	Language_info Info   `json:"language_info" bson:"language_info"`
}

// metadata object within each 'cell': usually empty from what I've seen
// An empty struct allows us to get the '{}' field in JSON
type cellmetadata struct{}

// Foobar is an example of an embedded struct;
type Foobar struct {
	Name string
	Age  int
}

// Notebook a struct modeling the structure of the underlying JSON of a Jupyter Notebook file
// NOTE: capitalize fields so they are exported (as go Unmarshal requires they are to automatically populate)
type Notebook struct {
	Cells          []Cell `json:"cells" bson:"cells"`
	Metadata       Meta   `json:"metadata" bson:"metadata"`
	Nbformat       int    `json:"nbformat" bson:"nbformat"`
	Nbformat_minor int    `json:"nbformat_minor" bson:"nbformat_minor"`
	//Foobar         Foobar `gorm:"embedded" bson:"embedded"`
}
