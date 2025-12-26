package migration

import (
	"encoding/json"
	"os"
)

func getNewsFileDump(fileLocation string) *newsData {

	b, err := os.ReadFile(fileLocation)
	if err != nil {
		panic(err)
	}

	d := newsData{}
	err = json.Unmarshal(b, &d)
	if err != nil {
		panic(err)
	}
	return &d
}
