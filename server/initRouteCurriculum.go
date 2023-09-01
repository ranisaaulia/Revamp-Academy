package server

import (
	controllers "codeid.revampacademy/controllers/curriculumControllers"
	"github.com/gin-gonic/gin"
)

func InitRouterCurriculum(router *gin.Engine, controllerMrg *controllers.ControllerManager) *gin.Engine {

	progentityRoute := router.Group("/curriculum")
	{
		progentityRoute.GET("/progentity", controllerMrg.ProgEntityController.GetListProgEntity)
		progentityRoute.GET("/progentity/:id", controllerMrg.ProgEntityController.GetProgEntity)
		progentityRoute.POST("/progentity", controllerMrg.ProgEntityController.CreateProgEntity)
		progentityRoute.PUT("/progentity/:id", controllerMrg.ProgEntityController.UpdateProgEntity)
		progentityRoute.DELETE("/progentity/:id", controllerMrg.ProgEntityController.DeleteProgEntity)

		// progentityRoute.POST("/withSection", controllerMrg.ProgEntityController.CreateProgEntityWithSection)

		progentityRoute.GET("/sections", controllerMrg.ProgEntityController.GetListSection)
		progentityRoute.GET("/sections/:id", controllerMrg.ProgEntityController.GetSection)
		progentityRoute.POST("/sections", controllerMrg.ProgEntityController.CreateSection)
		//progentityRoute.PUT("/sections/:id", controllerMrg.ProgEntityController.UpdateSection)

		progentityRoute.GET("/sectiondetail", controllerMrg.ProgEntityController.GetListSectionDetail)

		progentityRoute.GET("/mastercategory", controllerMrg.ProgEntityController.GetListMasterCategory)

		progentityRoute.GET("/gabung", controllerMrg.ProgEntityController.GetListGabung)
		progentityRoute.GET("/gabung/:id", controllerMrg.ProgEntityController.GetGabung)
		progentityRoute.POST("/createallgabung", controllerMrg.ProgEntityController.CreateGabung)

	}

	progentitydescRoute := router.Group("/curriculum")
	{
		progentitydescRoute.GET("/progEntityDesc", controllerMrg.ProgEntityDescController.GetListProgEntityDesc)
		progentitydescRoute.GET("/progEntityDesc/:id", controllerMrg.ProgEntityDescController.GetProgEntityDesc)
		progentitydescRoute.POST("/progEntityDesc", controllerMrg.ProgEntityDescController.CreateProgEntityDesc)
		progentitydescRoute.DELETE("/progEntityDesc/:id", controllerMrg.ProgEntityDescController.DeleteProgEntityDesc)

	}
	progreviewsRoute := router.Group("/curriculum")
	{
		progreviewsRoute.GET("/progReviews", controllerMrg.ProgReviewsController.GetListProgReviews)
		progreviewsRoute.GET("/progReviews/:id", controllerMrg.ProgReviewsController.GetProgramReviews)
	}
	sectionDetailmaterialRoute := router.Group("/curriculum")
	{
		sectionDetailmaterialRoute.GET("/sectionDetailMaterial", controllerMrg.SectionDetailMaterialController.GetListSectionDetailMaterial)
		sectionDetailmaterialRoute.GET("/sectionDetailMaterial/:id", controllerMrg.SectionDetailMaterialController.GetSectionDetailMaterial)
		sectionDetailmaterialRoute.POST("/sectionDetailMaterial", controllerMrg.SectionDetailMaterialController.CreatesectiondetailMaterial)
		sectionDetailmaterialRoute.PUT("/sectionDetailMaterial/:id", controllerMrg.SectionDetailMaterialController.UpdateSectionDetailMaterial)
		sectionDetailmaterialRoute.DELETE("/sectionDetailMaterial/:id", controllerMrg.SectionDetailMaterialController.DeleteSectionDetailMaterial)

	}
	return router

}
