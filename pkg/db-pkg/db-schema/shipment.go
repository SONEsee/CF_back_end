package dbschema

import "time"

type ShipmentDBSchema struct {
	ID             int        `db:"id" json:"id"`
	OrderID        int        `db:"order_id" json:"order_id"`
	CourierName    *string    `db:"courier_name" json:"courier_name"`
	TrackingNumber *string    `db:"tracking_number" json:"tracking_number"`
	LabelPdfURL    *string    `db:"label_pdf_url" json:"label_pdf_url"`
	ShippingStatus string     `db:"shipping_status" json:"shipping_status"`
	ShippedAt      *time.Time `db:"shipped_at" json:"shipped_at"`
	DeliveredAt    *time.Time `db:"delivered_at" json:"delivered_at"`
}
