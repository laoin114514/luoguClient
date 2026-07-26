package luogusdk

// RecordStatus 提交状态
type RecordStatus int

// Language 编程语言
type Language int

// --- 提交状态常量 ---

const (
	StatusWaiting             RecordStatus = 0  // 等待评测
	StatusRunningJudging      RecordStatus = 1  // 评测中
	StatusCompiling           RecordStatus = 2  // 编译中
	StatusOutputLimitExceeded RecordStatus = 3  // 输出超限
	StatusMemoryLimitExceeded RecordStatus = 4  // 内存超限
	StatusTimeLimitExceeded   RecordStatus = 5  // 时间超限
	StatusWrongAnswer         RecordStatus = 6  // 答案错误
	StatusRuntimeError        RecordStatus = 7  // 运行错误
	StatusCompileError        RecordStatus = 8  // 编译错误
	StatusAccepted            RecordStatus = 12 // 通过
	StatusUnaccepted          RecordStatus = 14 // 未通过（部分分/有测试点未通过）
)

// --- 语言常量 ---

const (
	LangC      Language = 1  // C
	LangCPP    Language = 2  // C++98
	LangPascal Language = 3  // Pascal
	LangJava   Language = 4  // Java
	LangPython Language = 7  // Python 3
	LangGo     Language = 14 // Go
	LangCPP14  Language = 28 // C++14
)
