package stackOverflow

type User struct{
	userId string
	reputation int
}

func(u User) GetUserId()(string){
	return u.userId
}

func (u User) GetReputation()(int){
	return u.reputation
}