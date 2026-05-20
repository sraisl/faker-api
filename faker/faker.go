package faker

import (
	"github.com/go-faker/faker/v4"
)

func FakeName() string {
	return faker.Name()
}

func FakeEmail() string {
	return faker.Email()
}