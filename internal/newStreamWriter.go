package internal

import (
	"fmt"
	"gerador-excel/models"
	"log/slog"

	"github.com/xuri/excelize/v2"
)

func NewStreamWriter(allocationsChan <-chan models.AllocationKeyReport, errChan <-chan error) error {
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
	for allocation := range allocationsChan {
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

	if err, ok := <-errChan; ok && err != nil {
		slog.Error("Erro recebido do canal de erro:", err)
		return err
	}

	if err := sw.Flush(); err != nil {
		slog.Error(fmt.Sprintf("%v", err))
		return err
	}

	if err := f.SaveAs("AllocationKeyReport.xlsx"); err != nil {
		slog.Error(fmt.Sprintf("%v", err))
		return err
	}

	return nil
}
