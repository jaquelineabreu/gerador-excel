package datagenerator

import (
	"fmt"
	"gerador-excel/models"
	"gerador-excel/repository"
	"time"

	"github.com/jmoiron/sqlx"
	"golang.org/x/exp/slog"
)

func Gerador(db *sqlx.DB) {
	marketing := "Marketing"
	digital := "Digital"
	deskA := "Desk A"
	sales := "Sales"
	regional := "Regional"
	deskB := "Desk B"
	campaignAnalysis := "Campaign Analysis"
	analyzing := "Analyzing the effectiveness of the latest marketing campaign"
	key123 := "Key123"
	newProductLaunch := "New Product Launch"
	launching := "Launching the new product line in Q3"
	teamInvolved := "Team involved in the project"
	analysis := "Analysis"
	launch := "Launch"
	marketingExpense := "Marketing Expense"
	createdBy := "uuid-user-123"
	updatedBy := "uuid-user-456"
	cycleId := int64(1)
	originAreaId := int64(10)
	originSubAreaId := int64(20)
	originDeskId := int64(30)
	allocatedDeskId := int64(40)
	activityId := int64(50)
	projectId := int64(60)
	headCountId := int64(70)
	currencyId := int64(1)
	finalValue := float64(12000.50)
	allocatedValue := float64(6000.25)
	deskPercent := float64(0.5)

	for i := 0; i < 1000000; i++ {
		baseReport := models.AllocationKeyReport{
			CycleId:                 cycleId + int64(i),
			OriginAreaId:            &originAreaId,
			OriginSubAreaId:         &originSubAreaId,
			OriginDeskId:            &originDeskId,
			AllocatedDeskId:         &allocatedDeskId,
			ActivityId:              &activityId,
			ProjectId:               &projectId,
			HeadCountId:             &headCountId,
			OriginArea:              &marketing,
			OriginSubArea:           &digital,
			OriginDesk:              &deskA,
			AllocatedArea:           &sales,
			AllocatedSubArea:        &regional,
			AllocatedDesk:           &deskB,
			Activity:                &campaignAnalysis,
			ActivityDescription:     &analyzing,
			PredefinedAllocationKey: &key123,
			Project:                 &newProductLaunch,
			ProjectDescription:      &launching,
			HeadCountDescription:    &teamInvolved,
			FinalValue:              &finalValue,
			AllocatedValue:          &allocatedValue,
			ActivityType:            &analysis,
			ProjectType:             &launch,
			CurrencyId:              &currencyId,
			CostSubCategory:         &marketingExpense,
			DeskPercent:             &deskPercent,
			CreatedAt:               time.Now(),
			UpdatedAt:               &time.Time{},
			DeletedAt:               &time.Time{},
			CreatedBy:               &createdBy,
			UpdatedBy:               &updatedBy,
		}

		if err := repository.CreateAllocationkeyReport(baseReport, db); err != nil {
			slog.Error(fmt.Sprintf("Erro ao inserir registro %d: %v", i+1, err))
		}
	}

	fmt.Println("Registros inseridos com sucesso!")
}
