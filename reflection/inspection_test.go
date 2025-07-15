package reflection

import "testing"


type structType struct {
	id int
	name string
	height float32
}

type mapType map[string]interface{}

func TestVarDump(t *testing.T) {
	sliceStr := []string{"A", "B", "c"}
	structObj := structType {
		id: 1,
		name: "John",
		height: 12.8,
	}
	mapObj := mapType{
		"name": "Anna",
		"Age": 12,
	}
	sliceStruct := []structType {
		{id: 2, name: "Cl1", height: 12.3},
		{id: 3, name: "Cl2", height: 19.3},
	}

	VarDump(sliceStr)
	VarDump(structObj)
	VarDump(mapObj)
	VarDump(sliceStruct)
}