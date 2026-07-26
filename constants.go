package luogusdk

// RecordStatus 提交状态
type RecordStatus int

// Language 编程语言
type Language int

// --- 提交状态常量 ---
// 通过实际抓包确认的：
const (
	StatusCompiling  RecordStatus = 2  // 编译中（实测：score=null 的记录）
	StatusAccepted   RecordStatus = 12 // 通过（实测：score=100 的记录全是 12）
	StatusUnaccepted RecordStatus = 14 // 未通过/部分分（实测：score<100 的记录全是 14）
)

// --- 语言常量 ---
// 通过实际抓包确认的：
const (
	LangGo    Language = 14 // Go（实测：含 Go 源码的记录 language=14）
	LangCPP14 Language = 28 // C++14（实测：部分记录 language=28）
)
