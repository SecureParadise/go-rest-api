package sqlite

import (
	"database/sql"
	"fmt"

	"github.com/SecureParadise/students-api/internal/config"
	"github.com/SecureParadise/students-api/internal/types"
	_ "github.com/mattn/go-sqlite3"
	// since sqlit3 will not be used directly so use _
)

// for postgres
// make struct of postgres
type Sqlite struct {
	Db *sql.DB
}

// there is no concept of construtor in go
// we use hack to get work of constructor
// By convention we use New function for the constructor

// implement storage interface for postgres ,if you want to use postgres
func New(cfg *config.Config) (*Sqlite, error) {
	db, err := sql.Open("sqlite3", cfg.StoragePath)

	if err != nil {
		return nil, err
	}
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS students(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT,
	email TEXT,
	age INTEGER)`)
	if err != nil {
		return nil, err
	}
	return &Sqlite{
		Db: db,
	}, nil
}

func (s *Sqlite) CreateStudent(name string, email string, age int) (uint64, error) {
	// use (?,?,?) for prevention of sqlite injection
	stmt, err := s.Db.Prepare("INSERT INTO students (name,email,age) VALUES(?,?,?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(name, email, age)
	if err != nil {
		return 0, err
	}
	lastId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint64(lastId), nil
}

// neeed sqlite driver
// go get github.com/mattn/go-sqlite3

func (s *Sqlite) GetStudentById(id uint64) (types.Student, error) {
	stmt, err := s.Db.Prepare("SELECT id,name,email,age FROM students WHERE id = ? LIMIT 1")
	if err != nil {
		return types.Student{}, err
	}
	defer stmt.Close()
	var student types.Student
	err = stmt.QueryRow(int64(id)).Scan(&student.Id, &student.Name, &student.Email, &student.Age)
	if err != nil {
		if err == sql.ErrNoRows {
			return types.Student{}, fmt.Errorf("no student found with id %d", id)
		}
		return types.Student{}, fmt.Errorf("Query error: %w", err)
	}
	return student, nil
}

func (s *Sqlite) GetStudents() ([]types.Student, error) {
	stmt, err := s.Db.Prepare("SELECT id,name,email,age FROM students")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var students []types.Student
	for rows.Next() {
		var student types.Student
		err := rows.Scan(&student.Id, &student.Name, &student.Email, &student.Age)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}
	return students, nil
}
