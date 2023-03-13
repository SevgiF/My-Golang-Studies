package models

type Config struct { //büyük harfle başlayan public olur
	Database database
}

type database struct { //küçük harfle başlayan private
	Host     string
	Port     int
	Database string
	User     string
	Password string
}
