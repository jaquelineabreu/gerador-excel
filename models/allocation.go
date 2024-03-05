package models

import (
	"time"
)

type AllocationKeyReport struct {
	Id                      int64      `json:"id" db:"id"`
	CycleId                 int64      `json:"cycle_id" db:"cycleId"`
	OriginAreaId            *int64     `json:"origin_area_id,omitempty" db:"originAreaId"`
	OriginSubAreaId         *int64     `json:"origin_sub_area_id,omitempty" db:"originSubAreaId"`
	OriginDeskId            *int64     `json:"origin_desk_id,omitempty" db:"originDeskId"`
	AllocatedDeskId         *int64     `json:"allocated_desk_id,omitempty" db:"allocatedDeskId"`
	ActivityId              *int64     `json:"activity_id,omitempty" db:"activityId"`
	ProjectId               *int64     `json:"project_id,omitempty" db:"projectId"`
	HeadCountId             *int64     `json:"head_count_id,omitempty" db:"headCountId"`
	OriginArea              *string    `json:"origin_area,omitempty" db:"originArea"`
	OriginSubArea           *string    `json:"origin_sub_area,omitempty" db:"originSubArea"`
	OriginDesk              *string    `json:"origin_desk,omitempty" db:"originDesk"`
	AllocatedArea           *string    `json:"allocated_area,omitempty" db:"allocatedArea"`
	AllocatedSubArea        *string    `json:"allocated_sub_area,omitempty" db:"allocatedSubArea"`
	AllocatedDesk           *string    `json:"allocated_desk,omitempty" db:"allocatedDesk"`
	Activity                *string    `json:"activity,omitempty" db:"activity"`
	ActivityDescription     *string    `json:"activity_description,omitempty" db:"activityDescription"`
	PredefinedAllocationKey *string    `json:"predefined_allocation_key,omitempty" db:"predefinedAllocationKey"`
	Project                 *string    `json:"project,omitempty" db:"project"`
	ProjectDescription      *string    `json:"project_description,omitempty" db:"projectDescription"`
	HeadCountDescription    *string    `json:"headCount_description,omitempty" db:"headCountDescription"`
	FinalValue              *float64   `json:"final_value,omitempty" db:"finalValue"`
	AllocatedValue          *float64   `json:"allocated_value,omitempty" db:"allocatedValue"`
	ActivityType            *string    `json:"activity_type,omitempty" db:"activityType"`
	ProjectType             *string    `json:"project_type,omitempty" db:"projectType"`
	CurrencyId              *int64     `json:"currency_id,omitempty" db:"currencyId"`
	CostSubCategory         *string    `json:"cost_sub_category,omitempty" db:"costSubCategory"`
	DeskPercent             *float64   `json:"desk_percent,omitempty" db:"deskPercent"`
	CreatedAt               time.Time  `json:"created_at" db:"createdAt"`
	UpdatedAt               *time.Time `json:"updated_at,omitempty" db:"updatedAt"`
	DeletedAt               *time.Time `json:"deleted_at,omitempty" db:"deletedAt"`
	CreatedBy               *string    `json:"created_by,omitempty" db:"createdBy"`
	UpdatedBy               *string    `json:"updated_by,omitempty" db:"updatedBy"`
}
