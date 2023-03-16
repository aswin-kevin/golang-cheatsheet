package get_html

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Get_html() {
	urls := []string{"https://securin.io"}
	for _, url := range urls {
		fmt.Printf("%s is going in ...\n", url)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("Got some error in step 1")
			return
		}
		body, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Println("Got some error in step 2")
			return
		}
		fmt.Printf("%s", body)
	}
}
