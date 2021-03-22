package model

import (
	"database/sql/driver"
	"fmt"
	"time"
)

// 别名方式 (不推荐)
//使用别名方式后时类型转换 time.Time(t)，而且 Scan 方法中不能直接通过类型断言 v.(TimeNormal) 将接口转换到 TimeNormal。
//另外，设置别名后，TimeNormal 并不能直接使用原始类型 time.Time 的各种方法和成员，需要先进行类型转换。
//显然，通过结构体匿名嵌入的方式并不存在这样的不便，这种方式可以很好的保持对象的原有性质。
//type TimeNormal  time.Time

// 内嵌方式（推荐）
type TimeNormal struct {
	time.Time
}

func (t TimeNormal) MarshalJSON() ([]byte, error) {
	// tune := fmt.Sprintf(`"%s"`, t.Format("2006-01-02 15:04:05"))
	tune := t.Format(`"2006-01-02 15:04:05"`)
	return []byte(tune), nil
}

// Value insert timestamp into mysql need this function.
func (t TimeNormal) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// Scan valueof time.Time
func (t *TimeNormal) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = TimeNormal{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

type BaseModel struct {
	// gorm.Model
	ID        uint        `gorm:"primary_key" json:"id"`
	CreatedAt TimeNormal  `json:"createdAt"`
	UpdatedAt TimeNormal  `json:"updatedAt"`
	DeletedAt *TimeNormal `sql:"index" json:"-"`
}
