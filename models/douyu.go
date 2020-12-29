package models

import (
	"encoding/json"
	"fmt"
)

const (
	DouyuBaseUrl = "http://capi.douyucdn.cn/api/v1"
)

var douyuHeader map[string]string

func init() {
	douyuHeader = make(map[string]string)
	douyuHeader["User-Agent"] = "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_0_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36"
}

type DouyuResponse struct {
	Error int                      `json:"error"`
	Data  []map[string]interface{} `json:"data"`
}

type DouyuCategory struct {
	CateID    string `json:"cate_id"`
	CateName  string `json:"cate_name"`
	ShortName string `json:"short_name"`
}

type DouyuSubCategory struct {
	TagID   string `json:"tag_id"`
	TagName string `json:"tag_name"`
	PicName string `json:"pic_name"`
	Count   int    `json:"count"`
}

type DouyuRoom struct {
	RoomID   string `json:"room_id"`
	RoomName string `json:"room_name"`
	GameName string `json:"game_name"`
	NickName string `json:"nickname"`
	PicName  string `json:"pic_name"`
	Online   int    `json:"online"`
}

type DouyuSource struct {
	RoomID string `json:"room_id"`
	Source string `json:"source"`
}

func ListDouyuRoomByTagID(tagID string, page *Page) ([]*DouyuRoom, error) {
	res, err := HttpGet(fmt.Sprintf("%s/live/%s?offset=%d&limit=%d", DouyuBaseUrl, tagID, page.Offset, page.Limit), douyuHeader)
	if err != nil {
		return nil, err
	}

	var rsp DouyuResponse
	if err := json.Unmarshal(res, &rsp); err != nil {
		return nil, err
	}

	dcs := []*DouyuRoom{}
	for _, data := range rsp.Data {
		dcs = append(dcs, &DouyuRoom{
			RoomID:   data["room_id"].(string),
			RoomName: data["room_name"].(string),
			GameName: data["game_name"].(string),
			NickName: data["nickname"].(string),
			PicName:  data["room_src"].(string),
			Online:   int(data["online"].(float64)),
		})
	}

	return dcs, nil
}

func ListDouyuRoomByCateID(cateID string, page *Page) ([]*DouyuRoom, error) {
	res, err := HttpGet(fmt.Sprintf("%s/getColumnRoom/%s?offset=%d&limit=%d", DouyuBaseUrl, cateID, page.Offset, page.Limit), douyuHeader)
	if err != nil {
		return nil, err
	}

	var rsp DouyuResponse
	if err := json.Unmarshal(res, &rsp); err != nil {
		return nil, err
	}

	dcs := []*DouyuRoom{}
	for _, data := range rsp.Data {
		dcs = append(dcs, &DouyuRoom{
			RoomID:   data["room_id"].(string),
			RoomName: data["room_name"].(string),
			GameName: data["game_name"].(string),
			NickName: data["nickname"].(string),
			PicName:  data["room_src"].(string),
			Online:   int(data["online"].(float64)),
		})
	}

	return dcs, nil
}

func ListDouyuRoom(page *Page) ([]*DouyuRoom, error) {
	res, err := HttpGet(fmt.Sprintf("%s/live?offset=%d&limit=%d", DouyuBaseUrl, page.Offset, page.Limit), douyuHeader)
	if err != nil {
		return nil, err
	}

	var rsp DouyuResponse
	if err := json.Unmarshal(res, &rsp); err != nil {
		return nil, err
	}

	dcs := []*DouyuRoom{}
	for _, data := range rsp.Data {
		dcs = append(dcs, &DouyuRoom{
			RoomID:   data["room_id"].(string),
			RoomName: data["room_name"].(string),
			GameName: data["game_name"].(string),
			NickName: data["nickname"].(string),
			PicName:  data["room_src"].(string),
			Online:   int(data["online"].(float64)),
		})
	}

	return dcs, nil
}

func ListDouyuSubCategory(shortName string) ([]*DouyuSubCategory, error) {
	res, err := HttpGet(fmt.Sprintf("%s/getColumnDetail?shortName=%s", DouyuBaseUrl, shortName), douyuHeader)
	if err != nil {
		return nil, err
	}

	var rsp DouyuResponse
	if err := json.Unmarshal(res, &rsp); err != nil {
		return nil, err
	}

	dcs := []*DouyuSubCategory{}
	for _, data := range rsp.Data {
		dcs = append(dcs, &DouyuSubCategory{
			TagID:   data["tag_id"].(string),
			TagName: data["tag_name"].(string),
			PicName: data["pic_name"].(string),
			Count:   int(data["count"].(float64)),
		})
	}

	return dcs, nil
}

func ListDouyuCategory() ([]*DouyuCategory, error) {
	res, err := HttpGet(fmt.Sprintf("%s/getColumnList", DouyuBaseUrl), douyuHeader)
	if err != nil {
		return nil, err
	}

	var rsp DouyuResponse
	if err := json.Unmarshal(res, &rsp); err != nil {
		return nil, err
	}

	dcs := []*DouyuCategory{}
	for _, data := range rsp.Data {
		dcs = append(dcs, &DouyuCategory{
			CateID:    data["cate_id"].(string),
			CateName:  data["cate_name"].(string),
			ShortName: data["short_name"].(string),
		})
	}

	return dcs, nil
}

/*
func FindDouyuRoomByRoomID(db *bolt.DB, id string) (*DouyuRoom, error) {
	dr := DouyuRoom{ID: id}
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(DatabaseDouyu))
		v := b.Get([]byte(id))
		dr.Url = string(v)
		return nil
	})
	return &dr, err
}
*/
