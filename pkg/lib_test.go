package pkg

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestF_Init(t *testing.T) {
	f := NewF("../fixtures/slicestore")

	fmt.Println(f.Finfo.Size())

	for {
		b, err := f.Read()
		if err != nil {
			fmt.Println(string(b), err)
			if f.offset != 3448407 {
				t.Fatalf("Error reading file:%v", f.offset)
			}
			break
		}
	}
}

func TestIdea(t *testing.T) {
	f := NewF("../fixtures/slicestore")

	reg := NewReg()
	var regLooper = &reg

	Looper(f, regLooper)

	fp, _ := os.Open("../fixtures/slicestore")
	defer fp.Close()

	b := make([]byte, 3448407)
	fp.Read(b)
	lines := strings.Split(string(b), "\n")
	for i, v := range reg.idx {
		if strings.ContainsAny(lines[v], "UTC 2019") != true {
			t.Fatalf("Should have found date: %s, %d\n", lines[v], i)
		}
	}

	// Here's how we get a date value

	s, _ := ReadIdx("../fixtures/slicestore", reg.seg[14].offset,
		reg.seg[14].size)
	if s[reg.seg[14].idx[1]] != "Sat Apr  6 22:08:02 UTC 2019" {
		t.Fatalf("Didn't get date:%v\n", s[reg.seg[14].idx[1]])
	}
	fmt.Println(s[reg.seg[14].idx[1]])

}

func SkipTestParseData(t *testing.T) {

	//file := "../fixtures/fixture_10000"

	file := "/Users/mchirico/go_tempStuff/src/github.com/mchirico/go_tempStuff/1/northlake.txt"
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
	ParseData(data)
}
