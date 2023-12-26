----------------------------------------------------------------
-- Insert default users sample data
----------------------------------------------------------------
INSERT INTO emn.mst_user (
  mst_user_id,
  user_name,
  user_password,
  administrative
) VALUES
  ('XXXXXXXX', '', '$2a$10$rse6B2kQ/qRHPoYZ9Bcyk.KNUtcaFNv/wZFjyuuCtfN5y3TdEwJSi', FALSE),
  ('00000000', 'administrator', '$2a$10$Ek3yuH8Sfh4Z.P63tIuX6.B54P3zSzE1hvGquzDjOx9Q3Y7Av4rj6', TRUE),
  ('00100001', 'katoryo', '$2a$10$5Cs0Yb/180HwVKGCBkUzgu2feaVVA8oqO6cB7dodxErW2JDvMNw56', FALSE),
  ('99999999', 'test', '$2a$10$n0hVT.iZUBZ2VCyUGo152O1GQBXuOXNt24mleDzFFbWFbQqeAnPPG', FALSE)
;

INSERT INTO emn.table_information (
  table_information_id
) VALUES
  ('XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX'),
  ('XXXXXXXX-XXXX-XXXX-XXXX-000000000001'), -- item's table information id
  ('XXXXXXXX-XXXX-XXXX-XXXX-000000000002'), -- item's table information id
  ('XXXXXXXX-XXXX-XXXX-XXXX-000000000003'), -- item's table information id
  ('XXXXXXXX-XXXX-XXXX-XXXX-000000000004'), -- item's table information id
  ('XXXXXXXX-XXXX-XXXX-XXXX-000000000005'), -- item's table information id
  ('XXXXXXXX-XXXX-XXXX-XXXX-000000000006'), -- item's table information id
  ('XXXXXXXX-XXXX-XXXX-XXXX-000000000007'), -- item's table information id
  ('XXXXXXXX-XXXX-XXXX-XXXX-000000000008'), -- item's table information id
  ('XXXXXXXX-XXXX-XXXX-XXXX-000000000009'), -- item's table information id
  ('XXXXXXXX-XXXX-XXXX-XXXX-000000000011'), -- item category's table information id
  ('XXXXXXXX-XXXX-XXXX-XXXX-000000000012'), -- item category's table information id
  ('XXXXXXXX-XXXX-XXXX-XXXX-000000000021'), -- item status's table information id
  ('XXXXXXXX-XXXX-XXXX-XXXX-000000000022'), -- item status's table information id
  ('XXXXXXXX-XXXX-XXXX-XXXX-000000000023'), -- item status's table information id
  ('XXXXXXXX-XXXX-XXXX-XXXX-000000000024'), -- item status's table information id
  ('XXXXXXXX-XXXX-XXXX-XXXX-000000000025'), -- item status's table information id
  ('XXXXXXXX-XXXX-XXXX-XXXX-000000000026'), -- item status's table information id
  ('XXXXXXXX-XXXX-XXXX-XXXX-000000000031'), -- item unit's table information id
  ('XXXXXXXX-XXXX-XXXX-XXXX-000000000032'), -- item unit's table information id
  ('XXXXXXXX-XXXX-XXXX-XXXX-000000000033'), -- item unit's table information id
  ('XXXXXXXX-XXXX-XXXX-XXXX-000000000034'), -- item unit's table information id
  ('XXXXXXXX-XXXX-XXXX-XXXX-000000000035'), -- item unit's table information id
  ('XXXXXXXX-XXXX-XXXX-XXXX-000000000036'), -- item unit's table information id
  ('XXXXXXXX-XXXX-XXXX-XXXX-000000000041'), -- item process's table information id
  ('XXXXXXXX-XXXX-XXXX-XXXX-000000000042'), -- item process's table information id
  ('XXXXXXXX-XXXX-XXXX-XXXX-000000000043'), -- item process's table information id
  ('XXXXXXXX-XXXX-XXXX-XXXX-000000000044'), -- item process's table information id
  ('XXXXXXXX-XXXX-XXXX-XXXX-000000000045'), -- item process's table information id
  ('XXXXXXXX-XXXX-XXXX-XXXX-000000000046'), -- item process's table information id
  ('XXXXXXXX-XXXX-XXXX-XXXX-000000000047'), -- item process's table information id
  ('XXXXXXXX-XXXX-XXXX-XXXX-000000000048'), -- item process's table information id
  ('XXXXXXXX-XXXX-XXXX-XXXX-000000000049'), -- item process's table information id
  ('XXXXXXXX-XXXX-XXXX-XXXX-00000000004a'), -- item process's table information id
  ('XXXXXXXX-XXXX-XXXX-XXXX-00000000004b'), -- item process's table information id
  ('XXXXXXXX-XXXX-XXXX-XXXX-00000000004c'), -- item process's table information id
  ('XXXXXXXX-XXXX-XXXX-XXXX-00000000004d'), -- item process's table information id
  ('XXXXXXXX-XXXX-XXXX-XXXX-00000000004e'), -- item process's table information id
  ('XXXXXXXX-XXXX-XXXX-XXXX-000000000051'), -- machine's table information id
  ('XXXXXXXX-XXXX-XXXX-XXXX-000000000052'), -- machine's table information id
  ('XXXXXXXX-XXXX-XXXX-XXXX-000000000053'), -- machine's table information id
  ('XXXXXXXX-XXXX-XXXX-XXXX-000000000054'), -- machine's table information id
  ('XXXXXXXX-XXXX-XXXX-XXXX-000000000055'), -- machine's table information id
  ('XXXXXXXX-XXXX-XXXX-XXXX-000000000056'), -- machine's table information id
  ('XXXXXXXX-XXXX-XXXX-XXXX-000000000057'), -- machine's table information id
  ('XXXXXXXX-XXXX-XXXX-XXXX-000000000058') -- machine's table information id
