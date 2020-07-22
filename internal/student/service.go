package student

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
