package structPack

import (
	"time"
)

//根据需要的字段来填充数据库
type Anote struct {
	Id             int       `json:"id" xorm:"pk autoincr" `
	Cover          string    `json:"cover"`
	Title          string    `json:"title"`
	Content        string    `json:"content"`
	Commentnum     int       `json:"commentnum"`
	Uid            int       `json:"uid"`
	Newslist_tpl   int       `json:"newslist_tpl"`
	Ctime          time.Time `json:"ctime"  xorm:"created"`
	Is_open        int       `json:"is_open"`
	Password_code  string    `json:"password_code"`
	Url            string    `json:"url"`
	Is_auto        int       `json:"is_auto"`
	Cateid         int       `json:"cateid"`
	Is_auto_page   int       `json:"is_auto_page"`
	Is_auto_source string    `json:"is_auto_source"`
}
