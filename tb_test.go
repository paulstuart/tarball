package tarball

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
)

func TestCreate(t *testing.T) {
	buf := &bytes.Buffer{}
	if err := Create("tb*.go", buf); err != nil {
		t.Fatal(err)
	}
	w := ioutil.Discard
	if testing.Verbose() {
		w = os.Stdout
	}
	if err := ReadAll(buf, w); err != nil {
		t.Fatal(err)
	}
}

func TestReadFile(t *testing.T) {
	const filename = "tb_test.go"
	buf := &bytes.Buffer{}
	if err := Create("tb*.go", buf); err != nil {
		t.Fatal(err)
	}
	onDisk, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatal(err)
	}
	out := &bytes.Buffer{}
	err = ReadFile(buf, out, filename)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(onDisk, out.Bytes()) {
		t.Fatal(err)
	}
}
