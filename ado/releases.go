package ado

import (
	"encoding/json"
	"fmt"
	"time"
)

type ReleaseDefinition struct {
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	Path       string    `json:"path"`
}

type Release struct {
	Id                int               `json:"id"`
	Url               string            `json:"url"`
	Name              string            `json:"name"`
	Status            string            `json:"status"`
	CreatedOn         time.Time         `json:"createdOn"`
	ModifiedOn        time.Time         `json:"modifiedOn"`
	ReleaseDefinition ReleaseDefinition `json:"releaseDefinition"`
}

type ReleaseDetails struct {
	Id                int               `json:"id"`
	Url               string            `json:"url"`
	Name              string            `json:"name"`
	Status            string            `json:"status"`
	CreatedOn         time.Time         `json:"createdOn"`
	ModifiedOn        time.Time         `json:"modifiedOn"`
	Artifacts         []Artifact        `json:"artifacts"`
	ReleaseDefinition ReleaseDefinition `json:"releaseDefinition"`
}

type Artifact struct {
	Type                string              `json:"type"`
	Alias               string              `json:"alias"`
	DefinitionReference DefinitionReference `json:"definitionReference"`
}

type DefinitionReference struct {
	Branch  Reference `json:"branch"`
	Version Reference `json:"version"`
}

type Reference struct {
	Id         int       `json:"id"`
	Name       string    `json:"name"`
}

func (c Client) GetReleasesForProject(pn string) ([]Release, error) {
	path := fmt.Sprintf("/%s/_apis/release/releases", pn)
	resp, err := c.doRequestForBody("GET", path, nil)
	if err != nil {
		return nil, err
	}
	list := &[]Release{}
	je := json.Unmarshal(resp, list)
	return *list, je
}

func (c Client) GetReleaseDetails(pn string, id int) (*ReleaseDetails, error) {
	path := fmt.Sprintf("/%s/_apis/release/releases/%d", pn, id)
	resp, err := c.doRequestForBody("GET", path, nil)
	if err != nil {
		return nil, err
	}
	rel := &ReleaseDetails{}
	je := json.Unmarshal(resp, rel)
	return rel, je
}
