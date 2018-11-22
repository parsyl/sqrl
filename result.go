package sqrl

import (
	"database/sql"
)

func RowsAffected(res sql.Result, err error) (int64, error) {
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

func LastInsertId(res sql.Result, err error) (int64, error) {
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}
