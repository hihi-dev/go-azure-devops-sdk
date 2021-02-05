package ado

import (
	"encoding/json"
	"fmt"
	"time"
)

type Repository struct {
	Id            string  `json:"id"`
	Url           string  `json:"url"`
	Name          string  `json:"name"`
	Size          int64   `json:"size"`
	SshUrl        string  `json:"sshUrl"`
	WebUrl        string  `json:"webUrl"`
	Project       Project `json:"project"`
	RemoteUrl     string  `json:"remoteUrl"`
	DefaultBranch string  `json:"defaultBranch"`
}

type Project struct {
	Id             string    `json:"id"`
	Url            string    `json:"url"`
	Name           string    `json:"name"`
	State          string    `json:"state"`
	Revision       int       `json:"revision"`
	Visibility     string    `json:"visibility"`
	Description    string    `json:"description"`
	LastUpdateTime time.Time `json:"lastUpdateTime"`
}

type repoListBody struct {
	Value []Repository `json:"value"`
}

func (c *Client) ListRepositoriesInProject(p string) ([]Repository, error) {
	path := fmt.Sprintf("/%s/_apis/git/repositories", p)
	return c.fetchRepoList(path)
}

func (c *Client) ListAllRepositories() ([]Repository, error) {
	return c.fetchRepoList("/_apis/git/repositories")
}

func (c *Client) GetRepositoryById(id string) (*Repository, error) {
	path := fmt.Sprintf("/_apis/git/repositories/%s", id)
	return c.fetchGitRepo(path)
}

func (c *Client) GetRepositoryByName(project, repo string) (*Repository, error) {
	path := fmt.Sprintf("/%s/_apis/git/repositories/%s", project, repo)
	return c.fetchGitRepo(path)
}

func (c *Client) fetchGitRepo(path string) (*Repository, error) {
	resp, err := c.doRequestForBody(BaseUrl, "GET", path, nil)
	if err != nil {
		return nil, err
	}
	r := &Repository{}
	je := json.Unmarshal(resp, r)
	return r, je
}

func (c *Client) fetchRepoList(path string) ([]Repository, error) {
	resp, err := c.doRequestForBody(BaseUrl, "GET", path, nil)
	if err != nil {
		return nil, err
	}
	r := &repoListBody{}
	je := json.Unmarshal(resp, r)
	return r.Value, je
}
