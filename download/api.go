package download

import (
	"io"
	"net/http"
)

func getApiResult(url string) ([]byte, error) {
	// urlにGETリクエストを送信し、結果を返す
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// リクエストヘッダーを設定
	req.Header.Set("Content-Type", "application/json")

	// リクエストを送信
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 結果を読み込む
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
