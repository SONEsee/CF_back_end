INSERT INTO public.order_status_logs (order_id,from_status,to_status,changed_by_type,changed_by_id,note,changed_at) VALUES
	 (1,NULL,'UNPAID','STAFF'::public."changed_by_type_enum",1,'ສ້າງອໍເດີ','2026-07-12 09:20:01.416472+07'),
	 (1,'UNPAID','PAYMENT_PENDING_VERIFY','STAFF'::public."changed_by_type_enum",1,'ອັບໂຫຼດສະລິບແລ້ວ','2026-07-12 09:24:20.093154+07'),
	 (1,'PAYMENT_PENDING_VERIFY','PAID','STAFF'::public."changed_by_type_enum",1,'ຢືນຢັນສະລິບແລ້ວ','2026-07-12 09:24:20.103971+07'),
	 (3,NULL,'UNPAID','STAFF'::public."changed_by_type_enum",1,'ສ້າງອໍເດີ','2026-07-12 09:24:57.954831+07'),
	 (3,'UNPAID','CANCELLED','STAFF'::public."changed_by_type_enum",1,'ລູກຄ້າຍົກເລີກ','2026-07-12 09:25:12.90825+07');
