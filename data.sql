CREATE TABLE IF NOT EXISTS wallet
(
    id integer primary key NOT NULL,
    address text,
    balance real
);

CREATE TABLE IF NOT EXISTS transaction_history
(
    id integer primary key NOT NULL,
    from_id integer NOT NULL,
    to_id integer NOT NULL,
    transfer_amount real not null,
    transfer_time datetime not null,
    FOREIGN KEY(from_id) REFERENCES wallet(id)
);

insert into wallet (address, balance) VALUES ('e240d825d255af751f5f55af8d9671beabdf2236c0a3b4e2639b3e182d994c88', 3.50);

insert into wallet (address, balance) VALUES ('e240d825d255af751f5f55af8d9671beabdf389123ewqedq34dd3q231efwer23', 2.00)