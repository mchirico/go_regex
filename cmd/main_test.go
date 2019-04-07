package main

import (
	"flag"
	"fmt"
	"github.com/mchirico/date/parse"
	"io"
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

	s := ""
	for idx, v := range recs {
		r := re.FindString(v)
		if r != "" {

			tt, _ := parse.DateTimeParse(v).GetTime()
			nic := recs[idx+1]

			msg := strings.Trim(recs[idx+6], " ")
			msg = strings.Replace(msg, ":", ",", -1)
			msg = strings.Replace(msg, " ", ",", -1)

			dateF := tt.Format("2006-01-02 15:04:05")

			//fmt.Printf("%s, bond0, %s, %s \n",dateF,nic,msg)
			s += fmt.Sprintf("%s, bond0, %s, %s \n", dateF, nic, msg)

			//fmt.Printf("    %s\n",recs[idx+1])
			//fmt.Printf("    %s\n",recs[idx+2])
			//fmt.Printf("    %s\n",recs[idx+6])
			//fmt.Printf("    %s\n",recs[idx+7])
			//fmt.Printf("    %s\n",recs[idx+20])
		}

	}

	d1 := []byte(s)
	ioutil.WriteFile("out.csv", d1, 0644)

}

// fixture_10000

func TestReadFixture1000(t *testing.T) {

	file := "../fixtures/fixture_10000"
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

	rereject := regexp.MustCompile(`.*REJECT.*`)

	s := ""
	for idx, v := range recs {

		r := re.FindString(v)
		if r != "" {

			tt, _ := parse.DateTimeParse(v).GetTime()
			nic := recs[idx+1]

			msg := strings.Trim(recs[idx+6], " ")
			msg = strings.Replace(msg, ":", ",", -1)
			msg = strings.Replace(msg, " ", ",", -1)

			dateF := tt.Format("2006-01-02 15:04:05")

			//fmt.Printf("%s, bond0, %s, %s \n",dateF,nic,msg)
			s += fmt.Sprintf("%s, bond0, %s, %s", dateF, nic, msg)

			//fmt.Printf("    %s\n",recs[idx+1])
			//fmt.Printf("    %s\n",recs[idx+2])
			//fmt.Printf("    %s\n",recs[idx+6])
			//fmt.Printf("    %s\n",recs[idx+7])
			//fmt.Printf("    %s\n",recs[idx+20])
		}
		r2 := rereject.FindString(v)
		if r2 != "" {
			ss := strings.Split(recs[idx], " ")
			if len(ss) > 10 {
				//fmt.Println(ss[0], ss[2])
				s += fmt.Sprintf("%s,%s\n", ss[0], ss[2])

			}
		}

	}

	d1 := []byte(s)
	ioutil.WriteFile("out.csv", d1, 0644)

}

func TestQuick(t *testing.T) {
	s := "Thu Mar 21 19:07:52 UTC 2019"
	tt, err := parse.DateTimeParse(s).GetTimeLocSquish()
	if 1 == 2 {
		fmt.Printf("_>%s<_ %v %v\n", s, tt, err)
	}

}

func NoTestRead(t *testing.T) {

	const BufferSize = 1000000
	file, err := os.Open("/Users/mchirico/ibm/stormgr/northlake.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	buffer := make([]byte, BufferSize)

	for {
		//bytesread, err := file.Read(buffer)
		_, err := file.Read(buffer)

		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}

			break
		}

		//fmt.Println("bytes read: ", bytesread)

	}
	//fmt.Println("bytestream to string: ", string(buffer[:1000000]))

}
