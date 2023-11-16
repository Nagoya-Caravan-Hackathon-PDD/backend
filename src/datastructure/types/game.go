package types

import "time"

type CreateGame struct {
	GameID            string       `json:"game_id"`
	OwnerID           string       `json:"owner_id"`
	CreatedAt         time.Time    `json:"created_at"`
	OwnerGitmonStatus GitmonStatus `json:"gitmon_status"`
}

type JoinGame struct {
	GameID             string       `json:"game_id"`
	UserID             string       `json:"owner_id"`
	UserIDGitmonStatus GitmonStatus `json:"gitmon_status"`
}

type GitmonStatus struct {
	GitmonID string  `json:"gitmon_id"`
	Name     string  `json:"name"`
	HP       int     `json:"hp"`
	Attack   int     `json:"attack"`
	Defence  int     `json:"defence"`
	Speed    int     `json:"speed"`
	Buf      Buf     `json:"buf"`
	Skills   []Skill `json:"skills"`
}

type Buf struct {
	// バフについて通常時は -1, バフかけたらターン数が加算される
	// 毎ターン開始に -1する。　毎ターン終了後にチェックを行い0になっていたら-1 にする。
	BufAttack    int `json:"buf_attack"`
	BufDefence   int `json:"buf_defence"`
	DebufAttack  int `json:"debuf_attack"`
	DebufDefence int `json:"debuf_defence"`
}

type SkillType string

const (
	TypeAttack  SkillType = "attack"
	TypeDefence SkillType = "defence"
	TypeBuf     SkillType = "buf"
	TypeDebuf   SkillType = "debuf"
	TypeHeal    SkillType = "heal"
)

type Skill struct {
	// BP制にする。毎ターン+2もらえる
	SkillID     int       `json:"skill_id"`
	RequiredBP  int       `json:"required_bp"`
	SkillName   string    `json:"skill_name"`
	Description string    `json:"description"`
	SkillType   SkillType `json:"skill_type"`
	SkillValue  float32   `json:"skill_value"`
}
