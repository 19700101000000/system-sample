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

	rows, err := db.Query("SELECT `u`.`name` AS `user`, `g`.`id` AS `number`, `s`.`rate` AS `rate`, `c`.`param` AS `param`, `g`.`image` AS `image`, `g`.`timestamp` AS `datetime` FROM `gallery` `g` INNER JOIN `user` `u` ON `g`.`user` = `u`.`id` LEFT OUTER JOIN (SELECT `target`, COUNT(*) AS `param` FROM `monitor` GROUP BY `target`) `c` ON `u`.`id` = `c`.`target` LEFT OUTER JOIN (SELECT `target`, SUM(`rate`) AS `rate` FROM `monitor` GROUP BY `target`) `s` ON `u`.`id` = `s`.`target` ORDER BY `datetime` DESC")
	if err != nil {
		fmt.Printf("error by db.Galleries:: %v\n", err)
		return
	}

	galleries := make([]Gallery, 0)
	for rows.Next() {
		gallery := Gallery{}
		if err := rows.Scan(&gallery.Username, &gallery.Number, &gallery.EvalSum, &gallery.EvalParam, &gallery.Imagepath, &gallery.Datetime); err != nil {
			fmt.Printf("error by db.Galleries:: %v\n", err)
			continue
		}
		galleries = append(galleries, gallery)
	}
	result["galleries"] = galleries
	return
}
