INSERT INTO public.payments (order_id,shop_bank_account_id,payment_method,slip_image_path,bank_trans_ref_id,verified_amount,is_valid_slip,paid_at,created_at) VALUES
	 (1,NULL,'SLIP'::public."payment_method_enum",'https://example.com/slip1.jpg','TXN-0001',442500.00,true,'2026-07-12 09:28:08.297374+07','2026-07-12 09:28:08.276239+07');
