package store

import (
	"database/sql"
	"reflect"
)

func ScanRowsToMaps(rows *sql.Rows) ([]map[string]interface{}, error) {
	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	buff := make([]interface{}, len(cols))
	for k := range cols {
		var v interface{}
		buff[k] = v
	}

	var data []map[string]interface{}
	for rows.Next() {
		if err := rows.Scan(buff...); err != nil {
			return nil, err
		}

		line := make(map[string]interface{}, len(cols))
		for k1, v1 := range cols {
			if buff[k1] == nil {
				continue
			}

			//reflect
			value := reflect.Indirect(reflect.ValueOf(buff[k1]))
			line[v1] = value.Interface()
		}
		data = append(data, line)
	}
	return data, nil
}
