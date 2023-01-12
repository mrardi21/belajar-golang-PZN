package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (userSlice UserSlice) Len() int {
	return len(userSlice)
}

func (userSlice UserSlice) Less(i, j int) bool {
	return userSlice[i].Age < userSlice[j].Age
}

func (userSlice UserSlice) Swap(i, j int) {
	userSlice[i], userSlice[j] = userSlice[j], userSlice[i]
}

func main() {
	users := []User{
		{"Ardi1", 26},
		{"Ardi2", 25},
		{"Ardi3", 22},
		{"Ardi4", 23},
		{"Ardi5", 24},
		{"Ardi6", 21},
	}

	sort.Sort(UserSlice(users))
	fmt.Println(users)

	sort.Slice(users, func(i, j int) bool {
		return users[i].Age > users[j].Age
	})
	fmt.Println(users)
}
