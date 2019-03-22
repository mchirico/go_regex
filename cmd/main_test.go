package main

import (
	"flag"
	"fmt"
	"github.com/mchirico/date/parse"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	SetupFunction()
	retCode := m.Run()
	TeardownFunction()
	os.Exit(retCode)
}

func SetupFunction() {
	flag.Parse()
}

func TeardownFunction() {

}

func TestReadFixtures(t *testing.T) {


	file := "../fixtures/fixture_a"
	raw, err := ioutil.ReadFile(file)
	if err != nil {
		t.FailNow()
	}
	if len(raw) > 200 {
		t.Logf("✅  Success!\n")
	} else {
		t.Logf("Could not read file")
	}

	data := string(raw)
	recs := strings.Split(data, "\n")
	re := regexp.MustCompile(`.*UTC 2019.*`)
	for idx, v:= range recs {
		r := re.FindString(v)
		if r != "" {

			tt, _ := parse.DateTimeParse(v).GetTime()
			fmt.Printf("%q %v \n", re.FindString(v),tt)

			fmt.Printf("    %s\n",recs[idx+1])
			fmt.Printf("    %s\n",recs[idx+2])
			fmt.Printf("    %s\n",recs[idx+6])
			fmt.Printf("    %s\n",recs[idx+7])
		}

	}
}


func TestQuick(t *testing.T) {
	s := "Thu Mar 21 19:07:52 UTC 2019"
	tt, err := parse.DateTimeParse(s).GetTimeLocSquish()
	fmt.Printf("_>%s<_ %v %v\n", s,tt,err)

}