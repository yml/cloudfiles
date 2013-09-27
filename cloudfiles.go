package main

import (
	"flag"
	"fmt"
	"github.com/ncw/swift"
	"os"
	"time"
)

const (
	ERROR_CODE = 1
)

var timeout_str string
var targz_path string
var container string

func init() {
	flag.StringVar(&timeout_str, "timeout", "10s", "Timeout in seconds")
	flag.StringVar(&targz_path, "file", "", "Complete path to the tar.gz file")
	flag.StringVar(&container, "container", "", "Targetted cloud file container")

}

func main() {
	flag.Parse()
	fmt.Println("Connecting to rackspace:", os.Getenv("UserName"))
	// Create a connection
	c := swift.Connection{
		UserName: os.Getenv("UserName"),
		ApiKey:   os.Getenv("ApiKey"),
		AuthUrl:  "https://identity.api.rackspacecloud.com/v1.0",
	}
	timeout, err := time.ParseDuration(timeout_str)
	if err != nil {
		fmt.Println("An errorr occured while parsing the duration for the connection timeout\n", err)
		os.Exit(ERROR_CODE)
	}
	c.Timeout = timeout
	fmt.Println("Connection Timeout set to ", c.Timeout)
	// Authenticate
	err = c.Authenticate()
	if err != nil {
		fmt.Println("An error occured while authenticated the User", c.UserName)
		os.Exit(ERROR_CODE)
	}

	// List all the containers
	containers, err := c.ContainerNames(nil)
	fmt.Println(containers)

	if targz_path == "" {
		fmt.Println("You must specify a .tar.gz file to upload")
		os.Exit(ERROR_CODE)
	}

	if container == "" {
		fmt.Println("You must specify a target container")
		os.Exit(ERROR_CODE)
	}

	file, err := os.Open(targz_path)
	if err != nil {
		fmt.Println("An error occured while opening the tar.gz file")
		os.Exit(ERROR_CODE)
	}
	defer file.Close()

	// Bulk upload to rackspace
	bulkUploadResult, err := c.BulkUpload(container, file, swift.UploadTarGzip, nil)
	if err != nil {
		fmt.Println("An error occured while bulk uploading the tar.gz\n", err)
		os.Exit(ERROR_CODE)
	}

	fmt.Println("Uploaded: ", bulkUploadResult)
}
