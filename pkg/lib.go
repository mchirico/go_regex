package pkg

import (
	"errors"
	"fmt"
	"github.com/mchirico/date/parse"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func findRet(b []byte, n int) (int, error) {
	p := 0
	for i := n - 1; n >= 0; i-- {
		if b[i] == '\n' {
			p = i
			break
		}
	}
	if p == 0 {
		return 0, errors.New("No return found. (findRet)")
	}

	return p, nil
}

type F struct {
	File   string
	size   int
	B      []byte
	P      int // Pointer to return
	N      int // Number of bytes return from read
	O      int64 // New offset
	offset int64  // Offset one behind
	Finfo  os.FileInfo
}

func NewF(file string) F {
	f := F{}
	f.File = file
	f.size = 200000
	f.init()
	return f
}

func (f *F) init() error {
	finfo, err := os.Stat(f.File)
	if err != nil {
		return err
	}
	f.Finfo = finfo
	return nil
}

func (f *F) Read() ([]byte, error) {

	f.offset += int64(f.P)
	if f.offset+1 >= f.Finfo.Size() {
		return nil, errors.New("EOF")
	}

	fp, err := os.Open(f.File)
	defer fp.Close()

	if err != nil {
		return []byte{}, err
	}

	o, err := fp.Seek(f.offset, 0)
	if err != nil {
		return []byte{}, err
	}

	b := make([]byte, f.size)
	n, err := fp.Read(b)

	p, err := findRet(b, n)

	f.B = b
	f.N = n  // Number of bytes on read
	f.O = o  // New offset
	f.P = p  // Pointer to return

	return f.B[0:f.P], err

}

func ReadIdx(file string, offset int64, size int) ([]string, error) {
	f, err := os.Open(file)
	defer f.Close()
	if err != nil {
		return nil, err
	}
	_, err = f.Seek(offset, 0)
	if err != nil {
		return nil, err
	}

	b := make([]byte, size)
	_, err = f.Read(b)
	if err != nil {
		return nil, err
	}

	return strings.Split(string(b), "\n"), nil

}

type Seg struct {
	size   int
	lines  int
	offset int64
	idx    []int
}

type Reg struct {
	idx        []int
	reg        string
	seg        []Seg
	lastLength int
}

func NewReg() Reg {
	r := Reg{}
	r.idx = []int{}
	r.seg = []Seg{}
	r.reg = `^[A-Z][a-z][a-z] [A-Z][a-z][a-z][ ]{1,}[0-9]{1,}[ ][0-9]{1,}:[0-9]{1,}:[0-9]{1,} UTC 20[0-9]{2}$`
	return r
}

func (r *Reg) FindDateHeading(b []byte) {
	re := regexp.MustCompile(r.reg)
	result := strings.Split(string(b), "\n")
	lines := len(result)
	size := len(b)
	seg := Seg{}
	seg.lines = lines
	if len(r.seg) >= 1 {
		seg.offset = int64(r.seg[len(r.seg)-1].size) + r.seg[len(r.seg)-1].offset
	}

	seg.size = size
	for i, v := range result {
		result := re.FindAllStringSubmatchIndex(v, -1)
		if len(result) == 1 {
			r.idx = append(r.idx, i+r.lastLength)
			seg.idx = append(seg.idx, i)
		}
	}

	r.seg = append(r.seg, seg)
	r.lastLength += lines - 1
}

func (r *Reg) cmd(b []byte) {
	r.FindDateHeading(b)
}

type looper interface {
	cmd([]byte)
}

func Looper(f F, l looper) {
	for {
		b, err := f.Read()
		l.cmd(b)
		if err != nil {
			break
		}
	}
}

func ParseData(data string) {

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
			s += fmt.Sprintf("%s, bond0, %s, %s", dateF, nic, msg)

		}
		r2 := rereject.FindString(v)
		if r2 != "" {
			ss := strings.Split(recs[idx], " ")
			if len(ss) > 10 {
				s += fmt.Sprintf("%s,%s\n", ss[0], ss[2])

			}
		}

	}

	d1 := []byte(s)
	ioutil.WriteFile("out.csv", d1, 0644)

}
