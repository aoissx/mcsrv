package download

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/aoissx/mcsrv/config"
	"github.com/aoissx/mcsrv/model"
)

const (
	url       = "https://api.papermc.io/v2/projects/"
	serverJar = "server.jar"
)

func DownloadServerJar(c model.ConfigModel) error {

	name := c.Server.Name
	version := c.Server.Version

	// バージョンの整合性確認
	versions, err := getVersions(name)
	if err != nil {
		return err
	}
	// バージョンが存在するか確認
	if version == "latest" {
		version = versions[len(versions)-1]
	} else {
		if !contains(versions, version) {
			return fmt.Errorf("Version %s is not found.", version)
		}
	}

	// 最新ビルド番号取得
	build, err := getLatestBuild(name, version)

	// ビルド番号の整合性確認
	if err != nil {
		return err
	}
	config.LogInfo(fmt.Sprintf("Project: %s, Version: %s, Build: %d", name, version, build))
	// jarnameを設定
	jarname := fmt.Sprintf("%s-%s-%d.jar", name, version, build)

	// jarをダウンロード
	err = downloadJar(name, version, build, jarname)
	if err != nil {
		return err
	}

	return nil
}

// 最新ビルド番号取得
func getLatestBuild(name string, version string) (int, error) {
	api := url + name + "/versions/" + version
	data, err := getApiResult(api)
	if err != nil {
		return 0, err
	}

	// dataをパースしてビルド番号を取得
	var builds model.PaperBuildsAPIResponse
	err = json.Unmarshal(data, &builds)
	if err != nil {
		return 0, err
	}

	if builds.Error != "" {
		return 0, err
	}

	// 最新ビルド番号を返す
	// builds.Buildsの最後の要素を取得
	latestBuild := builds.Builds[len(builds.Builds)-1]
	return latestBuild, nil
}

// 最新バージョン取得
func getVersions(name string) ([]string, error) {
	api := url + name
	data, err := getApiResult(api)
	if err != nil {
		return nil, err
	}

	// dataをパースしてバージョンを取得
	var versions model.VersionsApiResponse
	err = json.Unmarshal(data, &versions)
	if err != nil {
		return nil, err
	}

	if versions.Error != "" {
		return nil, err
	}

	return versions.Versions, nil
}

// バージョンが存在するか確認
func contains(versions []string, version string) bool {
	for _, v := range versions {
		if v == version {
			return true
		}
	}
	return false
}

// jarをダウンロード
func downloadJar(name string, version string, build int, jarname string) error {
	// api := url + name + "/versions/" + version + "/builds/" + string(build) + "/downloads/" + jarname
	api := fmt.Sprintf("%s/%s/versions/%s/builds/%d/downloads/%s", url, name, version, build, jarname)

	file, err := os.Create(serverJar)
	if err != nil {
		return err
	}
	defer file.Close()

	resp, err := http.Get(api)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// jarを保存
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}
	return nil
}
