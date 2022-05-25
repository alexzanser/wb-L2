package repository


type Repository struct {
	Calendar	Calendar
}

func New() *Repository {
	return &Repository{
		Calendar: NewCalendar(),
	}
}
