package data

import (
	"context"
	"database/sql"
	"time"
)

type Subscriber struct {
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

type SubscriberModel struct {
	DB *sql.DB
}

func (m SubscriberModel) GetAll(id, name string) ([]*Subscriber, error) {
	query := `
		SET @P1 = ?;
		SET @P2 = ?;
		SELECT MBAnalit.Dop3 AS tabno,
			MBAnalit.Dop AS fio,
			MBAnalitSprPodr.NameAn AS podrazd,
			MBAnalit.Stroka AS dolzhn,
			MBAnalit.Dop4 AS vnuttel,
			MBAnalit.Email AS email,
			MBAnalit.GorTel AS gortel,
			MBAnalit.SotTel AS sotov,
			MBAnalit.DomTel AS hometel
		FROM MBAnalit AS MBAnalit
			LEFT OUTER JOIN MBAnalitSpr AS MBAnalitSprPodr ON MBAnalit.Podr = MBAnalitSprPodr.Analit
		WHERE (MBAnalit.Vid = 288)
			AND (ISNULL(MBAnalit.OurFirm, 38838) = 38838)
			AND (MBAnalit.Dop4 IS NOT NULL)
			AND (MBAnalit.Sost <> 'З')
			AND (MBAnalit.Dop3 LIKE '%' + @P1 + '%' OR @P1 = '')
			AND (MBAnalit.Dop LIKE '%' + @P2 + '%' OR @P2 = '')
		ORDER BY fio`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query, id, name)
	if err != nil {
		return nil, err
	}

	subscribers := []*Subscriber{}

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

		var s Subscriber
		s.EmployeeId = input.employeeId
		s.Name, _ = Win1251ToUTF8(input.name)
		s.Department, _ = Win1251ToUTF8(input.department.String)
		s.Employ, _ = Win1251ToUTF8(input.employ.String)
		s.InternalPhone = input.internalPhone.String
		s.Email = input.email.String
		s.Landline = input.landline.String
		s.MobilePhone = input.mobilePhone.String
		s.DomesticPhone = input.domesticPhone.String

		subscribers = append(subscribers, &s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return subscribers, nil
}
