package service

type UserDomainInterface interface {
	GetEmail() string
	GetFirstName() string
	GetLastName() string
	GetAge() int8
	SetId(string)
	GetId() string
}

func NewUserDomain(firstName, lastName, email string, age int8) UserDomainInterface {
	return &userDomain{
		id:        "",
		firstName: firstName,
		lastName:  lastName,
		email:     email,
		age:       age,
	}
}
