package xlsx

// Xml xml文件结构接口
type Xml interface {

	// FilePath 返回文件路径
	FilePath() string

	// Marshal 序列化文件的xml内容
	Marshal() ([]byte, error)
}
