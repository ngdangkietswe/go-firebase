/**
 * Author : ngdangkietswe
 * Since  : 10/21/2025
 */

package util

import (
	"go-firebase/internal/data/ent"
	"go-firebase/pkg/request"
	"go-firebase/pkg/response"

	"entgo.io/ent/dialect/sql"
)

const (
	OrderAsc  = "asc"
	OrderDesc = "desc"
)

const (
	DefaultPage     = 0
	DefaultPageSize = 10
	MaxPageSize     = 100
	DefaultSort     = "updated_at"
	DefaultOrder    = OrderDesc
)

func NormalizePaginationRequest(pagination *request.PaginateRequest) {
	if pagination.Page <= 0 {
		pagination.Page = DefaultPage
	} else {
		pagination.Page = pagination.Page - 1
	}

	if pagination.PageSize <= 0 {
		pagination.PageSize = DefaultPageSize
	} else if pagination.PageSize > MaxPageSize {
		pagination.PageSize = MaxPageSize
	}

	if pagination.Sort == "" {
		pagination.Sort = DefaultSort
	}

	if pagination.Order == "" || (pagination.Order != OrderAsc && pagination.Order != OrderDesc) {
		pagination.Order = DefaultOrder
	}
}

func AsPageMeta(pagination *request.PaginateRequest, totalItems int) *response.PageMeta {
	pageMeta := &response.PageMeta{
		TotalItems: totalItems,
	}

	totalPages := totalItems / pagination.PageSize
	if totalItems%pagination.PageSize != 0 {
		totalPages++
	}

	pageMeta.TotalPages = totalPages
	pageMeta.CurrentPage = pagination.Page + 1
	pageMeta.PageSize = pagination.PageSize
	pageMeta.HasPrevious = pagination.Page > 0
	pageMeta.HasNext = pagination.Page+1 < totalPages

	return pageMeta
}

func ToSortOrder(request *request.PaginateRequest) func(*sql.Selector) {
	if request.Order == OrderAsc {
		return ent.Asc(request.Sort)
	}
	return ent.Desc(request.Sort)
}
