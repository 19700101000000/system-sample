
/*社員テーブル*/
CREATE TABLE `employee`(
 `id` CHAR(7), 
 `password` VARCHAR(255), 
 `name` VARCHAR(50), 
 `sex` BOOLEAN, 
 PRIMARY KEY(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*車テーブル*/ 
CREATE TABLE `car_information`(
 `id` INT AUTO_INCREMENT, 
 `name` VARCHAR(15), 
 `type` INT, 
 `formula_year` CHAR(2), 
 `displacement` CHAR(4), 
 `model_year` CHAR(2), 
 `grade` VARCHAR(20), 
 `fuel` BOOLEAN, 
 `dealer` BOOLEAN, 
 `parallel` BOOLEAN, 
 `right_handle` BOOLEAN, 
 `left_handle` BOOLEAN, 
 `model` VARCHAR(16), 
 `ridingcapacity` INT, 
 `drivesystem` BOOLEAN, 
 `doors` INT, 
 `shape` CHAR(2), 
 `carbody_no` VARCHAR(20), 
 `loadingcapacity` CHAR(6), 
 `mileage` CHAR(6), 
 `exteriorcolor` VARCHAR(10), 
 `exteriorcolor_no` VARCHAR(10), 
 `interiorcolor` VARCHAR(10), 
 `interiorcolor_no` VARCHAR(10), 
 `newcar_warrantycard` BOOLEAN, 
 `instructionmanual` BOOLEAN, 
 `mission_shift` INT, 
 `mission_atormt` BOOLEAN, 
 `mission_fast` INT, 
 `equipment_ps` BOOLEAN, 
 `equipment_pw` BOOLEAN, 
 `equipment_abs` BOOLEAN, 
 `equipment_aw` BOOLEAN, 
 `equipment_cassette` BOOLEAN, 
 `equipment_cd` BOOLEAN, 
 `equipment_md` BOOLEAN, 
 `equipment_tv` BOOLEAN, 
 `equipment_navi` BOOLEAN, 
 `equipment_leathersheet` BOOLEAN, 
 `airbag_single` BOOLEAN, 
 `airbag_w` BOOLEAN, 
 `airbag_tv` BOOLEAN, 
 `airbag_navi` BOOLEAN, 
 `airbag_aircondition` BOOLEAN, 
 `airbag_audio` BOOLEAN, 
 `aircondition` INT, 
 `sunroof` VARCHAR(10), 
 `8number_type` INT, 
 `nox_deadline` DATE, 
 `inspectiondeadline` DATE, 
 `registration_no_plate` VARCHAR(8), 
 `registration_no_classification` CHAR(3), 
 `registration_no_kana` CHAR(1), 
 `registration_no_serial` CHAR(4), 
 `name_change_date` DATE, 
 `car_history` INT, 
 `model_designation_no` CHAR(5), 
 `classification_entry_no` CHAR(3), 
 PRIMARY KEY(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*車修理場所名テーブル*/ 
CREATE TABLE `car_information_repair_type`(
 `id` INT AUTO_INCREMENT, 
 `name` VARCHAR(30), 
 PRIMARY KEY(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*車修理場所テーブル*/
CREATE TABLE `car_information_repair`(
 `car_information_id` INT, 
 `car_information_repair_type_id` INT, 
 `memo` TEXT, 
 PRIMARY KEY(`car_information_id`,`car_information_repair_type_id`),
 FOREIGN KEY(`car_information_id`) REFERENCES `car_information`(`id`),
 FOREIGN KEY(`car_information_repair_type_id`) REFERENCES `car_information_repair_type`(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*金融機関テーブル*/
CREATE TABLE `financial_institution`(
 `id` CHAR(4), 
 `name` VARCHAR(50),
 PRIMARY KEY(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*金融機関支店テーブル*/
CREATE TABLE `financial_institution_branch`(
 `id` INT AUTO_INCREMENT, 
 `financial_institution_id` CHAR(4),
 `branch_id` CHAR(3), 
 `name` VARCHAR(50), 
 PRIMARY KEY(`id`),
 FOREIGN KEY(`financial_institution_id`) REFERENCES `financial_institution`(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*仕入先テーブル*/
CREATE TABLE `suppliers`(
 `id` INT AUTO_INCREMENT, 
 `financial_institution_branch_id` INT,
 `name` VARCHAR(50), 
 `address` VARCHAR(70), 
 `tel` VARCHAR(20), 
 `account_deposittype` CHAR(1), 
 `account_no` CHAR(7), 
 `account_holder` VARCHAR(40),
 PRIMARY KEY(`id`),
 FOREIGN KEY(`financial_institution_branch_id`) REFERENCES `financial_institution_branch`(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*顧客テーブル*/ 
CREATE TABLE `client`(
 `id` INT AUTO_INCREMENT, 
 `name` VARCHAR(50), 
 `tel` VARCHAR(20), 
 `address` VARCHAR(70), 
 `financial_institution_branch_id` INT, 
 `account_deposittype` CHAR(1),
 `account_no` CHAR(7), 
 `account_holder` VARCHAR(40),
 PRIMARY KEY(`id`),
 FOREIGN KEY(`financial_institution_branch_id`) REFERENCES `financial_institution_branch`(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*買注文受注テーブル*/
CREATE TABLE `buy_orders`( 
 `id` INT AUTO_INCREMENT,
 `client_id` INT,
 `no` INT,
 `estimate_flag` BOOLEAN,
 `manufacturer` VARCHAR(30),
 `car_name` VARCHAR(30),
 `car_model_year` CHAR(2),
 `car_budget` VARCHAR(15),
 `car_exterior_color` VARCHAR(20),
 `car_interior_color` VARCHAR(20),
 `car_model_no` VARCHAR(20),
 `insert_date` TIMESTAMP,
 `delivery_flag` BOOLEAN,
 PRIMARY KEY(`id`),
 FOREIGN KEY(`client_id`) REFERENCES `client`(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*買注文仕入テーブル*/
CREATE TABLE `buy_purchase`(
 `id` INT AUTO_INCREMENT, 
 `buy_orders_id` INT, 
 `purchase_date` DATE, 
 `suppliers_id` INT, 
 `purchase_amount` VARCHAR(15), 
 `expense` VARCHAR(15), 
 `car_name_deadline` DATE, 
 `purchase_cost` VARCHAR(15), 
 `consumption_tax` VARCHAR(15), 
 `sales_date` DATE, 
 `listing_no` CHAR(5), 
 `fee` VARCHAR(15), 
 `total_sales` VARCHAR(15), 
 `net_income` VARCHAR(15), 
 `tax_included` VARCHAR(15), 
 `car_name_deadline_finish` DATE, 
 `purchase_availability` BOOLEAN, 
 `document_flag` BOOLEAN, 
 `payment_flag` BOOLEAN, 
 `quantity` INT, 
 `car_information_id` INT, 
 `insert_date` TIMESTAMP,
 PRIMARY KEY(`id`),
 FOREIGN KEY(`buy_orders_id`) REFERENCES `buy_orders`(`id`),
 FOREIGN KEY(`suppliers_id`) REFERENCES `suppliers`(`id`),
 FOREIGN KEY(`car_information_id`) REFERENCES `car_information`(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*買注文請求テーブル*/ 
CREATE TABLE `buy_claim`(
 `id` INT AUTO_INCREMENT, 
 `buy_purchase_id` INT, 
 `deadline` DATE, 
 `billed` BOOLEAN,
 `from_date` CHAR(6),
 `to_date` CHAR(6),
 `other_banks_ticket` CHAR(10),
 `payment_classification` CHAR(1),
 `edi_information` CHAR(20),
 `insert_date` TIMESTAMP,
 PRIMARY KEY(`id`),
 FOREIGN KEY(`buy_purchase_id`) REFERENCES `buy_purchase`(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*買注文回収テーブル*/
CREATE TABLE `buy_recovery`(
 `id` INT AUTO_INCREMENT, 
 `buy_claim_id` INT, 
 `recovery_price` VARCHAR(15), 
 `recovery_date` DATE, 
 `insert_date` TIMESTAMP,
 PRIMARY KEY(`id`),
 FOREIGN KEY(`buy_claim_id`) REFERENCES `buy_claim`(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*変更ログテーブル*/
CREATE TABLE `log_change`(
 `data_type` INT,
 `data_id` INT,
 `employee_id` CHAR(7),
 `data_comment` TEXT,
 PRIMARY KEY(`data_type`,`data_id`),
 FOREIGN KEY(`employee_id`) REFERENCES `employee`(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
