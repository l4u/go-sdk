package main

import (
	"context"
	"fmt"

	"github.com/transloadit/go-sdk"
)

func main() {

	// Create client
	options := transloadit.DefaultConfig
	options.AuthKey = "TRANSLOADIT_KEY"
	options.AuthSecret = "TRANSLOADIT_SECRET"
	client, err := transloadit.NewClient(options)
	if err != nil {
		panic(err)
	}

	// Initialize new assembly
	assembly := transloadit.NewAssembly()

	// Add a file to upload
	assembly.AddFile("image", "../../fixtures/lol_cat.jpg")

	// Instructions will be read from the template
	// with specified id stored on Transloadit's servers.
	assembly.TemplateId = "TRANSLOADIT_TEMPLATE_ID"

	// Start the upload
	info, err := client.StartAssembly(context.Background(), assembly)
	if err != nil {
		panic(err)
	}

	// All files have now been uploaded and the assembly has started but no
	// results are available yet since the conversion has not finished.
	// The AssemblyWatcher provides functionality for polling until the assembly
	// has ended.
	waiter := client.WaitForAssembly(context.Background(), info.AssemblyUrl)
	info = <-waiter.Response

	fmt.Printf("You can view the result at: %s\n", info.Results["resize"][0].Url)
}
