package callInHandle

import (
	log "go-commweb/log"
	"go-commweb/mysql"
)

type callInLimit struct {
	limit int
}

func getCallInLimit (compId string) (limit int) {
	sql := "SELECT max_callin_cnt FROM pbx.tx_company where compid=?"
	db := mysql.GetMysqlDb().DB

	row := db.QueryRow(sql, compId)
	callInLimit := new(callInLimit)
	if err := row.Scan(&callInLimit.limit); err != nil {
		log.LOGGER.Error("scan failed, err:%v", err)
		return
	}
	limit = callInLimit.limit
	return
}