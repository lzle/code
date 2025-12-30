package callInHandle

import (
	"sync"
	"time"
)

var	(
	compIdToCompany = make(map[string]*company)
	compMutex  sync.RWMutex
)

type company struct {
	compId       string
	callContain  []string
	maxLimit     int
}

func (company *company) UpdateLimit() {
	limit := getCallInLimit(company.compId)
	if limit != 0 {
		company.maxLimit = limit
	}
}

func UpdateCompany () {
	for {
		for _, company := range compIdToCompany {
			company.UpdateLimit()
		}
		time.Sleep(time.Second * 60)
	}
}

func (company *company) canCall() (res bool) {
	if len(company.callContain) < company.maxLimit {
		return true
	}
	return false
}


func getCompany (compId string) *company {
	compMutex.Lock()
	defer compMutex.Unlock()

	if company, ok := compIdToCompany[compId]; ok {
		return company
	}
	return nil
}

func setCompany (company *company)  {
	compMutex.Lock()
	defer compMutex.Unlock()
	compIdToCompany[company.compId] = company
}

func newCompany (compId string) *company {
	company := new(company)
	company.compId = compId
	company.UpdateLimit()
	setCompany(company)
	return company
}




