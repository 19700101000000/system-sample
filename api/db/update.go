package db

import (
	"fmt"
)

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
		"UPDATE `work_request` SET `establish` = ?, `alive` = ? WHERE `user` = ? AND `wanted` = ? AND `id` = ?",
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
