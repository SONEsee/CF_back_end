INSERT INTO public.orders (shop_id,customer_id,live_session_id,discount_id,order_number,current_status,items_total_amount,discount_amount,shipping_fee,net_payable_amount,note,created_at,updated_at) VALUES
	 (3,1,NULL,1,'ORD-3-20260712092001-2173','PAID'::public."order_status_enum",475000.00,47500.00,15000.00,442500.00,NULL,'2026-07-12 09:20:01.416472+07','2026-07-12 09:24:20.114124+07'),
	 (3,1,NULL,NULL,'ORD-3-20260712092457-1171','CANCELLED'::public."order_status_enum",285000.00,0.00,0.00,285000.00,NULL,'2026-07-12 09:24:57.954831+07','2026-07-12 09:25:12.911723+07');
