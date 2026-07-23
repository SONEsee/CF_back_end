INSERT INTO public.chat_messages (conversation_id,sender_type,message_type,message_body,attachment_url,is_read,sent_at) VALUES
	 (1,'CUSTOMER'::public."sender_type_enum",'TEXT'::public."message_type_enum",'ສະບາຍດີ ຢາກຖາມລາຄາເສື້ອຢືດ',NULL,false,'2026-07-12 09:58:31.086597+07'),
	 (1,'CUSTOMER'::public."sender_type_enum",'TEXT'::public."message_type_enum",'ມີໄຊສ໌ L ບໍ່?',NULL,false,'2026-07-12 09:58:31.113382+07'),
	 (1,'STAFF_AGENT'::public."sender_type_enum",'TEXT'::public."message_type_enum",'ມີຄ່ະ ລາຄາ 95,000 ກີບ',NULL,false,'2026-07-12 09:59:09.108965+07'),
	 (2,'CUSTOMER'::public."sender_type_enum",'TEXT'::public."message_type_enum",'ຍັງຢູ່ບໍ່?',NULL,false,'2026-07-12 09:59:28.898159+07');
