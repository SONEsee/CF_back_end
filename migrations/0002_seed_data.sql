-- Seed ຂໍ້ມູນພື້ນຖານ (idempotent — ໃຊ້ ON CONFLICT DO NOTHING, ປອດໄພ run ຊ້ຳໄດ້)
-- ລຳດັບຕາມ FK dependency: Zone1 -> Zone2 -> Zone3 -> Zone4 -> Zone5 -> Zone6 -> Zone7
-- login ຫຼັງ seed: username=admin, password=Admin@123

BEGIN;

-- ===== Zone 1: Tenant/Billing =====
INSERT INTO subscription_plans (id, plan_name, price_monthly, max_users, max_products, features)
VALUES (1, 'Starter', 99.00, 5, 200, '{"live_selling": true}'::jsonb)
ON CONFLICT (id) DO NOTHING;

INSERT INTO shops (id, shop_name, phone, status, timezone)
VALUES (1, 'ຮ້ານທົດລອງ (Seed Shop)', '02099999999', 'ACTIVE', 'Asia/Vientiane')
ON CONFLICT (id) DO NOTHING;

INSERT INTO shop_subscriptions (shop_id, plan_id, start_date, status)
SELECT 1, 1, CURRENT_DATE, 'ACTIVE'
WHERE NOT EXISTS (SELECT 1 FROM shop_subscriptions WHERE shop_id = 1 AND plan_id = 1);

INSERT INTO shop_bank_accounts (shop_id, bank_name, account_number, account_name, is_active)
SELECT 1, 'BCEL', '010-11-00-00000-1', 'ຮ້ານທົດລອງ', true
WHERE NOT EXISTS (SELECT 1 FROM shop_bank_accounts WHERE shop_id = 1);

INSERT INTO shop_settings (shop_id, currency, vat_rate)
SELECT 1, 'LAK', 7
WHERE NOT EXISTS (SELECT 1 FROM shop_settings WHERE shop_id = 1);

-- ===== Zone 2: Auth/RBAC =====
INSERT INTO roles (id, shop_id, role_name, description)
VALUES (1, NULL, 'Admin', 'ຜູ້ດູແລລະບົບສູງສຸດ')
ON CONFLICT (id) DO NOTHING;

INSERT INTO modules (id, module_name, display_order)
VALUES (1, 'ຈັດການລະບົບ', 1)
ON CONFLICT (id) DO NOTHING;

INSERT INTO main_menus (id, module_id, menu_name, icon_class)
VALUES (1, 1, 'ຜູ້ໃຊ້', 'fa-users')
ON CONFLICT (id) DO NOTHING;

INSERT INTO sub_menus (id, main_menu_id, submenu_name, route_path)
VALUES (1, 1, 'ລາຍຊື່ຜູ້ໃຊ້', '/users')
ON CONFLICT (id) DO NOTHING;

INSERT INTO permissions (role_id, submenu_id, can_view, can_create, can_update, can_delete)
SELECT 1, 1, true, true, true, true
WHERE NOT EXISTS (SELECT 1 FROM permissions WHERE role_id = 1 AND submenu_id = 1);

-- password = Admin@123 (bcrypt hash generated ດ້ວຍ golang.org/x/crypto/bcrypt)
INSERT INTO users (id, shop_id, role_id, username, password_hash, full_name, email, is_active)
VALUES (1, 1, 1, 'admin', '$2a$10$YVwY.5iVXULXmLckdDMX4.kPkZibhd3iyYCOGmEtmpdfjOes1tmeK', 'ຜູ້ດູແລລະບົບ', 'admin@example.com', true)
ON CONFLICT (id) DO NOTHING;

UPDATE shops SET owner_user_id = 1 WHERE id = 1 AND owner_user_id IS NULL;

-- ===== Zone 3: Product catalog =====
INSERT INTO product_categories (id, shop_id, name, sort_order)
VALUES (1, 1, 'ເສື້ອຜ້າ', 1)
ON CONFLICT (id) DO NOTHING;

INSERT INTO products (id, shop_id, category_id, product_name, description)
VALUES (1, 1, 1, 'ເສື້ອຢືດຄໍມົນ', 'ຜ້າຝ້າຍ 100%')
ON CONFLICT (id) DO NOTHING;

INSERT INTO product_variants (id, product_id, variant_name, sku_code, price, cost_price, weight_grams)
VALUES (1, 1, 'ໄຊສ໌ M ສີດຳ', 'SEED-TSHIRT-M-BLACK', 89000, 50000, 200)
ON CONFLICT (id) DO NOTHING;

INSERT INTO inventories (product_variant_id, actual_qty, reserved_qty, available_qty, reorder_level)
SELECT 1, 100, 0, 100, 10
WHERE NOT EXISTS (SELECT 1 FROM inventories WHERE product_variant_id = 1);

