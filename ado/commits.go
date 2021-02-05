package ado

import (
	"encoding/json"
	"fmt"
)

type commitsResponseBody struct {
	Count int      `json:"count"`
	Value []Commit `json:"value"`
}

type Commit struct {
	Id        string `json:"commitId"`
	Author    Author `json:"author"`
	Committer Author `json:"committer"`
	Comment   string `json:"comment"`
	Url       string `json:"url"`
	RemoteUrl string `json:"remoteUrl"`
}

type Author struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Date  string `json:"date"`
}

type ChangeCount struct {
	Add    int `json:"add"`
	Edit   int `json:"edit"`
	Delete int `json:"delete"`
}

func (c *Client) GetCommitsForRepositoryById(id string) ([]Commit, error) {
	path := fmt.Sprintf("/_apis/git/repositories/%s/commits", id)
	return c.fetchCommits(path)
}

func (c *Client) GetCommitsForRepositoryByName(project, repo string) ([]Commit, error) {
	path := fmt.Sprintf("/%s/_apis/git/repositories/%s/commits", project, repo)
	return c.fetchCommits(path)
}

func (c *Client) GetDiffBetweenVersionsById(id, source, target string) ([]Commit, error) {
	path := fmt.Sprintf("/_apis/git/repositories/%s/commits?searchCriteria.compareVersion.version=%s&searchCriteria.itemVersion.version=%s", id, source, target)
	return c.fetchCommits(path)
}

func (c *Client) GetDiffBetweenVersionsByName(project, repo, source, target string) ([]Commit, error) {
	path := fmt.Sprintf("/%s/_apis/git/repositories/%s/commits?searchCriteria.compareVersion.version=%s&searchCriteria.itemVersion.version=%s", project, repo, source, target)
	return c.fetchCommits(path)
}

func (c *Client) fetchCommits(path string) ([]Commit, error) {
	resp, err := c.doRequestForBody(BaseUrl, "GET", path, nil)
	if err != nil {
		return nil, err
	}
	r := &commitsResponseBody{}
	je := json.Unmarshal(resp, r)
	return r.Value, je
}
