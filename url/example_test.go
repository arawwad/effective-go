package url_test

import (
	"fmt"
	"log"

	"github.com/arawwad/effective-go/url"
)

func ExampleURL() {
	u, err := url.Parse("http://foo.com/go")

	if err != nil {
		log.Fatal(err)
	}

	u.Scheme = "https"
	fmt.Println(u)
	// output:
	// https://foo.com/go
}
