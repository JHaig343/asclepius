package main

import (
	// "fmt"
	// "net/http"
	// Local package + remote package from github
	"fmt"
	"github.com/JHaig343/asclepius/ipynbparser"
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
	fmt.Println("Now retrieving notebook from DB")
	nb := ipynbparser.RetrieveNotebook(ctx, client)
	nb.Encode("mongo.json")
	ipynbparser.CloseConnection(ctx, client)

}
