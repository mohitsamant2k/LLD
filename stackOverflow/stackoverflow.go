package stackOverflow

import (
	"fmt"
	"sync"
)

var (
	instance *StackOverFlow
	once     sync.Once
)

type StackOverFlow struct {
	questions []*Question
	answers   []*Answer
	// comment []*Comment
	users []*User
	// tag we can add tag but that is many to many relation with the question
	// question , tag, tagquestion make three entity iterate inb tagquestion entity
}

func (s *StackOverFlow) AddUSer(user *User) {
	s.users = append(s.users, user)
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

func (s *StackOverFlow) AddQuestion(userId string, content string, id string) {
	question := &Question{
		VotablePost: &VotablePost{
			CommentablePost: &CommentablePost{
				Post: &Post{
					postId:  id,
					content: content,
					author:  s.GetUser(userId),
				},
			},
			votes: make(map[string]int),
		},
	}
	s.questions = append(s.questions, question)
}

func (s *StackOverFlow) AddAnswer(userId string, questionId string, content string, id string) {
	if s.GetQuestion(questionId) != nil {
		answer := &Answer{
			VotablePost: &VotablePost{
				CommentablePost: &CommentablePost{
					Post: &Post{
						postId:  id,
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

func (s *StackOverFlow) Getpost(postId string) *VotablePost {
	for _, v := range s.questions {
		if v.postId == postId {
			return v.VotablePost
		}
	}

	for _, v := range s.answers {
		if v.postId == postId {
			return v.VotablePost
		}
	}
	return nil
}

func (s *StackOverFlow) AddVotes(userId string, postId string, votes int) {
	v := s.Getpost(postId)
	if v != nil {
		v.AddVotes(s.GetUser(userId), votes)
		return
	}
	fmt.Println("no post found")
}

func GetInstance() (*StackOverFlow, error) {
	once.Do(func() {
		instance = &StackOverFlow{}
	})
	return instance, nil
}
