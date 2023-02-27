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