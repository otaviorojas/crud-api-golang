package repository

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"log"

	"github.com/crud_golang/database"
	"github.com/crud_golang/models"
)

type MySqlStudentRepo struct {
	Students map[int64]models.Student
}

const table = "students"

/******************************************************************************************************/

func (repo MySqlStudentRepo) GetAll() ([]models.Student, error) {

	var students []models.Student

	query := "SELECT * FROM " + table + ";"
	rows, err := database.Query(query)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	for rows.Next() {

		var student models.Student

		err := rows.Scan(
			&student.ID,
			&student.Name,
			&student.Age,
			&student.Email,
			&student.Course)

		if err != nil {
			return nil, err
		}

		students = append(students, student)
	}

	return students, nil
}

/******************************************************************************************************/

func (repo MySqlStudentRepo) Get(id int64) (*models.Student, error) {

	query := "SELECT * FROM " + table + " WHERE id =?;"
	rows, err := database.Query(query, id)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	for rows.Next() {

		var student models.Student

		err := rows.Scan(&student.ID, &student.Name, &student.Age, &student.Email, &student.Course)

		if err != nil {
			return nil, err
		}

		repo.Students[student.ID] = student
	}

	student, ok := repo.Students[id]

	if !ok {
		return nil, models.ErrStudentNotFound
	}

	return &student, nil

}

/******************************************************************************************************/

func (repo MySqlStudentRepo) Save(student *models.Student) (*models.Student, error) {

	if student == nil || student.Name == "" || student.Course == "" || student.Age == 0 {
		return nil, models.ErrValidateInput
	}

	query := "INSERT INTO " + table + " (name, age, email, course) VALUES (?,?,?,?);"

	result, err := database.Exec(query, student.Name, student.Age, student.Email, student.Course)

	if err != nil {
		fmt.Println("mysqlStudent.Save() - An error occured while saving the student")
		return nil, err
	}

	resultValue := *result
	student.ID, _ = resultValue.LastInsertId()
	studentObj, err := repo.Get(student.ID)

	if err != nil {
		return nil, err
	}

	return studentObj, nil
}

/******************************************************************************************************/

func (repo MySqlStudentRepo) Delete(id int64) (*models.Student, error) {

	student, err := repo.Get(id)

	if err == models.ErrStudentNotFound {
		return nil, models.ErrStudentNotFound
	}

	if err != nil {
		log.Printf("myssqlStudent.Delete() - An error occurred while deleting the student")
		return nil, err
	}

	query := "delete from " + table + " where id = ?;"
	_, err = database.Exec(query, student.ID)

	if err != nil {
		log.Println("myssqlStudent.Delete() - An error occured while deleting the student")
		return nil, err
	}

	return student, nil

}

/******************************************************************************************************/

func (repo MySqlStudentRepo) Update(student *models.Student) (*models.Student, error) {

	var query = "update " + table + " set "

	var searchParams []string

	if student.Name != "" {
		searchParams = append(searchParams, "name = '"+student.Name+"' ")
	}
	if student.Age != 0 {
		searchParams = append(searchParams, "age = "+strconv.FormatInt(student.Age, 10)+"")
	}
	if student.Email != "" {
		searchParams = append(searchParams, "email = '"+student.Email+"'")
	}
	if student.Course != "" {
		searchParams = append(searchParams, "course = '"+student.Course+"'")
	}

	if len(searchParams) <= 0 {

		return nil, models.ErrUpdateStudent

	}

	query = query + strings.Join(searchParams, ", ") + " where id = ?"
	_, err := database.Exec(query, student.ID)

	if err != nil {
		return nil, errors.New("myssqlStudent.Update() - An error occured while update the student")
	}

	ret, err := repo.Get(student.ID)

	return ret, nil

}
