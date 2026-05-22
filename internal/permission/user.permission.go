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

func ReadUserPerm() *model.Permission {
	return &model.Permission{
		Action:   constant.ActionRead,
		Resource: ResourceUser,
	}
}

func CreateUserPerm() *model.Permission {
	return &model.Permission{
		Action:   constant.ActionCreate,
		Resource: ResourceUser,
	}
}

func UpdateUserPerm() *model.Permission {
	return &model.Permission{
		Action:   constant.ActionUpdate,
		Resource: ResourceUser,
	}
}

func DeleteUserPerm() *model.Permission {
	return &model.Permission{
		Action:   constant.ActionDelete,
		Resource: ResourceUser,
	}
}
