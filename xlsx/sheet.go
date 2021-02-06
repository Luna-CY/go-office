package xlsx

import (
    "strconv"
    "sync"
)

// Sheet 工作表定义
type Sheet struct {
    // meta 元数据
    meta struct {
        // id 工作表ID
        id string

        // name 工作表名
        name string

        // relationshipId 关联的定义ID
        relationshipId string
    }

    rm sync.RWMutex
    // rows 行
    rows []*Row
}

// SetName 设置数据表名
func (s *Sheet) SetName(name string) *Sheet {
    s.meta.name = name

    return s
}

// AddRow 添加一行数据
func (s *Sheet) AddRow() *Row {
    r := new(Row)
    r.id = strconv.Itoa(len(s.rows) + 1)

    s.rm.Lock()
    defer s.rm.Unlock()

    s.rows = append(s.rows, r)

    return r
}
