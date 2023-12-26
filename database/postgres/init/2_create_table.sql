--
-- Project: emn
-- Author: ktr
-- 
-- Use PostgreSQL
-- Created at ...
--

-- [Sammary of types]
--    Master:       The tables that records basic information for carrying out business.
--    Transaction:  The tables that records detailed information that has occurred in carrying out business.
--    Temporary:    The tables that can only survive during connected sessions.
--    Meta:         The tables that contains the contents of the database itself.
--    Intersection: The tables relating the many-to-many table.
--    Basis:        The basis table for relying on data integrity constraints with foreign keys.
--    Normal:       Just a tables that is neither of the above.

-- [Naming conventions]
--    1-1. Prefix master tables name with "MST_".
--    1-2. Prefix transaction tables name with "TRN_".
--    1-3. Prefix temporary tables name with "TMP_".
--    1-4. Prefix meta tables name with "META_".
--    1-5. Prefix intersection tables name with "BSS_".
--    1-6. Prefix basis tables name with "BSS_".
--    2-1. Use both UpperCamelCase and snake_case for primary key, forign key and table name.

-- [Configuration]
--    1-1. Use CHAR type for system specific settings or something that can be classified into a finite number of pieces.
--           For example, users' id, unit types, list of things consumed in making things, and more.
--    1-2. Use VARCHAR for somethings that are domain based or cannot be classified into a finite number of pieces.
--           For example, items, clients, inventories, and more.
--

-- [Other rules]
--    1-1. Between tables that have a one-to-many or many-to-many relationship, the prefix of the referenced side takes precedence.
--

----------------------------------------------------------------
-- Create tables
----------------------------------------------------------------

--
-- Users' information for signing in.
--
CREATE TABLE emn.mst_user (
  mst_user_id     CHAR(8)     NOT NULL,
  user_name       VARCHAR(16) NOT NULL,
  user_password   VARCHAR(64) NOT NULL,
  administrative  BOOLEAN     NOT NULL DEFAULT FALSE,
  stop_using      DATE        NOT NULL DEFAULT '9999-12-31',
  created_at      TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP,
  created_by      CHAR(8)     NOT NULL DEFAULT '00000000',
  PRIMARY KEY (mst_user_id)
);

--
--
--
CREATE TABLE emn.table_information (
  table_information_id  CHAR(36)  NOT NULL,
  created_at            TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  created_by            CHAR(8)   NOT NULL DEFAULT '00000000',
  updated_at            TIMESTAMP NOT NULL DEFAULT '9999-12-31',
  updated_by            CHAR(8)   NOT NULL DEFAULT 'XXXXXXXX',
  PRIMARY KEY (table_information_id)
);

--
-- Express items' category.
--
CREATE TABLE emn.mst_item_category (
  mst_item_category_id  CHAR(2)       NOT NULL,
  category_name         VARCHAR(16)   NOT NULL,
  remark                VARCHAR(1024) NOT NULL DEFAULT '',
  stop_using            DATE          NOT NULL DEFAULT '9999-12-31',
  table_information_id  CHAR(36)      NOT NULL DEFAULT 'XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX',
  PRIMARY KEY (mst_item_category_id)
);

--
-- Express items' status.
--
CREATE TABLE emn.mst_item_status (
  mst_item_status_id    CHAR(2)       NOT NULL,
  status_name           VARCHAR(16)   NOT NULL,
  remark                VARCHAR(1024) NOT NULL DEFAULT '',
  stop_using            DATE          NOT NULL DEFAULT '9999-12-31',
  table_information_id  CHAR(36)      NOT NULL DEFAULT 'XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX',
  PRIMARY KEY (mst_item_status_id)
);

--
--
--
CREATE TABLE emn.mst_item_unit (
  mst_item_unit_id      CHAR(2)       NOT NULL,
  unit_name             VARCHAR(16)   NOT NULL,
  remark                VARCHAR(1024) NOT NULL DEFAULT '',
  stop_using            DATE          NOT NULL DEFAULT '9999-12-31',
  table_information_id  CHAR(36)      NOT NULL DEFAULT 'XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX',
  PRIMARY KEY (mst_item_unit_id)
);

--
--
--
CREATE TABLE emn.mst_address (
  mst_address_id            VARCHAR(8)    NOT NULL,
  building                  VARCHAR(32)   NOT NULL DEFAULT '',
  street                    VARCHAR(32)   NOT NULL DEFAULT '',
  city                      VARCHAR(32)   NOT NULL DEFAULT '',
  state                     VARCHAR(32)   NOT NULL DEFAULT '',
  province                  VARCHAR(32)   NOT NULL DEFAULT '',
  region                    VARCHAR(32)   NOT NULL DEFAULT '',
  prefecture                VARCHAR(32)   NOT NULL DEFAULT '',
  zip                       VARCHAR(16)   NOT NULL DEFAULT '',
  postal_code               VARCHAR(8)    NOT NULL DEFAULT '',
  country                   VARCHAR(32)   NOT NULL DEFAULT '',
  bss_phone_number_list_id  CHAR(10)      NOT NULL DEFAULT 'XXXXXXXXXX',
  remark                    VARCHAR(1024) NOT NULL DEFAULT '',
  stop_using                DATE          NOT NULL DEFAULT '9999-12-31',
  table_information_id      CHAR(36)      NOT NULL DEFAULT 'XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX',
  PRIMARY KEY (mst_address_id)
);

