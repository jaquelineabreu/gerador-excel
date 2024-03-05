package repository

import (
	"database/sql"
	"gerador-excel/models"
	"log/slog"

	"github.com/jmoiron/sqlx"
)

func CreateAllocationkeyReport(report models.AllocationKeyReport, db *sqlx.DB) error {
	slog.Info("CreateAllocationkeyReport...")
	_, err := db.Exec(`
 		INSERT INTO allocation_key_report (
 			cycleId, originAreaId, originSubAreaId, originDeskId, allocatedDeskId,
 			activityId, projectId, headCountId, originArea, originSubArea, originDesk,
 			allocatedArea, allocatedSubArea, allocatedDesk, activity, activityDescription,
 			predefinedAllocationKey, project, projectDescription, headCountDescription,
 			finalValue, allocatedValue, activityType, projectType, createdBy, updatedBy, currencyId, costSubCategory, deskPercent
 		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		report.CycleId, report.OriginAreaId, report.OriginSubAreaId, report.OriginDeskId, report.AllocatedDeskId,
		report.ActivityId, report.ProjectId, report.HeadCountId, report.OriginArea, report.OriginSubArea, report.OriginDesk,
		report.AllocatedArea, report.AllocatedSubArea, report.AllocatedDesk, report.Activity, report.ActivityDescription,
		report.PredefinedAllocationKey, report.Project, report.ProjectDescription, report.HeadCountDescription,
		report.FinalValue, report.AllocatedValue, report.ActivityType, report.ProjectType, report.CreatedBy,
		report.UpdatedBy, report.CurrencyId, report.CostSubCategory, report.DeskPercent,
	)
	if err != nil {
		return err
	}

	return nil
}

func SelectAllAllocationKeyReports(db *sqlx.DB) (*sql.Rows, error) {
	query := `
 		SELECT
 			id, cycleId, originAreaId, originSubAreaId, originDeskId, allocatedDeskId,
 			activityId, projectId, headCountId, originArea, originSubArea, originDesk,
 			allocatedArea, allocatedSubArea, allocatedDesk, activity, activityDescription,
 			predefinedAllocationKey, project, projectDescription, headCountDescription,
 			finalValue, allocatedValue, activityType, projectType, createdBy, updatedBy,
 			currencyId, costSubCategory, deskPercent, updatedAt, createdAt, deletedAt
 		FROM allocation_key_report
 	`

	rows, err := db.Query(query)
	if err != nil {
		slog.Error("Erro ao consultar dados: ", err)
		return nil, err
	}

	return rows, nil
}

func StreamAllAllocationKeyReports(db *sqlx.DB) (<-chan models.AllocationKeyReport, <-chan error) {
	allocationsChan := make(chan models.AllocationKeyReport)
	errChan := make(chan error, 1)

	go func() {
		defer close(allocationsChan)
		defer close(errChan)

		query := `
            SELECT
                id, cycleId, originAreaId, originSubAreaId, originDeskId, allocatedDeskId,
                activityId, projectId, headCountId, originArea, originSubArea, originDesk,
                allocatedArea, allocatedSubArea, allocatedDesk, activity, activityDescription,
                predefinedAllocationKey, project, projectDescription, headCountDescription,
                finalValue, allocatedValue, activityType, projectType, createdBy, updatedBy,
                currencyId, costSubCategory, deskPercent, updatedAt, createdAt, deletedAt
            FROM allocation_key_report
        `

		rows, err := db.Queryx(query)
		if err != nil {
			errChan <- err
			return
		}
		defer rows.Close()

		for rows.Next() {
			var allocation models.AllocationKeyReport
			if err := rows.StructScan(&allocation); err != nil {
				errChan <- err
				return
			}
			allocationsChan <- allocation
		}

		if err = rows.Err(); err != nil {
			errChan <- err
		}
	}()

	return allocationsChan, errChan
}
