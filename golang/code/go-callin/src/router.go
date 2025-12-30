package src

import (
	"database/sql"
	"go-callin/core"
	"strconv"
)

type Router struct {
	db *sql.DB
}

func (r *Router) Init() {
	r.db = core.MYSQL.GetDB()
	core.ROUTER = r
}

// 根据接续号码获取企业
func (r *Router) GetCompId(dh string) (compId string) {
	sql := "select compid from pbx.tx_access_num where number = ? and enable = 1 limit 1"
	row := r.db.QueryRow(sql, dh)
	err := row.Scan(&compId)
	if err != nil {
		core.LOGGER.Error("scan failed, %v", err)
		return
	}
	return
}

// 获取号码类型  0-2 坐席、分机  3 手机
func (r *Router) GetRouteType(trunkIp string, trunkPort string) (routeType int) {
	sql := "select type from pbx.tx_server where ip = ? and port = ?"
	row := r.db.QueryRow(sql, trunkIp, trunkPort)
	err := row.Scan(&routeType)
	if err != nil {
		core.LOGGER.Error("scan failed, %v", err)
		return
	}
	return
}

func (r *Router) AccessMode(compId string, callee string) (mode string, defined string) {
	sql := "select process_mode,process_defined from pbx.tx_access_num where compid = ? and number = ?"
	row := r.db.QueryRow(sql, compId, callee)
	err := row.Scan(&mode, &defined)
	if err != nil {
		core.LOGGER.Error("scan failed, %v", err)
		return
	}
	return
}

func (r *Router) AccessModeNext(compId string, callee string) (mode string, defined string) {
	sql := "select process_mode2,process_defined2 from pbx.tx_access_num where compid = ? and number = ?"
	row := r.db.QueryRow(sql, compId, callee)
	err := row.Scan(&mode, &defined)
	if err != nil {
		core.LOGGER.Error("scan failed, %v", err)
		return
	}
	return
}

func (r *Router) BlackLists(blackId string) (lists [][2]string) {
	sql := "select type,value from pbx.tx_route_black_list where blackid = ?"
	rows, err := r.db.Query(sql, blackId)
	if err != nil {
		core.LOGGER.Error("scan failed, %v", err)
		return
	}
	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()

	for rows.Next() {
		bType,bValue := 0,""
		err := rows.Scan(&bType, &bValue)
		if err != nil {
			core.LOGGER.Info("scan failed, %v", err)
		}
		lists = append(lists, [2]string{strconv.Itoa(bType),bValue})
	}
	return
}

func (r *Router) MobileArea(prefix string, area string) bool {
	var count int
	sql := "select count(1) from tx_res_mobile_phonearea where prefix = ？and areacode = ? limit 1"
	row := r.db.QueryRow(sql, prefix, area)
	err := row.Scan(&count)
	if err != nil {
		core.LOGGER.Error("scan failed, %v", err)
	}
	if count > 0 {
		return true
	}
	return false
}

func (r *Router) BlackProcess(blackId string, compId string)(bMode string, bDefined string) {
	sql := "select process_mode,process_defined from tx_route_black where black_identify=? and compid=?"
	row := r.db.QueryRow(sql, blackId, compId)
	err := row.Scan(&bMode, &bDefined)
	if err != nil {
		core.LOGGER.Error("scan failed, %v", err)
		return
	}
	return
}
