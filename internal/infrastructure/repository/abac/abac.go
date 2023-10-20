package abac

import (
	"context"
	"errors"
	sq "github.com/Masterminds/squirrel"
	"saina.gitlab.yandexcloud.net/saina/backend/saina-api/internal/domain/entity"
	"saina.gitlab.yandexcloud.net/saina/backend/saina-api/internal/domain/storage/dto"
	"saina.gitlab.yandexcloud.net/saina/backend/saina-api/internal/domain/valueobject"
)

func AirDataConstraint(builder sq.SelectBuilder, restrictions UserPermissions) sq.SelectBuilder {
	if restrictions.Role.HigherOrEqualThan(valueobject.UserRoleDeveloper) { // if user is developer status
		return builder
	} else { // if user has lower permissions than developer
		return builder.Where(sq.Eq{"airl_code": restrictions.AirlCode})
	}
}

func GetRestrictions(ctx context.Context) (entity.Restrictions, error) {
	val, ok := ctx.Value("restrictions").(entity.Restrictions)
	if !ok {
		return entity.Restrictions{}, errors.New("error no restrictions was found")
	}
	return val, nil
}

func WrapperOfAirDataConstraint(builder *sq.SelectBuilder, ctx context.Context) {
	restrictions, err := GetRestrictions(ctx)
	if err != nil {
		panic("no restriction found in context")
	}
	*builder = AirDataConstraint(*builder, UserPermissions{
		AirlCode: restrictions.AirlCode,
		Role:     restrictions.UserRole,
	})
}

func HasUserPermission(user UserPermissions, targetAction TargetActionRestrictions) bool {
	if user.Role.HigherOrEqualThan(valueobject.UserRoleAdmin) {
		return true
	}
	if targetAction.AirlCode == nil {
		return true
	}
	return user.AirlCode == *targetAction.AirlCode
}

func IdentityWhere(b sq.SelectBuilder, i dto.FlightIdentity) sq.SelectBuilder {
	return b.Where(sq.Eq{
		"airl_code": i.AirlCode,
		"number":    i.Number,
		"orig_iata": i.OriginIATA,
		"dest_iata": i.DestinationIATA,
		"flt_date":  i.FltDate,
	})
}

func IdentityWithoutFltDateWhere(b sq.SelectBuilder, i dto.FlightIdentity) sq.SelectBuilder {
	return b.Where(sq.Eq{
		"airl_code": i.AirlCode,
		"number":    i.Number,
		"orig_iata": i.OriginIATA,
		"dest_iata": i.DestinationIATA,
	})
}
