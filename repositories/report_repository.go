package repositories

import (
	"database/sql"
	"kasir-api/models"
)

type ReportRepository struct {
	db *sql.DB
}

func NewReportRepository(db *sql.DB) *ReportRepository {
	return &ReportRepository{db: db}
}

func (r *ReportRepository) GetDailyReport() (*models.DailyReport, error) {
	var report models.DailyReport

	// 1. Total Revenue & Total Transaksi
	// Menggunakan COALESCE untuk menangani jika tidak ada transaksi (NULL -> 0)
	querySummary := `
		SELECT COALESCE(SUM(total_amount), 0), COUNT(id)
		FROM transactions
		WHERE DATE(created_at) = CURRENT_DATE
	`
	err := r.db.QueryRow(querySummary).Scan(&report.TotalRevenue, &report.TotalTransaksi)
	if err != nil {
		return nil, err
	}

	// 2. Produk Terlaris
	queryBestSelling := `
		SELECT p.name, COALESCE(SUM(td.quantity), 0) as total_qty
		FROM transaction_details td
		JOIN transactions t ON td.transaction_id = t.id
		JOIN products p ON td.product_id = p.id
		WHERE DATE(t.created_at) = CURRENT_DATE
		GROUP BY p.name
		ORDER BY total_qty DESC
		LIMIT 1
	`
	err = r.db.QueryRow(queryBestSelling).Scan(&report.ProdukTerlaris.Nama, &report.ProdukTerlaris.QtyTerjual)
	if err != nil {
		if err == sql.ErrNoRows {
			report.ProdukTerlaris = models.BestSellingProduct{Nama: "-", QtyTerjual: 0}
		} else {
			return nil, err
		}
	}

	return &report, nil
}
