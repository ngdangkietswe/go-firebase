/**
 * Author : ngdangkietswe
 * Since  : 10/18/2025
 */

package model

type User struct {
	UserID    string    `json:"user_id"`
	Email     string    `json:"email"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Devices   []*Device `json:"devices"`
}
