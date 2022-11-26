package xlsx

import (
	"github.com/Luna-CY/go-office/common"
	"time"
)

func New() *Document {
	var d = new(Document)

	d.workbook = NewWorkbook()
	d.property.application.name = common.Application
	d.property.application.version = common.Version

	return d
}

type Document struct {
	workbook *Workbook // 工作簿核心对象，该属性是私有只读的，可以通过GetWorkbook()方法获取

	property struct {
		core struct {
			creator        string    // 创建用户名
			lastModifyUser string    // 最近一次修改的用户名
			created        time.Time // 创建时间
			updated        time.Time // 编辑时间
		}

		application struct {
			name    string // 应用名称
			version string // 应用版本号
		}
	}
}

// GetWorkbook 取到工作簿对象
func (s *Document) GetWorkbook() *Workbook {
	return s.workbook
}

// NewSheet 代理:新建数据表
func (s *Document) NewSheet() *Sheet {
	return s.workbook.NewSheet()
}

// SetCreator 设置创建文档的用户名
func (s *Document) SetCreator(user string) {
	s.property.core.creator = user
}

// SetModifyUser 设置最后修改的用户名
func (s *Document) SetModifyUser(user string) {
	s.property.core.lastModifyUser = user
}

// SetCreateTime 设置创建时间
func (s *Document) SetCreateTime(time time.Time) {
	s.property.core.created = time
}

// SetUpdateTime 设置最后一次更新时间
func (s *Document) SetUpdateTime(time time.Time) {
	s.property.core.updated = time
}
