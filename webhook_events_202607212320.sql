INSERT INTO public.webhook_events (social_account_id,event_type,raw_payload,processed,received_at) VALUES
	 (1,'message_received','{"text": "hello", "sender": "123"}',true,'2026-07-12 10:06:25.248581+07'),
	 (1,'page','{"entry": [{"id": "fbpage-12345", "time": 123}], "object": "page"}',false,'2026-07-12 17:19:49.813183+07'),
	 (2,'line_event','{"events": [{"type": "message"}], "destination": "line-channel-999"}',false,'2026-07-12 17:25:03.042398+07');
