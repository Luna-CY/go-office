package run

type BreakLineType string

const (
    // BreakLineTypeDefault 下一行开始
    // 此类型为默认类型
    BreakLineTypeDefault = BreakLineType("textWrapping")

    // BreakLineTypePage 从下一页开始
    // 设置为此类型时，新的内容将从下一页开始
    BreakLineTypePage = BreakLineType("page")

    // BreakLineTypePage 从下一列开始
    BreakLineTypeColumn = BreakLineType("column")
)

type BreakLineClearType string

const (
    // BreakLineClearTypeDefault 不设置clear
    // 这是默认值
    BreakLineClearTypeDefault = BreakLineClearType("none")

    // BreakLineClearTypeLeft 从左侧开始
    BreakLineClearTypeLeft = BreakLineClearType("left")

    // BreakLineClearTypeRight 从右侧开始
    BreakLineClearTypeRight = BreakLineClearType("right")

    // BreakLineClearTypeAll
    BreakLineClearTypeAll = BreakLineClearType("all")
)
