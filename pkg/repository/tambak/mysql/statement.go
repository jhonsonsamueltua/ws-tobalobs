package mysql

const (
	queryGetAllTambak                  = "SELECT tambak_id, nama_tambak, status FROM tambak WHERE user_id = ? ORDER BY tambak_id DESC"
	queryGetTambakByID                 = "SELECT tambak_id, nama_tambak, panjang, lebar, jenis_budidaya, tanggal_mulai_budidaya, usia_lobster, jumlah_lobster, jumlah_lobster_jantan, jumlah_lobster_betina, status FROM tambak WHERE user_id = ? && tambak_id = ?"
	queryGetLastMonitorTambak          = "SELECT t.tambak_id, t.nama_tambak, IFNULL(m.ph, 0) as ph, IFNULL(m.do, 0) as do, IFNULL(m.suhu,0) as suhu, IFNULL(m.waktu_tanggal,'') as waktu_tanggal, IFNULL(m.keterangan,'') as keterangan FROM tambak as t LEFT JOIN monitor_tambak as m ON t.tambak_id = m.tambak_id WHERE t.tambak_id = ? ORDER BY m.monitor_tambak_id DESC LIMIT 1"
	queryInsertMonitoringTambak        = `INSERT INTO monitor_tambak (tambak_id, ph, do, suhu,	waktu_tanggal, keterangan) VALUES (?, ?, ?, ?, ?, ?)`
	queryInsertNotifikasiKondisiTambak = `INSERT INTO notifikasi (tambak_id, penyimpangan_kondisi_tambak_id, tipe_notifikasi, keterangan, status_notifikasi, waktu_tanggal) VALUES (?, ?, ?, ?, ?, ?)`
	queryUpdateNotifikasiKondisiTambak = `UPDATE notifikasi SET status_notifikasi = ? WHERE notifikasi_id = ?`
	queryInsertTambak                  = `INSERT INTO tambak (user_id, nama_tambak, panjang, lebar, jenis_budidaya, tanggal_mulai_budidaya, usia_lobster, jumlah_lobster, jumlah_lobster_jantan, jumlah_lobster_betina, status, pakan_pagi, pakan_sore, ganti_air) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
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

	queryGetAllInfo = `SELECT info_id, judul, penjelasan FROM info_budidaya`
	queryInsertInfo = `INSERT INTO info_budidaya (judul, penjelasan) VALUES (?, ?)`
	QueryUpdateInfo = `
		UPDATE info_budidaya
		SET 
			judul = ?,
			penjelasan = ?
		WHERE info_id = ?
	`
	queryDeleteInfo = `DELETE FROM info_budidaya WHERE info_id = ?`

	queryGetAllPanduan = `SELECT panduan_aplikasi_id, judul, penjelasan FROM panduan_aplikasi`
	queryInsertPanduan = `INSERT INTO panduan_aplikasi (judul, penjelasan) VALUES (?, ?)`
	QueryUpdatePanduan = `
		UPDATE panduan_aplikasi
		SET 
			judul = ?,
			penjelasan = ?
		WHERE panduan_aplikasi_id = ?
	`
	queryDeletePanduan = `DELETE FROM panduan_aplikasi WHERE panduan_aplikasi_id = ?`

	queryGetAllSchedule = `SELECT id, minutes, hours, day_of_month, months, day_of_week, type_guideline, tambak_id, description FROM scheduling WHERE enabled = 1`
)
