package main

import (
	"flag"
	"io/ioutil"
	"os"
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
	data, err := ioutil.ReadFile(file)
	if err != nil {
		t.FailNow()
	}
	if len(data) > 200 {
		t.Logf("✅  Success!\n")
	} else {
		t.Logf("Could not read file")
	}
}