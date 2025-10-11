package stackOverflow

import "fmt"

// RunDemo builds a small scenario with users, questions, answers, votes and prints reputations.
func RunDemo() {
	so, _ := GetInstance()

	// create users
	u1 := &User{userId: "u1"}
	u2 := &User{userId: "u2"}
	u3 := &User{userId: "u3"}
	so.AddUSer(u1)
	so.AddUSer(u2)
	so.AddUSer(u3)

	// questions
	so.AddQuestion("u1", "How to write a goroutine?", "q1")
	so.AddQuestion("u2", "Difference between slice and array?", "q2")

	// answers
	so.AddAnswer("u2", "q1", "Use go keyword before the function call.", "a1")
	so.AddAnswer("u3", "q1", "Launch with go and sync using WaitGroup.", "a2")
	so.AddAnswer("u1", "q2", "Arrays have fixed length, slices are dynamic views.", "a3")

	// votes (u2 upvotes q1, u3 upvotes q1, u3 downvotes q2, u1 upvotes a1, u1 upvotes a2 then switch to downvote, u2 upvotes a3)
	so.AddVotes("u2", "q1", 1)  // +1 to u1
	so.AddVotes("u3", "q1", 1)  // +1 to u1 (u1 rep = 2)
	so.AddVotes("u3", "q2", -1) // -1 to u2 (u2 rep = -1)
	so.AddVotes("u1", "a1", 1)  // +1 to u2 (u2 rep = 0)
	so.AddVotes("u1", "a2", 1)  // +1 to u3 (u3 rep = 1)
	so.AddVotes("u1", "a2", -1) // switch vote: a2 author u3 gets -2 (1 -> -1)
	so.AddVotes("u2", "a3", 1)  // +1 to u1 (u1 rep = 3)

	fmt.Println("User reputations:")
	fmt.Printf("u1: %d\n", u1.reputation) // expect 3
	fmt.Printf("u2: %d\n", u2.reputation) // expect 0
	fmt.Printf("u3: %d\n", u3.reputation) // expect -1
}
