package sto

import (
	"github.com/achillescres/pkg/object/oid"
	"saina.gitlab.yandexcloud.net/saina/backend/saina-api/internal/domain/entity"
)

type BusinessRuleGroupTable struct {
	ID            oid.ID                 `json:"id" binding:"required"`
	Name          string                 `json:"name" binding:"required"`
	BusinessRules []*entity.BusinessRule `json:"businessRules" binding:"required"`
}

func NewBusinessRuleGroupTable(ID oid.ID, name string, BusinessRules []*entity.BusinessRule) *BusinessRuleGroupTable {
	return &BusinessRuleGroupTable{ID: ID, Name: name, BusinessRules: BusinessRules}
}