;
----------------------------------------------------------------
-- Insert sample data
----------------------------------------------------------------
INSERT INTO emn.bss_phone_number_list (
  bss_phone_number_list_id
) VALUES
  ('XXXXXXXXXX'),
  ('0000001001')
;

INSERT INTO emn.bss_machine_list (
  bss_machine_list_id
) VALUES
  ('XXXXXXXXXX'),
  ('1001000001'),
  ('2001000101'),
  ('2001000201')
;

INSERT INTO emn.bss_end_user_list (
  bss_end_user_list_id
) VALUES
  ('XXXXXXXXXX')
;

INSERT INTO emn.bss_consumption_list (
  bss_consumption_list_id
) VALUES
  ('XXXXXXXXXX'),
  ('4c8bda02-4b27-460a-b46f-d65b4e4f5e48')
;

INSERT INTO emn.mst_item_category (
  mst_item_category_id,
  category_name,
  table_information_id
) VALUES
  ('01', '加工品', 'XXXXXXXX-XXXX-XXXX-XXXX-000000000011'),
  ('02', '購入品', 'XXXXXXXX-XXXX-XXXX-XXXX-000000000021')
;

INSERT INTO emn.mst_item_status (
  mst_item_status_id,
  status_name,
  table_information_id
) VALUES
  ('00', '材料', 'XXXXXXXX-XXXX-XXXX-XXXX-000000000021'),
  ('01', '完成品', 'XXXXXXXX-XXXX-XXXX-XXXX-000000000022'),
  ('02', '半製品', 'XXXXXXXX-XXXX-XXXX-XXXX-000000000023'),
  ('03', '中間品', 'XXXXXXXX-XXXX-XXXX-XXXX-000000000024'),
  ('04', '部品', 'XXXXXXXX-XXXX-XXXX-XXXX-000000000025'),
  ('05', '仕掛品', 'XXXXXXXX-XXXX-XXXX-XXXX-000000000026')
;

INSERT INTO emn.mst_item_unit (
  mst_item_unit_id,
  unit_name,
  table_information_id
) VALUES
  ('01', 'mg', 'XXXXXXXX-XXXX-XXXX-XXXX-000000000031'),
  ('02', 'g', 'XXXXXXXX-XXXX-XXXX-XXXX-000000000032'),
  ('03', 'kg', 'XXXXXXXX-XXXX-XXXX-XXXX-000000000033'),
  ('04', 'ml', 'XXXXXXXX-XXXX-XXXX-XXXX-000000000034'),
  ('05', 'l', 'XXXXXXXX-XXXX-XXXX-XXXX-000000000035'),
  ('06', '個', 'XXXXXXXX-XXXX-XXXX-XXXX-000000000036')
;

INSERT INTO emn.mst_arrival_category (
  mst_arrival_category_id,
  arrival_category_name
) VALUES
  ('0001', '完納'),
  ('0002', '分納')
;

INSERT INTO emn.mst_operator (
  mst_operator_id,
  operator_name
) VALUES
  ('XXXXXXXX', 'default value'),
  ('00010001', '大石 真'),
  ('00010002', '岩崎 慶一'),
  ('00010003', '坂本 宏弥'),
  ('00020001', '五十嵐 雅宏'),
  ('00020002', '岡 宏')
