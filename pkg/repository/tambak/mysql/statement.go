package mysql

const (
	queryGetAllTambak                  = "SELECT tambak_id, nama_tambak, status FROM tambak WHERE user_id = ? ORDER BY tambak_id DESC"
	queryGetTambakByID                 = "SELECT tambak_id, nama_tambak, panjang, lebar, jenis_budidaya, tanggal_mulai_budidaya, usia_lobster, jumlah_lobster, jumlah_lobster_jantan, jumlah_lobster_betina, status FROM tambak WHERE user_id = ? && tambak_id = ?"
	queryGetLastMonitorTambak          = "SELECT t.tambak_id, t.nama_tambak, IFNULL(m.ph, 0) as ph, IFNULL(m.do, 0) as do, IFNULL(m.suhu,0) as suhu, IFNULL(m.waktu_tanggal,'') as waktu_tanggal, IFNULL(m.keterangan,'') as keterangan FROM tambak as t LEFT JOIN monitor_tambak as m ON t.tambak_id = m.tambak_id WHERE t.tambak_id = ? ORDER BY m.monitor_tambak_id DESC LIMIT 1"
	queryInsertMonitoringTambak        = `INSERT INTO monitor_tambak (tambak_id, ph, do, suhu,	waktu_tanggal, keterangan) VALUES (?, ?, ?, ?, ?, ?)`
	queryInsertNotifikasiKondisiTambak = `INSERT INTO notifikasi (tambak_id, penyimpangan_kondisi_tambak_id, tipe_notifikasi, keterangan, status_notifikasi, waktu_tanggal) VALUES (?, ?, ?, ?, ?, ?)`
	queryUpdateNotifikasiKondisiTambak = `UPDATE notifikasi SET status_notifikasi = "pending" WHERE notifikasi_id = ?`
	queryInsertTambak                  = `INSERT INTO tambak (user_id, nama_tambak, panjang, lebar, jenis_budidaya, tanggal_mulai_budidaya, usia_lobster, jumlah_lobster, jumlah_lobster_jantan, jumlah_lobster_betina, status) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	queryGetAllInfo                    = `SELECT info_id, judul, penjelasan FROM info_budidaya`
	queryGetAllPanduan                 = `SELECT panduan_aplikasi_id, judul, penjelasan FROM panduan_aplikasi`
	queryGetMonitorTambak              = `SELECT monitor_tambak_id, ph, do, suhu, DATE_FORMAT(waktu_tanggal,'%h:%i:%s %p') waktu_tanggal, keterangan FROM monitor_tambak WHERE tambak_id = ? AND DATE(waktu_tanggal) = ? ORDER BY waktu_tanggal ASC`
	QueryUpdateTambak                  = `
		UPDATE tambak
		SET 
			nama_tambak = ?,
			panjang = ?,
			lebar = ?,
			jenis_budidaya = ?,
			usia_lobster = ?,
			jumlah_lobster = ?,
			jumlah_lobster_jantan = ?,
			jumlah_lobster_betina = ?
		WHERE tambak_id = ?
	`
	queryGetAllTambakID    = `SELECT tambak_id, user_id, nama_tambak FROM tambak WHERE STATUS = 'Aktif'`
	queryGetUserIDByTambak = `SELECT user_id FROM tambak WHERE tambak_id = ?`
)
