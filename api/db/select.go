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

func User(name string) (result map[string]interface{}) {
	result = make(map[string]interface{})
	user := StructUser{}

	err := db.QueryRow(
		"SELECT `name`, `show_name`, `alive` FROM `user` WHERE `name` = ?",
		name,
	).Scan(&user.Name, &user.ShowName, &user.Alive)
	if err != nil {
		fmt.Printf("error by db.Auth:: %v\n", err)
	}
	result["user"] = user
	return
}

func WorksWantedlies(name string) (result map[string]interface{}) {
	result = make(map[string]interface{})
	wantedlies := make([]StructWanted, 0)

	rows, err := db.Query(
		"SELECT `u`.`name` AS `name`, `w`.`id` AS `id`, `w`.`title` AS `title`, `w`.`description` AS `description`, `w`.`price` AS `price` FROM `work_wanted` `w` INNER JOIN `user` `u` ON `u`.`id` = `w`.`user` WHERE `u`.`name` = ?",
		name,
	)
	if err != nil {
		fmt.Printf("error by db.WorksWantedlies:: %v\n", err)
		return
	}

	for rows.Next() {
		wanted := StructWanted{}
		if err := rows.Scan(
			&wanted.Username,
			&wanted.Number,
			&wanted.Title,
			&wanted.Description,
			&wanted.Price,
		); err != nil {
			fmt.Printf("error by db.WorksWantedlies:: %v\n", err)
			continue
		}
		wantedlies = append(wantedlies, wanted)
	}
	result["wantedlies"] = wantedlies
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

	rows, err := db.Query("SELECT `u`.`name` AS `user`, `g`.`id` AS `number`, `r`.`rate` AS `rate`, `p`.`param` AS `param`, `f`.`favo` AS `favo`, `c`.`comm` AS `comm`, `g`.`image` AS `image`, `g`.`timestamp` AS `datetime` FROM `gallery` `g` INNER JOIN `user` `u` ON `g`.`user` = `u`.`id` LEFT OUTER JOIN (SELECT `creator`, `gallery`, COUNT(*) AS `favo` FROM `gallery_favorite` GROUP BY `creator`, `gallery`) `f` ON `f`.`creator` = `u`.`id` AND `f`.`gallery` = `g`.`id` LEFT OUTER JOIN (SELECT `creator`, `gallery`, COUNT(*) AS `comm` FROM `gallery_comment` GROUP BY `creator`, `gallery`) `c` ON `c`.`creator` = `u`.`id`AND `c`.`gallery` = `g`.`id` LEFT OUTER JOIN (SELECT `target`, COUNT(*) AS `param` FROM `monitor` GROUP BY `target`) `p` ON `u`.`id` = `p`.`target` LEFT OUTER JOIN (SELECT `target`, SUM(`rate`) AS `rate` FROM `monitor` GROUP BY `target`) `r` ON `u`.`id` = `r`.`target` ORDER BY `datetime` DESC")
	if err != nil {
		fmt.Printf("error by db.Galleries:: %v\n", err)
		return
	}

	galleries := make([]Gallery, 0)
	for rows.Next() {
		gallery := Gallery{}
		if err := rows.Scan(
			&gallery.Username,
			&gallery.Number,
			&gallery.EvalSum,
			&gallery.EvalParam,
			&gallery.Favorite,
			&gallery.Comment,
			&gallery.Imagepath,
			&gallery.Datetime,
		); err != nil {
			fmt.Printf("error by db.Galleries:: %v\n", err)
			continue
		}
		gallery.Imagepath = "/api/images/" + gallery.Imagepath
		if gallery.EvalSum == nil {
			i := 0
			gallery.EvalSum = &i
		}
		if gallery.Favorite == nil {
			i := 0
			gallery.Favorite = &i
		}
		if gallery.Comment == nil {
			i := 0
			gallery.Comment = &i
		}
		galleries = append(galleries, gallery)
	}
	result["galleries"] = galleries
	return
}
