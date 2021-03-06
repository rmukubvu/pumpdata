insert into public.pump_company (id, pump_id, company_id, created_date) values (1, 1, 1, 1503683839);
insert into public.pump_sensor (id, type_id, pump_id, s_value, created_date) values (1, 1, 1, 3000, 1592701444);
insert into public.pump_sensor (id, type_id, pump_id, s_value, created_date) values (2, 10, 1, 3000, 1592701842);
insert into public.pump_sensor (id, type_id, pump_id, s_value, created_date) values (3, 10, 1, 4000, 1592701993);
insert into public.pump_sensor (id, type_id, pump_id, s_value, created_date) values (4, 8, 1, 40, 1592702022);
insert into public.pump_sensor (id, type_id, pump_id, s_value, created_date) values (5, 8, 1, 40, 1592732294);
insert into public.pump_sensor (id, type_id, pump_id, s_value, created_date) values (6, 8, 1, 40, 1592732368);
insert into public.pump_sensor (id, type_id, pump_id, s_value, created_date) values (7, 8, 1, 40, 1592732835);
insert into public.pump_sensor (id, type_id, pump_id, s_value, created_date) values (8, 8, 1, 40, 1592732975);
insert into public.pump_sensor (id, type_id, pump_id, s_value, created_date) values (9, 8, 1, 40, 1592755472);
insert into public.pump_sensor (id, type_id, pump_id, s_value, created_date) values (10, 8, 1, 809, 1592755824);
insert into public.pump_sensor (id, type_id, pump_id, s_value, created_date) values (11, 8, 1, 657, 1592755851);
insert into public.pump_sensor (id, type_id, pump_id, s_value, created_date) values (12, 5, 1, 657, 1592845973);
insert into public.pump_type (id, name) values (1, 'Electrical');
insert into public.pump_type (id, name) values (2, 'Mechanical');
insert into public.pump_type (id, name) values (3, 'Actuator');
insert into public.pump_type (id, name) values (4, 'Test Jig');
insert into public.pump (id, type_id, serial_number, nick_name, lat, lng, created_date) values (1, 1, 123456789, 'Fourways-Pump', -26.018908, 28.004368, 1590957563);
insert into public.pump (id, type_id, serial_number, nick_name, lat, lng, created_date) values (3, 1, 135792468, 'Benoni Mall', -26.153149, 28.315287, 1591039477);
insert into public.pump (id, type_id, serial_number, nick_name, lat, lng, created_date) values (5, 1, 135792469, 'Springs Mall', -26.184183, 28.325602, 1591039575);
insert into public.pump (id, type_id, serial_number, nick_name, lat, lng, created_date) values (6, 1, 135792470, 'Village View Mall', -26.164742, 28.154837, 1591039769);
insert into public.sensor_alarms_contacts (id, company_id, cell_phone, email_address, created_date) values (1, 1, 27719084111, 'rmukubvu@gmail.com', 1592842350);
insert into public.sensor_alarms_contacts (id, company_id, cell_phone, email_address, created_date) values (2, 1, 27604050091, 'mukubvua@gmail.com', 1592842424);
insert into public.sensor_alarms (id, type_id, min_value, max_value, alert_message, created_date) values (3, 3, 30, 90, 'need quick attention', 1592728176);
insert into public.sensor_alarms (id, type_id, min_value, max_value, alert_message, created_date) values (1, 1, 3000, 4000, 'need quick attention', 1592728176);
insert into public.sensor_alarms (id, type_id, min_value, max_value, alert_message, created_date) values (8, 8, 500, 800, 'need quick attention', 1592728176);
insert into public.sensor_alarms (id, type_id, min_value, max_value, alert_message, created_date) values (13, 13, 15, 35, 'need quick attention', 1592728176);
insert into public.sensor_alarms (id, type_id, min_value, max_value, alert_message, created_date) values (6, 6, 19, 56, 'need quick attention', 1592728176);
insert into public.sensor_alarms (id, type_id, min_value, max_value, alert_message, created_date) values (7, 7, 50, 100, 'need quick attention', 1592728176);
insert into public.sensor_alarms (id, type_id, min_value, max_value, alert_message, created_date) values (15, 15, 14, 45, 'need quick attention', 1592728176);
insert into public.sensor_alarms (id, type_id, min_value, max_value, alert_message, created_date) values (14, 14, 56, 78, 'need quick attention', 1592728176);
insert into public.sensor_alarms (id, type_id, min_value, max_value, alert_message, created_date) values (11, 11, 30, 67, 'need quick attention', 1592728176);
insert into public.sensor_alarms (id, type_id, min_value, max_value, alert_message, created_date) values (2, 2, 25, 40, 'need quick attention', 1592728176);
insert into public.sensor_alarms (id, type_id, min_value, max_value, alert_message, created_date) values (12, 12, 44, 88, 'need quick attention', 1592728176);
insert into public.sensor_alarms (id, type_id, min_value, max_value, alert_message, created_date) values (9, 9, 90, 180, 'need quick attention', 1592728176);
insert into public.sensor_alarms (id, type_id, min_value, max_value, alert_message, created_date) values (5, 5, 50, 200, 'need quick attention', 1592728176);
insert into public.sensor_alarms (id, type_id, min_value, max_value, alert_message, created_date) values (4, 4, 60, 120, 'need quick attention', 1592728176);
insert into public.sensor_alarms (id, type_id, min_value, max_value, alert_message, created_date) values (10, 10, 45, 78, 'need quick attention', 1592728176);
insert into public.sensor_data (id, serial_number, type_id, s_value, update_date, type_text) values (1, 123456789, 1, 3000, 1592701444, 'magnetic_pickup');
insert into public.sensor_data (id, serial_number, type_id, s_value, update_date, type_text) values (4, 123456789, 10, 4000, 1592701993, 'low_engine_water_level');
insert into public.sensor_data (id, serial_number, type_id, s_value, update_date, type_text) values (6, 123456789, 8, 657, 1592755851, 'fuel_level_backup');
insert into public.sensor_data (id, serial_number, type_id, s_value, update_date, type_text) values (14, 123456789, 5, 657, 1592845973, 'block_temperature');
