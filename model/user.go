package model

import (
	"fmt"
)

type User struct {
	id			string `db:"id"`
	username 	string `db:"username"`
	password	string `db:"password"`
	name		string `db:"name"`
	email		string `db:"email"`
	telp		string `db:"telp"`
	nik			string `db:"nik"`
	iImage 		string `db:"identity_image"`
	ipImage		string `db:"identity_person_image"`
	ixImage		string `db:"identity_extra_image"`
	googleId	string `db:"google_id"`
	facebookId	string `db:"facebook_id"`
	status		string `db:"status"`
}

func login(id string) User {
	db := InitDb()

	rows, err := db.Query("SELECT * FROM user where id = ?", id)
	if err != nil {
        fmt.Println(err.Error())
	}

	result := User{}
	for rows.Next() {
		var uid string
		var username string
		var name string
		var email string
		var telp string
		var nik string
		var iImage string
		var ipImage string
		var ixImage string
		var googleId string
		var facebookId string
		var status string

		err = rows.Scan(&uid, &username, &password, &name, &email, &telp, &nik, &iImage, &ipImage, &ixImage, &googleId, &facebookId, &status)
		if err != nil {
			fmt.Println(err.Error())
		}
		result.id = &uid 
		result.username = &username
		result.name = &name
		result.email = &email
		result.telp = &telp
		result.nik = &nik
		result.iImage = &iImage
		result.ipImage = &ipImage
		result.ixImage = &ixImage
		result.googleId = &googleId
		result.facebookId = &facebookId 
		result.status = &status
	}

	fmt.Println("Login")
	return result
}