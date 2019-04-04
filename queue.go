package main

import (
	"errors"
)

func getQueue(q *Queue) *Queue {
	if q.Links == nil {
		q.Links = []string{}
	}
	return q
}

func (q *Queue) enqueue(links []string) {
	q.lock.Lock()
	q.Links = append(q.Links, links...)
	q.lock.Unlock()
}

func (q *Queue) getLinks(c int) ([]string, error) {
	if q.Links == nil || len(q.Links) == 0 {
		return nil, errors.New("queue is empty or is not initialized")
	}

	q.lock.Lock()

	if c > len(q.Links) {
		c = len(q.Links)
	}

	links := q.Links[:c]
	q.Links = q.Links[c:len(q.Links)]

	q.lock.Unlock()

	return links, nil
}
