package luogusdk

// CaptchaSolver 验证码求解器，接收 JPEG 图片字节，返回识别结果
type CaptchaSolver func(image []byte) (string, error)

// LoginRequest 登录请求体
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Captcha  string `json:"captcha"`
}

// LoginResponse 登录响应（会话由 HTTP cookie 维护，此处字段供参考）
type LoginResponse struct {
	UID      int    `json:"uid"`
	ClientID string `json:"client_id"`
}

// Problem 题目详情（匹配洛谷 SSR 页面中 lentille-context 的实际结构）
type Problem struct {
	PID        string         `json:"pid"`
	Title      string         `json:"name"`
	Difficulty int            `json:"difficulty"`
	Tags       []int          `json:"tags"`
	Samples    [][]string     `json:"samples"`
	Limits     ProblemLimits  `json:"limits"`
	Provider   UserInfo       `json:"provider"`
	Content    ProblemContent `json:"contenu"`
}

// DescText 返回题目描述的 Markdown 文本
func (p *Problem) DescText() string { return p.Content.Description }

// InputText 返回输入格式描述
func (p *Problem) InputText() string { return p.Content.InputFormat }

// OutputText 返回输出格式描述
func (p *Problem) OutputText() string { return p.Content.OutputFormat }

// HintText 返回说明/提示
func (p *Problem) HintText() string { return p.Content.Hint }

// TimeLimit 返回时间限制（ms），取第一个值
func (p *Problem) TimeLimit() int {
	if len(p.Limits.Time) > 0 {
		return p.Limits.Time[0]
	}
	return 0
}

// MemoryLimit 返回内存限制（KB），取第一个值
func (p *Problem) MemoryLimit() int {
	if len(p.Limits.Memory) > 0 {
		return p.Limits.Memory[0]
	}
	return 0
}

// ProblemContent 题目内容（从 contenu/content 字段提取）
type ProblemContent struct {
	Description string `json:"description"`
	InputFormat string `json:"formatI"`
	OutputFormat string `json:"formatO"`
	Hint        string `json:"hint"`
	Background  string `json:"background"`
}

// ProblemLimits 时空限制
type ProblemLimits struct {
	Time   []int `json:"time"`
	Memory []int `json:"memory"`
}

// SearchParams 题目搜索参数
type SearchParams struct {
	Keyword  string
	Page     int
	PageSize int
}

// SearchResult 搜索结果
type SearchResult struct {
	Problems []ProblemSummary
	Total    int
	Page     int
	PerPage  int
}

// ProblemSummary 题目摘要
type ProblemSummary struct {
	PID        string `json:"pid"`
	Title      string `json:"name"`
	Difficulty int    `json:"difficulty"`
	Tags       []int  `json:"tags"`
}

// Solution 题解详情
type Solution struct {
	ID      string   `json:"lid"`
	Author  UserInfo `json:"author"`
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Likes   int      `json:"upvote"`
}

// SolutionList 题解列表
type SolutionList struct {
	Solutions []SolutionSummary
	Total     int
	Page      int
	PerPage   int
}

// SolutionSummary 题解摘要
type SolutionSummary struct {
	ID     string   `json:"lid"`
	Title  string   `json:"title"`
	Author UserInfo `json:"author"`
	Likes  int      `json:"upvote"`
}

// Translation 题目翻译
type Translation struct {
	Language     string `json:"-"`
	Title        string `json:"name"`
	Description  string `json:"description"`
	InputFormat  string `json:"formatI"`
	OutputFormat string `json:"formatO"`
	Hint         string `json:"hint"`
	Background   string `json:"background"`
}

// UserInfo 用户信息
type UserInfo struct {
	UID    int    `json:"uid"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

// --- 记录 (Record) ---

// RecordListParams 记录搜索参数
type RecordListParams struct {
	User    int          // 用户 UID（0 表示不限制）
	Problem string       // 题目 PID（空表示不限制）
	Status  RecordStatus // 状态过滤（0 表示全部，可用 StatusAccepted 等常量）
	Page    int          // 页码
}

// RecordList 记录列表
type RecordList struct {
	Records []RecordSummary
	Count   int
}

// RecordSummary 记录摘要
type RecordSummary struct {
	ID               int          `json:"id"`
	Status           RecordStatus `json:"status"`
	Score            int          `json:"score"`
	Time             int          `json:"time"`
	Memory           int          `json:"memory"`
	SourceCodeLength int          `json:"sourceCodeLength"`
	SubmitTime       int64        `json:"submitTime"`
	Language         Language     `json:"language"`
	EnableO2         bool         `json:"enableO2"`
	Problem          ProblemRef   `json:"problem"`
	User             UserInfo     `json:"user"`
}

// ProblemRef 记录中的题目引用
type ProblemRef struct {
	PID        string `json:"pid"`
	Title      string `json:"title"`
	Difficulty int    `json:"difficulty"`
	FullScore  int    `json:"fullScore"`
	Type       string `json:"type"`
}

// RecordDetail 记录详情（含源代码和评测详情）
type RecordDetail struct {
	RecordSummary
	SourceCode string      `json:"sourceCode"`
	Detail     JudgeDetail `json:"detail"`
}

// JudgeDetail 评测详情
type JudgeDetail struct {
	CompileResult CompileResult `json:"compileResult"`
	JudgeResult   JudgeResult   `json:"judgeResult"`
}

// CompileResult 编译结果
type CompileResult struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// JudgeResult 评测结果
type JudgeResult struct {
	Subtasks          []SubtaskResult `json:"subtasks"`
	FinishedCaseCount int             `json:"finishedCaseCount"`
	Status            RecordStatus    `json:"status"`
	Time              int             `json:"time"`
	Memory            int             `json:"memory"`
	Score             int             `json:"score"`
}

// SubtaskResult 子任务结果
type SubtaskResult struct {
	ID        int                       `json:"id"`
	Score     int                       `json:"score"`
	Status    RecordStatus              `json:"status"`
	Time      int                       `json:"time"`
	Memory    int                       `json:"memory"`
	TestCases map[string]TestCaseResult `json:"testCases"`
}

// TestCaseResult 单个测试点结果
type TestCaseResult struct {
	ID          int          `json:"id"`
	Status      RecordStatus `json:"status"`
	Time        int          `json:"time"`
	Memory      int          `json:"memory"`
	Score       int          `json:"score"`
	Signal      int          `json:"signal"`
	ExitCode    int          `json:"exitCode"`
	Description string       `json:"description"`
	SubtaskID   int          `json:"subtaskID"`
}
