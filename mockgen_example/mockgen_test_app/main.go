package main

import "my/mockgen_test_app/database"

func main() {
	database.Setup()
	user := database.NewUser(23, "utkarshmani1997@gmail.com", "Utkarsh", "Mani Tripathi")
	insertUser(user)
	database.TearDown()
}

func insertUser(user database.Users) (int, error) {
	id, err := user.Insert()
	if err != nil {
		return 0, err
	}
	return id, nil
}
