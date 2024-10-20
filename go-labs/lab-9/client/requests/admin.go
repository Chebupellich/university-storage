package requests

import (
	"bytes"
	"client/exceptions"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func AdminAccess(secret string) {
	credentials := map[string]string{"secret": secret}
	body, _ := json.Marshal(credentials)

	req, err := http.NewRequest("POST", baseURL+"/admin-access", bytes.NewBuffer(body))
	if err != nil {
		exceptions.BadRequestError(err)
		return
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		exceptions.BadRequestError(err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		var result map[string]string
		body, _ := io.ReadAll(resp.Body)
		err = json.Unmarshal(body, &result)
		if err != nil {
			exceptions.InternalServerError(err)
			return
		}

		accessToken = result["token"]

		fmt.Println(" * Get root access * ")
	} else {
		exceptions.GeneralError(nil)
	}
}
