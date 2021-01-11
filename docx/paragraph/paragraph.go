package paragraph

import (
    "github.com/Luna-CY/go-office/docx/run"
    "sync"
)

// Paragraph 段落结构
// 每个段落都包含自己的样式属性定义以及任意个 Run 结构
type Paragraph struct {

    // ppr 段落样式属性定义
    ppr *PPr

    // rpr 段落内容统一样式
    rpr *run.RPr

    rm sync.RWMutex
    // runs Run 结构列表
    runs []*run.Run
}

func (p *Paragraph) GetProperties() *PPr {
    if nil == p.ppr {
        p.ppr = new(PPr)
    }

    return p.ppr
}

func (p *Paragraph) GetRunProperties() *run.RPr {
    if nil == p.rpr {
        p.rpr = new(run.RPr)
    }

    return p.rpr
}

func (p *Paragraph) GetRuns() []*run.Run {
    p.rm.RLock()
    defer p.rm.RUnlock()

    return p.runs
}

// AddRun 新增一个内容结构
func (p *Paragraph) AddRun() *run.Run {
    r := new(run.Run)

    p.rm.Lock()
    p.runs = append(p.runs, r)
    p.rm.Unlock()

    return r
}

// AddBreakLine 添加换行符
func (p *Paragraph) AddBreakLine(breakLineType run.BreakLineType, breakLineClearType run.BreakLineClearType) {
    r := new(run.Run)

    p.rm.Lock()
    p.runs = append(p.runs, r)
    p.rm.Unlock()

    r.AddBreakLine(breakLineType, breakLineClearType)
}
