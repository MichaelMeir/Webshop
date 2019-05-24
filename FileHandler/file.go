package FileHandler

import "io/ioutil"

func GetFile(root string, filename string) string {
	content, err := ioutil.ReadFile(root + filename)
	if err != nil {
		return err.Error()
	}
	return string(content)
}
