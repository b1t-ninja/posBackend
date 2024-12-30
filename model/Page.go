package model

import "time"

type Response struct {
	Object  string `json:"object"`
	Results []Page `json:"results"`
}

type Page struct {
	Object         string    `json:"object"`
	Id             string    `json:"id"`
	CreatedTime    time.Time `json:"created_time"`
	LastEditedTime time.Time `json:"last_edited_time"`
	CreatedBy      struct {
		Object string `json:"object"`
		Id     string `json:"id"`
	} `json:"created_by"`
	LastEditedBy struct {
		Object string `json:"object"`
		Id     string `json:"id"`
	} `json:"last_edited_by"`
	Cover  interface{} `json:"cover"`
	Icon   interface{} `json:"icon"`
	Parent struct {
		Type       string `json:"type"`
		DatabaseId string `json:"database_id"`
	} `json:"parent"`
	Archived   bool `json:"archived"`
	InTrash    bool `json:"in_trash"`
	Properties struct {
		Picture struct {
			Id    string `json:"id"`
			Type  string `json:"type"`
			Files []struct {
				Name string `json:"name"`
				Type string `json:"type"`
				File struct {
					Url        string    `json:"url"`
					ExpiryTime time.Time `json:"expiry_time"`
				} `json:"file"`
			} `json:"files"`
		} `json:"Picture"`
		Price struct {
			Id     string  `json:"id"`
			Type   string  `json:"type"`
			Number float64 `json:"number"`
		} `json:"Price"`
		Size struct {
			Id     string `json:"id"`
			Type   string `json:"type"`
			Select struct {
				Id    string `json:"id"`
				Name  string `json:"name"`
				Color string `json:"color"`
			} `json:"select"`
		} `json:"Size"`
		Ingredients struct {
			Id       string `json:"id"`
			Type     string `json:"type"`
			RichText []struct {
				Type string `json:"type"`
				Text struct {
					Content string      `json:"content"`
					Link    interface{} `json:"link"`
				} `json:"text"`
				Annotations struct {
					Bold          bool   `json:"bold"`
					Italic        bool   `json:"italic"`
					Strikethrough bool   `json:"strikethrough"`
					Underline     bool   `json:"underline"`
					Code          bool   `json:"code"`
					Color         string `json:"color"`
				} `json:"annotations"`
				PlainText string      `json:"plain_text"`
				Href      interface{} `json:"href"`
			} `json:"rich_text"`
		} `json:"Ingredients"`
		Name struct {
			Id    string `json:"id"`
			Type  string `json:"type"`
			Title []struct {
				Type string `json:"type"`
				Text struct {
					Content string      `json:"content"`
					Link    interface{} `json:"link"`
				} `json:"text"`
				Annotations struct {
					Bold          bool   `json:"bold"`
					Italic        bool   `json:"italic"`
					Strikethrough bool   `json:"strikethrough"`
					Underline     bool   `json:"underline"`
					Code          bool   `json:"code"`
					Color         string `json:"color"`
				} `json:"annotations"`
				PlainText string      `json:"plain_text"`
				Href      interface{} `json:"href"`
			} `json:"title"`
		} `json:"Name"`
	} `json:"properties"`
	Url       string      `json:"url"`
	PublicUrl interface{} `json:"public_url"`
}
