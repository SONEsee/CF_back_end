INSERT INTO public.stock_movements (product_variant_id,movement_type,qty_change,balance_after,ref_type,ref_id,note,created_by,created_at) VALUES
	 (1,'IN'::public."movement_type_enum",50,50,'purchase',NULL,'ຮັບເຄື່ອງເຂົ້າສາງ',1,'2026-07-12 08:28:52.233133+07'),
	 (1,'OUT'::public."movement_type_enum",-20,30,'order',NULL,'ຂາຍອອກ',1,'2026-07-12 08:29:12.927088+07'),
	 (1,'OUT'::public."movement_type_enum",-10,20,'stock_reservation',1,'ຢືນຢັນການຂາຍຈາກສະຕັອກທີ່ຈອງໄວ້',1,'2026-07-12 08:34:18.461303+07'),
	 (1,'OUT'::public."movement_type_enum",-5,15,'stock_reservation',4,'ຢືນຢັນການຂາຍຈາກສະຕັອກທີ່ຈອງໄວ້',1,'2026-07-12 09:24:20.103971+07');
