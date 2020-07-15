package models

type Tambak struct {
	TambakID             int64   `json:"tambakID,omitempty"`
	UserID               int64   `json:"userID,omitempty"`
	NamaTambak           string  `json:"namaTambak,omitempty"`
	Panjang              float64 `json:"panjang,omitempty"`
	Lebar                float64 `json:"lebar,omitempty"`
	JenisBudidaya        string  `json:"jenisBudidaya,omitempty"`
	TanggalMulaiBudidaya string  `json:"tanggalMulaiBudidaya,omitempty"`
	UsiaLobster          int     `json:"usiaLobster,omitempty"`
	JumlahLobster        int     `json:"jumlahLobster,omitempty"`
	JumlahLobsterJantan  int     `json:"jumlahLobsterJantan,omitempty"`
	JumlahLobsterBetina  int     `json:"jumlahLobsterBetina,omitempty"`
	Status               string  `json:"status,omitempty"`
	PakanPagi            string  `json:"pakanPagi,omitempty"`
	PakanSore            string  `json:"pakanSore,omitempty"`
	GantiAir             string  `json:"gantiAir,omitempty"`
}

type MonitorTambak struct {
	MonitorTambakId int64   `json:"monitorTambakID,omitempty"`
	TambakId        int64   `json:"tambakID"`
	NamaTambak      string  `json:"namaTambak"`
	PH              float64 `json:"ph"`
	DO              float64 `json:"do"`
	Suhu            float64 `json:"suhu"`
	WaktuTanggal    string  `json:"waktuTanggal"`
	Keterangan      string  `json:"keterangan"`
}

type Notifikasi struct {
	NotifikasiID                int64  `json:"notifikasiID,omitempty"`
	TambakID                    int64  `json:"tambakID,omitempty"`
	UserID                      int64  `json:"userID,omitempty"`
	GuidelineID                 int64  `json:"guidelineID,omitempty"`
	PenyimpanganKondisiTambakID int64  `json:"penyimpanganKondisiTambakId,omitempty"`
	Keterangan                  string `json:"keterangan,omitempty"`
	StatusNotifikasi            string `json:"statusNotifikasi,omitempty"`
	WaktuTanggal                string `json:"waktuTanggal,omitempty"`
	TipeNotifikasi              string `json:"tipeNotifikasi,omitempty"`
	NamaTambak                  string `json:"namaTambak,omitempty"`
	AksiPenyimpangan            string `json:"aksiPenyimpangan,omitempty"`
	KondisiPenyimpangan         string `json:"kondisiPenyimpangan,omitempty"`
	AksiGuideline               string `json:"aksiGuideline,omitempty"`
	KondisiGuideline            string `json:"kondisiGuideline,omitempty"`
}

type MessagePushNotif struct {
	ID               string `json:"notifikasiID,omitempty"`
	Title            string `json:"title,omitempty"`
	Body             string `json:"body,omitempty"`
	StatusNotifikasi string `json:"statusNotifikasi,omitempty"`
	TipeNotifikasi   string `json:"tipeNotifikasi,omitempty"`
	WaktuTanggal     string `json:"waktuTanggal,omitempty"`
}

type NotifikasiKondisi struct {
	NotifikasiID     int64  `json:"notifikasiID"`
	NamaTambak       string `json:"namaTambak,omitempty"`
	Keterangan       string `json:"keterangan,omitempty"`
	Kondisi          string `json:"kondisi,omitempty"`
	AksiPenyimpangan string `json:"aksiPenyimpangan,omitempty"`
	WaktuTanggal     string `json:"waktuTanggal,omitempty"`
}

type Info struct {
	InfoID     int64  `json:"infoID"`
	Judul      string `json:"judul"`
	Penjelasan string `json:"penjelasan"`
}

type Panduan struct {
	PanduanAplikasiID int64  `json:"panduanAplikasiID"`
	Judul             string `json:"judul"`
	Penjelasan        string `json:"penjelasan"`
}

type Guideline struct {
	GuidelineID   int64  `json:"guidelineID"`
	AksiGuideline string `json:"aksiGuideline"`
	Notifikasi    string `json:"notifikasi"`
	TipeBudidaya  string `json:"tipeBudidaya"`
	TipeJadwal    string `json:"tipeJadwal"`
	Interval      string `json:"interval"`
	Waktu         string `json:"waktu"`
}
