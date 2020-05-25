package models

import "github.com/dgrijalva/jwt-go"

type User struct {
	UserID       int64
	Username     string `json:username`
	Password     string `json:password`
	Nama         string `json:nama`
	Alamat       string `json:alamat`
	NoHp         string `json:noHp`
	TanggalLahir string `json:tanggalLahir`
	Role         string `jsom:role`
}

type AuthResponse struct {
	Token    string `json:"token"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

type Token struct {
	UserId int64
	jwt.StandardClaims
}
