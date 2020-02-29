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
}

type MonitorTambak struct {
	MonitorTambakId int64   `json:"monitorTambakID,omitempty"`
	TambakId        int64   `json:"tambakID,omitempty"`
	NamaTambak      string  `json:"namaTambak,omitempty"`
	PH              float64 `json:"ph,omitempty"`
	DO              float64 `json:"do,omitempty"`
	Suhu            float64 `json:"suhu,omitempty"`
	WaktuTanggal    string  `json:"waktuTanggal,omitempty"`
	Keterangan      string  `json:"keterangan,omitempty"`
}
