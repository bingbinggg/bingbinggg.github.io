package entity

// 数据库表明自定义，默认为model的复数形式，比如这里默认为 users
func (Spot) TableName() string {
	return "sys_spot"
}

type Spot struct {
	Id      int     `json:"id"`
	Name    string  `json:"name"`
	Picture string  `json:"picture"`
	Score   float32 `json:"score"`
	Address string  `json:"address"`
	Mobile  string  `json:"mobile"`
	Text    string  `json:"text"`
}
