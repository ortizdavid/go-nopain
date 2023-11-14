package filemanager

import (
	"testing"
)

type folderTest struct {
	folderName string
	expected bool
}

type fileTest struct {
	fileName string
	folderName string
	content string
	expected bool
}

var fileManager FileManager

func TestCreateSingleFolder(t *testing.T) {
	testCases := [] folderTest {
		{ folderName: "Files", expected: true },
		{ folderName: "Moved", expected: true },
		{ folderName: "Single-Files", expected: true },
		{ folderName: "Many-Files", expected: true },
		{ folderName: "For-Remove", expected: true },
	}
	for _, test := range testCases {
		got := fileManager.CreateSingleFolder(test.folderName)
		if test.expected != got {
			t.Error(test.folderName, " not created. Got: ", got, ". Expected: ",  test.expected)
		}
	}
}

func TestCreateManyFolders(t *testing.T) {
	testCases := [] folderTest {
		{ folderName: "Folder/f1/f2/f3", expected: true },
		{ folderName: "Musics/mp3", expected: true },
		{ folderName: "Documents/pdf", expected: true },
		{ folderName: "Documents/docx", expected: true },
		{ folderName: "Images/jpg/", expected: true },
		{ folderName: "Images/png", expected: true },
		{ folderName: "Others/sub1/sub2/sub3", expected: true },
	}
	for _, test := range testCases {
		got := fileManager.CreateManyFolders(test.folderName)
		if test.expected != got {
			t.Error(test.folderName, " not created. Got: ", got, ". Expected: ",  test.expected)
		}
	}
}

func TestCreateSingleFile(t *testing.T) {
	dirName := "Files"
	testCases := [] fileTest {
		{ fileName: "file.csv", folderName: dirName, expected: true },
		{ fileName: "file.log", folderName: dirName, expected: true },
		{ fileName: "file.txt", folderName: dirName, expected: true },
		{ fileName: "file.json", folderName: dirName, expected: true },
		{ fileName: "remove1.txt", folderName: "For-Remove", expected: true },
		{ fileName: "remove2.txt", folderName: "For-Remove", expected: true },
	}
	for _, test := range testCases {
		got := fileManager.CreateSingleFile(test.folderName, test.fileName)
		if test.expected != got {
			t.Error(test.fileName, " not created. Got: ", got, ". Expected: ",  test.expected)
		}
	}
}

func TestCreateManyFiles(t *testing.T) {
	dirName := "Many-Files"
	got := fileManager.CreateManyFiles(dirName, "file.txt", "file.log", "file.csv", "file.docx", "file.json")
	expected := true
	if expected != got {
		t.Error("Files not created. Got: ", got, ". Expected: ",  expected)
	}
}

func TestWriteFile(t *testing.T) {
	dirName := "Files"
	testCases := [] fileTest {
		{ fileName: "file.txt",  folderName: dirName, content: "Hello World", expected: true },
		{ fileName: "file.log",  folderName: dirName, content: " 2023-01-01 19:00:04, INFO:  User admin created!", expected: true },
		{ fileName: "file.csv",  folderName: dirName, content: "John, 2022-02-02, Luanda", expected: true },
		{ fileName: "file.json",  folderName: dirName, content: `{ "message": "Hello World", "error": true }`, expected: true },
	}
	for _, test := range testCases {
		got := fileManager.WriteFile(test.folderName, test.fileName, test.content)
		if test.expected != got {
			t.Error(test.fileName, " not created. Got: ", got, ". Expected: ",  test.expected)
		}
	}
}

func TestRemoveFolder(t *testing.T) {
	testCases := [] folderTest {
		{ folderName: "Musics/mp3", expected: true },
		{ folderName: "Images/jpg/", expected: true },
		{ folderName: "Images/png", expected: true },
		{ folderName: "Others/sub1/sub2/sub3", expected: true },
	}
	for _, test := range testCases {
		got := fileManager.RemoveFolder(test.folderName)
		if test.expected != got {
			t.Error(test.folderName, " not removed. Got: ", got, ". Expected: ",  test.expected)
		}
	}
}


func TestRemoveFile(t *testing.T) {
	dirName := "For-Remove"
	testCases := [] fileTest {
		{ fileName: "remove1.txt", folderName: dirName, expected: true },
		{ fileName: "remove2.txt", folderName: dirName, expected: true },
	}
	for _, test := range testCases {
		got := fileManager.RemoveFile(test.folderName, test.fileName)
		if test.expected != got {
			t.Error(test.fileName, " not removed. Got: ", got, ". Expected: ",  test.expected)
		}
	}
}

func TestMoveFile(t *testing.T) {
	origin := "Files"
	destination := "Moved"
	testCases := [] fileTest {
		{ fileName: "remove1.txt", folderName: origin, expected: true },
		{ fileName: "remove2.txt", folderName: origin, expected: true },
	}
	for _, test := range testCases {
		got := fileManager.MoveFile(test.fileName, test.folderName, destination)
		if test.expected != got {
			t.Error(test.fileName, " not moved. Got: ", got, ". Expected: ",  test.expected)
		}
	}
}
