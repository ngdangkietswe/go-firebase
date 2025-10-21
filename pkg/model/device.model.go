/**
 * Author : ngdangkietswe
 * Since  : 10/18/2025
 */

package model

type Device struct {
	DeviceID string `json:"device_id"`
	UserID   string `json:"user_id"`
	Token    string `json:"token"`
	Platform string `json:"platform"`
}
