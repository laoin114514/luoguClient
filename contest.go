package luoguclient

import (
	"fmt"
	"net/http"
)

// ContestService 比赛服务
type ContestService struct {
	client *Client
}

// GetList 获取比赛列表
func (c *ContestService) GetList(page int) (*ContestList, error) {
	if page <= 0 {
		page = 1
	}
	path := fmt.Sprintf("/contest/list?page=%d", page)
	resp, err := c.client.get(path)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("get contest list: status %d", resp.StatusCode)
	}

	var result struct {
		Data struct {
			Contests struct {
				Result  []ContestSummary `json:"result"`
				Count   int              `json:"count"`
				PerPage int              `json:"perPage"`
			} `json:"contests"`
		} `json:"data"`
	}
	if err := parseLentilleContext(resp, &result); err != nil {
		return nil, err
	}
	return &ContestList{
		Contests: result.Data.Contests.Result,
		Count:    result.Data.Contests.Count,
	}, nil
}

// GetDetail 获取比赛详情
func (c *ContestService) GetDetail(id int) (*ContestDetail, error) {
	path := fmt.Sprintf("/contest/%d", id)
	resp, err := c.client.get(path)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("get contest %d: status %d", id, resp.StatusCode)
	}

	var result struct {
		Data struct {
			Contest ContestDetail `json:"contest"`
			Joined int          `json:"joined"`
		} `json:"data"`
	}
	if err := parseLentilleContext(resp, &result); err != nil {
		return nil, err
	}
	result.Data.Contest.Joined = result.Data.Joined
	return &result.Data.Contest, nil
}
