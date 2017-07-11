package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Trying the ifttt maker channel.")

	var key = flag.String("key", "", "API key for the ifttt maker channel.")
	flag.Parse()

	maker := NewMaker(*key)
	maker.Send("test4", map[string]string{"foo": "bar"})
}

const URL = "https://maker.ifttt.com/trigger/%s/with/key/%s"

type Maker struct {
	key string
}

func NewMaker(key string) *Maker {
	if key == "" {
		fmt.Println("Can't have zero length key")
		os.Exit(1)
	}
	return &Maker{key}
}

func (m *Maker) Send(event string, data map[string]string) {
	url := fmt.Sprintf(URL, event, m.key)
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(data)
	resp, err := http.Post(url, "application/json", b)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}
