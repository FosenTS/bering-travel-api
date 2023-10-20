package abac

import (
	"saina.gitlab.yandexcloud.net/saina/backend/saina-api/internal/domain/valueobject"
)

type UserPermissions struct {
	AirlCode string
	Role     valueobject.UserRole
}

func NewUserPermissions(airlCode string, role valueobject.UserRole) *UserPermissions {
	return &UserPermissions{AirlCode: airlCode, Role: role}
}

type TargetActionRestrictions struct {
	AirlCode *string
}

type Restriction func(t *TargetActionRestrictions)

func NewTargetActionRestrictions(restrictions ...Restriction) *TargetActionRestrictions {
	t := &TargetActionRestrictions{}
	for _, r := range restrictions {
		r(t)
	}

	return t
}

func WithAirlCodeRestriction(airlCode string) Restriction {
	return func(t *TargetActionRestrictions) {
		t.AirlCode = &airlCode
	}
}
