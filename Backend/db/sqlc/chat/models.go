// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"time"
)

type Record struct {
	TargetID  int32              `json:"target_id"`
	RoleType  string             `json:"role_type"`
	MediaType string             `json:"media_type"`
	Message   string        `json:"message"`
	Media     string        `json:"media"`
	CreatedAt time.Time `json:"created_at"`
	Isread    bool             `json:"isread"`
}
