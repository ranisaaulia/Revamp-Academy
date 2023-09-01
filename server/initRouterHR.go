package server

import (
	"codeid.revampacademy/controllers/hrController"
	"github.com/gin-gonic/gin"
)

func InitRouterHR(routers *gin.Engine, controllerMgr *hrController.ControllerManager) *gin.Engine {

	// Register routes here
	clientContractRoute := routers.Group("/clientcontract")
	{
		// routers endpoint/url http category
		clientContractRoute.GET("", controllerMgr.ClientContractController.GetListClientContract)
		clientContractRoute.GET("/:id", controllerMgr.ClientContractController.GetClientContract)
		clientContractRoute.POST("", controllerMgr.ClientContractController.CreateClientContract)
		clientContractRoute.PUT("/:id", controllerMgr.ClientContractController.UpdateClientContract)
		clientContractRoute.DELETE("/:id", controllerMgr.ClientContractController.DeleteClientContract)
	}

	departmentHistoryRoute := routers.Group("/departmenthistory")
	{
		// routers endpoint/url http category
		departmentHistoryRoute.GET("", controllerMgr.DepartmentHistoryController.GetListDepartmentHistory)
		departmentHistoryRoute.GET("/:id", controllerMgr.DepartmentHistoryController.GetDepartmentHistory)
		departmentHistoryRoute.POST("", controllerMgr.DepartmentHistoryController.CreateDepartmentHistory)
		departmentHistoryRoute.PUT("/:id", controllerMgr.DepartmentHistoryController.UpdateDepartmentHistory)
		departmentHistoryRoute.DELETE("/:id", controllerMgr.DepartmentHistoryController.DeleteDepartmenHistory)
	}

	departmentRoute := routers.Group("/department")
	{
		// routers endpoint/url http category
		departmentRoute.GET("", controllerMgr.DepartmentController.GetListDepartment)
		departmentRoute.GET("/:id", controllerMgr.DepartmentController.GetDepartment)
		departmentRoute.POST("", controllerMgr.DepartmentController.CreateDepartment)
		departmentRoute.PUT("/:id", controllerMgr.DepartmentController.UpdateDepartment)
		departmentRoute.DELETE("/:id", controllerMgr.DepartmentController.DeleteDepartment)
	}

	employeeRouter := routers.Group("/employee")
	{
		// routers endpoint/url http category
		employeeRouter.GET("", controllerMgr.EmployeeController.GetListEmployee)
		employeeRouter.GET("/:id", controllerMgr.EmployeeController.GetEmployee)
		employeeRouter.POST("", controllerMgr.EmployeeController.CreateEmployee)
		employeeRouter.PUT("/:id", controllerMgr.EmployeeController.UpdateEmployee)
		employeeRouter.DELETE("/:id", controllerMgr.EmployeeController.DeleteEmployee)
	}

	payHistoryRoute := routers.Group("/payhistory")
	{
		// routers endpoint/url http category
		payHistoryRoute.GET("", controllerMgr.PayHistoryController.GetListPayHistory)
		payHistoryRoute.GET("/:id", controllerMgr.PayHistoryController.GetListPayHistory)
		payHistoryRoute.POST("", controllerMgr.PayHistoryController.CreatePayHistory)
		payHistoryRoute.PUT("/:id", controllerMgr.PayHistoryController.UpdatePayHistory)
		payHistoryRoute.DELETE("/:id", controllerMgr.PayHistoryController.DeletePayHistory)
	}

	talentDetailRoute := routers.Group("/api/hr")
	{
		// routers endpoint/url http category
		talentDetailRoute.GET("/talentdetail", controllerMgr.TalentsDetailMockupController.GetListTalentDetailMockupDetail)
		talentDetailRoute.GET("/talentdetail/:id", controllerMgr.TalentsDetailMockupController.GetTalentDetail)

	}

	talentRoute := routers.Group("/api/hr")
	{
		// routers endpoint/url http category
		talentRoute.GET("/talent", controllerMgr.TalentsMockupController.GetListTalentMockup)
		talentRoute.GET("/talent/search", controllerMgr.TalentsMockupController.SearchTalent)
		talentRoute.GET("/talent/paging", controllerMgr.TalentsMockupController.PagingTalent)
	}
	return routers
}
