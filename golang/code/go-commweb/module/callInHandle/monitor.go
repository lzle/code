package callInHandle

import (
	log "go-commweb/log"
	"time"
)

func Monitor() {
	for {
		for _, company := range compIdToCompany {
			log.LOGGER.Info("monintor compid[%s] limit[%d] calls[%d]", company.compId, company.maxLimit, len(company.callContain))
		}
		time.Sleep(time.Second* 60)
	}
}
