package service

import (
	"context"
	"fmt"

	pb "github.com/SudarshanZone/Open_Pos/internal/generated"
	"gorm.io/gorm"
)

type Server struct {
	Db *gorm.DB
	pb.UnimplementedFnoPositionServiceServer
}

// FnoPosition represents the structure for the database query result
type FnoPosition struct {
	Contract          string  `gorm:"column:Contract"`
	Position          string  `gorm:"column:Position"`
	TotalQty          int32   `gorm:"column:TotalQty"`
	AvgCostPrice      float64 `gorm:"column:AvgCostPrice"`
	OpenPositionValue float64 `gorm:"column:OpenPositionValue"`
}

func (s *Server) GetFNOPosition(ctx context.Context, req *pb.FnoPositionRequest) (*pb.FnoPositionResponse, error) {
	var pb2 []FnoPosition

	// Define the query with GORM, including the match condition
	query := `
		SELECT
			CASE
			WHEN FCP_PRDCT_TYP = 'F' THEN
				'FUT-' || FCP_UNDRLYNG || '-' || TO_CHAR(FCP_EXPRY_DT, 'DD-Mon-YYYY')
			WHEN FCP_PRDCT_TYP = 'O' THEN
				'OPT-' || FCP_UNDRLYNG || '-' || TO_CHAR(FCP_EXPRY_DT, 'DD-Mon-YYYY') || '-' ||
				FCP_STRK_PRC || '-' ||
				CASE
				WHEN FCP_OPT_TYP = 'C' THEN 'CE'
				WHEN FCP_OPT_TYP = 'P' THEN 'PE'
				ELSE ''
				END
			ELSE
				FCP_UNDRLYNG || ' ' || TO_CHAR(FCP_EXPRY_DT, 'DD-Mon-YYYY')
			END AS "Contract",
			CASE
			WHEN FCP_OPNPSTN_FLW = 'B' THEN 'BUY'
			WHEN FCP_OPNPSTN_FLW = 'S' THEN 'SELL'
			WHEN FCP_OPNPSTN_FLW = 'N' THEN ''
			ELSE FCP_OPNPSTN_FLW
			END AS "Position",
			ABS(FCP_OPNPSTN_QTY) AS "TotalQty",
			COALESCE(FCP_AVG_PRC, 0) AS "AvgCostPrice",
			COALESCE(FCP_OPNPSTN_VAL, 0) AS "OpenPositionValue"
		FROM
			FCP_FO_SPN_CNTRCT_PSTN
		WHERE
			(FCP_OPNPSTN_QTY != 0 OR FCP_IBUY_QTY != 0 OR FCP_ISELL_QTY != 0)
			AND FCP_CLM_MTCH_ACCNT = ?;
	`

	// Execute the query with the provided account number
	if err := s.Db.Raw(query, req.GetFcpClmMtchAccnt()).Scan(&pb2).Error; err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	for _, pos := range pb2 {
		fmt.Printf("Contract: %s, Position: %s, Total Qty: %d, Avg Cost Price: %.2f, Open Position Value: %.2f\n",
			pos.Contract, pos.Position, pos.TotalQty, pos.AvgCostPrice, pos.OpenPositionValue)
	}

	// Populate the response
	response := &pb.FnoPositionResponse{}
	for _, pos := range pb2 {
		response.FfoContract = append(response.FfoContract, pos.Contract)
		response.FfoPstn = append(response.FfoPstn, pos.Position)
		response.FfoQty = append(response.FfoQty, pos.TotalQty)
		response.FfoAvgPrc = append(response.FfoAvgPrc, float32(pos.AvgCostPrice))
	}

	return response, nil
}
