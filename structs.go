package main

import "sync"

//DownloadRequest holds request data and also holds validation data
type DownloadRequest struct {
	Links []string `form:"links" json:"links" binding:"required,dive,url"`
}

//Queue holds queue data through application lifetime
type Queue struct {
	Links []string
	lock  sync.Mutex
}
