/**
 * Author : ngdangkietswe
 * Since  : 10/26/2025
 */

package permission

import (
	"go-firebase/pkg/constant"
	"go-firebase/pkg/model"
)

const ResourceUser = "user"

func CreateUserPerm() *model.Permission {
	return &model.Permission{
		Action:   constant.ActionCreate,
		Resource: ResourceUser,
	}
}

func ReadUserPerm() *model.Permission {
	return &model.Permission{
		Action:   constant.ActionRead,
		Resource: ResourceUser,
	}
}
