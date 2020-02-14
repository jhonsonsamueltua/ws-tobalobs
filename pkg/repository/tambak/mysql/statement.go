package mysql

const (
	queryGetAllTambak = "SELECT id, name, location, description FROM tambak"

	queryInsertMonitoringTambak = `INSERT INTO monitor_tambak (tambak_id, ph, do, suhu,	waktu_tanggal, keterangan) VALUES (?, ?, ?, ?, ?, ?)`
)
