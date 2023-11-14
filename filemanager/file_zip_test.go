package filemanager

import (
	"testing"
)

type fileOrFolder struct {
	source string
	target string
	expected any
}

var fileZip FileZip

func TestZipFile(t *testing.T) {
	testCases := [] fileOrFolder {
		{ source: "Files/file.csv", target: "Files/file.csv.zip", expected: nil },
		{ source: "Files/file.txt", target: "Files/file.txt.zip", expected: nil },
		{ source: "Files/file.json", target: "Files/file.json.zip", expected: nil },
		{ source: "Files/file.log", target: "Files/file.log.zip", expected: nil },
	}
	for _, test := range testCases {
		got := fileZip.ZipFileOrFolder(test.source, test.target)
		if test.expected  != got {
			t.Error(test.source, " not Zipped. Got: ", got, "Expected: ",  test.expected)
		}
	}
}

func TestZipFolder(t *testing.T) {
	testCases := [] fileOrFolder {
		{ source: "Documents", target: "Documents.zip", expected: nil },
		{ source: "Images", target: "Images.zip", expected: nil },
		{ source: "Musics", target: "Musics.zip", expected: nil },
		{ source: "Others", target: "Others.zip", expected: nil },
	}
	for _, test := range testCases {
		got := fileZip.ZipFileOrFolder(test.source, test.target)
		if test.expected != got {
			t.Error(test.source, " not Zipped. Got: ", got, "Expected: ",  test.expected)
		}
	}
}