;

INSERT INTO emn.phone_number_list (
  bss_phone_number_list_id,
  phone_number_list_no,
  phone_number
) VALUES 
  ('XXXXXXXXXX', 1, 'XX-XXXX-XXXX'),
  ('0000001001', 1, '0267-29-0001'),
  ('0000001001', 2, '0267-29-0002'),
  ('0000001001', 3, '0267-29-0003')
;
-- INSERT INTO emn.mst_address (
--   mst_address_id,
--   building,
--   street,
--   city,
--   prefecture,
--   postal_code,
--   phone_number_1
-- ) VALUES
--   ('XXXXXXXX', 'default value', '', '', '', '', ''),
--   ('00000010', '本社ビル', '1-1-1', '小諸市', '長野県', '384-0000', '0267-29-0001'),
--   ('00000011', '', '1-1-2', '小諸市', '長野県', '384-0000', '0267-29-0011'),
--   ('00000012', '', '1-1-3', '小諸市', '長野県', '384-0000', '0267-29-0012'),
--   ('00000013', '', '1-1-4', '小諸市', '長野県', '384-0000', '0267-29-0013'),
--   ('00000021', '', '1-2-1', '小諸市', '長野県', '384-0000', '0267-29-0021'),
--   ('00000022', '', '1-2-2', '小諸市', '長野県', '384-0000', '0267-29-0022'),
--   ('00000101', '本社工場', '1-99-1', '佐久市', '長野県', '384-0300', '0267-60-0001'),
--   ('00000201', '本社工場', '3-1-1', '上田市', '長野県', '386-0000', '0267-20-0001'),
--   ('00100001', 'NGNビル1階', '2-2-1', '長野市', '長野県', '380-0000', '0262-29-0001'),
--   ('00100002', '', '2-2-2', '長野市', '長野県', '380-0000', '0262-29-0001'),
--   ('00100003', '', '2-2-3', '松本市', '長野県', '390-0000', '0263-29-0001'),
--   ('00200001', 'SNJKビル3階', '1-1', '新宿区', '東京都', '101-9999', '03-3200-0001'),
--   ('01000001', '', '9-1', '中央区', '東京都', '111-9999', '03-3300-0001'),
--   ('01000002', '', '1-9', '東区', '大阪府', '530-0000', '03-3300-0001')
-- ;
INSERT INTO emn.mst_address (
  mst_address_id,
  building,
  street,
  city,
  prefecture,
  postal_code
) VALUES
  ('XXXXXXXX', 'default value', '', '', '', ''),
  ('00000010', '本社ビル', '1-1-1', '小諸市', '長野県', '384-0000'),
  ('00000011', '', '1-1-2', '小諸市', '長野県', '384-0000'),
  ('00000012', '', '1-1-3', '小諸市', '長野県', '384-0000'),
  ('00000013', '', '1-1-4', '小諸市', '長野県', '384-0000'),
  ('00000021', '', '1-2-1', '小諸市', '長野県', '384-0000'),
  ('00000022', '', '1-2-2', '小諸市', '長野県', '384-0000'),
  ('00000101', '本社工場', '1-99-1', '佐久市', '長野県', '384-0300'),
  ('00000201', '本社工場', '3-1-1', '上田市', '長野県', '386-0000'),
  ('00100001', 'NGNビル1階', '2-2-1', '長野市', '長野県', '380-0000'),
  ('00100002', '', '2-2-2', '長野市', '長野県', '380-0000'),
  ('00100003', '', '2-2-3', '松本市', '長野県', '390-0000'),
  ('00200001', 'SNJKビル3階', '1-1', '新宿区', '東京都', '101-9999'),
  ('01000001', '', '9-1', '中央区', '東京都', '111-9999'),
  ('01000002', '', '1-9', '東区', '大阪府', '530-0000'),
  ('10001001', '', '1-1', '長野市', '長野県', '380-9001'),
  ('10002001', '', '2-1', '松本市', '長野県', '390-9001')

;

INSERT INTO emn.mst_company (
  mst_company_id,
  company_name,
  mst_address_id
) VALUES
  ('XXXXXXXX', 'default value', 'XXXXXXXX'),
  ('0001', '小諸工業 株式会社', '00000010'),
  ('0002', '佐久総合化学 株式会社', '00000101'),
  ('0003', '上田技工 株式会社', '00000201'),
  ('0004', '長野技研 株式会社', '00100001'),
  ('0101', '東京工業商会 株式会社', '00200001'),
  ('1001', '東京工業機械 株式会社', '01000001'),
  ('1002', '大阪工機 株式会社', '01000002'),
  ('2001', '長野自動車 株式会社', '10001001'),
  ('3001', '松本オートモービル 株式会社', '10002001')
