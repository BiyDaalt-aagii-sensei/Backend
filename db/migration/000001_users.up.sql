-- Table "users"

BEGIN;

CREATE TABLE "users" (
    "Id" BIGSERIAL NOT NULL,
    "Username" VARCHAR NOT NULL,
    "Password" VARCHAR NOT NULL,
    "FirstName" VARCHAR NOT NULL,
    "LastName" VARCHAR NOT NULL,
    "Email" VARCHAR NOT NULL,
    "Password_At_Change" TIMESTAMPTZ NOT NULL DEFAULT '0001-01-01 00:00:00Z',
    "Created_At" TIMESTAMPTZ NOT NULL DEFAULT (now()),
    CONSTRAINT "User_Pk" PRIMARY KEY ("Id"),
    CONSTRAINT "User_Unique" UNIQUE ("Username","Email")
)TABLESPACE pg_default;

COMMIT;