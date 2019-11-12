package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const (
	EnvRepositoryID = "SUNSTONE_REPOSITORY_ID"
	EnvUsername     = "SUNSTONE_USERNAME"
	EnvAPIKey       = "SUNSTONE_API_KEY"
	EnvTemplatePath = "SUNSTONE_TEMPLATE_PATH"
)

// var strTemplatePath = flag.String("template", "", "Path to the template file")

// var strCommand = flag.String("template", "", "Webhook destination (https://example.com/webhooks)")

func main() {

	// flag.Usage = func() {
	// 	fmt.Printf("Usage of %s:\n", os.Args[0])
	// 	fmt.Printf("    actions --template app-deployment.yaml \n")
	// 	flag.PrintDefaults()
	// }

	if len(os.Args) == 1 {
		fmt.Println("usage:")
		fmt.Println("")
		fmt.Println("   sunstone-action USERNAME API_KEY REPOSITORY_ID TEMPLATE_PATH")
		os.Exit(1)
	}

	// flag.Parse()

	// if *strTemplatePath == "" && os.Getenv(EnvTemplatePath) == "" {
	// 	fmt.Println("both --template flag and SUNSTONE_TEMPLATE_PATH env variable cannot be empty, usage:")
	// 	flag.PrintDefaults()
	// 	os.Exit(1)
	// }

	// if os.Getenv(EnvUsername) == "" {
	// 	fmt.Printf("env variable %s not set", EnvUsername)
	// 	os.Exit(1)
	// }

	// if os.Getenv(EnvUsername) == "" {
	// 	fmt.Printf("env variable %s not set", EnvUsername)
	// 	os.Exit(1)
	// }

	// if os.Getenv(EnvAPIKey) == "" {
	// 	fmt.Printf("env variable %s not set", EnvAPIKey)
	// 	os.Exit(1)
	// }
	// if os.Getenv(EnvRepositoryID) == "" {
	// 	fmt.Printf("env variable %s not set", EnvRepositoryID)
	// 	os.Exit(1)
	// }

	if len(os.Args) != 5 {
		fmt.Println("incorrect number of arguments, wants 4 but got: ", len(os.Args))
		os.Exit(1)
	}

	username := os.Args[1]
	apiKey := os.Args[2]
	repoID := os.Args[3]
	templatePath := os.Args[4]

	templateFile, err := ioutil.ReadFile(templatePath)
	if err != nil {
		fmt.Printf("failed to read template '%s' file, error: %s\n", templatePath, err)

		fmt.Println("files in the current dir:")
		files, err := ioutil.ReadDir("./")
		if err != nil {
			log.Fatal(err)
		}

		for _, f := range files {
			fmt.Println(f.Name())
		}

		os.Exit(1)
	}

	patchRequest := privateRepositoryPatchRequest{
		Template: string(templateFile),
	}

	bts, err := json.Marshal(&patchRequest)
	if err != nil {
		fmt.Printf("failed to marshal template patch request, error: %s\n", err)
		os.Exit(1)
	}

	client := http.DefaultClient

	req, err := http.NewRequest(http.MethodPatch, "https://apps.sunstone.dev/v1/private-repositories/"+repoID, bytes.NewBuffer(bts))
	if err != nil {
		fmt.Printf("failed to construct request, error: %s\n", err)
		os.Exit(1)
	}

	req.SetBasicAuth(username, apiKey)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("update request failed, error: %s\n", err)
		os.Exit(1)
	}

	if resp.StatusCode != 200 {
		fmt.Printf("unexpected status code %d\n", resp.StatusCode)
	}

	fmt.Println("::set-output name=status::updated")
}

type privateRepositoryPatchRequest struct {
	Template string `json:"template"`
}
