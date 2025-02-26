package database

type Table struct {
	TableName		string
	PrimaryKey		string
	AffectedRows	int64
	AffectedCols	int64
	LastInsertedId	int64
}

type Fields map[string]interface{}
type Conditions map[string]interface{}

func (tb *Table) Insert(data Fields)  error {
	return nil
}

func (tb *Table) Update(data Fields, conditions Conditions)  error {
	return nil
}

func (tb *Table) Delete(conditions Conditions)  error {
	return nil
}

