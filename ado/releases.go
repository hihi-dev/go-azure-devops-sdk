package ado

import (
	"encoding/json"
	"fmt"
	"time"
)

type releaseBody struct {
	Count int       `json:"count"`
	Value []Release `json:"value"`
}

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
	Branch     Reference `json:"branch"`
	Version    Reference `json:"version"`
	Project    Reference `json:"project"`
	Definition Reference `json:"definition"`
}

type Reference struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func (c Client) GetReleasesForProject(pn string) ([]Release, error) {
	path := fmt.Sprintf("/%s/_apis/release/releases", pn)
	resp, err := c.doRequestForBody(BaseUrlLegacy, "GET", path, nil)
	if err != nil {
		return nil, err
	}
	list := &releaseBody{}
	je := json.Unmarshal(resp, list)
	return list.Value, je
}

func (c Client) GetReleaseDetails(pn string, id int) (*ReleaseDetails, error) {
	path := fmt.Sprintf("/%s/_apis/release/releases/%d", pn, id)
	resp, err := c.doRequestForBody(BaseUrlLegacy, "GET", path, nil)
	if err != nil {
		return nil, err
	}
	rel := &ReleaseDetails{}
	je := json.Unmarshal(resp, rel)
	return rel, je
}