--
--
--
CREATE TABLE emn.mst_company (
  mst_company_id        VARCHAR(8)    NOT NULL,
  company_name          VARCHAR(32)   NOT NULL,
  mst_address_id        VARCHAR(8)    NOT NULL,
  remark                VARCHAR(1024) NOT NULL DEFAULT '',
  stop_using            DATE          NOT NULL DEFAULT '9999-12-31',
  table_information_id  CHAR(36)      NOT NULL DEFAULT 'XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX',
  PRIMARY KEY (mst_company_id)
);

--
--
--
CREATE TABLE emn.mst_warehouse (
  mst_warehouse_id      VARCHAR(8)    NOT NULL,
  warehouse_name        VARCHAR(32)   NOT NULL,
  mst_company_id        VARCHAR(8)    NOT NULL,
  mst_address_id        VARCHAR(8)    NOT NULL,
  remark                VARCHAR(1024) NOT NULL DEFAULT '',
  stop_using            DATE          NOT NULL DEFAULT '9999-12-31',
  table_information_id  CHAR(36)      NOT NULL DEFAULT 'XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX',
  PRIMARY KEY (mst_warehouse_id)
);

--
--
--
CREATE TABLE emn.mst_factory (
  mst_factory_id        VARCHAR(8)    NOT NULL,
  factory_name          VARCHAR(32)   NOT NULL,
  mst_company_id        VARCHAR(8)    NOT NULL,
  mst_address_id        VARCHAR(8)    NOT NULL,
  remark                VARCHAR(1024) NOT NULL DEFAULT '',
  stop_using            DATE          NOT NULL DEFAULT '9999-12-31',
  table_information_id  CHAR(36)      NOT NULL DEFAULT 'XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX',
  PRIMARY KEY (mst_factory_id)
);

--
--
--
CREATE TABLE emn.mst_machine (
  mst_machine_id        VARCHAR(8)    NOT NULL,
  machine_name          VARCHAR(16)   NOT NULL,
  mst_factory_id        VARCHAR(8)    NOT NULL,
  maker_id              VARCHAR(8)    NOT NULL DEFAULT 'XXXXXXXX',
  remark                VARCHAR(1024) NOT NULL DEFAULT '',
  stop_using            DATE          NOT NULL DEFAULT '9999-12-31',
  table_information_id  CHAR(36)      NOT NULL DEFAULT 'XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX',
  PRIMARY KEY (mst_machine_id)
);

--
--
--
CREATE TABLE emn.mst_arrival_category (
  mst_arrival_category_id VARCHAR(4)    NOT NULL,
  arrival_category_name   VARCHAR(16)   NOT NULL,
  remark                  VARCHAR(1024) NOT NULL DEFAULT '',
  stop_using              DATE          NOT NULL DEFAULT '9999-12-31',
  table_information_id    CHAR(36)      NOT NULL DEFAULT 'XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX',
  PRIMARY KEY (mst_arrival_category_id)
);

--
--
--
CREATE TABLE emn.mst_operator (
  mst_operator_id       VARCHAR(8)    NOT NULL,
  operator_name         VARCHAR(16)   NOT NULL,
  remark                VARCHAR(1024) NOT NULL DEFAULT '',
  stop_using            DATE          NOT NULL DEFAULT '9999-12-31',
  table_information_id  CHAR(36)      NOT NULL DEFAULT 'XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX',
  PRIMARY KEY (mst_operator_id)
);

--
--
--
CREATE TABLE emn.mst_process (
  mst_process_id        VARCHAR(8)    NOT NULL,
  process_name          VARCHAR(16)   NOT NULL,
  mst_factory_id        VARCHAR(8)    NOT NULL,
  remark                VARCHAR(1024) NOT NULL DEFAULT '',
  stop_using            DATE          NOT NULL DEFAULT '9999-12-31',
  table_information_id  CHAR(36)      NOT NULL DEFAULT 'XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX',
  PRIMARY KEY (mst_process_id)
);

--
-- Products' master data
--
CREATE TABLE emn.mst_item (
  mst_item_id                 VARCHAR(16)     NOT NULL,
  item_number                 VARCHAR(32)     NOT NULL,
  item_name                   VARCHAR(32)     NOT NULL,
  abbreviation                VARCHAR(32)     NOT NULL DEFAULT '',
  mst_item_category_id        CHAR(2)         NOT NULL,
  mst_item_status_id          CHAR(2)         NOT NULL,
  mst_item_unit_id            CHAR(2)         NOT NULL,
  bss_machine_list_id         CHAR(10)        NOT NULL DEFAULT 'XXXXXXXXXX',
  mst_warehouse_id            VARCHAR(8)      NOT NULL DEFAULT 'XXXXXXXX',
  lower_limit_inventory_qty   NUMERIC(10,2)   NOT NULL DEFAULT 0.00,
  safety_inventory_qty        NUMERIC(10,2)   NOT NULL DEFAULT 0.00,
  upper_limit_inventory_qty   NUMERIC(10,2)   NOT NULL DEFAULT 0.00,
  min_lot_qty                 NUMERIC(10,2)   NOT NULL DEFAULT 0.00,
  max_lot_qty                 NUMERIC(10,2)   NOT NULL DEFAULT 0.00,
  unit_price                  NUMERIC(10,2)   NOT NULL DEFAULT 0.00,
  validity_days               INTEGER         NOT NULL DEFAULT 365,
  mst_client_id               VARCHAR(8)      NOT NULL DEFAULT 'XXXXXXXX',
  mst_delivery_destination_id VARCHAR(8)      NOT NULL DEFAULT 'XXXXXXXX',
  bss_end_user_list_id        CHAR(10)        NOT NULL DEFAULT 'XXXXXXXXXX',
  rank                        VARCHAR(4)      NOT NULL DEFAULT '',
  remark                      VARCHAR(1024)   NOT NULL DEFAULT '',
  stop_using                  DATE            NOT NULL DEFAULT '9999-12-31',
  table_information_id        CHAR(36)        NOT NULL DEFAULT 'XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX',
  PRIMARY KEY (mst_item_id)
);

