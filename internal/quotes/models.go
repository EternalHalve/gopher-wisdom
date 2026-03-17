package quotes

type Quote struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	Content   string `json:"content"`
	Anime     string `json:"anime"`
	Character string `json:"character"`
}
