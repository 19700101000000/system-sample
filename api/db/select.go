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

/* works requests */
func WorksRequests(name string, auth bool) (result map[string]interface{}) {
	result = make(map[string]interface{})
	requests := make([]StructRequest, 0)

	sql := "SELECT `u1`.`name` AS `owner`, `r`.`wanted` AS `wanted_id`, `r`.`id` AS `request_id`, `u2`.`name` AS `requester`, `r`.`title` AS `title`, `r`.`description` AS `description`, `r`.`price` AS `price`, `r`.`establish` AS `establish`, `r`.`alive` AS `alive` FROM `work_request` `r` INNER JOIN `user` `u1` ON `r`.`user` = `u1`.`id` INNER JOIN `user` `u2` ON `r`.`requester` =	`u2`.`id` WHERE `u2`.`name` = ?"
	if !auth {
		sql += " AND `r`.`alive` = true"
	}
	sql += " ORDER BY `r`.`create_at` DESC, `r`.`alive` DESC, `r`.`establish` ASC"
	rows, err := db.Query(sql, name)
	if err != nil {
		fmt.Printf("error by db.WorksReqests:: %v\n", err)
		return
	}

	for rows.Next() {
		request := StructRequest{}
		if err := rows.Scan(
			&request.OwnerName,
			&request.WantedID,
			&request.Number,
			&request.UserName,
			&request.Title,
			&request.Description,
			&request.Price,
			&request.Establish,
			&request.Alive,
		); err != nil {
			fmt.Printf("error by db.WorksReqests:: %v\n", err)
			continue
		}
		requests = append(requests, request)
	}
	result["requests"] = requests
	return
}

/* works wanteds */
func WorksWanteds(name string, auth bool) (result map[string]interface{}) {
	result = make(map[string]interface{})
	wanteds := make([]StructWanted, 0)

	sql := "SELECT `u`.`name` AS `name`, `w`.`id` AS `id`, `w`.`title` AS `title`, `w`.`description` AS `description`, `w`.`price` AS `price`, `w`.`alive` AS `alive` FROM `work_wanted` `w` INNER JOIN `user` `u` ON `u`.`id` = `w`.`user` WHERE `u`.`name` = ?"
	if !auth {
		sql += " AND  `w`.`alive` = true"
	}
	sql += " ORDER BY `w`.`create_at` DESC, `w`.`alive` DESC"
	rows, err := db.Query(
		sql,
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
			&wanted.Alive,
		); err != nil {
			fmt.Printf("error by db.WorksWantedlies:: %v\n", err)
			continue
		}
		wanteds = append(wanteds, wanted)
	}
	result["wanteds"] = wanteds
	return
}

/* categories */
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