;

INSERT INTO emn.mst_warehouse (
  mst_warehouse_id,
  warehouse_name,
  mst_company_id,
  mst_address_id
) VALUES
  ('XXXXXXXX', 'default value', '0001', 'XXXXXXXX'),
  ('0001', '第一工場材料置場', '0001', '00000011'),
  ('0002', '第二工場材料置場', '0001', '00000012'),
  ('0003', '第一工場製品置場', '0001', '00000011'),
  ('0004', '第二工場製品置場', '0001', '00000012'),
  ('0005', '第一倉庫', '0001', '00000021'),
  ('0006', '第二倉庫', '0001', '00000022')
;

INSERT INTO emn.mst_factory (
  mst_factory_id,
  factory_name,
  mst_company_id,
  mst_address_id
) VALUES
  ('XXXXXXXX', 'default value', '0001', 'XXXXXXXX'),
  ('0001', '第一工場', '0001', '00000011'),
  ('0002', '第二工場', '0001', '00000012'),
  ('0003', '検査棟', '0001', '00000013'),
  ('0004', '本社工場', '0002', '00000101'),
  ('0005', '本社第一工場', '0003', '00000201'),
  ('0006', '長野工場', '0004', '00000012'),
  ('0007', '松本工場', '0004', '00000012')
;

INSERT INTO emn.mst_process (
  mst_process_id,
  process_name,
  mst_factory_id,
  table_information_id
) VALUES
  ('XXXXXXXX', 'default value', 'XXXXXXXX', 'XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX'),
  ('0001', '圧造', '0001', 'XXXXXXXX-XXXX-XXXX-XXXX-000000000041'),
  ('0002', '圧延', '0006', 'XXXXXXXX-XXXX-XXXX-XXXX-000000000042'),
  ('0003', '組付', '0007', 'XXXXXXXX-XXXX-XXXX-XXXX-000000000043'),
  ('0004', '転造', '0006', 'XXXXXXXX-XXXX-XXXX-XXXX-000000000044'),
  ('0005', 'タップ', '0002', 'XXXXXXXX-XXXX-XXXX-XXXX-000000000045'),
  ('0006', 'ネジ検査', '0003', 'XXXXXXXX-XXXX-XXXX-XXXX-000000000046'),
  ('0007', '外観検査', '0003', 'XXXXXXXX-XXXX-XXXX-XXXX-000000000047'),
  ('0008', '梱包', '0003', 'XXXXXXXX-XXXX-XXXX-XXXX-000000000048'),
  ('1011', '焼準', '0004', 'XXXXXXXX-XXXX-XXXX-XXXX-000000000049'),
  ('1012', '焼鈍', '0004', 'XXXXXXXX-XXXX-XXXX-XXXX-00000000004a'),
  ('1021', '三価クロメート', '0004', 'XXXXXXXX-XXXX-XXXX-XXXX-00000000004b'),
  ('1022', '三価黒クロメート', '0004', 'XXXXXXXX-XXXX-XXXX-XXXX-00000000004c'),
  ('1023', 'クロムメッキ', '0004', 'XXXXXXXX-XXXX-XXXX-XXXX-00000000004d'),
  ('1024', 'ユニクロメッキ', '0004', 'XXXXXXXX-XXXX-XXXX-XXXX-00000000004e')
;

INSERT INTO emn.mst_machine (
  mst_machine_id,
  machine_name,
  mst_factory_id,
  maker_id,
  table_information_id
) VALUES
  ('XXXXXXXX', 'default value', 'XXXXXXXX', 'XXXXXXXX', 'XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX'),
  ('0001001', 'A01', '0001', '1001', 'XXXXXXXX-XXXX-XXXX-XXXX-000000000051'),
  ('0001002', 'B01', '0001', '1001', 'XXXXXXXX-XXXX-XXXX-XXXX-000000000052'),
  ('0001003', 'B02', '0001', '1001', 'XXXXXXXX-XXXX-XXXX-XXXX-000000000053'),
  ('0001004', 'X01C', '0001', '1001', 'XXXXXXXX-XXXX-XXXX-XXXX-000000000054'),
  ('0001005', 'X02C', '0001', '1001', 'XXXXXXXX-XXXX-XXXX-XXXX-000000000055'),
  ('0001006', 'X03C', '0001', '1001', 'XXXXXXXX-XXXX-XXXX-XXXX-000000000056'),
  ('0002001', 'C1号機', '0002', '1002', 'XXXXXXXX-XXXX-XXXX-XXXX-000000000057'),
  ('0002002', 'C2号機', '0002', '1002', 'XXXXXXXX-XXXX-XXXX-XXXX-000000000058')
