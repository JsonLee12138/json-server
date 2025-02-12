package core

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/JsonLee12138/json-server/internal/global"
	"github.com/JsonLee12138/json-server/pkg/utils"
)

var (
	baseUrl    = "https://api.apifox.com/v1"
	ImportPath = "/projects/{projectId}/import-openapi"
)

type Apifox struct{}

type ImportOptions struct {
	TargetEndpointFolderId        string `json:"targetEndpointFolderId"`
	TargetSchemaFolderId          string `json:"targetSchemaFolderId"`
	EndpointOverwriteBehavior     string `json:"endpointOverwriteBehavior"`
	SchemaOverwriteBehavior       string `json:"schemaOverwriteBehavior"`
	UpdateFolderOfChangedEndpoint bool   `json:"updateFolderOfChangedEndpoint"`
	PrependBasePath               bool   `json:"prependBasePath"`
}
type ImportDTO struct {
	Input   string        `json:"input"`
	Options ImportOptions `json:"options"`
}

func (a *Apifox) getURL(path, projectId, locale string) string {
	u, err := url.Parse(baseUrl)
	if err != nil {
		return ""
	}
	u.Path += strings.Replace(path, "{projectId}", projectId, 1)
	q := u.Query()
	q.Set("locale", locale)
	u.RawQuery = q.Encode()
	return u.String()
}

func (a *Apifox) Import(data []byte) error {
	locale := utils.DefaultIfEmpty(global.Config.Apifox.Locale, "zh-CN")
	url := a.getURL(ImportPath, global.Config.Apifox.ProjectId, locale)
	importDTO := ImportDTO{
		Input: string(data),
		Options: ImportOptions{
			TargetEndpointFolderId:        global.Config.Apifox.TargetEndpointFolderId,
			TargetSchemaFolderId:          global.Config.Apifox.TargetSchemaFolderId,
			EndpointOverwriteBehavior:     string(global.Config.Apifox.EndpointOverwriteBehavior),
			SchemaOverwriteBehavior:       string(global.Config.Apifox.SchemaOverwriteBehavior),
			UpdateFolderOfChangedEndpoint: global.Config.Apifox.UpdateFolderOfChangedEndpoint,
			PrependBasePath:               global.Config.Apifox.PrependBasePath,
		},
	}

	jsonData, err := json.Marshal(importDTO)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", global.Config.Apifox.Token))
	req.Header.Set("X-Apifox-Api-Version", global.Config.Apifox.ApiVersion)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	if res == nil {
		return errors.New("请求错误")
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("import failed, status code: %d, body: %s", res.StatusCode, string(body))
	}
	return nil
}