--
--
--
CREATE TABLE emn.mst_process_order (
  mst_item_id           VARCHAR(16)   NOT NULL,
  process_order_no      SMALLSERIAL   NOT NULL,
  is_selectable         BOOLEAN       NOT NULL, -- make it possible to choose whether production should be done in-house or outsourced.
  is_parallel           BOOLEAN       DEFAULT FALSE, -- make it possible to choose whether work on multiple processes in parallel.
  mst_process_id        VARCHAR(8)    NOT NULL,
  remark                VARCHAR(1024) NOT NULL DEFAULT '',
  stop_using            DATE          NOT NULL DEFAULT '9999-12-31',
  table_information_id  CHAR(36)      NOT NULL DEFAULT 'XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX',
  PRIMARY KEY (mst_item_id, process_order_no, is_selectable)
);

--
--
--
CREATE TABLE emn.mst_bom (
  mst_sub_item_id       VARCHAR(16)   NOT NULL,
  bom_no                SMALLSERIAL   NOT NULL,
  start_to_use          DATE          NOT NULL DEFAULT CURRENT_TIMESTAMP,
  mst_basic_item_id     VARCHAR(16)   NOT NULL,
  required_qty          NUMERIC(10,2) NOT NULL DEFAULT 0.00,
  mst_item_unit_id      CHAR(2)       NOT NULL,
  remark                VARCHAR(1024) NOT NULL NOT NULL DEFAULT '',
  stop_using            DATE          NOT NULL DEFAULT '9999-12-31',
  table_information_id  CHAR(36)      NOT NULL DEFAULT 'XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX',
  PRIMARY KEY (mst_sub_item_id, bom_no, start_to_use)
);

--
--
--
CREATE TABLE emn.bss_phone_number_list (
  bss_phone_number_list_id CHAR(10) NOT NULL,
  PRIMARY KEY (bss_phone_number_list_id)
);

--
--
--
CREATE TABLE emn.bss_machine_list (
  bss_machine_list_id CHAR(10) NOT NULL,
  PRIMARY KEY (bss_machine_list_id)
);

--
--
--
CREATE TABLE emn.bss_consumption_list (
  bss_consumption_list_id CHAR(36) NOT NULL,
  PRIMARY KEY (bss_consumption_list_id)
);

--
--
--
CREATE TABLE emn.bss_end_user_list (
  bss_end_user_list_id  CHAR(10) NOT NULL,
  PRIMARY KEY (bss_end_user_list_id)
);

--
--
--
CREATE TABLE emn.int_outsourcing_order_shipment (
  int_outsourcing_order_shipment_id   CHAR(10)    NOT NULL,
  trn_outsourcing_order_id            CHAR(10)    NOT NULL,
  outsourcing_order_no                SMALLSERIAL NOT NULL,
  trn_shipment_id                     VARCHAR(16) NOT NULL,
  PRIMARY KEY (int_outsourcing_order_shipment_id)
);

--
--
--
CREATE TABLE emn.int_received_order_delivery (
  int_received_order_delivery_id  CHAR(10)    NOT NULL,
  trn_received_order_id           CHAR(10)    NOT NULL,
  received_order_no               SMALLSERIAL NOT NULL,
  trn_delivery_id                 CHAR(10)    NOT NULL,
  delivery_no                     SMALLSERIAL NOT NULL,
  PRIMARY KEY (int_received_order_delivery_id)
);

--
--
--
CREATE TABLE emn.phone_number_list (
  bss_phone_number_list_id  CHAR(10)    NOT NULL,
  phone_number_list_no      SMALLSERIAL NOT NULL,
  phone_number              VARCHAR(16) NOT NULL,
  table_information_id      CHAR(36)    NOT NULL DEFAULT 'XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX',
  PRIMARY KEY (bss_phone_number_list_id, phone_number_list_no)
);

--
--
--
CREATE TABLE emn.machine_list (
  bss_machine_list_id   CHAR(10)    NOT NULL,
  machine_list_no       SMALLSERIAL NOT NULL,
  mst_machine_id        VARCHAR(8)  NOT NULL,
  table_information_id  CHAR(36)    NOT NULL DEFAULT 'XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX',
  PRIMARY KEY (bss_machine_list_id, machine_list_no)
);

--
-- The consumption list of a production.
--
CREATE TABLE emn.consumption_list (
  bss_consumption_list_id CHAR(36)        NOT NULL,
  consumption_list_no     SMALLSERIAL     NOT NULL,
  mst_warehouse_id        VARCHAR(8)      NOT NULL,
  mst_item_id             VARCHAR(16)     NOT NULL,
  mst_process_id          VARCHAR(8)      NOT NULL DEFAULT 'XXXXXXXX',
  lot                     VARCHAR(16)     NOT NULL,
  branch                  VARCHAR(16)     NOT NULL DEFAULT '',
  non_defective_qty       NUMERIC(10,2)   NOT NULL DEFAULT 0.00,
  defective_qty           NUMERIC(10,2)   NOT NULL DEFAULT 0.00,
  suspended_qty           NUMERIC(10,2)   NOT NULL DEFAULT 0.00,
  is_used_up              BOOLEAN         NOT NULL DEFAULT FALSE,
  transaction_type        VARCHAR(32)     NOT NULL, -- Store the usecase name in this column.
  PRIMARY KEY (bss_consumption_list_id, consumption_list_no)
);

