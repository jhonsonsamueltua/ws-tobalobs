package mysql

const (
	queryGetAllTambak           = "SELECT tambak_id, nama_tambak, status FROM tambak WHERE user_id = ?"
	queryGetTambakByID          = "SELECT tambak_id, nama_tambak, panjang, lebar, jenis_budidaya, tanggal_mulai_budidaya, usia_lobster, jumlah_lobster, jumlah_lobster_jantan, jumlah_lobster_betina, status FROM tambak WHERE user_id = ? && tambak_id = ?"
	queryGetLastMonitorTambak   = "SELECT t.tambak_id, t.nama_tambak, m.ph, m.do, m.suhu, m.waktu_tanggal, m.keterangan FROM tambak as t JOIN monitor_tambak as m ON t.tambak_id = m.tambak_id WHERE t.tambak_id = ? ORDER BY m.monitor_tambak_id DESC LIMIT 1"
	queryInsertMonitoringTambak = `INSERT INTO monitor_tambak (tambak_id, ph, do, suhu,	waktu_tanggal, keterangan) VALUES (?, ?, ?, ?, ?, ?)`
	queryInsertTambak           = `INSERT INTO tambak (user_id, nama_tambak, panjang, lebar, jenis_budidaya, tanggal_mulai_budidaya, usia_lobster, jumlah_lobster, jumlah_lobster_jantan, jumlah_lobster_betina, status) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
)
