package service

import "sync"

var infoService *InfoService
var infoServiceOnce sync.Once

type InfoService struct {
	ApplicantService
}

func GetInfoService() *InfoService {
	infoServiceOnce.Do(func() {
		infoService = &InfoService{}
		applicantService := new(ApplicantServiceImpl)

		infoService.ApplicantService = applicantService
	})
	return infoService
}
