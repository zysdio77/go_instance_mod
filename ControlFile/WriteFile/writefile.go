package WriteFile

import (
	"io/ioutil"
)

func WriteFile(filename string,data []byte)  error{
	err := ioutil.WriteFile(filename,data,0644)
	return err
}