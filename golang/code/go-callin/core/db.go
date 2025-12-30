package core

import "database/sql"

type Mysql interface {
	GetDB() *sql.DB
}

type Router interface {
	GetCompId(dh string) (compId string)

	GetRouteType(trunkIp string, trunkPort string) (routeType int)

	AccessMode(compId string, callee string) (mode string, defined string)

	BlackLists(blackId string) (lists [][2]string)

	MobileArea(prefix string, area string) bool

	BlackProcess(blackId string, compId string)(bMode string, bDefined string)

	AccessModeNext(compId string, callee string) (mode string, defined string)
}