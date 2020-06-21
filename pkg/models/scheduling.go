package models

type Scheduling struct {
	ID            int64
	Enabled       int
	Minutes       string
	Hours         string
	DayOfMonth    string
	Months        string
	DayOfWeek     string
	TypeGuideline string
	TambakID      int64
	Description   string
}
