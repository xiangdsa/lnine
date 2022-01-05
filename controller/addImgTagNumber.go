package controller

import "updateLnineDownloadMorepic/dao"

//computer
func AddComputerImgTagNumber() {
	maxID := dao.QueryComputerMaxID()
	for id := int64(0); id < maxID; id += 300 {
		bzSeoEcmsDownloads:=dao.QueryComputerMessage(id)
		for i,_:=range bzSeoEcmsDownloads{
			newContent:=dao.ComputerImgTagNumber(*bzSeoEcmsDownloads[i])
			dao.UpdateComputerSoftsay(newContent,bzSeoEcmsDownloads[i].ID)
		}
	}
}


//Az
func AddAzImgTagNumber(){
	maxID := dao.QueryAZMaxID()
	for id := int64(67694); id < maxID; id += 300 {
		bzSeoEcmsAzData1:=dao.QueryAZMessage(id)
		for i,_:=range bzSeoEcmsAzData1{
			newstext:=dao.AZImgTagNumber(*bzSeoEcmsAzData1[i])
			dao.UpdateAzNewsText(newstext,bzSeoEcmsAzData1[i].ID)
		}
	}
}

