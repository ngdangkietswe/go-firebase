/**
 * Author : ngdangkietswe
 * Since  : 10/18/2025
 */

package constant

import "time"

const CtxTimeOut = time.Duration(10000) * time.Second

const CtxFirebaseUIDKey = "ctx_firebase_uid"
const CtxSysUIDKey = "ctx_system_uid"

const AuthHeaderPrefixBearer = "Bearer "
const DefaultPassword = "123qwe!@#"

const PreloadDeviceTokens = "device_tokens"

const NotificationTopicTech = "tech"
const NotificationTopicAlgo = "algo"
