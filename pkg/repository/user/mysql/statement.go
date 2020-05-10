package mysql

const (
	queryInsertUser       = `INSERT INTO user (username, password, nama, alamat, no_hp, tanggal_lahir) VALUES (?, ?, ?, ?, ?, ?)`
	querySelectUser       = `SELECT user_id, username, password, nama, alamat, no_hp, tanggal_lahir FROM user WHERE username=?`
	querySelectDetailUser = `SELECT user_id, username, password, nama, alamat, no_hp, tanggal_lahir FROM user WHERE user_id=?`
)
