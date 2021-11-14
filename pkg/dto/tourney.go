package dto

type GroupSlot struct {
	UserID string `json:"user_id"`
	Team   Team   `json:"team"`
}

type Group struct {
	Name      string      `json:"name"`
	TeamSlots []GroupSlot `json:"team_slots"`
}

type Tourney struct {
	ID            string  `json:"id"`
	GroupsCount   int     `json:"groups_count"`
	TeamsPerGroup int     `json:"teams_per_group"`
	Groups        []Group `json:"groups"`
}
