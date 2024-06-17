CREATE TABLE "items" (
  "id" bigserial PRIMARY KEY,
  "item" varchar NOT NULL
);

CREATE INDEX ON "items" ("item");