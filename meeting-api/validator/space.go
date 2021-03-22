package validator

//获取所有地点列表验证规则
type GetAllSpaces struct {
	SortBy string `form:"sortBy" validate:"required,oneof=ASC DESC" label:"排序顺序"` //eq=0|eq=1 , one of=0 1
	Order  string `form:"order" validate:"required,max=30" label:"排序字段"`
}

//获取地点列表验证规则
type GetSpaces struct {
	Page     int64  `form:"page" validate:"required,numeric,min=0" label:"分页查询页码"`
	PageSize int64  `form:"pageSize" validate:"required,numeric,min=1,max=100" label:"分页查询每页数量"`
	SortBy   string `form:"sortBy" validate:"required,oneof=ASC DESC" label:"排序顺序"` //eq=0|eq=1 , one of=0 1
	Order    string `form:"order" validate:"required,max=30" label:"排序字段"`
}

//创建地点验证规则
type CreateSpace struct {
	Name   string  `json:"name" validate:"required,min=2,max=30" label:"地点名称"`
	Lng    float64 `json:"lng" validate:"required,numeric,lte=130" label:"经度"`
	Lat    float64 `json:"lat" validate:"required,numeric,lte=130" label:"纬度"`
	Status string  `json:"status" validate:"oneof=0 1" label:"启用状态"` //eq=0|eq=1 , one of=0 1
}

//更新地点验证规则
type UpdateSpace struct {
	Id     int64   `json:"id" validate:"required" label:"地点ID"`
	Name   string  `json:"name" validate:"required,min=2,max=30" label:"地点名称"`
	Lng    float64 `json:"lng" validate:"required,numeric,lte=130" label:"经度"`
	Lat    float64 `json:"lat" validate:"required,numeric,lte=130" label:"纬度"`
	Status string  `json:"status" validate:"oneof=0 1" label:"启用状态"` //eq=0|eq=1 , one of=0 1
}

//启用/禁用地点验证规则
type UpdateSpaceStatus struct {
	Id     int64  `json:"id" validate:"required" label:"地点ID"`
	Status string `json:"status" validate:"oneof=0 1" label:"启用状态"` //eq=0|eq=1 , one of=0 1
}
