package stackOverflow

type VoteType int

const (
	UpVote VoteType = iota
	DownVote
)

type Post struct {
	postId  string
	content string
	author  *User
}

type Comment struct {
	*Post
}

type CommentablePost struct {
	*Post
	comments []*Comment
}

func (cp *CommentablePost) AddCommentToPost(comment *Comment) {
	cp.comments = append(cp.comments, comment)
}

type VotablePost struct {
	*CommentablePost
	votes map[string]int
	// observer []  we can user observer patter to update the reputation for the user
}

func (vp *VotablePost) AddObserver() {
	// to do
}

func (vp *VotablePost) AddVotes(user *User, vote int) {
	if user == nil {
		return
	}
	if vote != 1 && vote != -1 { // enforce only up/down
		return
	}
	value, exists := vp.votes[user.userId]
	change := 0

	if exists {
		if value != vote { // switching vote
			if value == -1 { // -1 -> +1
				change = 2
			} else { // +1 -> -1
				change = -2
			}
		} else { // same vote again: no change
			change = 0
		}
	} else { // first time vote
		change = vote
	}

	if change != 0 {
		vp.author.reputation += change // could use observer pattern
	}
	// persist / update the vote value (even if same) for future switches
	vp.votes[user.userId] = vote
}

type Answer struct {
	*VotablePost
	questionId string
}

type Question struct {
	*VotablePost
}
