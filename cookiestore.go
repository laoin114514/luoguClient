package luogusdk

import (
	"encoding/json"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"path/filepath"
)

// exportableCookie 可序列化的 cookie 结构
type exportableCookie struct {
	Name     string `json:"name"`
	Value    string `json:"value"`
	Domain   string `json:"domain"`
	Path     string `json:"path"`
	Secure   bool   `json:"secure,omitempty"`
	HttpOnly bool   `json:"http_only,omitempty"`
}

// ExportableCookieJar 包装 cookiejar.Jar，支持持久化
type ExportableCookieJar struct {
	jar      *cookiejar.Jar
	savePath string // 非空时 SetCookies 自动写回文件
}

func newExportableCookieJar(savePath string) (*ExportableCookieJar, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}
	return &ExportableCookieJar{jar: jar, savePath: savePath}, nil
}

// setSavePath 更新持久化路径（空则禁用自动保存）
func (j *ExportableCookieJar) setSavePath(path string) {
	j.savePath = path
}

// SetCookies 设置 cookie，若 savePath 非空则自动持久化到文件
func (j *ExportableCookieJar) SetCookies(u *url.URL, cookies []*http.Cookie) {
	j.jar.SetCookies(u, cookies)
	if j.savePath != "" {
		// 忽略写文件错误 — http.CookieJar 接口不支持返回 error
		_ = saveCookies(j, j.savePath)
	}
}

func (j *ExportableCookieJar) Cookies(u *url.URL) []*http.Cookie {
	return j.jar.Cookies(u)
}

// Clear 清空内存中的 cookie 并删除持久化文件
func (j *ExportableCookieJar) Clear() error {
	newJar, err := cookiejar.New(nil)
	if err != nil {
		return err
	}
	j.jar = newJar
	if j.savePath != "" {
		_ = os.Remove(j.savePath)
	}
	return nil
}

// Export 将所有 cookie 导出为 JSON 字节
func (j *ExportableCookieJar) Export() ([]byte, error) {
	u, err := url.Parse(luoguBaseURL)
	if err != nil {
		return nil, err
	}
	cookies := j.jar.Cookies(u)
	exported := make([]exportableCookie, 0, len(cookies))
	for _, c := range cookies {
		exported = append(exported, exportableCookie{
			Name:     c.Name,
			Value:    c.Value,
			Domain:   c.Domain,
			Path:     c.Path,
			Secure:   c.Secure,
			HttpOnly: c.HttpOnly,
		})
	}
	return json.Marshal(exported)
}

// Import 从 JSON 字节导入 cookie
func (j *ExportableCookieJar) Import(data []byte) error {
	var cookies []exportableCookie
	if err := json.Unmarshal(data, &cookies); err != nil {
		return err
	}
	u, err := url.Parse(luoguBaseURL)
	if err != nil {
		return err
	}
	for _, c := range cookies {
		j.jar.SetCookies(u, []*http.Cookie{{
			Name:     c.Name,
			Value:    c.Value,
			Domain:   c.Domain,
			Path:     c.Path,
			Secure:   c.Secure,
			HttpOnly: c.HttpOnly,
		}})
	}
	return nil
}

// defaultCookiePath 返回默认 cookie 文件路径
func defaultCookiePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".luogu", "cookies.json"), nil
}

// saveCookies 保存 cookie 到文件
func saveCookies(jar *ExportableCookieJar, filePath string) error {
	data, err := jar.Export()
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(filePath), 0700); err != nil {
		return err
	}
	return os.WriteFile(filePath, data, 0600)
}

// loadCookies 从文件加载 cookie
func loadCookies(jar *ExportableCookieJar, filePath string) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	return jar.Import(data)
}
