package luogusdk

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// RecordService 记录服务
type RecordService struct {
	client *Client
}

// GetList 获取记录列表
func (r *RecordService) GetList(params RecordListParams) (*RecordList, error) {
	q := url.Values{}
	if params.User > 0 {
		q.Set("user", strconv.Itoa(params.User))
	}
	if params.Problem != "" {
		q.Set("pid", params.Problem)
	}
	if params.Status > 0 {
		q.Set("status", strconv.Itoa(int(params.Status)))
	}
	page := params.Page
	if page <= 0 {
		page = 1
	}
	q.Set("page", strconv.Itoa(page))

	path := "/record/list?" + q.Encode()
	resp, err := r.client.get(path)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("get record list: status %d", resp.StatusCode)
	}

	var result struct {
		CurrentData struct {
			Records struct {
				Result  []RecordSummary `json:"result"`
				Count   int             `json:"count"`
				PerPage int             `json:"perPage"`
			} `json:"records"`
		} `json:"currentData"`
	}
	if err := parseFeInjection(resp, &result); err != nil {
		return nil, err
	}
	return &RecordList{
		Records: result.CurrentData.Records.Result,
		Count:   result.CurrentData.Records.Count,
	}, nil
}

// GetDetail 获取记录详情（含源代码和评测结果）
func (r *RecordService) GetDetail(rid int) (*RecordDetail, error) {
	path := fmt.Sprintf("/record/%d", rid)
	resp, err := r.client.get(path)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("get record %d: status %d", rid, resp.StatusCode)
	}

	var result struct {
		CurrentData struct {
			Record RecordDetail `json:"record"`
		} `json:"currentData"`
	}
	if err := parseFeInjection(resp, &result); err != nil {
		return nil, err
	}
	return &result.CurrentData.Record, nil
}

// parseFeInjection 从 HTML 页面中提取 window._feInjection 内的 JSON 数据（调用方负责关闭 resp.Body）
func parseFeInjection(resp *http.Response, v interface{}) error {
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return fmt.Errorf("parse HTML: %w", err)
	}

	var encoded string
	doc.Find("script").Each(func(_ int, s *goquery.Selection) {
		if encoded != "" {
			return // 已找到
		}
		text := s.Text()
		idx := strings.Index(text, "_feInjection")
		if idx < 0 {
			return
		}
		// 提取 decodeURIComponent("...") 中的引号内容
		after := text[idx:]
		q1 := strings.Index(after, `"`)
		if q1 < 0 {
			return
		}
		q2 := strings.Index(after[q1+1:], `"`)
		if q2 < 0 {
			return
		}
		encoded = after[q1+1 : q1+1+q2]
	})

	if encoded == "" {
		return fmt.Errorf("_feInjection not found in page")
	}

	decoded, err := url.QueryUnescape(encoded)
	if err != nil {
		return fmt.Errorf("decode _feInjection: %w", err)
	}

	if err := json.Unmarshal([]byte(decoded), v); err != nil {
		return fmt.Errorf("unmarshal _feInjection: %w", err)
	}
	return nil
}
