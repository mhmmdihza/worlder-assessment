package storage

import (
	"bytes"
	"github.com/jmoiron/sqlx"
	"time"
)

type Sensor struct{
	ID1 string `db:"id1"`
	ID2 string `db:"id2"`
	SensorValue int `db:"sensor_value"`
	TimeStamp time.Time `db:"timestamp"`
}

type SensorReq struct{
	ID1 string
	ID2 string
	From int64
	To int64
}

func Insert(s *sqlx.DB,sensor Sensor) error {
	_, err := s.NamedExec(`INSERT INTO sensor (id1,id2,sensor_value) VALUES (:id1,:id2,:sensor_value)`,
		sensor)
	return err
}

func Retrieve(s *sqlx.DB, sensor SensorReq) (*[]Sensor, error) {
	sen := []Sensor{}
	stmt, param := searchCriteria(sensor, append([]interface{}{}))
	err := s.Select(&sen, "SELECT * FROM sensor where 1=1 " + stmt+ " ORDER BY timestamp ASC",param...)
	return &sen,err
}

func searchCriteria(req SensorReq, args []interface{}) (string, []interface{}) {
	var QueryBuilder bytes.Buffer
	if req.ID1 != "" {
		args = append(args, req.ID1)
		QueryBuilder.WriteString(" AND id1 = ?")
	}
	if req.ID2 != "" {
		args = append(args, req.ID2)
		QueryBuilder.WriteString(" AND id2 = ?")
	}
	if req.From != 0 {
		args = append(args, req.From)
		QueryBuilder.WriteString(" AND UNIX_TIMESTAMP(timestamp) >= ?")
	}
	if req.To != 0 {
		args = append(args, req.To)
		QueryBuilder.WriteString(" AND UNIX_TIMESTAMP(timestamp) < ?")
	}

	return QueryBuilder.String(), args
}


