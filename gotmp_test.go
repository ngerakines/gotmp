package gotmp

import (
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path"
	"strconv"
	"testing"
)

var (
	counter = 1
)

func TestCreateDir(t *testing.T) {
	tfm := NewTemporaryFileManager()
	if tfm == nil {
		t.Error("tfm is nil")
	}
	base := makeTempDir()
	defer func() {
		os.RemoveAll(base)
	}()

	tmpFile := tfm.Create(createFile(base))
	for key, value := range tfm.List() {
		log.Println(key, value)
	}
	tmpFile.Release()
	for key, value := range tfm.List() {
		log.Println(key, value)
	}
}

func makeTempDir() string {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	counter = counter + 1
	base := path.Join(cwd, strconv.Itoa(counter))
	if err := os.Mkdir(base, 0777); err != nil {
		panic(err)
	}
	return base
}

func createFile(base string) string {
	handle, err := ioutil.TempFile(base, strconv.Itoa(counter))
	if err != nil {
		panic(err)
	}
	handle.WriteString(strconv.Itoa(rand.Int()))
	handle.Close()
	handle.Sync()
	return path.Join(base, handle.Name())
}
