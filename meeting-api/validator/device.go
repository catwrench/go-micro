package validator

type GetDevices struct {
	Name     string `form:"name" validate:"" label:"设备名称"`
	Page     int64  `form:"page" validate:"required,numeric,min=0" label:"分页查询页码"`
	PageSize int64  `form:"pageSize" validate:"required,numeric,min=1,max=100" label:"分页查询每页数量"`
	SortBy   string `form:"sortBy" validate:"required,oneof=ASC DESC" label:"排序顺序"`
	Order    string `form:"order" validate:"required,max=30" label:"排序字段"`
}

type CreateDevice struct {
	Name     string `json:"name" validate:"required,min=2,max=30" label:"设备名称"`
	Sn       string `json:"sn" validate:"required,lte=30" label:"设备编号"`
	ImageUrl string `json:"image_url" validate:"required,lte=200" label:"设备图"`
}

type UpdateDevice struct {
	Id       int64  `json:"id" validate:"required" label:"设备ID"`
	Name     string `json:"name" validate:"required,min=2,max=30" label:"设备名称"`
	Sn       string `json:"sn" validate:"required,lte=30" label:"设备编号"`
	ImageUrl string `json:"image_url" validate:"required,lte=200" label:"设备图"`
}

type DelDevice struct {
	Id int64 `json:"id" validate:"required" label:"设备ID"`
}
