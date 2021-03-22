package validator

type GetRooms struct {
	SpaceId  int64  `form:"space_id" validate:"" label:"所属地点ID"`
	Page     int64  `form:"page" validate:"required,numeric,min=0" label:"分页查询页码"`
	PageSize int64  `form:"pageSize" validate:"required,numeric,min=1,max=100" label:"分页查询每页数量"`
	SortBy   string `form:"sortBy" validate:"required,oneof=ASC DESC" label:"排序顺序"`
	Order    string `form:"order" validate:"required,max=30" label:"排序字段"`
}

type CreateRoom struct {
	SpaceId     int64  `json:"space_id" validate:"required" label:"所属地点ID"`
	Name        string `json:"name" validate:"required,min=2,max=30" label:"会议室名称"`
	Status      string `json:"status" validate:"oneof=0 1" label:"启用状态"`
	ImageUrl    string `json:"image_url" validate:"required,lte=200" label:"会议室图片"`
	CapacityMin int64  `json:"capacity_min" validate:"required,gte=0" label:"建议最小使用人数"`
	CapacityMax int64  `json:"capacity_max" validate:"required,gte=0,gtefield=CapacityMin" label:"建议最大使用人数"`
	DeviceIds   string `json:"device_ids" label:"设备id"`
}

type UpdateRoom struct {
	Id          int64  `json:"id" validate:"required" label:"会议室ID"`
	SpaceId     int64  `json:"space_id" validate:"required" label:"所属地点ID"`
	Name        string `json:"name" validate:"required,min=2,max=30" label:"会议室名称"`
	Status      string `json:"status" validate:"oneof=0 1" label:"启用状态"`
	ImageUrl    string `json:"image_url" validate:"required,lte=200" label:"会议室图片"`
	CapacityMin int64  `json:"capacity_min" validate:"required,gte=0" label:"建议最小使用人数"`
	CapacityMax int64  `json:"capacity_max" validate:"required,gte=0,gtefield=CapacityMin" label:"建议最大使用人数"`
	DeviceIds   string `json:"device_ids" label:"设备id"`
}

type DeleteRoom struct {
	Id int64 `json:"id" validate:"required" label:"会议室ID"`
}
