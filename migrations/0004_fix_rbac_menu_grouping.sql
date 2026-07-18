-- ແກ້ໄຂຜົນຂ້າງຄຽງຈາກ 0003: id=2 (module ແລະ main_menu) ຊ້ຳກັບຂໍ້ມູນ "ຈັດການສິນຄ້າ"/"ລາຍການສິນຄ້າ"
-- ທີ່ມີຢູ່ກ່ອນແລ້ວໃນ DB (ບໍ່ໄດ້ຢູ່ໃນ migrations ນີ້), ເຮັດໃຫ້ submenu "/role" ບໍ່ຖືກເພີ່ມ (ON CONFLICT skip)
-- ແລະ submenu /permission, /main-menu, /sub-menu ຖືກຈັດເຂົ້າກຸ່ມ "ລາຍການສິນຄ້າ" ຜິດແທນທີ່ຈະເປັນກຸ່ມ RBAC ຂອງມັນເອງ
-- Idempotent — ໃຊ້ ON CONFLICT DO NOTHING, ປອດໄພ run ຊ້ຳໄດ້

BEGIN;

-- ສ້າງກຸ່ມເມນູ "ສິດທິ ແລະ ເມນູ" ທີ່ຖືກຕ້ອງ (id ໃໝ່, ບໍ່ຊ້ຳກັບຂໍ້ມູນເດີມ)
INSERT INTO main_menus (id, module_id, menu_name, icon_class)
VALUES (4, 1, 'ສິດທິ ແລະ ເມນູ', 'mdi-shield-key')
ON CONFLICT (id) DO NOTHING;

-- ຍ້າຍ submenu /permission, /main-menu, /sub-menu ອອກຈາກກຸ່ມ "ລາຍການສິນຄ້າ" ມາຢູ່ກຸ່ມ RBAC ທີ່ຖືກຕ້ອງ
UPDATE sub_menus SET main_menu_id = 4
WHERE id IN (3, 4, 5) AND route_path IN ('/permission', '/main-menu', '/sub-menu');

-- ເພີ່ມ submenu "/role" ທີ່ຫາຍໄປ (ID 2 ຖືກໃຊ້ໄປແລ້ວໂດຍ "ລາຍການທັງໝົດ" /products/list)
INSERT INTO sub_menus (id, main_menu_id, submenu_name, route_path)
VALUES (9, 4, 'ຈັດການສິດການນຳໃຊ້', '/role')
ON CONFLICT (id) DO NOTHING;

-- ໃຫ້ສິດ Admin (role_id=1) ເຂົ້າເຖິງ submenu /role ແບບເຕັມ
INSERT INTO permissions (role_id, submenu_id, can_view, can_create, can_update, can_delete)
SELECT 1, 9, true, true, true, true
WHERE NOT EXISTS (SELECT 1 FROM permissions WHERE role_id = 1 AND submenu_id = 9);

SELECT setval(pg_get_serial_sequence('main_menus', 'id'), GREATEST((SELECT MAX(id) FROM main_menus), 1));
SELECT setval(pg_get_serial_sequence('sub_menus', 'id'), GREATEST((SELECT MAX(id) FROM sub_menus), 1));

COMMIT;
