package handler

import (
	"fmt"
	"github.com/19700101000000/system-sample/api/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

/* update request */
func UpdateRequestChecked(c *gin.Context) {
	// check auth
	name, ok := getAuth(c)
	if !ok {
		c.String(http.StatusForbidden, "forbidden")
		return
	}

	// get data
	var reqData struct {
		WantedID  int `json:"wanted"`
		RequestID int `json:"request"`
	}
	err := c.Bind(&reqData)
	if err != nil {
		fmt.Printf("error handler.UpdateRequestChecked:: %v\n", err)
		c.String(http.StatusBadRequest, "bad request")
		return
	}

	// query
	ok = db.UpdateRequestChecked(
		UserList[name].ID,
		reqData.WantedID,
		reqData.RequestID,
	)
	if !ok {
		c.String(http.StatusBadRequest, "bad request")
		return
	}
	c.String(http.StatusOK, "ok")
}
func UpdateRequestStatus(c *gin.Context) {
	// check auth
	name, ok := getAuth(c)
	if !ok {
		c.String(http.StatusForbidden, "forbidden")
		return
	}

	// get data
	var reqData struct {
		WantedID  int  `json:"wanted"`
		RequestID int  `json:"request"`
		Establish bool `json:"establish"`
		Alive     bool `json:"alive"`
	}
	err := c.Bind(&reqData)
	if err != nil {
		fmt.Printf("error handler.UpdateRequestStatus:: %v\n", err)
		c.String(http.StatusBadRequest, "bad request")
		return
	}

	// query
	ok = db.UpdateRequestStatus(
		UserList[name].ID,
		db.StructRequest{
			Wanted: db.StructWanted{
				Number: reqData.WantedID,
			},
			Number:    reqData.RequestID,
			Establish: reqData.Establish,
			Alive:     reqData.Alive,
		},
	)
	if !ok {
		c.String(http.StatusBadRequest, "bad request")
		return
	}
	c.String(http.StatusOK, "ok")
}
