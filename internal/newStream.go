package internal

import (
	"database/sql"
	"fmt"
	"gerador-excel/models"
	"log/slog"

	"github.com/xuri/excelize/v2"
)

func NewStreamWriter(rows *sql.Rows) error {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			slog.Error("Falha ao fechar o arquivo:", err)
		}
	}()
	sw, err := f.NewStreamWriter("Sheet1")
	if err != nil {
		slog.Error(fmt.Sprintf("%v", err))
		return err
	}

	defer rows.Close()

	headers := []string{
		"Id", "Cycle ID", "Origin Area ID", "Origin Sub Area ID", "Origin Desk ID", "Allocated Desk ID",
		"Activity ID", "Project ID", "Head Count ID", "Origin Area", "Origin Sub Area",
		"Origin Desk", "Allocated Area", "Allocated Sub Area", "Allocated Desk", "Activity",
		"Activity Description", "Predefined Allocation Key", "Project", "Project Description",
		"Head Count Description", "Final Value", "Allocated Value", "Activity Type", "Project Type",
		"Created By", "Updated By", "Currency ID", "Cost Sub Category", "Desk Percent", "Updated At", "Created At", "Deleted At",
	}

	if err := sw.SetRow("A1", convertToStringInterface(headers)); err != nil {
		slog.Error(fmt.Sprintf("%v", err))
		return err
	}

	rowIndex := 2
	for rows.Next() {
		var allocation models.AllocationKeyReport
		err := rows.Scan(
			&allocation.Id, &allocation.CycleId, &allocation.OriginAreaId, &allocation.OriginSubAreaId,
			&allocation.OriginDeskId, &allocation.AllocatedDeskId, &allocation.ActivityId, &allocation.ProjectId,
			&allocation.HeadCountId, &allocation.OriginArea, &allocation.OriginSubArea, &allocation.OriginDesk,
			&allocation.AllocatedArea, &allocation.AllocatedSubArea, &allocation.AllocatedDesk, &allocation.Activity,
			&allocation.ActivityDescription, &allocation.PredefinedAllocationKey, &allocation.Project,
			&allocation.ProjectDescription, &allocation.HeadCountDescription, &allocation.FinalValue,
			&allocation.AllocatedValue, &allocation.ActivityType, &allocation.ProjectType, &allocation.CreatedBy,
			&allocation.UpdatedBy, &allocation.CurrencyId, &allocation.CostSubCategory, &allocation.DeskPercent,
			&allocation.UpdatedAt, &allocation.CreatedAt, &allocation.DeletedAt,
		)
		if err != nil {
			slog.Error("Erro ao escanear dados: ", err)
			return err
		}

		row := []interface{}{
			allocation.Id, allocation.CycleId,
			treatsInt64(allocation.OriginAreaId),
			treatsInt64(allocation.OriginSubAreaId),
			treatsInt64(allocation.OriginDeskId),
			treatsInt64(allocation.AllocatedDeskId),
			treatsInt64(allocation.ActivityId),
			treatsInt64(allocation.ProjectId),
			treatsInt64(allocation.HeadCountId),
			treatsString(allocation.OriginArea),
			treatsString(allocation.OriginSubArea),
			treatsString(allocation.OriginDesk),
			treatsString(allocation.AllocatedArea),
			treatsString(allocation.AllocatedSubArea),
			treatsString(allocation.AllocatedDesk),
			treatsString(allocation.Activity),
			treatsString(allocation.ActivityDescription),
			treatsString(allocation.PredefinedAllocationKey),
			treatsString(allocation.Project),
			treatsString(allocation.ProjectDescription),
			treatsString(allocation.HeadCountDescription),
			treatsFloat64(allocation.FinalValue),
			treatsFloat64(allocation.AllocatedValue),
			treatsString(allocation.ActivityType),
			treatsString(allocation.ProjectType),
			treatsString(allocation.CreatedBy),
			treatsString(allocation.UpdatedBy),
			treatsInt64(allocation.CurrencyId),
			treatsString(allocation.CostSubCategory),
			treatsFloat64(allocation.DeskPercent),
			treatsTime(allocation.UpdatedAt),
			allocation.CreatedAt,
			treatsTime(allocation.DeletedAt),
		}

		cell, _ := excelize.CoordinatesToCellName(1, rowIndex)
		if err := sw.SetRow(cell, row); err != nil {
			slog.Error(fmt.Sprintf("%v", err))
			return err
		}
		rowIndex++
	}

	if err := sw.Flush(); err != nil {
		slog.Error(fmt.Sprintf("%v", err))
		return err
	}

	if err := f.SaveAs("AllocationSemCanal.xlsx"); err != nil {
		slog.Error(fmt.Sprintf("%v", err))
		return err
	}

	return nil
}
