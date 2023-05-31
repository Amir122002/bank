package models

type User struct {
	ID       int    `json:"id" gorm:"primaryKey;"`
	Login    string `json:"login"`
	Password string `json:"-"`
	//CreatedAt time.Time `json:"created_at"`
	//UpdatedAt time.Time `json:"updated_at"`

	Cell []Cell `json:"cells" gorm:"foreignKey:UserID"`
}

type Cell struct {
	Id    int `json:"id" gorm:"primaryKey;"`
	Money int `json:"money"`
	//CreatedAt time.Time `json:"created_at"`
	//UpdatedAt time.Time `json:"updated_at"`
	UserID uint `json:"userId" gorm:"index"`
}
