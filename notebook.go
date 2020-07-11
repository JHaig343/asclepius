package main

import (
	// "fmt"
	// "net/http"
	// Local package + remote package from github
	"github.com/JHaig343/asclepius/ipynbparser"
	// "example.com/user/hello/morestrings"
	// "github.com/google/go-cmp/cmp"
)

func main() {

	notebook := ipynbparser.DecodeNotebook()

	notebook.PrintMarkdownContent()
	// notebook.Encode("test.ipynb")

	//MongoDB stuff
	// connect to database
	ctx, client := ipynbparser.MakeConnection()
	//Test ping the db to make sure we are really connected
	ipynbparser.TestPing(ctx, client)
	//List all the available databases
	ipynbparser.ListDBs(ctx, client)
	//Insert the Notebook object into the ipynbparser collection
	ipynbparser.InsertNotebook(ctx, client, notebook)
	//Finally, close the DB connection
	ipynbparser.CloseConnection(ctx, client)

}
