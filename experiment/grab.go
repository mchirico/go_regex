package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

type T interface {
	FindString(string) string
}

func Check(t T, s string, v *string) {
	if t.FindString(s) != "" {
		*v = t.FindString(s)
	}
}

func main() {
	dat, err := ioutil.ReadFile("./fixtures/northStats")
	if err != nil {
		fmt.Println(err)
		return
	}
	recs := strings.Split(string(dat), "\n")

	date := regexp.MustCompile(`.*UTC 2019.*`)
	slice := regexp.MustCompile(`^slice.*`)
	//bond0 := regexp.MustCompile(`^bond0.*`)

	counter := 0

	rSlice := ""
	rDate := ""
	//rBond := ""

	for idx, v := range recs {

		if date.FindString(v) != "" {
			//rDate = date.FindString(v)
		}

		Check(slice, v, &rSlice)
		Check(date, v, &rDate)
		fmt.Println(rSlice, rDate)
		if rDate != "" {
			counter = -1
			fmt.Println(idx, rDate)
		}
		counter += 1
	}

	//fmt.Println(recs[0])
}
