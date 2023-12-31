package jobhireRepositories

import (
	"database/sql"
	"net/http"

	"codeid.revampacademy/models"
	"codeid.revampacademy/repositories/jobhireRepositories/dbContext"
	"github.com/gin-gonic/gin"
)

type JobHirePostRepo struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
	MasterRepo
}

func NewJobPostRepo(dbHandler *sql.DB) *JobHirePostRepo {
	return &JobHirePostRepo{
		dbHandler: dbHandler,
	}
}

func (jp JobHirePostRepo) GetListJobPost(ctx *gin.Context) ([]*models.JobhireJobPost, *models.ResponseError) {
	market := dbContext.New(jp.dbHandler)
	jobPost, err := market.GetListJobPost(ctx)

	listjobPost := make([]*models.JobhireJobPost, 0)

	for _, v := range jobPost {
		listJob := models.JobhireJobPost{
			JopoEntityID:       v.JopoEntityID,
			JopoNumber:         v.JopoNumber,
			JopoTitle:          v.JopoTitle,
			JopoStartDate:      v.JopoStartDate,
			JopoEndDate:        v.JopoEndDate,
			JopoMinSalary:      v.JopoMaxSalary,
			JopoMaxSalary:      v.JopoMaxSalary,
			JopoMinExperience:  v.JopoMinExperience,
			JopoMaxExperience:  v.JopoMaxExperience,
			JopoPrimarySkill:   v.JopoPrimarySkill,
			JopoSecondarySkill: v.JopoSecondarySkill,
			JopoPublishDate:    v.JopoPublishDate,
			JopoModifiedDate:   v.JopoModifiedDate,
			JopoEmpEntityID:    v.JopoEmpEntityID,
			JopoClitID:         v.JopoClitID,
			JopoJoroID:         v.JopoJoroID,
			JopoJotyID:         v.JopoJotyID,
			JopoJocaID:         v.JopoJocaID,
			JopoAddrID:         v.JopoAddrID,
			JopoWorkCode:       v.JopoWorkCode,
			JopoEduCode:        v.JopoEduCode,
			JopoInduCode:       v.JopoInduCode,
			JopoStatus:         v.JopoStatus,
		}
		listjobPost = append(listjobPost, &listJob)
	}

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return listjobPost, nil
}

func (jp JobHirePostRepo) GetListJobPostMerge(ctx *gin.Context) ([]*models.MergeJobAndMaster, *models.ResponseError) {
	market := dbContext.New(jp.dbHandler)
	jobPost, err := market.ListJobPost(ctx)

	listjobPost := make([]*models.MergeJobAndMaster, 0)

	for _, v := range jobPost {
		listJob := models.MergeJobAndMaster{
			MasterAddress: v.MasterAddress,
			JobHirePost:   v.JobHirePost,
			MasterCity:    v.MasterCity,
		}
		listjobPost = append(listjobPost, &listJob)
	}

	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return listjobPost, nil
}

// func (jp JobHirePostRepo) GetJobPostMerge(ctx *gin.Context) ([]*models.MergeJobAndMaster, *models.ResponseError) {
// 	//take list of job post
// 	post, err := jp.GetListJobPostMerge(ctx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	//take list of address
// 	address, err := jp.GetListMasterAddress(ctx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	addressRoadMap := make(map[int32]*models.MasterAddress)
// 	for _, a := range address {
// 		addressRoadMap[a.AddrID] = a
// 	}

// 	//take list of city
// 	city, err := jp.GetListMasterCity(ctx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	//make store for city by city id
// 	cityRoadMap := make(map[int32]*models.MasterCity)
// 	for _, v := range city {
// 		cityRoadMap[v.CityID] = v
// 	}

// 	//merge data jobpost, address, and city
// 	mergeJobDetail := make([]*models.MergeJobAndMaster, 0)
// 	for _, l := range post {
// 		a, ok := addressRoadMap[l.MasterAddress.AddrID]
// 		// cityRoadMap[a.AddrCityID.Int32]
// 		if !ok {
// 			continue
// 		}

// 		v, ok := cityRoadMap[a.AddrCityID.Int32]
// 		if !ok {
// 			continue
// 		}

// 		list := models.MergeJobAndMaster{
// 			JobHirePost:   *&models.JobhireJobPost{},
// 			MasterAddress: *a,
// 			MasterCity:    *v,
// 		}

// 		mergeJobDetail = append(mergeJobDetail, &list)
// 	}

// 	return mergeJobDetail, nil
// }
