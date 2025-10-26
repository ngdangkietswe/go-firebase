/**
 * Author : ngdangkietswe
 * Since  : 10/25/2025
 */

package model

type Principal struct {
	FirebaseUID string        `json:"firebase_uid"`
	SystemUID   string        `json:"system_uid"`
	Roles       []*Role       `json:"roles"`
	Permissions []*Permission `json:"permissions"`
}

type Role struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Permission struct {
	Action   string `json:"action"`
	Resource string `json:"resource"`
}

func (principal *Principal) HasPermission(perm *Permission) bool {
	for _, p := range principal.Permissions {
		if p.Action == perm.Action && p.Resource == perm.Resource {
			return true
		}
	}
	return false
}
