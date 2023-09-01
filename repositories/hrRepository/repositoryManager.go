package hrRepository

import "database/sql"

type RepositoryManager struct {
	ClientContractRepository
	DepartmentRepository
	DepartmentHistoryRepository
	EmployeeRepository
	PayHistoryRepository
	TalentsDetailMockupRepository
	TalentsMockupRepository
}

// constructor
func NewRepositoryManager(dbHandler *sql.DB) *RepositoryManager {
	return &RepositoryManager{
		*NewClientContractRepository(dbHandler),
		*NewDepartmentRepository(dbHandler),
		*NewDepartmentHistoryRepository(dbHandler),
		*NewEmployeeRepository(dbHandler),
		*NewPayHistoryRepository(dbHandler),
		*NewTalentDetailMockupRepository(dbHandler),
		*NewTalentMockupRepository(dbHandler),
	}
}
