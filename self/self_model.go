/**
*@Author: haoxiongxiao
*@Date: 2019/4/18
*@Description: CREATE GO FILE self
 */
package self

import "github.com/xhaoxiong/util"

//用于gorm的基础model
type Model struct {
	ID        uint           `gorm:"primary_key" json:"id"`
	CreatedAt util.JSONTime  `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt util.JSONTime  `json:"updatedAt"`
	DeletedAt *util.JSONTime `sql:"index" json:"-"`
}
