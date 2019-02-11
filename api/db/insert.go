package db

import (
	"fmt"
	"strconv"
)

/* insert user */
func InsertUser(name, password string) (id int, ok bool) {
	tx, err := db.Begin()
	if err != nil {
		fmt.Printf("error db.InsertUser:: %v\n", err)
		ok = false
		return
	}
	_, err = tx.Exec(
		"INSERT INTO `user`(`name`, `password`) VALUES (?, ?)",
		name,
		password,
	)
	if err != nil {
		fmt.Printf("error db.InsertUser:: %v\n", err)
		ok = false
		tx.Rollback()
		return
	}
	err = tx.QueryRow("SELECT `id` FROM `user` WHERE `name` = ?", name).Scan(&id)
	if err != nil {
		fmt.Printf("error db.InsertUser:: %v\n", err)
		ok = false
		tx.Rollback()
		return
	}
	ok = true
	tx.Commit()
	return
}

/* insert gallery */
func InsertGallery(userid int, filename string, categories []string) (ok bool) {
	tx, err := db.Begin()
	if err != nil {
		ok = false
		fmt.Printf("error db.InsertGallery:: %v\n", err)
		return
	}
	_, err = tx.Exec("INSERT INTO `gallery`(`user`, `image`) VALUES (?, ?)", userid, filename)
	if err != nil {
		ok = false
		fmt.Printf("error db.InsertGallery:: %v\n", err)
		tx.Rollback()
		return
	}
	var id int
	err = tx.QueryRow("SELECT COUNT(*) AS `id` FROM `gallery` WHERE `user` = ?", userid).Scan(&id)
	if err != nil {
		ok = false
		fmt.Printf("error db.InsertGallery:: %v\n", err)
		tx.Rollback()
		return
	}
	fmt.Println("insert id", id)
	for _, category := range categories {
		categoryid, err := strconv.Atoi(category)
		if err != nil {
			fmt.Printf("error db.InsertGallery:: %v\n", err)
			continue
		}

		_, err = tx.Exec("INSERT INTO `gallery_join_category`(`creator`, `gallery`, `category`) VALUES (?, ?, ?)", userid, id, categoryid)
		if err != nil {
			ok = false
			fmt.Printf("error db.InsertGallery:: %v\n", err)
			tx.Rollback()
			return
		}
	}

	tx.Commit()
	ok = true
	return
}

/* insert wanted */
func InsertWanted(userid int, data StructWanted) (ok bool) {
	tx, err := db.Begin()
	if err != nil {
		ok = false
		fmt.Printf("error db.InsertWanted:: %v\n", err)
		return
	}
	_, err = tx.Exec(
		"INSERT INTO `work_wanted`(`user`, `title`, `description`, `price`) VALUES (?, ?, ?, ?)",
		userid,
		data.Title,
		data.Description,
		data.Price,
	)
	if err != nil {
		ok = false
		fmt.Printf("error db.InsertWanted:: %v\n", err)
		tx.Rollback()
		return
	}
	tx.Commit()
	ok = true
	return
}

/* insert request */
func InsertRequest(userid int, data StructRequest) (ok bool) {
	tx, err := db.Begin()
	if err != nil {
		ok = false
		fmt.Printf("error db.InsertRequest:: %v\n", err)
		return
	}
	var ownerID int
	err = tx.QueryRow(
		"SELECT `w`.`user` AS `user` FROM `work_wanted` `w` INNER JOIN `user` `u` ON `u`.`id` = `w`.`user` WHERE `u`.`name` = ? AND `w`.`id` = ?",
		data.Wanted.Username,
		data.Wanted.Number,
	).Scan(&ownerID)
	if err != nil {
		ok = false
		fmt.Printf("error db.InsertRequest:: %v\n", err)
		tx.Rollback()
		return
	}
	_, err = tx.Exec(
		"INSERT INTO `work_request`(`user`, `wanted`, `requester`, `title`, `description`, `price`) VALUES (?, ?, ?, ?, ?, ?)",
		ownerID,
		data.Wanted.Number,
		userid,
		data.Title,
		data.Description,
		data.Price,
	)
	if err != nil {
		ok = false
		fmt.Printf("error db.InsertRequest:: %v\n", err)
		tx.Rollback()
		return
	}
	tx.Commit()
	ok = true
	return
}
