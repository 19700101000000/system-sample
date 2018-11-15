use `ih2018_db`;

/*社員テーブル*/
INSERT INTO `employee`(
    `id`, 
    `password`, 
    `name`, 
    `sex`
) VALUES (
    'GA18A01', 
    '1234567890', 
    '八木 明日香',
    1 
),(
    'SL18A01', 
    '1234567890', 
    '中尾 陽菜', 
    0
);

 /*車テーブル*/ 
INSERT INTO `car_information` (
    `name`,
    `type`,
    `formula_year`,
    `displacement`,
    `model_year`,
    `grade`,
    `fuel`,
    `dealer`,
    `parallel`,
    `right_handle`,
    `left_handle`,
    `model`,
    `ridingcapacity`,
    `drivesystem`,
    `doors`,
    `shape`,
    `carbody_no`,
    `loadingcapacity`,
    `mileage`,
    `exteriorcolor`,
    `exteriorcolor_no`,
    `interiorcolor`,
    `interiorcolor_no`,
    `newcar_warrantycard`,
    `instructionmanual`,
    `mission_shift`,
    `mission_atormt`,
    `mission_fast`,
    `equipment_ps`,
    `equipment_pw`,
    `equipment_abs`,
    `equipment_aw`,
    `equipment_cassette`,
    `equipment_cd`,
    `equipment_md`,
    `equipment_tv`,
    `equipment_navi`,
    `equipment_leathersheet`,
    `airbag_single`,
    `airbag_w`,
    `airbag_tv`,
    `airbag_navi`,
    `airbag_aircondition`,
    `airbag_audio`,
    `aircondition`,
    `sunroof`,
    `8number_type`,
    `nox_deadline`,
    `inspectiondeadline`,
    `registration_no_plate`,
    `registration_no_classification`,
    `registration_no_kana`,
    `registration_no_serial`,
    `name_change_date`,
    `car_history`,
    `model_designation_no`,
    `classification_entry_no`
    ) VALUES 
        ('トヨタ・86', 999, '28', '1998', '99', 'XXXXXXXXXXXXXXXXXXXX', 1, 0, 0, 1, 0, 'DBA-ZN6', 4, 1, 2, '99', '84917394672263496512', '1230', '60000', 'シルバー', '61K', 'ブラック', 'D4S', 0, 0, 6, 0, 6, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, '無', 599, '2019-01-31', '2019-01-31', 'ナニワ', '386', 'ハ', '8635', '2019-01-31', 1, '22317', '111'),
        ('スズキ・カプチーノ', 999, '8', '660', '95', 'スペースグレード', 1, 0, 0, 1, 0, 'E-EA21R', 2, 1, 2, '99', '94678723496092373512', '700', '990000', 'オレンジ', 'A7K', 'シルバー', '2BU', 0, 0, 5, 0, 5, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4, '無', 330, '2020-03-31', '2020-03-31', '大阪', '336', 'ロ', '8762', '2020-03-31', 1, '24397', '111'),
        ('トヨタ・ヘルファイヤ', 999, '25', '2500', '15', '2.5X', 1, 0, 0, 1, 0, 'DBA-AGH35W', 8, 1, 5, '99', '321417523905463496512', '1230', '60000', 'ブラック', 'A7K', 'グレー', 'F4S', 0, 0, 1, 1, 1, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 3, '無', 599, '2021-10-31', '2021-10-31', '筑豊', '156', 'ヘ', '2342', '2021-10-31', 1, '19346', '111'),
        ('N-BOX', 999, '29', '660', '17', 'G・L', 1, 0, 0, 1, 0, 'DBA-JF3', 4, 0, 5, '99', '72266491334951789462', '950', '300', 'イエロー', 'Y70P', 'ベージュブラウン', 'NH875P', 0, 0, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4, '無', 300, '2021-09-30', '2021-09-30', '兵庫', '306', 'ニ', '1329', '2021-09-30', 1, '29547', '111'
    );

 /*車修理場所名テーブル*/ 
