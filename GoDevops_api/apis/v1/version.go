package v1

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/yes5144/GoDevops/GoDevops_api/models"
	"github.com/yes5144/GoDevops/GoDevops_api/utils"
)

// fake data

// versions :=[

// 	["nz","wx","1.0.21,32"],
// 	["nz","hf","1.0.21.132"]

// ]

// GetAllVersions xxx
func GetAllVersions(c *gin.Context) {
	var version models.Version
	log.Println("get ids from api:---------------")
	data, _ := version.GetAll()
	utils.Success(c, gin.H{"data": data}, "all version")
}

func GetVersionsByIds(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg":  "version Ids",
		"data": gin.H{},
	})
}

func PostPackIds(c *gin.Context) {
	type apppackids struct {
		PackIds    string `json:"packids,omitempty"`
		ServerPath string `json:"serverpath,omitempty"`
		ClientPath string `json:"clientpath,omitempty"`
	}
	var postIds = apppackids{}
	e := c.BindJSON(&postIds)
	if e != nil {
		utils.Fail(c, nil, "parse data from web err")
	}
	ids := postIds.PackIds
	sPath := postIds.ServerPath
	cPath := postIds.ClientPath
	log.Println("postids: ", ids, sPath, cPath)

	var version models.Version
	log.Println("get ids from api:---------------")
	// log.Println(version.GetAll())
	log.Println("get ids from api:---------------")
	vs, err := version.GetByIds(ids)
	if err != nil {
		utils.Fail(c, nil, "err")
	}
	log.Println(vs)
	for _, v := range vs {
		log.Printf("vvv %T, vvv %v", v, v)
		tmpVer := v.Version
		tmpProj := v.Project
		tmpZone := v.Zone
		newVersionList := strings.Split(tmpVer, ".")
		lastVer, _ := strconv.Atoi(newVersionList[len(newVersionList)-1])
		nLVer := fmt.Sprintf("%d", lastVer+1)
		newVersionList[len(newVersionList)-1] = nLVer
		tmpNewVer := strings.Join(newVersionList[:], ".")
		dest := fmt.Sprintf("d:/20200624/%s/%s/%s_%s_v%s_date", tmpProj, tmpZone, tmpProj, tmpZone, tmpNewVer)
		// copy
		utils.CopyDir(sPath+"/svn_proc", dest)
		utils.CopyDir(sPath+"/wktest.sh", dest)
		zipFile := dest + ".zip"
		log.Println(tmpVer, tmpProj, tmpZone, tmpNewVer, dest, zipFile)
		// zip
		// utils.ZipDir(dest, zipFile)
		// update version

	}
	// version client server pack

	utils.Success(c, gin.H{"data": vs}, "data success")
}
