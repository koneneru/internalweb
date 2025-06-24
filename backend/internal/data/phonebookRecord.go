package data

import (
	"context"
	"database/sql"
	"time"
)

type phonebookRecord struct {
	employeeId    int64  `json:"employeeId"`    // Табельный номер
	name          string `json:"name"`          // ФИО
	department    string `json:"department"`    // Подразделение
	employ        string `json:"employ"`        // Должность
	internalPhone string `json:"internalPhone"` // Внутренний телефон
	email         string `json:"email"`
	landline      string `json:"landline"`      // Городской телефон
	mobilePhone   string `json:"mobilePhone"`   // Мобильный телефон
	domesticPhone string `json:"domesticPhone"` // Домашний телефон
}

type PhonebookRecordModel struct {
	DB *sql.DB
}

func (m PhonebookRecordModel) GetAll() ([]*phonebookRecord, error) {
	query := `SELECT MBAnalit.Dop3 AS tabno, MBAnalit.Dop AS fio, MBAnalitSprPodr.NameAn AS podrazd, MBAnalit.Stroka AS dolzhn, MBAnalit.Dop4 AS vnuttel,
		MBAnalit.Email AS email, MBAnalit.GorTel AS gortel, MBAnalit.SotTel AS sotov, MBAnalit.ICQ AS icq, MBAnalit.DomTel AS hometel
		FROM MBAnalit MBAnalit
		LEFT OUTER JOIN MBAnalitSpr MBAnalitSprPodr ON MBAnalit.Podr = MBAnalitSprPodr.Analit 
		WHERE (MBAnalit.Vid = 288) AND (ISNULL(MBAnalit.OurFirm, 38838) = 38838) AND (MBAnalit.Dop4 IS NOT NULL  OR MBAnalit.Email IS NOT NULL OR MBAnalit.SotTel IS NOT NULL
		OR MBAnalit.DomTel IS NOT NULL OR MBAnalit.ICQ IS NOT NULL) AND (MBAnalit.Sost <> 'З') AND (MBAnalit.Dop3 <> '2644') AND (MBAnalit.Dop3 <> '1761')
		GROUP BY podrazd
		ORDER BY podrazd, fio`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	records := []*phonebookRecord{}

	for rows.Next() {
		var record phonebookRecord

		err := rows.Scan(
			&record.employeeId,
			&record.name,
			&record.department,
			&record.employ,
			&record.internalPhone,
			&record.email,
			&record.landline,
			&record.mobilePhone,
			&record.domesticPhone,
		)
		if err != nil {
			return nil, err
		}

		records = append(records, &record)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return records, nil
}
