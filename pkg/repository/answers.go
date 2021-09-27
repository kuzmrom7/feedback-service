package repository

type Answer struct {
	ID        uint   `json:"id" gorm:"primarykey; "`
	Answer    string `json:"answer"`
	CreatedAt string `json:"created_at"`
	SourceId  string `json:"source_id"`
	StatusId  string `json:"status_id"`
	ReviewId  uint   `json:"-"`
}
