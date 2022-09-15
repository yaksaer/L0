DROP TABLE IF EXISTS orders, items, payments,deliveries, order_info;

CREATE TABLE orders (
    order_uid           varchar(100) PRIMARY KEY,
    track_number        varchar(100) UNIQUE,
    entry               varchar(20),
    locale              varchar(10),
    internal_signature  varchar(10),
    customer_id         varchar(10),
    delivery_service    varchar(10),
    shardkey            varchar(10),
    sm_id               integer,
    date_created        timestamp,
    oof_shard           varchar(10)
);

CREATE TABLE deliveries (
    id                  SERIAL PRIMARY KEY,
    name                varchar(100),
    phone               varchar(30),
    zip                 varchar(30),
    city                varchar(50),
    address             varchar(200),
    region              varchar(50),
    email               varchar(30)
);

CREATE TABLE payments (
    transaction         varchar(100) PRIMARY KEY REFERENCES orders(order_uid),
    request_id          varchar(100),
    currency            varchar(10),
    provider            varchar(50),
    amount              integer,
    payment_dt          bigint,
    bank                varchar(100),
    delivery_cost       integer,
    goods_total         integer,
    custom_fee          integer
);

CREATE TABLE items (
    chrt_id             integer,
    track_number        varchar(100) REFERENCES orders(track_number),
    price               integer,
    rid                 varchar(100),
    name                varchar(100),
    sale                integer,
    size                varchar(50),
    total_price         integer,
    nm_id               integer,
    Brand               varchar(100),
    Status              integer
);

CREATE TABLE order_info (
    DelID               SERIAL UNIQUE,
    orderID             varchar(100) PRIMARY KEY,
    trackNumber         varchar(100) REFERENCES orders(track_number)
);