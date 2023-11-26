package adapter

import "go-api-echo/internal/pkg/helpers/pagination"

type (
	ListEmployeeReq struct {
		Search string `query:"search"`
		pagination.Pagination
	}

	UpsertEmployeeReq struct {
		ID         int    `param:"id" db:"id"`
		FullName   string `json:"fullname" form:"fullname" db:"fullname"`
		LeaveQuota int    `json:"leaveQuota" form:"leaveQuota" db:"leavequota"`
	}
)
