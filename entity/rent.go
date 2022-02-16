package entity

type Rent struct {
	HorseID int    `json:"horse_id" db:"horse_id" binding:"required"`
	Since   string `json:"since" db:"since" binding:"required"`
	Till    string `json:"till" db:"till" binding:"required"`
}
