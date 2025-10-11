package stackOverflow

type VoteType int

const (
	UpVote VoteType = iota
	DownVote
)

type Post struct{
	postId string
	content string
	author *User 
}

type Comment struct {
	*Post
}

type CommentablePost struct{
	*Post
	comments []*Comment
}

func (cp *CommentablePost) AddCommentToPost(comment *Comment) {
	cp.comments = append(cp.comments, comment)
}

type VotablePost struct{
	*CommentablePost
	votes map[string]int
	// observer []  we can user observer patter to update the reputation for the user
}

func (vp *VotablePost) AddObserver(){
	
}

func (vp *VotablePost) AddVotes(user *User, vote int ){
	// user is not nil
	value , exists := vp.votes[user.userId]
	change := 0

	if exists {
		if value != vote {
			if value == -1 {
				change = 2
			} else {
				change = -2
			}
		}
	} else {
		change = vote
	}

	vp.author.reputation += change // this can be iplementated using the observer patter 
}

type Answer struct {
	*VotablePost
	questionId string
}

type Question struct{
	*VotablePost
}
