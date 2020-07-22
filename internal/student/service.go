package student

import "fmt"

type Service struct{}

func Newservice() *Service {
	return &Service{}
}

func (s Service) NumberOfStudents(room string) (int, error) {
	if room == "1" {
		return 10, nil
	}
	return 0, nil
}

func (s Service) CalulateGrade(point int) (string, error) {

	if point > 100 {
		return "", fmt.Errorf("Invalid grade!!!")
	}
	if point >= 80 {
		return "A", nil
	}
	if point >= 70 {
		return "B", nil
	}
	if point >= 60 {
		return "C", nil
	}
	if point >= 50 {
		return "D", nil
	}

	return "F", nil
}
