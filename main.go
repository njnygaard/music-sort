package main

import (
	"io/ioutil"
	"path"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/njnygaard/music-sort/parse"
)

const LIBRARY_PATH = "/mnt/sdb1/Unorganized/Processed/Takeout/Google Play Music/Tracks"

type MusicTree map[string]map[string]uint

func main() {
	logger := logrus.New()

	files, err := ioutil.ReadDir(LIBRARY_PATH)
	if err != nil {
		logger.Fatal(err)
	}

	ArtistAlbums := make(MusicTree, 0)

	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".csv") {
			album, err := parse.ParseCSV(path.Join(LIBRARY_PATH, f.Name()))
			if err != nil {
				panic(err)
			}

			if ArtistAlbums[album[0].Artist] == nil {
				ArtistAlbums[album[0].Artist] = make(map[string]uint, 0)
			}
			ArtistAlbums[album[0].Artist][album[0].Album]++
		}
	}

	// logger.Info(spew.Sdump(ArtistAlbums["Steely Dan"]))

	display(ArtistAlbums)
}

func display(collection MusicTree) {

	logger := logrus.New()

	for artist, albums := range collection {
		logger.Info(artist)
		for album, expectedSongCount := range albums {
			logger.Warnf("\nAlbum: %s\n\tExpected Song Count: %d", album, expectedSongCount)
		}
	}
}
