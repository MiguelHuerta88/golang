package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// built in function that takes type and empty spaces to be initialized with
	/*bs := make([]byte, 99999)
	resp.Body.Read(bs)
	fmt.Println(string(bs))*/

	// another method to do the same as above
	//io.Copy(os.Stdout, resp.Body)

	// custom writer interface implementation
	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

// by us passing logWriter we are meeting the criteria for the interface since we are
// definiing the Write method
func (logWriter) Write(bs []byte) (int, error) {
	// this is incorrect since we didnt code it to how we expected. But GO let us do it because
	// we met the requirements for return types
	//return 1, nil

	/* CORRECT way to do it */
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
