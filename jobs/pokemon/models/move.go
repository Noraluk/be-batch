package models

type Move struct {
	Move                Common               `json:"move"`
	VersionGroupDetails []VersionGroupDetail `json:"version_group_details"`
}
