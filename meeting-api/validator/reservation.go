package validator

type GetReservations struct {
	RoomId    int64  `form:"room_id" validate:"required" label:"会议室id"`
	StartDate string `form:"start_date" validate:"required" label:"开始日期"`
	EndDate   string `form:"end_date" validate:"required" label:"结束日期"`
	Page      int64  `form:"page" validate:"required,numeric,min=0" label:"分页查询页码"`
	PageSize  int64  `form:"pageSize" validate:"required,numeric,min=1,max=100" label:"分页查询每页数量"`
	SortBy    string `form:"sortBy" validate:"required,oneof=ASC DESC" label:"排序顺序"`
	Order     string `form:"order" validate:"required,max=30" label:"排序字段"`
}

type CreateReservation struct {
	RoomId         int64  `json:"room_id" validate:"required" label:"会议室id"`
	Title          string `json:"title" validate:"required,min=2,max=30" label:"会议主题"`
	Content        string `json:"content" validate:"required,max=500" label:"会议内容"`
	InitiatorId    int64  `json:"initiator_id" validate:"required" label:"发起人id"`
	Date           string `json:"date" validate:"required" label:"会议日期"`
	StartTime      string `json:"start_time" validate:"required" label:"会议开始时间"`
	EndTime        string `json:"end_time" validate:"required" label:"会议结束时间"`
	ParticipantIds string `json:"participant_ids" validate:"required" label:"会议参与人"`
	CId            int64  `json:"c_id" validate:"required" label:"公司id"`
}
