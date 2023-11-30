package adapter

import (
	"go-api-echo/internal/services/employee/entity"
	"sort"
)

type SortDate []string

func (a SortDate) Len() int           { return len(a) }
func (a SortDate) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortDate) Less(i, j int) bool { return a[i] < a[j] }

func TransformLeaveSubmissionRequest(rawReq LeaveSubmissionReq) *[]entity.LeaveSubmission {
	req := make([]entity.LeaveSubmission, 0)
	sort.Sort(SortDate(rawReq.Dates))

	for _, date := range rawReq.Dates {
		req = append(req, entity.LeaveSubmission{ID: rawReq.ID, Date: date})
	}

	return &req
}