--
--
--
CREATE TABLE emn.end_user_list (
  bss_end_user_list_id  CHAR(10)      NOT NULL,
  end_user_list_no      SMALLSERIAL   NOT NULL,
  end_user_id           VARCHAR(8)    NOT NULL,
  remark                VARCHAR(1024) NOT NULL DEFAULT '',
  stop_using            DATE          NOT NULL DEFAULT '9999-12-31',
  table_information_id  CHAR(36)      NOT NULL DEFAULT 'XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX',
  PRIMARY KEY (bss_end_user_list_id, end_user_list_no)
);

--
-- Produced inventories.
--
CREATE TABLE emn.inventory (
  mst_item_id               VARCHAR(16)     NOT NULL,
  mst_process_id            VARCHAR(8)      NOT NULL,
  lot                       VARCHAR(16)     NOT NULL,
  branch                    VARCHAR(16)     NOT NULL DEFAULT '',
  mst_warehouse_id          VARCHAR(8)      NOT NULL,
  non_defective_qty         NUMERIC(10,2)   NOT NULL DEFAULT 0.00,
  defective_qty             NUMERIC(10,2)   NOT NULL DEFAULT 0.00,
  suspended_qty             NUMERIC(10,2)   NOT NULL DEFAULT 0.00,
  expiration_date           DATE            NOT NULL DEFAULT '9999-12-31',
  is_used                   BOOLEAN         NOT NULL DEFAULT FALSE,
  is_used_up                BOOLEAN         NOT NULL DEFAULT FALSE,
  bss_regulate_inventory_id CHAR(10)        NOT NULL DEFAULT 'XXXXXXXXXX', -- This column has not yet been constrained.
  -- trn_inventory_count_id  CHAR(10)        NOT NULL DEFAULT 'XXXXXXXXXX', -- This column has not yet been constrained.
  PRIMARY KEY (mst_item_id, mst_process_id, lot, branch)
);

--
--
--
CREATE TABLE emn.bss_regulate_inventory (
  bss_regulate_inventory_id CHAR(10)  NOT NULL,
  PRIMARY KEY (bss_regulate_inventory_id)
);

--
--
--
CREATE TABLE emn.regulate_inventory (
  bss_regulate_inventory_id CHAR(10)        NOT NULL,
  regulate_inventory_no     SMALLSERIAL     NOT NULL,
  mst_item_id               VARCHAR(16)     NOT NULL,
  mst_process_id            VARCHAR(8)      NOT NULL,
  lot                       VARCHAR(16)     NOT NULL,
  branch                    VARCHAR(16)     NOT NULL DEFAULT '',
  mst_warehouse_id_before   VARCHAR(16)     NOT NULL,
  mst_warehouse_id_after    VARCHAR(16)     NOT NULL,
  nondefective_qty_before   NUMERIC(10,2)   NOT NULL,
  nondefective_qty_after    NUMERIC(10,2)   NOT NULL,
  defective_qty_before      NUMERIC(10,2)   NOT NULL,
  defective_qty_after       NUMERIC(10,2)   NOT NULL,
  suspended_qty_before      NUMERIC(10,2)   NOT NULL,
  suspended_qty_after       NUMERIC(10,2)   NOT NULL,
  remark                    VARCHAR(1024)   NOT NULL DEFAULT '',
  table_information_id      CHAR(36)        NOT NULL DEFAULT 'XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX',
  PRIMARY KEY (bss_regulate_inventory_id, regulate_inventory_no)
);

--
-- Produced results.
--
CREATE TABLE emn.trn_production_instruction (
  trn_production_instruction_id CHAR(36)        NOT NULL,
  mst_item_id                   VARCHAR(16)     NOT NULL,
  mst_process_id                VARCHAR(8)      NOT NULL DEFAULT 'XXXXXXXX',
  mst_machine_id                VARCHAR(8)      NOT NULL,
  mst_operator_id               VARCHAR(8)      NOT NULL,
  required_qty                  NUMERIC(10,2)   NOT NULL DEFAULT 0.00,
  start_to_produce              DATE            NOT NULL,
  end_producing                 DATE            NOT NULL,
  remark                        VARCHAR(1024)   NOT NULL DEFAULT '',
  is_canceled                   BOOLEAN         NOT NULL DEFAULT FALSE,
  table_information_id          CHAR(36)        NOT NULL DEFAULT 'XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX',
  PRIMARY KEY (trn_production_instruction_id)
);

--
-- Produced results.
--
CREATE TABLE emn.trn_production (
  trn_production_id             CHAR(36)      NOT NULL,
  trn_production_instruction_id CHAR(36)      NOT NULL DEFAULT 'XXXXXXXXXX',
  mst_item_id                   VARCHAR(16)   NOT NULL,
  mst_process_id                VARCHAR(8)    NOT NULL DEFAULT 'XXXXXXXX',
  lot                           VARCHAR(16)   NOT NULL,
  branch                        VARCHAR(16)   NOT NULL DEFAULT '',
  mst_machine_id                VARCHAR(8)    NOT NULL,
  mst_operator_id               VARCHAR(8)    NOT NULL,
  bss_consumption_list_id       CHAR(36)      NOT NULL,
  non_defective_qty             NUMERIC(10,2) NOT NULL DEFAULT 0.00,
  defective_qty                 NUMERIC(10,2) NOT NULL DEFAULT 0.00,
  suspended_qty                 NUMERIC(10,2) NOT NULL DEFAULT 0.00,
  produced_at                   DATE          NOT NULL DEFAULT CURRENT_TIMESTAMP,
  information                   VARCHAR(1024) NOT NULL DEFAULT '',
  remark                        VARCHAR(1024) NOT NULL DEFAULT '',
  is_canceled                   BOOLEAN       NOT NULL DEFAULT FALSE,
  table_information_id          CHAR(36)      NOT NULL DEFAULT 'XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX',
  PRIMARY KEY (trn_production_id)
);

