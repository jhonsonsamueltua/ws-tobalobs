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
)
