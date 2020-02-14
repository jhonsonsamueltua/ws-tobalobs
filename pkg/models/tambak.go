package models

type Tambak struct {
	ID          int64
	Name        string
	Location    string
	Description string
}

type MonitorTambak struct {
	MonitorTambakId int64
	TambakId        int64
	PH              float64
	DO              float64
	Suhu            float64
	WaktuTanggal    string
	Keterangan      string
}
