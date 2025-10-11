package stackOverflow

import (
	"fmt"
	"sync"
	"time"
)

var (
	instance *StackOverFlow
	mu       sync.Mutex
)

type StackOverFlow struct {
	questions []*Question
	answers   []*Answer
	// comment []*Comment
	users []*User
	// tag we can add tag but that is many to many relation with the question
	// question , tag, tagquestion make three entity iterate inb tagquestion entity
}

func (s *StackOverFlow) GetUser(userId string) *User {
	for _, v := range s.users {
		if v.userId == userId {
			return v
		}
	}
	return nil
}

func (s *StackOverFlow) GetQuestion(postId string) *Question {
	for _, v := range s.questions {
		if v.postId == postId {
			return v
		}
	}
	return nil
}

func (s *StackOverFlow) GetAnswer(postId string) *Answer {
	for _, v := range s.answers {
		if v.postId == postId {
			return v
		}
	}
	return nil
}

func (s *StackOverFlow) AddQuestion(userId string, content string) {
	question := &Question{
		VotablePost: &VotablePost{
			CommentablePost: &CommentablePost{
				Post: &Post{
					postId:  userId + time.Now().Format(time.RFC3339),
					content: content,
					author:  s.GetUser(userId),
				},
			},
			votes: make(map[string]int),
		},
	}
	s.questions = append(s.questions, question)
}

func (s *StackOverFlow) AddAnswer(userId string, questionId string, content string) {
	if s.GetQuestion(questionId) != nil {
		answer := &Answer{
			VotablePost: &VotablePost{
				CommentablePost: &CommentablePost{
					Post: &Post{
						postId:  userId + time.Now().Format(time.RFC3339),
						content: content,
						author:  s.GetUser(userId),
					},
				},
				votes: make(map[string]int),
			},
			questionId: questionId,
		}
		s.answers = append(s.answers, answer)
		return
	}
	fmt.Println("No Question exist")

}

func GetInstance() (*StackOverFlow, error) {
	mu.Lock()
	defer mu.Unlock()
	if instance == nil {
		instance = &StackOverFlow{}
	}
	return instance, nil
}
