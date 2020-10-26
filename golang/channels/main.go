package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	// create channel
	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c) // use go routine to call checkLink
	}

	for l := range c {
		//time.Sleep(5 * time.Second) // this will block the main routine
		//go checkLink(l, c)

		// we decided to use a function literal instead. This will be more better so we
		// dont add the sleep inside checkLink
		go func(link string) {
			// channel gotcha. We have to pass arg to func otherwise parm to checkLink
			// will not be correct
			time.Sleep(time.Second * 5)
			checkLink(link, c)
		}(l)
	}
	// have to wrap channel calls in a loop to prevent blocking behaviour
	/*for {
		//fmt.Println(<-c)
		go checkLink(<-c, c)
	}*/
	/* receiving a message from a channel is a bloking call */
	//fmt.Println(<-c) // receive message from the channel
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " might be down")
		//c <- "Might be down I think" // we are sending message back to channel
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	//c <- "Yep its up"
	c <- link
}
