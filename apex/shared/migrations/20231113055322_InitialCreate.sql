-- Create "apex_customers" table
CREATE TABLE "apex_customers" ("id" character varying NOT NULL, "event_raised_at" timestamptz NOT NULL, "name" character varying NULL, "given_name" character varying NULL, "middle_name" character varying NULL, "family_name" character varying NULL, "photo_url" character varying NULL, "photo_url_24" character varying NULL, "photo_url_32" character varying NULL, "photo_url_48" character varying NULL, "photo_url_72" character varying NULL, "photo_url_192" character varying NULL, "photo_url_512" character varying NULL, "created_at" timestamptz NOT NULL, "modified_at" timestamptz NULL, "deleted_at" timestamptz NULL, PRIMARY KEY ("id"));
-- Create index "apexcustomer_deleted_at" to table: "apex_customers"
CREATE INDEX "apexcustomer_deleted_at" ON "apex_customers" ("deleted_at");
-- Create "apex_customer_identities" table
CREATE TABLE "apex_customer_identities" ("id" character varying NOT NULL, "email" character varying NULL, "email_verified" boolean NULL, "created_at" timestamptz NOT NULL, "modified_at" timestamptz NULL, "deleted_at" timestamptz NULL, "apex_customer_identities" character varying NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "apex_customer_identities_apex_customers_identities" FOREIGN KEY ("apex_customer_identities") REFERENCES "apex_customers" ("id") ON DELETE CASCADE);
-- Create index "apexcustomeridentity_email" to table: "apex_customer_identities"
CREATE INDEX "apexcustomeridentity_email" ON "apex_customer_identities" ("email");
-- Create index "apexcustomeridentity_email_verified" to table: "apex_customer_identities"
CREATE INDEX "apexcustomeridentity_email_verified" ON "apex_customer_identities" ("email_verified");
-- Create index "apexcustomeridentity_apex_customer_identities" to table: "apex_customer_identities"
CREATE INDEX "apexcustomeridentity_apex_customer_identities" ON "apex_customer_identities" ("apex_customer_identities");
-- Create index "apexcustomeridentity_deleted_at" to table: "apex_customer_identities"
CREATE INDEX "apexcustomeridentity_deleted_at" ON "apex_customer_identities" ("deleted_at");
