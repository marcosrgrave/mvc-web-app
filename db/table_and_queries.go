package db

type Table struct {
	Name    string
	Columns []string
}

type Queries struct {
	SelectAll      string
	SelectByID     string
	UpdateByStruct string
	InsertByStruct string
	DeleteByID     string
}

func (t *Table) Queries() *Queries {
	var select_cols, update_cols, values_cols string

	for i, col := range t.Columns {
		if col == t.Columns[len(t.Columns)-1] {
			select_cols += col
			update_cols += col + "=?"
			values_cols += "?"
			break
		}
		if i != 0 {
			update_cols += col + "=?, "
		}
		values_cols += "?, "
		select_cols += col + ", "
	}

	select_all_query := "SELECT " + select_cols + " FROM " + t.Name
	select_query := "SELECT " + select_cols + " FROM " + t.Name + " WHERE " + t.Columns[0] + " = ?"
	update_query := "UPDATE " + t.Name + " SET " + update_cols + " WHERE " + t.Columns[0] + " = ?"
	insert_query := "INSERT INTO " + t.Name + " (" + select_cols + ") values (" + values_cols + ")"
	delete_query := "DELETE FROM " + t.Name + " WHERE " + t.Columns[0] + " = ?"

	return &Queries{
		SelectAll:      select_all_query,
		SelectByID:     select_query,
		UpdateByStruct: update_query,
		InsertByStruct: insert_query,
		DeleteByID:     delete_query,
	}
}
