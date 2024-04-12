package models

type VersionGroupDetail struct {
	LevelLearnedAt  int    `json:"level_learned_at"`
	MoveLearnMethod Common `json:"move_learn_method"`
	VersionGroup    Common `json:"version_group"`
}
