/**
 * Author : ngdangkietswe
 * Since  : 5/22/2026
 */

package request

type StartTOTPEnrollmentRequest struct {
	IDToken string `json:"id_token" validate:"required"`
}

type VerifyMFATOTPRequest struct {
	IDToken     string `json:"id_token" validate:"required"`
	SessionInfo string `json:"session_info" validate:"required"`
	TOTPCode    string `json:"totp_code" validate:"required,len=6"`
	DisplayName string `json:"display_name" validate:"required,len=8"`
}

type UnenrollMFARequest struct {
	FirebaseUID  string `json:"firebase_uid"  validate:"required"`
	EnrollmentID string `json:"enrollment_id" validate:"required"`
}

type MFASignInStartRequest struct {
	MFAPendingCredential string `json:"mfa_pending_credential" validate:"required"`
	EnrollmentID         string `json:"enrollment_id"         validate:"required"`
}

type MFASignInFinalizeRequest struct {
	MFAPendingCredential string `json:"mfa_pending_credential" validate:"required"`
	SessionInfo          string `json:"session_info"          validate:"required"`
	TOTPCode             string `json:"totp_code"             validate:"required,len=6"`
}
