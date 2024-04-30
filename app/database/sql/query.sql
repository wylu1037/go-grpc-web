-- name: InsertLedgerData :one
insert into ledger_record(data_id, transaction_hash, business_name, business_contract_address, protocol_name,
                          protocol_uri, created_at, updated_at)
values (?, ?, ?, ?, ?, ?, ?, ?)
returning *;

-- name: UpdateLedgerRecord :exec
update ledger_record
set data_id                   = ?,
    transaction_hash          = ?,
    business_name             = ?,
    business_contract_address = ?,
    protocol_name             = ?,
    protocol_uri              = ?,
    created_at                = ?,
    updated_at                = ?
where id = ?;

-- name: DeleteLedgerRecord :exec
delete
from ledger_record
where id = ?;

-- name: CountLedgerRecord :one
select count(1)
from ledger_record;

-- name: GetLedgerRecordById :one
select id,
       data_id,
       transaction_hash,
       business_name,
       business_contract_address,
       protocol_name,
       protocol_uri,
       created_at,
       updated_at
from ledger_record
where id = ?
limit 1;

-- name: GetLedgerRecordByTransactionHash :one
select id,
       data_id,
       transaction_hash,
       business_name,
       business_contract_address,
       protocol_name,
       protocol_uri,
       created_at,
       updated_at
from ledger_record
where transaction_hash = ?
limit 1;
