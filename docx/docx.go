package docx

import (
	"errors"
	"fmt"
	"os/user"
	"time"
)

// DocumentDefaultWidth 文档默认宽度
const DocumentDefaultWidth = 11906

// DocumentDefaultHeight 文档默认高度
const DocumentDefaultHeight = 16838

// New 新建一篇文档
func New() (*Document, error) {
	doc := new(Document)

	u, err := user.Current()
	if nil != err {
		return nil, errors.New(fmt.Sprintf("获取系统用户信息失败: %v", err))
	}

	// 初始化core.xml数据
	doc.core.CreateUser = u.Username
	doc.core.CreateTime = time.Now()
	doc.core.LastModifyUser = u.Username
	doc.core.LastModifyTime = time.Now()
	doc.core.Version = 0

	// 初始化app.xml
	doc.app.Template = "Normal.dotm"
	doc.app.Application = "Go Office Open Xml Library V1.0.0"
	doc.app.Security = 0

	// 初始化document.xml
	doc.GetSection().GetPageSize().SetWidth(DocumentDefaultWidth).SetHeight(DocumentDefaultHeight)
	doc.GetSection().GetPageMargin().SetMargin(1440, 1800, 1440, 1800)
	doc.GetSection().GetPageMargin().SetHeader(851).SetFooter(992)
	doc.GetSection().GetCols().SetSpace(425)

	// 默认样式
	doc.GetProperties().GetDefaultRunProperties().SetSize(21)

	return doc, nil
}
