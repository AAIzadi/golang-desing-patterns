package creational

type User struct {
	Name  string
	Email string
	Age   int
}

type UserBuilder struct {
	name  string
	email string
	age   int
}

func NewUserBuilder() *UserBuilder {
	return &UserBuilder{}
}

func (builder *UserBuilder) Name(name string) *UserBuilder {
	builder.name = name
	return builder
}

func (builder *UserBuilder) Email(email string) *UserBuilder {
	builder.email = email
	return builder
}

func (builder *UserBuilder) Age(age int) *UserBuilder {
	builder.age = age
	return builder
}

func (builder *UserBuilder) Build() User {
	return User{builder.name, builder.email, builder.age}
}
