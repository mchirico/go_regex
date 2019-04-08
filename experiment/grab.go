package main

import (
	"fmt"
	"github.com/mchirico/date/parse"
	"io/ioutil"
	"regexp"
	"strings"
)

type T interface {
	FindString(string) string
}

func Check(t T, s string, v *string) bool {
	if t.FindString(s) != "" {
		*v = t.FindString(s)
		return true
	}
	return false
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
	bond0 := regexp.MustCompile(`^bond0.*`)
	top := regexp.MustCompile(`^top -.*`)

	counter := 0

	rSlice := ""
	rDate := ""
	rBond := ""
	rTop := ""
	drop := ""
	ip := ""

	for _, v := range recs {

		if date.FindString(v) != "" {
			//rDate = date.FindString(v)
		}

		Check(slice, v, &rSlice)
		Check(date, v, &rDate)

		if Check(bond0, v, &rBond) {
			counter = -1
		}

		if Check(top, v, &rTop) {
			load := strings.Split(rTop, "load average:")
			drop = strings.ReplaceAll(drop, "\t", " ")
			drop = strings.ReplaceAll(drop, "  ", " ")
			drop = strings.ReplaceAll(drop, " ", ":")
			sdrop := strings.Split(drop, ":")

			tt, _ := parse.DateTimeParse(rDate).GetTime()
			// fmt.Println(sdrop,sdrop[0])
			fmt.Println(tt, ip, rSlice, sdrop[7], sdrop[9], sdrop[11], load[1])
		}

		counter += 1
		if counter == 1 {
			r := strings.Split(v, ":")
			r = strings.Split(r[1], " ")
			ip = r[0]

		}

		if counter == 4 {
			drop = v
		}
	}

	//fmt.Println(recs[0])
}
