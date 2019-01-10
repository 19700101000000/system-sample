package db

import (
	"fmt"
	"strconv"
)

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
