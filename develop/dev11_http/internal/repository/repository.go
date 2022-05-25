package repository

/*Repository is a wrapper for Calendar interface which can be 
modified according to data storage needs
*/
type Repository struct {
	Calendar	Calendar
}

//NewRepository return instance of type *Repository
func NewRepository() *Repository {
	return &Repository{
		Calendar: NewCalendar(),
	}
}
