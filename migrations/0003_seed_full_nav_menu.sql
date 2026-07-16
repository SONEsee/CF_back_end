-- Seed ຂໍ້ມູນ nav menu (module/main_menu/sub_menu) ໃຫ້ຄົບຕາມໜ້າທີ່ສ້າງໄວ້ໃນ CF_onLy frontend
-- Idempotent — ໃຊ້ ON CONFLICT DO NOTHING, ປອດໄພ run ຊ້ຳໄດ້
-- ອີງໃສ່ seed ເດີມໃນ 0002 (module id=1, main_menu id=1, sub_menu id=1, role id=1 "Admin")

BEGIN;

-- ແກ້ໄຂ route_path ຂອງ sub_menu ເດີມໃຫ້ກົງກັບ route ຈິງຂອງໜ້າ user (/user ບໍ່ແມ່ນ /users)
UPDATE sub_menus SET route_path = '/user' WHERE id = 1 AND route_path = '/users';

-- ===== Module 2: ຮ້ານຄ້າ =====
INSERT INTO modules (id, module_name, display_order)
VALUES (2, 'ຮ້ານຄ້າ', 2)
ON CONFLICT (id) DO NOTHING;

INSERT INTO main_menus (id, module_id, menu_name, icon_class)
VALUES (3, 2, 'ຮ້ານຄ້າ', 'mdi-store-cog')
ON CONFLICT (id) DO NOTHING;

INSERT INTO sub_menus (id, main_menu_id, submenu_name, route_path) VALUES
  (6, 3, 'ຈັດການຮ້ານຄ້າ', '/shop'),
  (7, 3, 'ຄ່າຕັ້ງຮ້ານຄ້າ', '/shop-setting'),
  (8, 3, 'ບັນຊີທະນາຄານຮ້ານຄ້າ', '/shop-bank-account')
ON CONFLICT (id) DO NOTHING;

-- ===== Module 1 (ເດີມ): ເພີ່ມ main_menu "ສິດທິ ແລະ ເມນູ" =====
INSERT INTO main_menus (id, module_id, menu_name, icon_class)
VALUES (2, 1, 'ສິດທິ ແລະ ເມນູ', 'mdi-shield-key')
ON CONFLICT (id) DO NOTHING;

INSERT INTO sub_menus (id, main_menu_id, submenu_name, route_path) VALUES
  (2, 2, 'ຈັດການສິດການນຳໃຊ້', '/role'),
  (3, 2, 'ຈັດການສິດອະນຸຍາດ', '/permission'),
  (4, 2, 'ຈັດການເມນູຫຼັກ', '/main-menu'),
  (5, 2, 'ຈັດການເມນູຍ່ອຍ', '/sub-menu')
ON CONFLICT (id) DO NOTHING;

-- ໃຫ້ສິດ Admin (role_id=1) ເຂົ້າເຖິງທຸກ submenu ໃໝ່ແບບເຕັມ (view/create/update/delete)
INSERT INTO permissions (role_id, submenu_id, can_view, can_create, can_update, can_delete)
SELECT 1, s.id, true, true, true, true
FROM sub_menus s
WHERE s.id IN (2, 3, 4, 5, 6, 7, 8)
  AND NOT EXISTS (
    SELECT 1 FROM permissions p WHERE p.role_id = 1 AND p.submenu_id = s.id
  );

-- ຣີເຊັດ sequence ໃຫ້ຕົງກັບ id ສູງສຸດທີ່ seed ໄວ້ (ປ້ອງກັນ id ຊ້ຳຮອບໜ້າ)
SELECT setval(pg_get_serial_sequence('modules', 'id'), GREATEST((SELECT MAX(id) FROM modules), 1));
SELECT setval(pg_get_serial_sequence('main_menus', 'id'), GREATEST((SELECT MAX(id) FROM main_menus), 1));
SELECT setval(pg_get_serial_sequence('sub_menus', 'id'), GREATEST((SELECT MAX(id) FROM sub_menus), 1));

COMMIT;
