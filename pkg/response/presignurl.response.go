/**
 * Author : ngdangkietswe
 * Since  : 10/26/2025
 */

package response

type PresignURLResponse struct {
	URL      string `json:"url"`
	ExpireAt string `json:"expire_at"`
}
