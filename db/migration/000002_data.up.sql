-- TABLE: "data"
BEGIN;

CREATE TABLE "data" (
    "gender" VARCHAR NOT NULL,
    "age" INT NOT NULL,
    "Investment_Avenues" VARCHAR NOT NULL,
    "Mutual_Funds" INT NOT NULL,
    "Equity_Market" INT NOT NULL,
    "Debentures" INT NOT NULL,
    "Government_Bonds"INT NOT NULL,
    "Fixed_Deposits" INT NOT NULL,
    "PPF" INT NOT NULL,
    "Gold" INT NOT NULL,
    "Stock_Marktet" VARCHAR NOT NULL,
    "Factor" VARCHAR NOT NULL,
    "Objective" VARCHAR NOT NULL,
    "Purpose" VARCHAR NOT NULL,
    "Duration" VARCHAR NOT NULL,
    "Invest_Monitor" VARCHAR NOT NULL,
    "Expect" VARCHAR NOT NULL,
    "Avenue" VARCHAR NOT NULL,
    "What are your savings objectives?" VARCHAR NOT NULL,
    "Reason_Equity" VARCHAR NOT NULL,
    "Reason_Mutual" VARCHAR NOT NULL,
    "Reason_Bonds" VARCHAR NOT NULL,
    "Reason_FD" VARCHAR NOT NULL,
    "Source" VARCHAR NOT NULL 
) TABLESPACE pg_default;

COMMIT;