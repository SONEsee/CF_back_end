INSERT INTO public.comment_intents (comment_raw_id,customer_id,matched_product_variant_id,parsed_qty,intent_status,processed_at) VALUES
	 (1,1,1,2,'CF_SUCCESS'::public."intent_status_enum",'2026-07-12 10:59:38.724355+07'),
	 (3,1,1,999,'OUT_OF_STOCK'::public."intent_status_enum",'2026-07-12 11:00:22.206935+07'),
	 (4,1,NULL,NULL,'INVALID_CODE'::public."intent_status_enum",'2026-07-12 11:00:59.067703+07'),
	 (5,1,1,1,'CF_SUCCESS'::public."intent_status_enum",'2026-07-12 17:53:05.793315+07'),
	 (6,NULL,NULL,NULL,'INVALID_CODE'::public."intent_status_enum",'2026-07-12 17:54:29.583908+07');
