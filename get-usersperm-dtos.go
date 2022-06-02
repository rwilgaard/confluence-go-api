package goconfluence

type GetAllGroupsWithAnyPermissionType struct {
	Groups     []string `json:"groups"`
	MaxResults int64    `json:"maxResults"`
	StartAt    int64    `json:"startAt"`
	Total      int64    `json:"total"`
}