INSERT INTO `car_information_repair_type`(
    `name`
) VALUES 
   ('左前タイヤ'),('右前タイヤ'),('左後タイヤ'),('右後タイヤ'),
   ('スペアタイヤ'),('Fバンパー左'),('左フェンダー前'),('左フェンダー中央①'),
   ('左フェンダー中央下'),('左フェンダー中央②'),('左フェンダー後'),('Rバンパー左'),
   ('Fバンパー中'),('ボンネット'),('前フロントガラス'),('ルーフ'),
   ('後フロントガラス'),('リアフード'),('Rバンパー中'),('Fバンパー右'),
   ('右側面前'),('右側面中央①'),('右側面中央下'),('右側面中央②'),
   ('左側面後'),('Rバンパー右'
);

/*車修理場所テーブル*/
INSERT INTO `car_information_repair`(
    `car_information_id`, 
    `car_information_repair_type_id`, 
    `memo`
) VALUES (
    '1',
    '1',
    'タイヤが擦り切れていた。'
),(
    '1',
    '11',
    '微かに傷があった。'
);

/*金融機関テーブル*/
INSERT INTO `financial_institution`(
    `id`, 
    `name`
) VALUES (
    '0001', 
    'みずほ銀行 （ﾐｽﾞﾎ）'
),(
    '0005', 
    '三菱ＵＦＪ銀行 （ﾐﾂﾋﾞｼﾕ-ｴﾌｼﾞｴｲ）'
),(
    '0009', 
    '三井住友銀行 （ﾐﾂｲｽﾐﾄﾓ）'
),(
    '0010', 
    'りそな銀行 （ﾘｿﾅ）'
),(
    '9900', 
    'ゆうちょ銀行 （ﾕｳﾁﾖ）'
);

/*金融機関支店テーブル*/
INSERT INTO `financial_institution_branch`(
    `financial_institution_id`,
    `branch_id`, 
    `name` 
) VALUES (
    '0001',
    '440',
    '大阪支店 （ｵｵｻｶ）'
),(
    '0001',
    '502',
    '大阪中央支店 （ｵｵｻｶﾁﾕｳｵｳ）'
),(
    '0005',
    '067',
    '大阪駅前支店 （ｵｵｻｶｴｷﾏｴ）'
),(
    '0009',
    '941',
    '大阪第一支店 （ｵｵｻｶﾁﾕｳｵｳ）'
),(
    '0010',
    '153',
    '荻窪支店 （ｵｷﾞｸﾎﾞ）'
),(
    '9900',
    '119',
    '一一九 （ｲﾁｲﾁｷﾕｳ）'
);

/*仕入先テーブル*/
INSERT INTO `suppliers`( 
    `financial_institution_branch_id`,
    `name`, 
    `address`, 
    `tel`, 
    `account_deposittype`, 
    `account_no`, 
    `account_holder`
) VALUES (
    '1',
    'ＵＳＳ大阪会場 ＵＳＳオートオークション', 
    '〒555-0041 大阪府大阪市西淀川区中島2丁目7-106', 
    '06-6476-1000', 
    '1', 
    '0123456', 
    'ハル タロウ'
),(
    '2',
    'ガリバー大阪中央出張車査定センター', 
    '〒537-0022 大阪府東成区中本1丁目11-24', 
    '0120-221-616', 
    '1', 
    '7890123', 
    'ナツ タロウ'
),(
    '3',
    'トヨタオートオークション近畿', 
    '〒572-0076 大阪府寝屋川市仁和寺本町3丁目1-22', 
    '072-826-3486', 
    '1', 
    '4567890', 
    'アキ タロウ'
);

/*顧客テーブル*/ 
INSERT INTO 'client'(
    `name`, 
    `tel`, 
    `address`, 
    `financial_institution_branch_id`, 
    `account_deposittype`,
    `account_no`, 
    `account_holder`
) VALUES (
 '高田 穂乃花', 
 '07-6576-1000', 
 '〒555-0041 大阪府大阪市西淀川区中島2丁目1-2', 
 '1', 
 '1',
 '0123456', 
 'タカダ ホノカ'
),(
 '前川 拓己', 
 '17-6576-1010', 
 '〒555-0041 大阪府大阪市西淀川区中島2丁目1-2', 
 '3', 
 '1',
 '0123456', 
 'マエカワ タクミ'
),(
 '河田 英毅', 
 '12-6676-1063', 
 '〒555-0041 大阪府大阪市西淀川区中島2丁目1-2', 
 '4', 
 '1',
 '0123456', 
 'カワタ ヒデキ'
);