--
-- Ordering to outside.
--
CREATE TABLE emn.trn_outsourcing_order (
  trn_outsourcing_order_id              CHAR(10)      NOT NULL,
  outsourcing_order_no                  SMALLSERIAL   NOT NULL,
  trn_outsourcing_order_instruction_id  CHAR(10)      NOT NULL DEFAULT 'XXXXXXXXXX', -- This column has not yet been constrained.
  mst_item_id                           VARCHAR(16)   NOT NULL,
  mst_process_id                        VARCHAR(8)    NOT NULL DEFAULT 'XXXXXXXX',
  lot                                   VARCHAR(16)   NOT NULL,
  branch                                VARCHAR(16)   NOT NULL DEFAULT '',
  non_defective_qty                     NUMERIC(10,2) NOT NULL DEFAULT 0.00,
  defective_qty                         NUMERIC(10,2) NOT NULL DEFAULT 0.00,
  suspended_qty                         NUMERIC(10,2) NOT NULL DEFAULT 0.00,
  mst_company_id                        VARCHAR(8)    NOT NULL,
  mst_next_process_id                   VARCHAR(8)    NOT NULL DEFAULT 'XXXXXXXX',
  mst_operator_id                       VARCHAR(8)    NOT NULL,
  order_unit_price                      NUMERIC(10,2) NOT NULL DEFAULT 0.00,
  order_price                           NUMERIC(10,2) NOT NULL DEFAULT 0.00,
  is_arrival_completed                  BOOLEAN       NOT NULL DEFAULT FALSE,
  remark                                VARCHAR(1024) NOT NULL DEFAULT '',
  is_canceled                           BOOLEAN       NOT NULL DEFAULT FALSE,
  table_information_id                  CHAR(36)      NOT NULL DEFAULT 'XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX',
  PRIMARY KEY (trn_outsourcing_order_id, outsourcing_order_no)
);

--
--
--
CREATE TABLE emn.trn_shipment (
  trn_shipment_id             VARCHAR(16)   NOT NULL,
  trn_shipment_instruction_id CHAR(10)      NOT NULL DEFAULT 'XXXXXXXXXX' , -- This column has not yet been constrained.
  bss_consumption_list_id     CHAR(10)      NOT NULL,
  shipment_date               TIMESTAMP     NOT NULL DEFAULT CURRENT_TIMESTAMP,
  mst_operator_id             VARCHAR(8)    NOT NULL,
  mst_company_id              VARCHAR(8)    NOT NULL,
  mst_factory_id              VARCHAR(8)    NOT NULL,
  remark                      VARCHAR(1024) NOT NULL DEFAULT '',
  is_canceled                 BOOLEAN       NOT NULL DEFAULT FALSE,
  table_information_id        CHAR(36)      NOT NULL DEFAULT 'XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX',
  PRIMARY KEY (trn_shipment_id)
);

--
--
--
CREATE TABLE emn.trn_arrival (
  trn_arrival_id          VARCHAR(16)   NOT NULL,
  trn_shipment_id         VARCHAR(16)   NOT NULL,
  arrival_date            TIMESTAMP     NOT NULL,
  arrival_qty             NUMERIC(10,2) NOT NULL DEFAULT 0.00,
  mst_arrival_category_id VARCHAR(4)    NOT NULL,
  mst_operator_id         VARCHAR(8)    NOT NULL,
  remark                  VARCHAR(1024) NOT NULL DEFAULT '',
  is_canceled             BOOLEAN       NOT NULL DEFAULT FALSE,
  table_information_id    CHAR(36)      NOT NULL DEFAULT 'XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX',
  PRIMARY KEY (trn_arrival_id)
);

--
--
-- Should it separate table...?
CREATE TABLE emn.trn_purchase (
  trn_purchase_id       VARCHAR(16)   NOT NULL,
  purchase_no           SMALLSERIAL   NOT NULL,
  mst_item_id           VARCHAR(16)   NOT NULL,
  mst_process_id        VARCHAR(8)    NOT NULL DEFAULT 'XXXXXXXX',
  lot                   VARCHAR(16)   NOT NULL,
  branch                VARCHAR(16)   NOT NULL DEFAULT '',
  ordere_item_qty       NUMERIC(10,2) NOT NULL DEFAULT 0.00,
  arrival_item_qty      NUMERIC(10,2) NOT NULL DEFAULT 0.00,
  mst_operator_id       VARCHAR(8)    NOT NULL DEFAULT 'XXXXXXXX',
  mst_company_id        VARCHAR(8)    NOT NULL,
  mst_warehouse_id      VARCHAR(8)    NOT NULL,
  unit_price            NUMERIC(10,2) NOT NULL DEFAULT 0.00,
  price                 NUMERIC(10,2) NOT NULL DEFAULT 0.00,
  is_arrival_completed  BOOLEAN       NOT NULL DEFAULT FALSE,
  remark                VARCHAR(1024) NOT NULL DEFAULT '',
  is_canceled           BOOLEAN       NOT NULL DEFAULT FALSE ,
  table_information_id  CHAR(36)      NOT NULL DEFAULT 'XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX',
  PRIMARY KEY (trn_purchase_id, purchase_no)
);

