-- ເພີ່ມ column ຮູບພາບຮ້ານຄ້າ (ໂລໂກ້) — ເກັບເປັນ URL string ຄືກັບ product_images.image_url,
-- ບໍ່ໄດ້ຈັດການ file upload ຢູ່ backend ນີ້ໂດຍກົງ
ALTER TABLE shops ADD COLUMN IF NOT EXISTS image_url TEXT;
