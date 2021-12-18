package parse

type AlbumCSV struct {

	//Title,Album,Artist,Duration (ms),Rating,Play Count,Removed
	//"hackensack","the complete london collection - cd 2","thelonious monk","479000","0","0",""

	Title     string `csv:"Title"`
	Album     string `csv:"Album"`
	Artist    string `csv:"Artist"`
	Duration  string `csv:"Duration (ms)"`
	Rating    string `csv:"Rating"`
	PlayCount string `csv:"Play Count"`
	Removed   string `csv:"Removed"`
}
