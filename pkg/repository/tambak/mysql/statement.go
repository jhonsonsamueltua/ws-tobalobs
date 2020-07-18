package mysql

const (
	queryGetAllTambak  = "SELECT tambak_id, nama_tambak, status, pakan_pagi, pakan_sore, ganti_air FROM tambak WHERE user_id = ? ORDER BY tambak_id DESC"
	queryGetTambakByID = "SELECT tambak_id, nama_tambak, panjang, lebar, jenis_budidaya, tanggal_mulai_budidaya, usia_lobster, jumlah_lobster, jumlah_lobster_jantan, jumlah_lobster_betina, status FROM tambak WHERE user_id = ? && tambak_id = ?"
	// queryGetLastMonitorTambak          = "SELECT t.tambak_id, t.nama_tambak, IFNULL(m.ph, 0) as ph, IFNULL(m.do, 0) as do, IFNULL(m.suhu,0) as suhu, IFNULL(m.waktu_tanggal,'') as waktu_tanggal, IFNULL(m.keterangan,'') as keterangan FROM tambak as t LEFT JOIN monitor_tambak as m ON t.tambak_id = m.tambak_id WHERE t.tambak_id = ? ORDER BY m.monitor_tambak_id DESC LIMIT 1"
	queryGetLastMonitorTambak          = "SELECT t.tambak_id, t.nama_tambak, m.ph, m.do, m.suhu, IFNULL(m.waktu_tanggal,'') as waktu_tanggal, IFNULL(m.keterangan,'') as keterangan FROM tambak as t LEFT JOIN monitor_tambak as m ON t.tambak_id = m.tambak_id WHERE t.tambak_id = ? ORDER BY m.monitor_tambak_id DESC LIMIT 1"
	queryInsertMonitoringTambak        = `INSERT INTO monitor_tambak (tambak_id, ph, do, suhu,	waktu_tanggal, keterangan) VALUES (?, ?, ?, ?, ?, ?)`
	queryInsertNotifikasiKondisiTambak = `INSERT INTO notifikasi (tambak_id, penyimpangan_kondisi_tambak_id, tipe_notifikasi, keterangan, status_notifikasi, waktu_tanggal) VALUES (?, ?, ?, ?, ?, ?)`
	queryUpdateNotifikasiKondisiTambak = `UPDATE notifikasi SET status_notifikasi = ? WHERE notifikasi_id = ?`
	queryInsertTambak                  = `INSERT INTO tambak (user_id, nama_tambak, panjang, lebar, jenis_budidaya, tanggal_mulai_budidaya, usia_lobster, jumlah_lobster, jumlah_lobster_jantan, jumlah_lobster_betina, status, pakan_pagi, pakan_sore, ganti_air) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	queryGetMonitorTambak              = `SELECT monitor_tambak_id, ph, do, suhu, DATE_FORMAT(waktu_tanggal,'%H:%i') waktu_tanggal, keterangan FROM monitor_tambak WHERE tambak_id = ? AND DATE(waktu_tanggal) = ? ORDER BY waktu_tanggal DESC`
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

	queryUpdateJadwalPakanPagi = `
		UPDATE tambak SET pakan_pagi = ? WHERE tambak_id = ?
	`
	queryUpdateJadwalPakanSore = `
		UPDATE tambak SET pakan_sore = ? WHERE tambak_id = ?
	`
	queryUpdateJadwalGantiAir = `
		UPDATE tambak SET ganti_air = ? WHERE tambak_id = ?
	`
	queryGetAllGuideline = `SELECT guideline_id, aksi_guideline, notifikasi, tipe_budidaya, tipe_jadwal, interval_value, waktu FROM guideline`
	queryAddGuideline    = `INSERT INTO guideline (aksi_guideline, notifikasi, tipe_budidaya, tipe_jadwal, interval_value, waktu) VALUES (?, ?, ?, ?, ?, ?)`
	queryUpdateGuideline = `
		UPDATE guideline
		SET 
			aksi_guideline = ?,
			notifikasi = ?,
			tipe_budidaya = ?,
			tipe_jadwal = ?,
			interval_value = ?,
			waktu = ?
		WHERE guideline_id = ?
	`
	queryGetTunnel    = "SELECT id, ip, port FROM tunnel WHERE id = 1"
	queryUpdateTunnel = `
		UPDATE tunnel
		SET 
			ip = ?,
			port = ?
		WHERE id = 1
	`
)
