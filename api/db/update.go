package db

import (
	"fmt"
)

/* update wanted */
func UpdateWantedStatus(userid, wantedid int, alive bool) (ok bool) {
	_, err := db.Exec(
		"UPDATE `work_wanted` SET `alive` = ? WHERE `user` = ? AND `id` = ?",
		alive,
		userid,
		wantedid,
	)
	if err != nil {
		ok = false
		return
	}
	ok = true
	return
}

/* update request */
func UpdateRequestChecked(userid, wantedid, requestid int) (ok bool) {
	_, err := db.Exec(
		"UPDATE `work_request` SET `check` = true WHERE `user` = ? AND `wanted` = ? AND `id` = ?",
		userid,
		wantedid,
		requestid,
	)
	if err != nil {
		fmt.Printf("error db.UpdateRequestChecked:: %v\n", err)
		ok = false
		return
	}
	ok = true
	return
}
func UpdateRequestStatus(userid int, data StructRequest) (ok bool) {
	_, err := db.Exec(
		"UPDATE `work_request` SET `establish` = ?, `alive` = ?, `check` = true WHERE `user` = ? AND `wanted` = ? AND `id` = ?",
		data.Establish,
		data.Alive,
		userid,
		data.Wanted.Number,
		data.Number,
	)
	if err != nil {
		fmt.Printf("error db.UpdateRequestChecked:: %v\n", err)
		ok = false
		return
	}
	ok = true
	return
}

/* update monitor( evaluate ) */
func UpdateEvaluate(userid int, data StructEvaluate) (ok bool) {
	tx, err := db.Begin()
	if err != nil {
		fmt.Printf("error db.UpdateEvaluate:: %v\n", err)
		ok = false
		return
	}

	// get target id
	var targetid int
	err = tx.QueryRow(
		"SELECT `id` FROM `user` WHERE `name` = ?",
		data.TargetName,
	).Scan(&targetid)
	if err != nil {
		fmt.Printf("error db.UpdateEvaluate:: %v\n", err)
		tx.Rollback()
		ok = false
		return
	}

	// insert monitor
	_, err = tx.Exec(
		"INSERT INTO `monitor`(`observer`, `target`, `rate`, `review`) VALUES(?, ?, ?, ?) ON DUPLICATE KEY UPDATE `rate` = ?, `review` = ?",
		userid,
		targetid,
		data.Rate,
		data.Review,
		data.Rate,
		data.Review,
	)
	if err != nil {
		fmt.Printf("error db.UpdateEvaluate:: %v\n", err)
		tx.Rollback()
		ok = false
		return
	}

	tx.Commit()
	ok = true
	return
}
