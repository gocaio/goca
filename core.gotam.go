/*
	Copyright © 2019 The Goca.io team

	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

		http://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/

package main

import (
	"strings"
)

// anal.gotam.go is the internal messaging queue userd for task dispatching.
// The GoTaM (Goca Task Manager) can handle multiple "topics" and multiple
// subscribers.

// Gotam is the main class for the MQ
type Gotam struct {
	topics []string            // List of topics for internal use
	Q      map[string][][]byte // The Task queue
}

// NewGotam initializes a new MQ
func NewGotam() *Gotam {
	return &Gotam{
		topics: []string{},
		Q:      make(map[string][][]byte),
	}
}

// // Subscribe to a new topic
// func (mq *Gotam) Subscribe(topic string) (ok bool) { return ok }

// // Allocate marks an element as in use
// func (mq *Gotam) Allocate(topic string) (ok bool) { return ok }

// Push sends data to gotam
func (mq *Gotam) Push(topic string, data []byte) {
	mq.addTopic(sanitizeMime(topic))
	mq.Q[topic] = append(mq.Q[topic], data)
}

// Get returns the first element on gotam without deleting the element
func (mq *Gotam) Get(topic string) (data []byte) {
	if mq.QLen(topic) < 1 {
		return data
	}
	return mq.Q[topic][0]
}

// Pop returns the first element on gotam and deletes it from the queue
func (mq *Gotam) Pop(topic string) (data []byte) {
	if mq.QLen(topic) < 1 {
		return data
	}
	data, mq.Q[topic] = mq.Q[topic][0], mq.Q[topic][1:]
	return data
}

// QLen returns the number of elements available in the given topic
func (mq *Gotam) QLen(topic string) (length int) { return len(mq.Q[topic]) }

// Len returns the number of registered topics
func (mq *Gotam) Len() (length int) { return len(mq.topics) }

// Mimes returns the mimetypes
func (mq *Gotam) Mimes() (mimes []string) { return mq.topics }

// ===========
// = Helpers =
// ===========

func (mq *Gotam) addTopic(topic string) {
	for _, ts := range mq.topics {
		if ts == topic {
			return
		}
	}

	mq.topics = append(mq.topics, topic)
}

func sanitizeMime(mime string) (sMime string) {
	mS := strings.Split(mime, ";")
	if len(mS) > 0 {
		sMime = mS[0]
	}

	return sMime
}
