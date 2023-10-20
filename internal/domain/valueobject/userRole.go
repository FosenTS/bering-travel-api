package valueobject

// ORDER IS VERY IMPORTANT! greatness means more rights | ПОРЯДОК ОЧЕНЬ ВАЖЕН !!!

type UserRole int

var UserRoles = map[string]UserRole{
	"tester":    -1,
	"manager":   0,
	"admin":     1,
	"developer": 2,
}

const (
	UserRoleTester = -1 + iota
	UserRoleManager
	UserRoleAdmin
	UserRoleDeveloper
)

func (r UserRole) HigherOrEqualThan(or UserRole) bool {
	return r >= or
}

func (r UserRole) LowerThan(or UserRole) bool {
	return r < or
}

func (r UserRole) Equal(or UserRole) bool {
	return r == or
}
