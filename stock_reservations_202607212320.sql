INSERT INTO public.stock_reservations (product_variant_id,customer_id,order_item_id,reserved_qty,expires_at,status,created_at) VALUES
	 (1,NULL,NULL,10,'2026-08-01 07:00:00+07','COMPLETED'::public."reservation_status_enum",'2026-07-12 08:34:04.495306+07'),
	 (1,NULL,NULL,5,'2026-08-01 07:00:00+07','EXPIRED'::public."reservation_status_enum",'2026-07-12 08:34:40.380133+07'),
	 (1,NULL,NULL,3,'2026-08-01 07:00:00+07','EXPIRED'::public."reservation_status_enum",'2026-07-12 09:13:01.785448+07'),
	 (1,1,1,5,'2026-07-13 09:20:01+07','COMPLETED'::public."reservation_status_enum",'2026-07-12 09:20:01.416472+07'),
	 (1,1,3,3,'2026-07-13 09:24:57+07','EXPIRED'::public."reservation_status_enum",'2026-07-12 09:24:57.954831+07'),
	 (1,1,NULL,2,'2026-07-13 10:59:38+07','HOLDING'::public."reservation_status_enum",'2026-07-12 10:59:38.724355+07'),
	 (1,1,NULL,1,'2026-07-13 17:53:05+07','HOLDING'::public."reservation_status_enum",'2026-07-12 17:53:05.793315+07');
