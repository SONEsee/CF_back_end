-- ເພີ່ມ column ຮູບໂປຣໄຟລ໌ຜູ້ໃຊ້ — ເກັບເປັນ URL string ຄືກັນກັບ shops.image_url
ALTER TABLE users ADD COLUMN IF NOT EXISTS profile_image TEXT;
