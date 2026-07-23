INSERT INTO public.discounts (shop_id,code,discount_type,discount_value,min_order,usage_limit,used_count,start_at,end_at,is_active) VALUES
	 (3,'SAVE10','PERCENT'::public."discount_type_enum",10.00,100000.00,100,1,NULL,NULL,true);
