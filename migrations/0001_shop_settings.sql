-- Zone 1 (Tenant/Billing) — ຕາຕະລາງທີ່ຂາດຫາຍ: ຄ່າຕັ້ງລະດັບຮ້ານ (1 ຮ້ານ = 1 ແຖວ)
CREATE TABLE shop_settings (
    id               SERIAL PRIMARY KEY,
    shop_id          INT NOT NULL UNIQUE REFERENCES shops(id) ON DELETE CASCADE,
    currency         VARCHAR(10) NOT NULL DEFAULT 'LAK',
    vat_rate         NUMERIC(5,2) NOT NULL DEFAULT 0,
    auto_reply_msg   TEXT,
    business_hours   JSONB NOT NULL DEFAULT '{}'::jsonb,
    created_at       TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at       TIMESTAMPTZ NOT NULL DEFAULT now()
);
