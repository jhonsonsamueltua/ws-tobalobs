package mysql

const (
	queryInsertUser       = `INSERT INTO user (username, password, nama, alamat, no_hp, tanggal_lahir, role) VALUES (?, ?, ?, ?, ?, ?, ?)`
	querySelectUser       = `SELECT user_id, username, password, nama, alamat, no_hp, tanggal_lahir, role FROM user WHERE username=?`
	querySelectUserByHP   = `SELECT user_id, username, password, nama, alamat, no_hp, tanggal_lahir, role FROM user WHERE no_hp=?`
	querySelectDetailUser = `SELECT user_id, username, password, nama, alamat, no_hp, tanggal_lahir FROM user WHERE user_id=?`
	QueryUpdateUser       = `
		UPDATE user
		SET 
			username = ?,
			nama = ?,
			alamat = ?,
			no_hp = ?,
			tanggal_lahir = ?
		WHERE user_id = ?
	`
	QueryUpdatePassword = `
		UPDATE user
		SET 
			password = ?
		WHERE user_id = ?
	`

	queryGetAllKondisi = `SELECT penyimpangan_kondisi_tambak_id, aksi_penyimpangan, kondisi, tipe, nilai FROM penyimpangan_kondisi_tambak`
	QueryUpdateKondisi = `
		UPDATE penyimpangan_kondisi_tambak
		SET 
			aksi_penyimpangan = ?,
			kondisi = ?,
			tipe = ?,
			nilai = ?
		WHERE penyimpangan_kondisi_tambak_id = ?
	`
	queryGetDeviceID    = `SELECT device_id FROM device_user WHERE user_id = ?`
	querySaveDeviceID   = `INSERT INTO device_user (user_id, device_id) VALUES (?, ?)`
	queryDeleteDeviceID = `DELETE FROM device_user WHERE user_id = ? AND device_id = ?`
)