/*買注文受注テーブル*/
INSERT INTO `buy_orders`( 
    `client_id`,
    `no`,
    `estimate_flag`,
    `manufacturer`,
    `car_name`,
    `car_model_year`,
    `car_budget`,
    `car_exterior_color`,
    `car_interior_color`,
    `car_model_no`,
) VALUES (
    '1',
    '1',
    '0',
    '0',
    'N-BOX',
    '25',
    '300,000',
    '0',
    '12',
    'DAA-ZWR80G'
),(
    '2',
    '1',
    '1',
    '1',
    'N-BOX',
    '26',
    '300,000',
    '1',
    '11',
    'DBA-ZWR81G'
),(
    '3',
    '1',
    '1',
    '2',
    'N-BOX',
    '25',
    '300,000',
    '2',
    '2',
    'AAA-ZQR80B'
),(
    '1',
    '2',
    '1',
    '0',
    'N-BOX',
    '27',
    '300,000',
    '5',
    '6',
    'DBA-ZNM80G'
),(
    '2',
    '2',
    '1',
    '1',
    'N-BOX',
    '28',
    '300,000',
    '9',
    '10',
    'LAA-ZLP81G'
);

/*買注文仕入テーブル*/
INSERT INTO `buy_purchase`(
    `buy_orders_id`, 
    `purchase_date`, 
    `suppliers_id`, 
    `purchase_amount`, 
    `expense`, 
    `car_name_deadline`, 
    `purchase_cost`, 
    `consumption_tax`, 
    `sales_date`, 
    `listing_no`, 
    `fee`, 
    `total_sales`, 
    `net_income`, 
    `tax_included`, 
    `car_name_deadline_finish`, 
    `purchase_availability`, 
    `document_flag`, 
    `payment_flag`, 
    `delivery_date`,
    `delivery_flag`,
    `car_information_id`
) VALUES (
    '2', 
    '', 
    '1', 
    '', 
    '', 
    '', 
    '', 
    '', 
    '', 
    '', 
    '', 
    '', 
    '', 
    '', 
    '', 
    '1', 
    '0', 
    '0', 
    '', 
    '1',
    '0' 
);

/*買注文請求テーブル*/ 
INSERT INTO `buy_claim`(
    `buy_purchase_id`, 
    `deadline`, 
    `billed`,
    `from_date`,
    `to_date`,
    `other_banks_ticket`,
    `payment_classification`,
    `edi_information`
) VALUES (
    '2', 
    '2018-11-30', 
    '0',
    from_date CHAR(6),
    to_date CHAR(6),
    other_banks_ticket CHAR(10),
    payment_classification CHAR(1),
    ''
),(
    '3', 
    '2018-11-30', 
    '0',
    from_date CHAR(6),
    to_date CHAR(6),
    other_banks_ticket CHAR(10),
    payment_classification CHAR(1),
    ''
),(
    '4', 
    '2018-11-30', 
    '0',
    from_date CHAR(6),
    to_date CHAR(6),
    other_banks_ticket CHAR(10),
    payment_classification CHAR(1),
    ''
);

/*買注文回収テーブル*/
INSERT INTO `buy_recovery`(
    `buy_claim_id`, 
    `recovery_price`, 
    `recovery_date`
) VALUES (
    '2', 
    '300,000', 
    '2018-10-25'
),(
    '3', 
    '130,000', 
    '2018-10-25'
);

/*変更ログテーブル*/
INSERT INTO `log_change`(
    `data_type`,
    `data_id`,
    `employee_id`,
    `data_comment`
) VALUES 
    ('0','1','SL18A01','新規登録'),('0','2','SL18A01','新規登録'),
    ('0','1','SL18A01','car_budget,変更'),
    ('0','3','SL18A01','新規登録'),('0','4','SL18A01','新規登録'),
    ('0','5','SL18A01','新規登録'),('1','1','SL18A01','新規登録'),
    ('1','2','SL18A01','新規登録'),('1','3','SL18A01','新規登録'),
    ('1','4','SL18A01','新規登録'),('2','1','SL18A01','新規登録'),
    ('2','2','SL18A01','新規登録'),('2','3','SL18A01','新規登録'),
    ('3','1','SL18A01','新規登録'),('3','2','SL18A01','新規登録'
);
