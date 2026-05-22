/**
 * Author : ngdangkietswe
 * Since  : 11/1/2025
 */

package util

import "strings"

func IsEmptyString(val *string) bool {
	return val == nil || len(strings.TrimSpace(*val)) == 0
}

func IsNotEmptyString(val *string) bool {
	return !IsEmptyString(val)
}
