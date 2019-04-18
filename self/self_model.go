/**
*@Author: haoxiongxiao
*@Date: 2019/4/18
*@Description: CREATE GO FILE self
 */
package self

import "github.com/xhaoxiong/util"

type Model struct {
	ID        uint           `gorm:"primary_key" json:"id"`
	CreatedAt util.JSONTime  `json:"createdAt"`
	UpdatedAt util.JSONTime  `json:"updatedAt"`
	DeletedAt *util.JSONTime `sql:"index" json:"-"`
}
