package tarball

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
)

func TestCreateGlob(t *testing.T) {
	buf := &bytes.Buffer{}
	if err := Create(buf, "tb*.go"); err != nil {
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

func TestCreateList(t *testing.T) {
	buf := &bytes.Buffer{}
	if err := Create(buf, "tb.go", "tb_test.go"); err != nil {
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
	if err := Create(buf, "tb*.go"); err != nil {
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

func TestList(t *testing.T) {
	buf := &bytes.Buffer{}
	if err := Create(buf, "tb*.go"); err != nil {
		t.Fatal(err)
	}
	list, err := FileList(buf)
	if err != nil {
		t.Fatal(err)
	}
	const want = 2
	if len(list) != want {
		t.Fatalf("want %d files -- got %d\n", want, len(list))
	}
}
