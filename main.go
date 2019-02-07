package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Job struct {
	color string
	name  string
	url   string
}

var result []Job

func main() {
//	Example:  url := "https://jenkinshost.com:8080/view/All/api/json"
        url := "https://localhost:8080/view/All/api/json"
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Fatal(err)
	}
	resp, _ := http.DefaultClient.Do(req)
	defer resp.Body.Close()
	body, berr := ioutil.ReadAll(resp.Body)
	if berr != nil {
		fmt.Println(berr)
	}
	m := map[string]interface{}{}
	if jerr := json.Unmarshal(body, &m); jerr != nil {
		panic(err)
	}

	for ky, vl := range m {
		if ky == "jobs" {
			jobs := vl.([]interface{})

			for _, v := range jobs {
				demo := v.(map[string]interface{})
				conversion(demo)
			}

		}
	}
	for _, j := range result {
			fmt.Println(j.name)
	}
}

func conversion(p map[string]interface{}) {
	var mine Job
	for x, y := range p {
		switch x {
		case "name":
			str := fmt.Sprintf("%v", y)
			mine.name = str
		case "color":
			str := fmt.Sprintf("%v", y)
			mine.color = str
		case "url":
			str := fmt.Sprintf("%v", y)
			mine.url = str
		}
	}
	result = append(result, mine)

}
