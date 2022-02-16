package entity

type Horse struct {
	ID            int    `json:"id" db:"id"`
	Name          string `json:"name" db:"name"`
	Gender        string `json:"gender" db:"gender"`
	Age           int    `json:"age" db:"age"`
	Weight        int    `json:"weight" db:"weight"`
	Breed         string `json:"breed" db:"breed"`
	Birthplace    string `json:"birthplace" db:"birthplace"`
	Habit         string `json:"habit" db:"habit"`
	Rentalprice   int    `json:"rentalprice" db:"rentalprice"`
	Purchaseprice int    `json:"purchaseprice" db:"purchaseprice"`
}
