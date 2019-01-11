package db

import (
	"crypto/sha256"
	"fmt"
)

func Auth(user, pass string) (id int, name string, ok bool) {
	err := db.QueryRow(
		"SELECT `id`, `name` FROM `user` WHERE `name` = ? AND `password` = ? AND alive = 1",
		user,
		fmt.Sprintf("%x", sha256.Sum256([]byte(pass))),
	).Scan(&id, &name)
	if err != nil {
		fmt.Printf("error by db.Auth:: %v\n", err)
		ok = false
		return
	}
	ok = true
	return
}

func Categories() (result map[string]interface{}) {
	result = make(map[string]interface{})

	rows, err := db.Query("SELECT `id`, `name` FROM `category`")
	if err != nil {
		fmt.Printf("error by db.Categories:: %v\n", err)
		return
	}

	options := make([]OptionCategory, 0)
	for rows.Next() {
		var id, name string
		if err := rows.Scan(&id, &name); err != nil {
			fmt.Printf("error by db.Categories:: %v\n", err)
			continue
		}
		options = append(options, OptionCategory{Text: name, Value: id})
	}
	result["options"] = options
	return
}

func Galleries() (result map[string]interface{}) {
	result = make(map[string]interface{})

	rows, err := db.Query("SELECT `u`.`name` AS `user`, `g`.`image` AS `image`, `g`.`timestamp` AS `datetime` FROM `gallery` `g` INNER JOIN `user` `u` ON `g`.`user` = `u`.`id` ORDER BY `datetime` DESC")
	if err != nil {
		fmt.Printf("error by db.Galleries:: %v\n", err)
		return
	}

	galleries := make([]Gallery, 0)
	for rows.Next() {
		gallery := Gallery{}
		if err := rows.Scan(&gallery.Username, &gallery.Imagepath, &gallery.Datetime); err != nil {
			fmt.Printf("error by db.Galleries:: %v\n", err)
			continue
		}
		galleries = append(galleries, gallery)
	}
	result["galleries"] = galleries
	return
}
