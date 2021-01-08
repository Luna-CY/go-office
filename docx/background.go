package docx

// themeTint 文档的背景配置
type Background struct {

    // color 背景色，16进制颜色值
    color string

    // themeColor 主题颜色的id，其指向的id在主题中定义
    themeColor string

    // themeShade 主题阴影颜色，0-255的16进制值
    themeShade string

    // themeTint 主题色调值，0-255的16进制值
    themeTint string
}