--
--
--
CREATE TABLE emn.trn_received_order (
  trn_received_order_id         VARCHAR(16)   NOT NULL,
  received_order_no             SMALLSERIAL   NOT NULL,
  received_order_date           TIMESTAMP     NOT NULL DEFAULT CURRENT_TIMESTAMP,
  mst_item_id                   VARCHAR(16)   NOT NULL,
  received_order_item_qty       NUMERIC(10,2) NOT NULL DEFAULT 0.00,
  mst_company_id                VARCHAR(8)    NOT NULL DEFAULT 'XXXXXXXX',
  shipping_date                 TIMESTAMP     NOT NULL DEFAULT CURRENT_TIMESTAMP,
  arrival_date                  TIMESTAMP     NOT NULL DEFAULT CURRENT_TIMESTAMP,
  mst_operator_id               VARCHAR(8)    NOT NULL DEFAULT 'XXXXXXXX',
  unit_price                    NUMERIC(10,2) NOT NULL DEFAULT 0.00,
  price                         NUMERIC(10,2) NOT NULL DEFAULT 0.00,
  is_delivery_completed         BOOLEAN       NOT NULL DEFAULT FALSE,
  remark                        VARCHAR(1024) NOT NULL DEFAULT '',
  is_canceled                   BOOLEAN       NOT NULL DEFAULT FALSE,
  table_information_id          CHAR(36)      NOT NULL DEFAULT 'XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX',
  PRIMARY KEY (trn_received_order_id, received_order_no)
);

--
--
--
CREATE TABLE emn.trn_delivery (
  trn_delivery_id             VARCHAR(16)   NOT NULL,
  delivery_no                 SMALLSERIAL   NOT NULL,
  trn_delivery_instruction_id CHAR(10)      NOT NULL DEFAULT 'XXXXXXXXXX' , -- This column has not yet been constrained.
  bss_consumption_list_id     CHAR(10)      NOT NULL,
  shipping_date               TIMESTAMP     NOT NULL DEFAULT CURRENT_TIMESTAMP,
  arrival_date                TIMESTAMP     NOT NULL DEFAULT CURRENT_TIMESTAMP,
  mst_operator_id             VARCHAR(8)    NOT NULL,
  remark                      VARCHAR(1024) NOT NULL DEFAULT '',
  is_canceled                 BOOLEAN       NOT NULL DEFAULT FALSE,
  table_information_id        CHAR(36)      NOT NULL DEFAULT 'XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX',
  PRIMARY KEY (trn_delivery_id, delivery_no)
);

------------------------------------------------
-- Alter tables
------------------------------------------------

--
--
--
ALTER TABLE emn.table_information
  ADD FOREIGN KEY (created_by)
    REFERENCES emn.mst_user (mst_user_id),
  ADD FOREIGN KEY (updated_by)
    REFERENCES emn.mst_user (mst_user_id)
;

--
--
--
ALTER TABLE emn.mst_item_category
  ADD FOREIGN KEY (table_information_id)
    REFERENCES emn.table_information (table_information_id)
;

--
--
--
ALTER TABLE emn.mst_item_status
  ADD FOREIGN KEY (table_information_id)
    REFERENCES emn.table_information (table_information_id)
;

--
--
--
ALTER TABLE emn.mst_item_unit
  ADD FOREIGN KEY (table_information_id)
    REFERENCES emn.table_information (table_information_id)
;

--
--
--
ALTER TABLE emn.mst_address
  ADD FOREIGN KEY (bss_phone_number_list_id)
    REFERENCES emn.bss_phone_number_list (bss_phone_number_list_id),
  ADD FOREIGN KEY (table_information_id)
    REFERENCES emn.table_information (table_information_id)
;

--
--
--
ALTER TABLE emn.mst_company
  ADD FOREIGN KEY (mst_address_id)
    REFERENCES emn.mst_address (mst_address_id),
  ADD FOREIGN KEY (table_information_id)
    REFERENCES emn.table_information (table_information_id)
;

--
--
--
ALTER TABLE emn.mst_warehouse
  ADD FOREIGN KEY (mst_company_id)
    REFERENCES emn.mst_company (mst_company_id),
  ADD FOREIGN KEY (mst_address_id)
    REFERENCES emn.mst_address (mst_address_id),
  ADD FOREIGN KEY (table_information_id)
    REFERENCES emn.table_information (table_information_id)
;

--
--
--
ALTER TABLE emn.mst_factory
  ADD FOREIGN KEY (mst_company_id)
    REFERENCES emn.mst_company (mst_company_id),
  ADD FOREIGN KEY (mst_address_id)
    REFERENCES emn.mst_address (mst_address_id),
  ADD FOREIGN KEY (table_information_id)
    REFERENCES emn.table_information (table_information_id)
;

--
--
--
ALTER TABLE emn.mst_machine
  ADD FOREIGN KEY (mst_factory_id)
    REFERENCES emn.mst_factory (mst_factory_id),
  ADD FOREIGN KEY (maker_id)
    REFERENCES emn.mst_company (mst_company_id),
  ADD FOREIGN KEY (table_information_id)
    REFERENCES emn.table_information (table_information_id)
;

--
--
--
ALTER TABLE emn.mst_arrival_category
  ADD FOREIGN KEY (table_information_id)
    REFERENCES emn.table_information (table_information_id)
;

