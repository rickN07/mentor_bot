package user

import (
	"database/sql"
	"encoding/json"
	"time"
)

type User struct {
	ID        int64           `db:"id"`
	Info      json.RawMessage `db:"info"`
	CreatedAt time.Time       `db:"created_at"`
	Mailing   json.RawMessage `db:"mailing"`
	SentLead  sql.NullTime    `db:"sent_lead"`
}
