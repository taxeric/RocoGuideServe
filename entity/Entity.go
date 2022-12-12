package entity

import "database/sql"

// Skills 技能集合
type Skills struct {
	TotalItem int     `json:"totalItem"`
	Data      []Skill `json:"data"`
}

type SkillType struct {
	Id   *int    `json:"id"`
	Name *string `json:"name"`
}

// Skill 技能原始数据
type Skill struct {
	Id                int64            `json:"id"`
	Name              string           `json:"name"`
	Description       string           `json:"description"`
	Value             int              `json:"value"`
	Amount            int              `json:"amount"`
	Speed             int              `json:"speed"`
	IsGenetic         bool             `json:"isGenetic"`
	AdditionalEffects string           `json:"additionalEffects"`
	IsBe              bool             `json:"isBe"`
	SkillType         SkillType        `json:"skillType"`
	Attributes        SpiritAttributes `json:"attributes"`
}

// Spirit 精灵
type Spirit struct {
	Id                  int64            `json:"id"`
	Avatar              string           `json:"avatar"`
	Name                string           `json:"name"`
	PrimaryAttributes   SpiritAttributes `json:"primaryAttributes"`
	SecondaryAttributes SpiritAttributes `json:"secondaryAttributes"`
	Group               SpiritGroup      `json:"groupId"`
	Series              Series           `json:"series"`
	Number              int              `json:"number"`
	Height              float32          `json:"height"`
	Weight              float32          `json:"weight"`
	Hobby               string           `json:"hobby"`
	Description         string           `json:"description"`
	RacePower           int              `json:"racePower"`
	RaceAttack          int              `json:"raceAttack"`
	RaceDefense         int              `json:"raceDefense"`
	RaceMagicAttack     int              `json:"raceMagicAttack"`
	RaceMagicDefense    int              `json:"raceMagicDefense"`
	RaceSpeed           int              `json:"raceSpeed"`
	Skills              []Skill          `json:"skills"`
}

type SpiritListItem struct {
	ID                  int64            `json:"id"`
	Number              string           `json:"number"`
	Avatar              string           `json:"avatar"`
	Name                string           `json:"name"`
	PrimaryAttributes   SpiritAttributes `json:"primaryAttributes"`
	SecondaryAttributes SpiritAttributes `json:"secondaryAttributes"`
}

type SpiritAttributes struct {
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
	Id   int     `json:"id"`
	Name *string `json:"name"`
}

type Series struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type SpiritSeries struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type SkillEnvironment struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	Introduce string `json:"introduce"`
	Effects   string `json:"effects"`
	Type      int    `json:"type"`
	Icon      string `json:"icon"`
}

type AbnormalState struct {
	Id        int64          `json:"id"`
	Name      string         `json:"name"`
	Introduce string         `json:"introduce"`
	Icon      sql.NullString `json:"icon"`
}

type AbnormalStateResponse struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	Introduce string `json:"introduce"`
	Icon      string `json:"icon"`
}
