package parse

import (
	"os"

	"github.com/gocarina/gocsv"
)

func ParseCSV(path string) (album []AlbumCSV, err error) {
	in, err := os.Open(path)
	if err != nil {
		return
	}
	defer in.Close()

	if err = gocsv.UnmarshalFile(in, &album); err != nil {
		return
	}

	return
}
