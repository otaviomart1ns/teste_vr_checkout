CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS transactions (
  id          UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  description VARCHAR(50) NOT NULL,
  date        DATE        NOT NULL,
  amount      NUMERIC(7,2) NOT NULL CHECK (amount > 0 AND amount <= 99999.99)
);