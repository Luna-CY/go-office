package run

// Underline 下划线配置结构
type Underline struct {
	// color 下划线的颜色
	color *string

	// lineType 下划线的类型
	lineType UnderlineType
}

type UnderlineType string

const (
	UnderlineDash            = UnderlineType("dash")
	UnderlineDashDotDotHeavy = UnderlineType("dashDotDotHeavy")
	UnderlineDashDotHeavy    = UnderlineType("dashDotHeavy")
	UnderlineDashedHeavy     = UnderlineType("dashedHeavy")
	UnderlineDashLong        = UnderlineType("dashLong")
	UnderlineDashLongHeavy   = UnderlineType("dashLongHeavy")
	UnderlineDotDash         = UnderlineType("dotDash")
	UnderlineDotDotDash      = UnderlineType("dotDotDash")
	UnderlineDotted          = UnderlineType("dotted")
	UnderlineDottedHeavy     = UnderlineType("dottedHeavy")
	UnderlineDouble          = UnderlineType("double")
	UnderlineNone            = UnderlineType("none")
	UnderlineSingle          = UnderlineType("single")
	UnderlineThick           = UnderlineType("thick")
	UnderlineWave            = UnderlineType("wave")
	UnderlineWavyDouble      = UnderlineType("wavyDouble")
	UnderlineWavyHeavy       = UnderlineType("wavyHeavy")
	UnderlineWords           = UnderlineType("words")
)
