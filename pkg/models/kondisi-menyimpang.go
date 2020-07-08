package models

type KondisiMenyimpang struct {
	ID               int64  `json:"id"`
	AksiPenyimpangan string `json:"aksiPenyimpangan,omitempty"`
	Kondisi          string `json:"kondisi,omitempty"`
	Tipe             string `json:"tipe,omitempty"`
	Nilai            string `json:"nilai,omitempty"`
}
