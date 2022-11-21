package entity

// Skills 技能集合
type Skills struct {
	TotalItem int     `json:"totalItem"`
	Data      []Skill `json:"data"`
}

// Skill 技能原始数据
type Skill struct {
	ID                int    `json:"id"`
	Name              string `json:"name"`
	AttributesID      int    `json:"attributes_id"`
	Description       string `json:"description"`
	SkillTypeID       int    `json:"skill_type_id"`
	Value             int    `json:"value"`
	Amount            int    `json:"amount"`
	IsGenetic         int    `json:"is_genetic"`
	Speed             int    `json:"speed"`
	IsBe              int    `json:"is_be"`
	AdditionalEffects string `json:"additional_effects"`
}

// Spirit 精灵
type Spirit struct {
	Avatar                string   `json:"avatar"`
	Name                  string   `json:"name"`
	PrimaryAttributesID   *int     `json:"primary_attributes_id"`
	SecondaryAttributesID *int     `json:"secondary_attributes_id"`
	GroupID               *int     `json:"group_id"`
	Number                int      `json:"number"`
	Height                float64  `json:"height"`
	Weight                float64  `json:"weight"`
	Hobby                 string   `json:"hobby"`
	Description           string   `json:"description"`
	RacePower             int      `json:"race_power"`
	RaceAttack            int      `json:"race_attack"`
	RaceDefense           int      `json:"race_defense"`
	RaceMagicAttack       int      `json:"race_magic_attack"`
	RaceMagicDefense      int      `json:"race_magic_defense"`
	RaceSpeed             int      `json:"race_speed"`
	Skills                []string `json:"skills"`
}

type SpiritListItem struct {
	ID                  int64            `json:"id"`
	Number              string           `json:"number"`
	Avatar              string           `json:"avatar"`
	Name                string           `json:"name"`
	PrimaryAttributes   GeniusAttributes `json:"primaryAttributes"`
	SecondaryAttributes GeniusAttributes `json:"secondaryAttributes"`
}

type GeniusAttributes struct {
	Id   *int    `json:"id"`
	Name *string `json:"name"`
}

// NewsList 新闻集合
type NewsList struct {
	TotalItem int    `json:"totalItem"`
	Data      []News `json:"data"`
}

// News 新闻
type News struct {
	Id         int64   `json:"id" `
	Title      string  `json:"title" `
	CreateTime string  `json:"createTime"`
	UpdateTime *string `json:"updateTime"`
	Url        string  `json:"url" `
	Type       int     `json:"type"`
	Content    string  `json:"content"`
}

type SpiritGroup struct {
	Id   *int    `json:"id"`
	Name *string `json:"name"`
}
