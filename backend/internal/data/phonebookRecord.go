package data

import (
	"context"
	"database/sql"
	"time"
)

type phonebookRecord struct {
	EmployeeId    string `json:"employeeId"`    // Табельный номер
	Name          string `json:"name"`          // ФИО
	Department    string `json:"department"`    // Подразделение
	Employ        string `json:"employ"`        // Должность
	InternalPhone string `json:"internalPhone"` // Внутренний телефон
	Email         string `json:"email"`         // Электронная почта
	Landline      string `json:"landline"`      // Городской телефон
	MobilePhone   string `json:"mobilePhone"`   // Мобильный телефон
	DomesticPhone string `json:"domesticPhone"` // Домашний телефон
}

type PhonebookRecordModel struct {
	DB *sql.DB
}

func (m PhonebookRecordModel) GetAll() ([]*phonebookRecord, error) {
	query := `SELECT MBAnalit.Dop3 AS tabno, MBAnalit.Dop AS fio, MBAnalitSprPodr.NameAn AS podrazd, MBAnalit.Stroka AS dolzhn, MBAnalit.Dop4 AS vnuttel,
		MBAnalit.Email AS email, MBAnalit.GorTel AS gortel, MBAnalit.SotTel AS sotov, MBAnalit.DomTel AS hometel
		FROM MBAnalit MBAnalit
		LEFT OUTER JOIN MBAnalitSpr MBAnalitSprPodr ON MBAnalit.Podr = MBAnalitSprPodr.Analit 
		WHERE (MBAnalit.Vid = 288) AND (ISNULL(MBAnalit.OurFirm, 38838) = 38838) AND (MBAnalit.Dop4 IS NOT NULL  OR MBAnalit.Email IS NOT NULL OR MBAnalit.SotTel IS NOT NULL
		OR MBAnalit.DomTel IS NOT NULL OR MBAnalit.ICQ IS NOT NULL) AND (MBAnalit.Sost <> 'З') AND (MBAnalit.Dop3 <> '2644') AND (MBAnalit.Dop3 <> '1761')
		ORDER BY podrazd, fio`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	records := []*phonebookRecord{}

	for rows.Next() {
		var input struct {
			employeeId    string
			name          string
			department    sql.NullString
			employ        sql.NullString
			internalPhone sql.NullString
			email         sql.NullString
			landline      sql.NullString
			mobilePhone   sql.NullString
			domesticPhone sql.NullString
		}

		err := rows.Scan(
			&input.employeeId,
			&input.name,
			&input.department,
			&input.employ,
			&input.internalPhone,
			&input.email,
			&input.landline,
			&input.mobilePhone,
			&input.domesticPhone,
		)
		if err != nil {
			return nil, err
		}

		var record phonebookRecord
		record.EmployeeId = input.employeeId
		record.Name, _ = Win1251ToUTF8(input.name)
		record.Department, _ = Win1251ToUTF8(input.department.String)
		record.Employ, _ = Win1251ToUTF8(input.employ.String)
		record.InternalPhone = input.internalPhone.String
		record.Email = input.email.String
		record.Landline = input.landline.String
		record.MobilePhone = input.mobilePhone.String
		record.DomesticPhone = input.domesticPhone.String

		records = append(records, &record)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return records, nil
}
