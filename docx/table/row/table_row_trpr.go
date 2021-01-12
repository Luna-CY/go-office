package row

// TrPr 行的属性结构定义
type TrPr struct {
    // cantSplit 设置为不可分割的行
    // 设置该标志位后，如果该行在一页内无法显示，会将此行放在下一页显示
    // 如果整页无法显示此行，将会在新页面顶头显示此行并延续至下一页
    cantSplit bool

    // hidden 设置此行为隐藏行
    // 但是需要注意，文档编辑软件可以有显示隐藏内容的设置
    hidden bool

    // horizontalAlignment 行的水平对齐方式
    // 此设置将会覆盖并独立与表格的水平对齐配置
    horizontalAlignment *HorizontalAlignmentType

    // cellSpacing 设置本行内单元格之间的间距
    // 单位仅支持点的二十分之一
    cellSpacing *int

    // header 设置该行为标题行
    // 标题行将在表格的开始以及表格每页的开始显示在第一行
    // 注意如果设置此标志的行不是表格首行，该标志将被忽略
    header bool

    // height 行的高度
    // 单位为点的二十分之一
    height *int

    // heightRule 行高度值的规则
    // 如果未指定此规则默认为 HeightRuleTypeExact
    // 如果未设置行高则此项无效
    heightRule *HeightRuleType
}

type HorizontalAlignmentType string

const (
    // HorizontalAlignmentStart
    // 行以文本段的左侧沿线对齐
    HorizontalAlignmentStart = HorizontalAlignmentType("start")

    // HorizontalAlignmentEnd
    // 行以文本段的右侧沿线对齐
    HorizontalAlignmentEnd = HorizontalAlignmentType("end")

    // HorizontalAlignmentCenter
    // 行以文本段的居中位置显示
    HorizontalAlignmentCenter = HorizontalAlignmentType("center")
)

type HeightRuleType string

const (
    // HeightRuleTypeAuto 设置为编辑软件自行定义
    HeightRuleTypeAuto = HeightRuleType("auto")

    // HeightRuleTypeAtLeast 设置为行高最少为指定的值
    HeightRuleTypeAtLeast = HeightRuleType("atLeast")

    // HeightRuleTypeExact 设置行高为指定的固定值
    HeightRuleTypeExact = HeightRuleType("exact")
)
