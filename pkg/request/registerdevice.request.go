/**
 * Author : ngdangkietswe
 * Since  : 10/17/2025
 */

package request

type RegisterDeviceRequest struct {
	DeviceToken string `json:"device_token" binding:"required"`
	Identifier  string `json:"identifier" binding:"required"` // e.g., user ID or email
	Platform    string `json:"platform" binding:"required"`   // e.g., "ios", "android", "web"
}
