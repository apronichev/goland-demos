package model

import "errors"

// todo paste TodoItem from json

type TodoItem struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
	Due   string `json:"due"`
}

func (i *TodoItem) Identity() int {
	return i.Id
}

func (i *TodoItem) Validate() error {
	if i.Title == "" {
		return errors.New("title is required")
	}
	if i.Due == "" {
		return errors.New("due is required")
	}
	return nil
}