package adapter

type (
	Employee struct {
		ID         int    `json:"id" db:"id"`
		FullName   string `json:"fullname" db:"fullname"`
		LeaveQuota int    `json:"leaveQuota" db:"leaveQuota"`
		OnLeave    int    `json:"onLeave" db:"onLeave"`
	}
)
