package db

type Conf struct {
	Host string
	Port string
	Username string
	Password string
	Name string
}

var Default = Conf{
	Host :"localhost",
	Port : "3306",
	Username: "root",
	Password : "root",
	Name : "staff",
}