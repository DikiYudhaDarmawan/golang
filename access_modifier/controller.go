package controller

type user struct {
	ID int `json:"id"`
	Name string `json:"Name"`
	Email string `json:"email"`
}

func ReadUserList() (users []user){
	users = make([]user, 2)

	users[0] = user{
		ID:  1,
		Name: "Diki Darmawan",
		Email: "dikdarmwn@gmail.com",
	}

	users[1] = user{
		ID:    2,
		Name:  "Darmawan Diki",
		Email:  "darmwndik@gmail.com",
	}

	return
}