-- ===== Zone 4: Customers =====
INSERT INTO customers (id, shop_id, customer_name, phone_number)
VALUES (1, 1, 'ລູກຄ້າທົດລອງ', '02055555555')
ON CONFLICT (id) DO NOTHING;

INSERT INTO customer_addresses (id, customer_id, recipient_name, phone, address, province, is_default)
VALUES (1, 1, 'ລູກຄ້າທົດລອງ', '02055555555', 'ບ້ານໂພນທັນ', 'ນະຄອນຫຼວງວຽງຈັນ', true)
ON CONFLICT (id) DO NOTHING;

UPDATE customers SET default_address_id = 1 WHERE id = 1 AND default_address_id IS NULL;

-- ===== Zone 5: Order/Payment =====
INSERT INTO discounts (id, shop_id, code, discount_type, discount_value, min_order, usage_limit)
VALUES (1, 1, 'WELCOME10', 'PERCENT', 10, 50000, 100)
ON CONFLICT (id) DO NOTHING;

-- ===== Zone 6: Social =====
INSERT INTO social_accounts (id, shop_id, platform, platform_account_id, account_name, is_active)
VALUES (1, 1, 'FACEBOOK_PAGE', 'seed-fbpage-001', 'ຮ້ານທົດລອງ Page', true)
ON CONFLICT (id) DO NOTHING;

INSERT INTO chat_templates (id, shop_id, trigger_keyword, response_body, is_active)
VALUES (1, 1, 'ລາຄາ', 'ສອບຖາມລາຄາ ພິມ CF ຕາມດ້ວຍລະຫັດສິນຄ້າ ແລະ ຈຳນວນ', true)
ON CONFLICT (id) DO NOTHING;

-- ===== Zone 7: Live-session =====
INSERT INTO live_sessions (id, social_account_id, session_title, status)
VALUES (1, 1, 'Live ທົດລອງ (Seed)', 'STREAMING')
ON CONFLICT (id) DO NOTHING;

INSERT INTO live_session_products (id, live_session_id, product_variant_id, live_price, cf_code_override)
VALUES (1, 1, 1, 79000, 'SEED01')
ON CONFLICT (id) DO NOTHING;

-- ຣີເຊັດ sequence ໃຫ້ຕົງກັບ id ສູງສຸດທີ່ seed ໄວ້ (ປ້ອງກັນ id ຊ້ຳຮອບໜ້າ)
SELECT setval(pg_get_serial_sequence('subscription_plans', 'id'), GREATEST((SELECT MAX(id) FROM subscription_plans), 1));
SELECT setval(pg_get_serial_sequence('shops', 'id'), GREATEST((SELECT MAX(id) FROM shops), 1));
SELECT setval(pg_get_serial_sequence('roles', 'id'), GREATEST((SELECT MAX(id) FROM roles), 1));
SELECT setval(pg_get_serial_sequence('modules', 'id'), GREATEST((SELECT MAX(id) FROM modules), 1));
SELECT setval(pg_get_serial_sequence('main_menus', 'id'), GREATEST((SELECT MAX(id) FROM main_menus), 1));
SELECT setval(pg_get_serial_sequence('sub_menus', 'id'), GREATEST((SELECT MAX(id) FROM sub_menus), 1));
SELECT setval(pg_get_serial_sequence('users', 'id'), GREATEST((SELECT MAX(id) FROM users), 1));
SELECT setval(pg_get_serial_sequence('product_categories', 'id'), GREATEST((SELECT MAX(id) FROM product_categories), 1));
SELECT setval(pg_get_serial_sequence('products', 'id'), GREATEST((SELECT MAX(id) FROM products), 1));
SELECT setval(pg_get_serial_sequence('product_variants', 'id'), GREATEST((SELECT MAX(id) FROM product_variants), 1));
SELECT setval(pg_get_serial_sequence('customers', 'id'), GREATEST((SELECT MAX(id) FROM customers), 1));
SELECT setval(pg_get_serial_sequence('customer_addresses', 'id'), GREATEST((SELECT MAX(id) FROM customer_addresses), 1));
SELECT setval(pg_get_serial_sequence('discounts', 'id'), GREATEST((SELECT MAX(id) FROM discounts), 1));
SELECT setval(pg_get_serial_sequence('social_accounts', 'id'), GREATEST((SELECT MAX(id) FROM social_accounts), 1));
SELECT setval(pg_get_serial_sequence('chat_templates', 'id'), GREATEST((SELECT MAX(id) FROM chat_templates), 1));
SELECT setval(pg_get_serial_sequence('live_sessions', 'id'), GREATEST((SELECT MAX(id) FROM live_sessions), 1));
SELECT setval(pg_get_serial_sequence('live_session_products', 'id'), GREATEST((SELECT MAX(id) FROM live_session_products), 1));

COMMIT;
