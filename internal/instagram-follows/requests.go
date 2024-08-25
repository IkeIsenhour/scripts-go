package insta

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/IkeIsenhour/scripts-go/pkg/networking"
)

type InstagramRequest struct {
	Client    *http.Client
	AccountId string
	XIgAppId  string
	Cookie    string
}

func NewInstagramRequest(c *http.Client, a string, x string, co string) *InstagramRequest {
	return &InstagramRequest{
		Client:    c,
		AccountId: a,
		XIgAppId:  x,
		Cookie:    co,
	}
}

func (i *InstagramRequest) GetAllFollowing() ([]Users, error) {
	headers := make(map[string]string)
	headers["X-IG-App-ID"] = i.XIgAppId
	headers["Cookie"] = i.Cookie

	var users []Users
	pageSize := 200
	maxId := "0"

	for pageSize == 200 {
		url := fmt.Sprintf("https://www.instagram.com/api/v1/friendships/%s/following/?count=200&max_id=%s", i.AccountId, maxId)

		res, err := networking.MakeRequest(i.Client, http.MethodGet, url, nil, &headers)
		if err != nil {
			return nil, err
		}

		bytes, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		f := &friendships{}
		json.Unmarshal(bytes, f)

		pageSize = f.PageSize
		maxId = f.NextMaxID

		users = append(users, f.Users...)
	}

	return users, nil
}

func (i *InstagramRequest) GetAllFollowers() ([]Users, error) {
	headers := make(map[string]string)
	headers["X-IG-App-ID"] = i.XIgAppId
	headers["Cookie"] = i.Cookie

	var users []Users
	var maxId string

	for {
		url := fmt.Sprintf("https://www.instagram.com/api/v1/friendships/%s/followers/?count=24&max_id=%s&search_surface=follow_list_page", i.AccountId, maxId)

		res, err := networking.MakeRequest(i.Client, http.MethodGet, url, nil, &headers)
		if err != nil {
			return nil, err
		}

		bytes, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		f := &friendships{}
		json.Unmarshal(bytes, f)

		if f.NextMaxID == "" {
			users = append(users, f.Users...)
			return users, nil
		}
		maxId = f.NextMaxID

		users = append(users, f.Users...)
	}
}
