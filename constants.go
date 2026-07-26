package luogusdk

// --- 提交状态常量 ---

const (
	StatusWaiting          = 0  // 等待评测
	StatusRunningJudging   = 1  // 评测中
	StatusCompiling        = 2  // 编译中
	StatusOutputLimitExceeded = 3  // 输出超限
	StatusMemoryLimitExceeded  = 4  // 内存超限
	StatusTimeLimitExceeded    = 5  // 时间超限
	StatusWrongAnswer          = 6  // 答案错误
	StatusRuntimeError         = 7  // 运行错误
	StatusCompileError         = 8  // 编译错误
	StatusAccepted             = 12 // 通过
	StatusUnaccepted           = 14 // 未通过（部分分/有测试点未通过）
)

// --- 语言常量 ---

const (
	LangC         = 1  // C
	LangCPP       = 2  // C++98
	LangPascal    = 3  // Pascal
	LangJava      = 4  // Java
	LangPython    = 7  // Python 3
	LangGo        = 14 // Go
	LangCPP14     = 28 // C++14
)
