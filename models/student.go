package models

import "errors"

type Student struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Age    int64  `json:"age"`
	Email  string `json:"email"`
	Course string `json:"course"`
}

var (
	ErrStudentNotFound = errors.New("Student not found")
	ErrValidateInput   = errors.New("Please, check if params are ok")
	ErrUpdateStudent   = errors.New("Please, send the params")
)
