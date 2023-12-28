package service

type userDomain struct {
	id        string
	firstName string
	lastName  string
	email     string
	age       int8
}

func (ud *userDomain) GetFirstName() string {
	return ud.firstName
}
func (ud *userDomain) GetLastName() string {
	return ud.lastName
}
func (ud *userDomain) GetEmail() string {
	return ud.email
}
func (ud *userDomain) GetAge() int8 {
	return ud.age
}
func (ud *userDomain) GetId() string {
	return ud.id
}
func (ud *userDomain) SetId(id string) {
	ud.id = id
}
