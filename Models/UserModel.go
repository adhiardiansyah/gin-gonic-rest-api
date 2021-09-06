package Models

type User struct {
	Id      uint   `json:"id"`
	Nama    string `json:"nama"`
	Email   string `json:"email"`
	Telepon string `json:"telepon"`
	Alamat  string `json:"alamat"`
}

func (b *User) TableName() string {
	return "user"
}
