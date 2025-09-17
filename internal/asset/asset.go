package asset

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/rumiani/rate-api/internal/db"
)

type Asset struct {
	ID           uuid.UUID `json:"id"`
	Code         string    `json:"code"`
	EnName       []string  `json:"enName"`
	FaName       []string  `json:"faName"`
	BuyCode      string    `json:"buyCode"`
	SellCode     string    `json:"sellCode"`
	Type         string    `json:"type"` // Could use custom enum later
	CurrentPrice float64   `json:"currentPrice"`
	Status       string    `json:"status"`
	UpdatedAt    string    `json:"updatedAt"` // ISO string for JSON
}

func GetAllAssets() ([]Asset, error) {
	rows, err := db.Pool.Query(context.Background(),
		`SELECT id, code, "enName", "faName", "buyCode", "sellCode", type, "currentPrice", status, "updatedAt"
		 FROM "Asset"
		 ORDER BY "updatedAt" DESC
		 LIMIT 100`) // Add pagination later
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var assets []Asset
	for rows.Next() {
		var a Asset
		var updatedAt time.Time
		err := rows.Scan(&a.ID, &a.Code, &a.EnName, &a.FaName, &a.BuyCode, &a.SellCode, &a.Type, &a.CurrentPrice, &a.Status, &updatedAt)
		if err != nil {
			return nil, err
		}
		a.UpdatedAt = updatedAt.Format(time.RFC3339)
		assets = append(assets, a)
	}
	return assets, nil
}

func GetAssetByCode(code string) (*Asset, error) {
	row := db.Pool.QueryRow(context.Background(),
		`SELECT id, code, "enName", "faName", "buyCode", "sellCode", type, "currentPrice", status, "updatedAt"
		 FROM "Asset" WHERE code=$1`, code)

	var a Asset
	var updatedAt time.Time
	err := row.Scan(&a.ID, &a.Code, &a.EnName, &a.FaName, &a.BuyCode, &a.SellCode, &a.Type, &a.CurrentPrice, &a.Status, &updatedAt)
	if err != nil {
		return nil, err
	}
	a.UpdatedAt = updatedAt.Format(time.RFC3339)
	return &a, nil
}