--
--
--
ALTER TABLE emn.mst_operator
  ADD FOREIGN KEY (table_information_id)
    REFERENCES emn.table_information (table_information_id)
;

--
--
--
ALTER TABLE emn.mst_process
  ADD FOREIGN KEY (mst_factory_id)
    REFERENCES emn.mst_factory (mst_factory_id),
  ADD FOREIGN KEY (table_information_id)
    REFERENCES emn.table_information (table_information_id)
;

--
--
--
ALTER TABLE emn.mst_item
  ADD FOREIGN KEY (bss_end_user_list_id)
    REFERENCES emn.bss_end_user_list (bss_end_user_list_id),
  ADD FOREIGN KEY (mst_item_category_id)
    REFERENCES emn.mst_item_category (mst_item_category_id),
  ADD FOREIGN KEY (mst_item_status_id)
    REFERENCES emn.mst_item_status (mst_item_status_id),
  ADD FOREIGN KEY (mst_item_unit_id)
    REFERENCES emn.mst_item_unit (mst_item_unit_id),
  ADD FOREIGN KEY (bss_machine_list_id)
    REFERENCES emn.bss_machine_list (bss_machine_list_id),
  ADD FOREIGN KEY (mst_warehouse_id)
    REFERENCES emn.mst_warehouse (mst_warehouse_id),
  ADD FOREIGN KEY (mst_client_id)
    REFERENCES emn.mst_company (mst_company_id),
  ADD FOREIGN KEY (mst_delivery_destination_id)
    REFERENCES emn.mst_company (mst_company_id),
  ADD FOREIGN KEY (bss_end_user_list_id)
    REFERENCES emn.bss_end_user_list (bss_end_user_list_id),
  ADD FOREIGN KEY (table_information_id)
    REFERENCES emn.table_information (table_information_id)
;

--
--
--
ALTER TABLE emn.mst_process_order
  ADD FOREIGN KEY (mst_item_id)
    REFERENCES emn.mst_item (mst_item_id),
  ADD FOREIGN KEY (mst_process_id)
    REFERENCES emn.mst_process (mst_process_id)
;

--
--
--
ALTER TABLE emn.mst_bom
  ADD FOREIGN KEY (mst_sub_item_id)
    REFERENCES emn.mst_item (mst_item_id),
  ADD FOREIGN KEY (mst_basic_item_id)
    REFERENCES emn.mst_item (mst_item_id),
  ADD FOREIGN KEY (mst_item_unit_id)
    REFERENCES emn.mst_item_unit (mst_item_unit_id),
  ADD FOREIGN KEY (table_information_id)
    REFERENCES emn.table_information (table_information_id)
;

--
--
--
ALTER TABLE emn.int_outsourcing_order_shipment
  ADD FOREIGN KEY (trn_outsourcing_order_id, outsourcing_order_no)
    REFERENCES emn.trn_outsourcing_order (trn_outsourcing_order_id, outsourcing_order_no),
  ADD FOREIGN KEY (trn_shipment_id)
    REFERENCES emn.trn_shipment (trn_shipment_id)
;

--
--
--
ALTER TABLE emn.int_received_order_delivery
  ADD FOREIGN KEY (trn_received_order_id, received_order_no)
    REFERENCES emn.trn_received_order (trn_received_order_id, received_order_no),
  ADD FOREIGN KEY (trn_delivery_id, delivery_no)
    REFERENCES emn.trn_delivery (trn_delivery_id, delivery_no)
;

--
--
--
ALTER TABLE emn.phone_number_list
  ADD FOREIGN KEY (bss_phone_number_list_id)
    REFERENCES emn.bss_phone_number_list (bss_phone_number_list_id)
;

--
--
--
ALTER TABLE emn.machine_list
  ADD FOREIGN KEY (bss_machine_list_id)
    REFERENCES emn.bss_machine_list (bss_machine_list_id)
;

--
--
--
ALTER TABLE emn.consumption_list
  ADD FOREIGN KEY (bss_consumption_list_id)
    REFERENCES emn.bss_consumption_list (bss_consumption_list_id),
  ADD FOREIGN KEY (mst_warehouse_id)
    REFERENCES emn.mst_warehouse (mst_warehouse_id),
  ADD FOREIGN KEY (mst_item_id)
    REFERENCES emn.mst_item (mst_item_id),
  ADD FOREIGN KEY (mst_process_id)
    REFERENCES emn.mst_process (mst_process_id)
;

--
--
--
ALTER TABLE emn.end_user_list
  ADD FOREIGN KEY (bss_end_user_list_id)
    REFERENCES emn.bss_end_user_list (bss_end_user_list_id),
  ADD FOREIGN KEY (end_user_id)
    REFERENCES emn.mst_company (mst_company_id)
;

--
--
--
ALTER TABLE emn.inventory
  ADD FOREIGN KEY (mst_item_id)
    REFERENCES emn.mst_item (mst_item_id),
  ADD FOREIGN KEY (mst_process_id)
    REFERENCES emn.mst_process (mst_process_id),
  ADD FOREIGN KEY (mst_warehouse_id)
    REFERENCES emn.mst_warehouse (mst_warehouse_id),
  ADD FOREIGN KEY (bss_regulate_inventory_id)
    REFERENCES emn.bss_regulate_inventory (bss_regulate_inventory_id)
;