;

INSERT INTO emn.machine_list (
  bss_machine_list_id,
  machine_list_no,
  mst_machine_id
) VALUES
  ('XXXXXXXXXX', 1, 'XXXXXXXX'),
  ('1001000001', 1, '0001004'),
  ('1001000001', 2, '0001005'),
  ('1001000001', 3, '0001006'),
  ('2001000101', 1, '0001002'),
  ('2001000101', 2, '0001003'),
  ('2001000201', 1, '0001002'),
  ('2001000201', 2, '0001003')
;
--
-- insert unfinished products.
--
INSERT INTO emn.mst_item (
  mst_item_id,
  item_number,
  item_name,
  mst_item_category_id,
  mst_item_status_id,
  mst_item_unit_id,
  bss_machine_list_id,
  mst_warehouse_id,
  table_information_id
) VALUES
  ('90001001', '514000-0950', 'S45C φ9.50', '02', '00', '03', 'XXXXXXXXXX', '0001', 'XXXXXXXX-XXXX-XXXX-XXXX-000000000001'),
  ('90002001', '330-0478', 'S45C φ4.78', '02', '00', '03', 'XXXXXXXXXX', '0001', 'XXXXXXXX-XXXX-XXXX-XXXX-000000000002'),
  ('90002002', '4XS000578', 'SCM440 15x20', '02', '00', '03', 'XXXXXXXXXX', '0001', 'XXXXXXXX-XXXX-XXXX-XXXX-000000000003'),
  ('10010000', 'NIA-100-1000', 'NIA Bolt 100-1000', '01', '03', '06', '1001000001', '0003', 'XXXXXXXX-XXXX-XXXX-XXXX-000000000004'),
  ('20010001', 'JPI100-01', 'JPI10001', '01', '04', '06', '2001000101', '0003', 'XXXXXXXX-XXXX-XXXX-XXXX-000000000005'),
  ('20010002', 'JPI100-02', 'JPI10002', '01', '04', '06', '2001000201', '0003', 'XXXXXXXX-XXXX-XXXX-XXXX-000000000006')
;

--
-- insert finished products.
--
INSERT INTO emn.mst_item (
  mst_item_id,
  item_number,
  item_name,
  mst_item_category_id,
  mst_item_status_id,
  mst_item_unit_id,
  bss_machine_list_id,
  mst_warehouse_id,
  mst_delivery_destination_id,
  table_information_id
) VALUES
  ('10010001', 'NIA-100-1000-SLV', 'NIA Bolt 100-1000 SLV', '01', '01', '06', 'XXXXXXXXXX', '0005', '0002', 'XXXXXXXX-XXXX-XXXX-XXXX-000000000007'),
  ('10010002', 'NIA-100-1000-BLK', 'NIA Bolt 100-1000 BLK', '01', '01', '06', 'XXXXXXXXXX', '0005', '0002', 'XXXXXXXX-XXXX-XXXX-XXXX-000000000008'),
  ('20010000', 'JPI100', 'JPI100', '01', '01', '06', 'XXXXXXXXXX', '0006', '0004', 'XXXXXXXX-XXXX-XXXX-XXXX-000000000009')
;

