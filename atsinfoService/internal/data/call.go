package data

import (
	"context"
	"database/sql"
	"time"
)

const timeFormat = "2006.01.02 15:04:05"

type Call struct {
	Id       string `json:"Id"`
	Ext      string
	Trk      int
	Calldate string
	Duration int
	Fg       string
	Cost     float32
	Ldate    string
	T        int

	Auth        string `json:"-"`
	Dialeddigit string
	AccountCode string
	ClipNumber  string
	ClipName    string
}

type CallModel struct {
	DB *sql.DB
}

func (m CallModel) GetAll(direction, subPhone, gateway, callPhone, intercity string, fromDate, toDate time.Time) ([]*Call, error) {
	// query := `SET DATEFORMAT ymd;
	// 	SELECT T, EXT, AUTH, TRK, convert(varchar(20),CALLDATE,20), DURATION, FG, DIALEDDIGIT, ACCOUNTCODE, COST, CLIPNUMBER, CLIPNAME, LDATE, ID
	// 	FROM ats_Calls as calls
	// 	WHERE (calls.CALLDATE BETWEEN ? AND ?) AND (calls.FG=? OR ?='')
	// 		AND (calls.EXT LIKE ? OR ?='') AND (calls.TRK LIKE ? OR ?='') AND (calls.DIALEDDIGIT LIKE ? OR calls.TRK LIKE ? OR ?='')
	// 		AND (LEN(calls.DIALEDDIGIT)>5 OR ?=0)
	// 	ORDER BY ID DESC`

	query := `SET DATEFORMAT ymd;
		SELECT T, EXT, AUTH, TRK, convert(varchar(20),CALLDATE,20), DURATION, FG, DIALEDDIGIT, ACCOUNTCODE, COST, CLIPNUMBER, CLIPNAME, LDATE, ID
		FROM ats_Calls as calls 
		WHERE (calls.CALLDATE BETWEEN ? AND ?) AND (calls.FG=? OR ?='')
			AND (calls.EXT LIKE ? OR ?='') AND (calls.TRK LIKE ? OR ?='') AND (calls.DIALEDDIGIT LIKE ? OR calls.TRK LIKE ? OR ?='')
			AND (LEN(calls.DIALEDDIGIT)>5 OR ?=0)
		ORDER BY ID DESC`

	args := []any{fromDate.Format(timeFormat), toDate.Format(timeFormat), direction, direction, subPhone, subPhone, gateway, gateway, callPhone, callPhone, callPhone, 0}
	//args := []any{fromDate.Format(timeFormat), toDate.Format(timeFormat), direction, subPhone, subPhone, gateway, gateway, callPhone, callPhone, callPhone, intercity}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	calls := []*Call{}

	for rows.Next() {
		var input struct {
			Id       string
			Ext      string
			Trk      int
			Calldate string
			Duration int
			Fg       string
			Cost     float32
			Ldate    string
			T        int

			Auth        sql.NullString
			Dialeddigit sql.NullString
			AccountCode sql.NullString
			ClipNumber  sql.NullString
			ClipName    sql.NullString
		}
		// var call Call

		// err := rows.Scan(
		// 	&call.T,
		// 	&call.Ext,
		// 	&call.Auth,
		// 	&call.Trk,
		// 	&call.Calldate,
		// 	&call.Duration,
		// 	&call.Fg,
		// 	&call.Dialeddigit,
		// 	&call.AccountCode,
		// 	&call.Cost,
		// 	&call.ClipNumber,
		// 	&call.ClipName,
		// 	&call.Ldate,
		// 	&call.Id,
		// )

		err := rows.Scan(
			&input.T,
			&input.Ext,
			&input.Auth,
			&input.Trk,
			&input.Calldate,
			&input.Duration,
			&input.Fg,
			&input.Dialeddigit,
			&input.AccountCode,
			&input.Cost,
			&input.ClipNumber,
			&input.ClipName,
			&input.Ldate,
			&input.Id,
		)
		if err != nil {
			return nil, err
		}
		var call = Call{
			T:           input.T,
			Ext:         input.Ext,
			Auth:        input.Auth.String,
			Trk:         input.Trk,
			Calldate:    input.Calldate,
			Duration:    input.Duration,
			Fg:          input.Fg,
			Dialeddigit: input.Dialeddigit.String,
			AccountCode: input.AccountCode.String,
			Cost:        input.Cost,
			ClipNumber:  input.ClipNumber.String,
			ClipName:    input.ClipName.String,
			Ldate:       input.Ldate,
			Id:          input.Id,
		}

		calls = append(calls, &call)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return calls, nil
}
