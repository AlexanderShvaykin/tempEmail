package test

import (
	"io/ioutil"
	"path"
	"runtime"
)

func Fixture(name string) (string, error) {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Dir(filename)
	file, err := ioutil.ReadFile(dir + "/fixtures/" + name)
	if err != nil {
		return "", err
	}
	return string(file), nil
}