INSERT INTO emn.mst_process_order (
  mst_item_id,
  process_order_no,
  is_selectable,
  is_parallel,
  mst_process_id
) VALUES
  ('10010000', 1, TRUE, FALSE, '0001'), -- 圧造
  ('10010000', 2, TRUE, FALSE, '0005'), -- タップ
  ('10010000', 3, FALSE, FALSE, '1012'), -- 焼鈍
  ('10010001', 1, FALSE, FALSE, '1021'), -- 三価クロメート
  ('10010001', 2, TRUE, TRUE, '0006'), -- ネジ検査: 外観検査と並列処理可能 かつ 社内/外注選択可能
  ('10010001', 3, TRUE, TRUE, '0007'), -- 外観検査: 外観検査と並列処理可能 かつ 社内/外注選択可能
  ('10010001', 4, FALSE, FALSE, '0008'), -- 梱包
  ('10010002', 1, FALSE, FALSE, '1022'), -- 三価黒クロメート
  ('10010002', 2, TRUE, TRUE, '0006'), -- ネジ検査: 外観検査と並列処理可能 かつ 社内/外注選択可能
  ('10010002', 3, TRUE, TRUE, '0007'), -- 外観検査: 外観検査と並列処理可能 かつ 社内/外注選択可能
  ('10010002', 4, FALSE, FALSE, '0008'), -- 梱包
  ('20010001', 1, TRUE, FALSE, '0001'), -- 圧造
  ('20010001', 2, TRUE, FALSE, '0005'), -- タップ
  ('20010001', 3, FALSE, FALSE, '1012'), -- 焼鈍
  ('20010001', 4, FALSE, FALSE, '1023'), -- クロムメッキ
  ('20010002', 1, TRUE, FALSE, '0001'), -- 圧造
  ('20010002', 2, TRUE, FALSE, '0004'), -- 転造
  ('20010002', 3, TRUE, FALSE, '0005'), -- タップ
  ('20010002', 4, FALSE, FALSE, '1023'), -- クロムメッキ
  ('20010000', 1, TRUE, FALSE, '0003'), -- 組付
  ('20010000', 2, TRUE, TRUE, '0006'), -- ネジ検査: 外観検査と並列処理可能 かつ 社内/外注選択可能
  ('20010000', 3, TRUE, TRUE, '0007'), -- 外観検査: 外観検査と並列処理可能 かつ 社内/外注選択可能
  ('20010000', 4, FALSE, FALSE, '0008') -- 梱包
;

INSERT INTO emn.mst_bom (
  mst_sub_item_id,
  bom_no,
  mst_basic_item_id,
  required_qty,
  mst_item_unit_id
) VALUES
  ('10010000', 1, '90001001', 8.90, '02'),
  ('20010001', 1, '90002002', 4.30, '02'),
  ('20010002', 1, '90002001', 14.50, '02'),
  ('10010001', 1, '10010000', 1, '06'),
  ('10010002', 1, '10010000', 1, '06'),
  ('20010000', 1, '20010001', 1, '06'),
  ('20010000', 2, '20010002', 2, '06')
;

INSERT INTO emn.trn_production_instruction (
  trn_production_instruction_id,
  mst_item_id,
  mst_process_id,
  mst_machine_id,
  mst_operator_id,
  required_qty,
  start_to_produce,
  end_producing
) VALUES
  ('01a2c8b4-e3ec-4df5-a027-b7ffe5c6ed5d', '10010000', '0001', '0001001', '00010003', 50000.00, '2021-05-10', '2021-05-14')
;

INSERT INTO emn.trn_production (
  trn_production_id,
  trn_production_instruction_id,
  mst_item_id,
  mst_process_id,
  lot,
  branch,
  mst_machine_id,
  mst_operator_id,
  bss_consumption_list_id,
  non_defective_qty,
  defective_qty,
  suspended_qty,
  produced_at,
  information
) VALUES
  ('eb2aa249-3614-4833-b979-99c87c4189a8', '01a2c8b4-e3ec-4df5-a027-b7ffe5c6ed5d', '10010000', '0001', '2105110001', '01', '0001001', '00010003', '4c8bda02-4b27-460a-b46f-d65b4e4f5e48', 15000.00, 450.00, 1000.00, '2021-05-11', '破損により金型を交換')
;

INSERT INTO emn.consumption_list (
  bss_consumption_list_id,
  consumption_list_no,
  mst_warehouse_id,
  mst_item_id,
  lot,
  non_defective_qty,
  transaction_type
) VALUES
  ('4c8bda02-4b27-460a-b46f-d65b4e4f5e48', 1, '0001', '90001001', '210426M001', 146850.00, 'production')
;

INSERT INTO emn.bss_regulate_inventory (
  bss_regulate_inventory_id
) VALUES
  ('XXXXXXXXXX')
;

INSERT INTO emn.inventory (
  mst_item_id,
  mst_process_id,
  lot,
  branch,
  mst_warehouse_id,
  non_defective_qty,
  defective_qty,
  suspended_qty,
  expiration_date
) VALUES
  ('90001001', 'XXXXXXXX', '210419M001', '', '0001', 95000.00, 0.00, 0.00, '2022-04-19'),
  ('90001001', 'XXXXXXXX', '210426M001', '', '0001', 395000.00, 0.00, 0.00, '2022-04-26')
;