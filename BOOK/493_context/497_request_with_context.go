package main

// INSTEAD OF Deprecated    transport.CancelRequest(req)
// We Use Request.WithContext !!!
// without any 'select' operator
// source: https://gist.github.com/superbrothers/dae0030c151d1f3c24311df77405169b

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

var (
	myUrl2 string
	delay2 int = 5
	w2     sync.WaitGroup
)

func connect2(t int) error {
	defer w2.Done()

	req, _ := http.NewRequest("GET", myUrl2, nil)

	ctx, cancel := context.WithTimeout(
		req.Context(), time.Duration(t)*time.Second)
	defer cancel()
	fmt.Println("AAA")

	req = req.WithContext(ctx)
	httpClient := http.DefaultClient
	response, err := httpClient.Do(req)

	if err != nil {
		fmt.Println(err)
		return err
	} else {
		realHTTPData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Error select:", err)
			return err
		}
		fmt.Printf("Server Response: %s\n", realHTTPData)
		return nil
	}
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Need a URL and a delay!")
		return
	}

	myUrl2 = os.Args[1]
	if len(os.Args) == 3 {
		t, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println(err)
			return
		}
		delay2 = t
	}

	fmt.Println("Delay:", delay2)

	//c := context.Background()
	//c, cancel := context.WithTimeout(c, time.Duration(delay2)*time.Second)
	//defer cancel()

	fmt.Printf("Connecting to %s \n", myUrl2)
	w2.Add(1)
	go connect2(delay2)
	w2.Wait()
	fmt.Println("Exiting...")
}