--
--
--
ALTER TABLE emn.regulate_inventory
  ADD FOREIGN KEY (mst_process_id)
    REFERENCES emn.mst_process (mst_process_id),
  ADD FOREIGN KEY (mst_warehouse_id_before)
    REFERENCES emn.mst_warehouse (mst_warehouse_id),
  ADD FOREIGN KEY (mst_warehouse_id_after)
    REFERENCES emn.mst_warehouse (mst_warehouse_id),
  ADD FOREIGN KEY (table_information_id)
    REFERENCES emn.table_information (table_information_id)
;

--
--
--
ALTER TABLE emn.trn_production_instruction
  ADD FOREIGN KEY (mst_item_id)
    REFERENCES emn.mst_item (mst_item_id),
  ADD FOREIGN KEY (mst_process_id)
    REFERENCES emn.mst_process (mst_process_id),
  ADD FOREIGN KEY (mst_machine_id)
    REFERENCES emn.mst_machine (mst_machine_id),
  ADD FOREIGN KEY (mst_operator_id)
    REFERENCES emn.mst_operator (mst_operator_id),
  ADD FOREIGN KEY (table_information_id)
    REFERENCES emn.table_information (table_information_id)
;

--
--
--
ALTER TABLE emn.trn_production
  ADD FOREIGN KEY (trn_production_instruction_id)
    REFERENCES emn.trn_production_instruction (trn_production_instruction_id),
  ADD FOREIGN KEY (mst_item_id)
    REFERENCES emn.mst_item (mst_item_id),
  ADD FOREIGN KEY (mst_process_id)
    REFERENCES emn.mst_process (mst_process_id),
  ADD FOREIGN KEY (mst_machine_id)
    REFERENCES emn.mst_machine (mst_machine_id),
  ADD FOREIGN KEY (mst_operator_id)
    REFERENCES emn.mst_operator (mst_operator_id),
  ADD FOREIGN KEY (bss_consumption_list_id)
    REFERENCES emn.bss_consumption_list (bss_consumption_list_id),
  ADD FOREIGN KEY (table_information_id)
    REFERENCES emn.table_information (table_information_id)
;

--
--
--
ALTER TABLE emn.trn_outsourcing_order
  ADD FOREIGN KEY (mst_item_id)
    REFERENCES emn.mst_item (mst_item_id),
  ADD FOREIGN KEY (mst_process_id)
    REFERENCES emn.mst_process (mst_process_id),
  ADD FOREIGN KEY (mst_company_id)
    REFERENCES emn.mst_company (mst_company_id),
  ADD FOREIGN KEY (mst_next_process_id)
    REFERENCES emn.mst_process (mst_process_id),
  ADD FOREIGN KEY (mst_operator_id)
    REFERENCES emn.mst_operator (mst_operator_id),
  ADD FOREIGN KEY (table_information_id)
    REFERENCES emn.table_information (table_information_id)
;

--
--
--
ALTER TABLE emn.trn_shipment
  ADD FOREIGN KEY (bss_consumption_list_id)
    REFERENCES emn.bss_consumption_list (bss_consumption_list_id),
  ADD FOREIGN KEY (mst_operator_id)
    REFERENCES emn.mst_operator (mst_operator_id),
  ADD FOREIGN KEY (mst_company_id)
    REFERENCES emn.mst_company (mst_company_id),
  ADD FOREIGN KEY (mst_factory_id)
    REFERENCES emn.mst_factory (mst_factory_id),
  ADD FOREIGN KEY (table_information_id)
    REFERENCES emn.table_information (table_information_id)
;

--
--
--
ALTER TABLE emn.trn_arrival
  ADD FOREIGN KEY (trn_shipment_id)
    REFERENCES emn.trn_shipment (trn_shipment_id),
  ADD FOREIGN KEY (mst_arrival_category_id)
    REFERENCES emn.mst_arrival_category (mst_arrival_category_id),
  ADD FOREIGN KEY (mst_operator_id)
    REFERENCES emn.mst_operator (mst_operator_id),
  ADD FOREIGN KEY (table_information_id)
    REFERENCES emn.table_information (table_information_id)
;

--
--
--
ALTER TABLE emn.trn_purchase
  ADD FOREIGN KEY (mst_item_id)
    REFERENCES emn.mst_item (mst_item_id),
  ADD FOREIGN KEY (mst_process_id)
    REFERENCES emn.mst_process (mst_process_id),
  ADD FOREIGN KEY (mst_operator_id)
    REFERENCES emn.mst_operator (mst_operator_id),
  ADD FOREIGN KEY (mst_company_id)
    REFERENCES emn.mst_company (mst_company_id),
  ADD FOREIGN KEY (mst_warehouse_id)
    REFERENCES emn.mst_warehouse (mst_warehouse_id),
  ADD FOREIGN KEY (table_information_id)
    REFERENCES emn.table_information (table_information_id)
;

--
--
--
ALTER TABLE emn.trn_received_order
  ADD FOREIGN KEY (mst_item_id)
    REFERENCES emn.mst_item (mst_item_id),
  ADD FOREIGN KEY (mst_company_id)
    REFERENCES emn.mst_company (mst_company_id),
  ADD FOREIGN KEY (mst_operator_id)
    REFERENCES emn.mst_operator (mst_operator_id),
  ADD FOREIGN KEY (table_information_id)
    REFERENCES emn.table_information (table_information_id)
;

--
--
--
ALTER TABLE emn.trn_delivery
  ADD FOREIGN KEY (bss_consumption_list_id)
    REFERENCES emn.bss_consumption_list (bss_consumption_list_id),
  ADD FOREIGN KEY (mst_operator_id)
    REFERENCES emn.mst_operator (mst_operator_id),
  ADD FOREIGN KEY (table_information_id)
    REFERENCES emn.table_information (table_information_id)
;
