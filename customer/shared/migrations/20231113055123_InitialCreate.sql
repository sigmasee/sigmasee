-- Create "customer_outboxes" table
CREATE TABLE "customer_outboxes" ("id" character varying NOT NULL, "timestamp" timestamptz NOT NULL, "topic" character varying NOT NULL, "key" bytea NOT NULL, "payload" bytea NOT NULL, "headers" jsonb NOT NULL, "retry_count" bigint NOT NULL, "status" character varying NOT NULL, "last_retry" timestamptz NULL, "processing_errors" jsonb NULL, PRIMARY KEY ("id"));
-- Create index "customeroutbox_last_retry_status" to table: "customer_outboxes"
CREATE INDEX "customeroutbox_last_retry_status" ON "customer_outboxes" ("last_retry", "status");
-- Create "customers" table
CREATE TABLE "customers" ("id" character varying NOT NULL, "designation" character varying NULL, "title" character varying NULL, "name" character varying NULL, "given_name" character varying NULL, "middle_name" character varying NULL, "family_name" character varying NULL, "photo_url" character varying NULL, "photo_url_24" character varying NULL, "photo_url_32" character varying NULL, "photo_url_48" character varying NULL, "photo_url_72" character varying NULL, "photo_url_192" character varying NULL, "photo_url_512" character varying NULL, "timezone" character varying NULL, "locale" character varying NULL, "created_at" timestamptz NOT NULL, "modified_at" timestamptz NULL, "deleted_at" timestamptz NULL, PRIMARY KEY ("id"));
-- Create index "customer_designation" to table: "customers"
CREATE INDEX "customer_designation" ON "customers" ("designation");
-- Create index "customer_title" to table: "customers"
CREATE INDEX "customer_title" ON "customers" ("title");
-- Create index "customer_name" to table: "customers"
CREATE INDEX "customer_name" ON "customers" ("name");
-- Create index "customer_given_name" to table: "customers"
CREATE INDEX "customer_given_name" ON "customers" ("given_name");
-- Create index "customer_middle_name" to table: "customers"
CREATE INDEX "customer_middle_name" ON "customers" ("middle_name");
-- Create index "customer_family_name" to table: "customers"
CREATE INDEX "customer_family_name" ON "customers" ("family_name");
-- Create index "customer_deleted_at" to table: "customers"
CREATE INDEX "customer_deleted_at" ON "customers" ("deleted_at");
-- Create "customer_settings" table
CREATE TABLE "customer_settings" ("id" character varying NOT NULL, "created_at" timestamptz NOT NULL, "modified_at" timestamptz NULL, "deleted_at" timestamptz NULL, "customer_customer_settings" character varying NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "customer_settings_customers_customer_settings" FOREIGN KEY ("customer_customer_settings") REFERENCES "customers" ("id") ON DELETE CASCADE);
-- Create index "customer_settings_customer_customer_settings_key" to table: "customer_settings"
CREATE UNIQUE INDEX "customer_settings_customer_customer_settings_key" ON "customer_settings" ("customer_customer_settings");
-- Create index "customersetting_customer_customer_settings" to table: "customer_settings"
CREATE UNIQUE INDEX "customersetting_customer_customer_settings" ON "customer_settings" ("customer_customer_settings");
-- Create index "customersetting_deleted_at" to table: "customer_settings"
CREATE INDEX "customersetting_deleted_at" ON "customer_settings" ("deleted_at");
-- Create "identities" table
CREATE TABLE "identities" ("id" character varying NOT NULL, "email" character varying NULL, "email_verified" boolean NULL, "created_at" timestamptz NOT NULL, "modified_at" timestamptz NULL, "deleted_at" timestamptz NULL, "customer_identities" character varying NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "identities_customers_identities" FOREIGN KEY ("customer_identities") REFERENCES "customers" ("id") ON DELETE CASCADE);
-- Create index "identity_email" to table: "identities"
CREATE INDEX "identity_email" ON "identities" ("email");
-- Create index "identity_email_verified" to table: "identities"
CREATE INDEX "identity_email_verified" ON "identities" ("email_verified");
-- Create index "identity_customer_identities" to table: "identities"
CREATE INDEX "identity_customer_identities" ON "identities" ("customer_identities");
-- Create index "identity_deleted_at" to table: "identities"
CREATE INDEX "identity_deleted_at" ON "identities" ("deleted_at");
