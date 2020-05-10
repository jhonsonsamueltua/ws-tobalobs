package mysql

const (
	// queryGetAllNotif = `
	// SELECT notif.notifikasi_penyimpangan_kondisi_tambak_id, monitor.keterangan, nyimpang.kondisi
	// FROM notifikasi_penyimpangan_kondisi_tambak as notif
	// LEFT JOIN monitor_tambak as monitor
	// ON notif.monitor_tambak_id = monitor.monitor_tambak_id
	// LEFT JOIN tambak as tambak
	// ON tambak.tambak_id = monitor.tambak_id
	// LEFT JOIN penyimpangan_kondisi_tambak as nyimpang
	// ON notif.penyimpangan_kondisi_tambak_id = nyimpang.penyimpangan_kondisi_tambak_id
	// WHERE tambak.tambak_id = ?
	// ORDER BY notif.notifikasi_penyimpangan_kondisi_tambak_id ASC
	// `
	queryGetAllNotif = `
		SELECT n.notifikasi_id, t.nama_tambak, n.keterangan, n.status_notifikasi, n.tipe_notifikasi
		FROM notifikasi as n
		LEFT JOIN tambak as t
		ON t.tambak_id = n.tambak_id
		WHERE t.user_id = ? AND status_notifikasi != "pending"
		ORDER BY n.notifikasi_id DESC
	`

	queryGetAllNotifPerTambak = `
		SELECT notifikasi_id, tipe_notifikasi, keterangan, status_notifikasi, tipe_notifikasi
		FROM notifikasi
		WHERE tambak_id = ? AND status_notifikasi != "pending"
		ORDER BY notifikasi_id DESC
	`

	queryGetAllNotifUnreadPerTambak = `
		SELECT notifikasi_id, tipe_notifikasi, keterangan, status_notifikasi, tipe_notifikasi
		FROM notifikasi
		WHERE tambak_id = ? AND status_notifikasi = "unread"
		ORDER BY notifikasi_id DESC
	`

	queryGetDetailNotif = `
		SELECT n.notifikasi_id, t.nama_tambak, n.tipe_notifikasi, n.keterangan, n.waktu_tanggal, n.status_notifikasi, IFNULL(pkt.aksi_penyimpangan, "") as aksi_penyimpangan, IFNULL(pkt.kondisi, "") as kondisi_peyimpangan, IFNULL(g.aksi_guideline, "") as aksi_guideline, IFNULL(g.notifikasi, "") as kondisi_guideline
		FROM notifikasi as n
		LEFT JOIN tambak as t
		ON n.tambak_id = t.tambak_id
		LEFT JOIN penyimpangan_kondisi_tambak as pkt
		ON n.penyimpangan_kondisi_tambak_id = pkt.penyimpangan_kondisi_tambak_id
		LEFT JOIN guideline as g
		ON n.guideline_id = g.guideline_id
		WHERE n.notifikasi_id = ?
	`

	queryUpdateStatusNotifikasi = `
		UPDATE notifikasi SET status_notifikasi = "read" WHERE notifikasi_id = ?
	`

	queryGetTotalNotifikasiUnread = `
		SELECT count(*) as totalNotif
		FROM notifikasi as n
		LEFT JOIN tambak as t
		on n.tambak_id = t.tambak_id
		WHERE t.user_id = ? AND n.status_notifikasi = "unread"
	`

	querySaveNotifGuideline = `INSERT INTO notifikasi (tambak_id, guideline_id, tipe_notifikasi, keterangan, status_notifikasi, waktu_tanggal) VALUES (?, ?, ?, ?, ?, ?)`
)
