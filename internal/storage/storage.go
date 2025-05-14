package storage

import (
	"github.com/SecureParadise/students-api/internal/types"
)

type Storage interface {
	CreateStudent(name string, email string, age int) (uint64, error)
	GetStudentById(id uint64) (types.Student, error)
	GetStudents() ([]types.Student, error)
}
