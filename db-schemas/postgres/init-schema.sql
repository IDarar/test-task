create table if not exists users(
    "id" serial,
    "name" VARCHAR(24) UNIQUE,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT "scan_tx_pk" PRIMARY KEY ("id")
);