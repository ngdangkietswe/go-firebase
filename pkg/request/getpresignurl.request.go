/**
 * Author : ngdangkietswe
 * Since  : 10/26/2025
 */

package request

type GetPresignURLRequest struct {
	BucketName  string `json:"bucket_name" binding:"required"`
	ObjectName  string `json:"object_name" binding:"required"`
	ContentType string `json:"content_type,omitempty"`
	ExpireTime  int64  `json:"expire_time,omitempty"` // in seconds
}
