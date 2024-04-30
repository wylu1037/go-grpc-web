-- 存证数据表
create table ledger_record
(
    -- 主键ID
    id                        integer primary key,
    -- 数据ID：用于从链上检索数据
    data_id                   varchar(66) not null,
    -- 交易哈希
    transaction_hash          varchar(66) not null,
    -- 业务名称
    business_name             varchar(60) not null,
    -- 业务合约地址： zltc_jNkDqNCKntZq5U4jX723r6b23tzULRD9s
    business_contract_address varchar(38) not null,
    -- 协议名称
    protocol_name             varchar(30) not null,
    -- 协议号
    protocol_uri              integer     not null,
    -- 创建时间
    created_at                date        not null,
    -- 更新时间
    updated_at                date        not null
);