package response

import helpers_pagination "go-api-echo/internal/pkg/helpers/pagination"

type HttpRes struct {
	Status                         int         `json:"status"`
	Message                        string      `json:"message"`
	Data                           interface{} `json:"data"`
	Errors                         []string    `json:"errors,omitempty"`
	*helpers_pagination.Pagination `json:"meta,omitempty"`
}

type HttpResPaginated struct {
	Status                         int         `json:"status"`
	Message                        string      `json:"message"`
	Data                           interface{} `json:"data,omitempty"`
	Errors                         []string    `json:"errors,omitempty"`
	*helpers_pagination.Pagination `json:"meta,omitempty"`
}
