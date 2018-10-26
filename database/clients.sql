CREATE OR REPLACE TABLE clients
(
  cpf                   VARCHAR(20),
  last_purchase_store   VARCHAR(20),
  most_frequent_store   VARCHAR(20),
  private               BOOLEAN,
  incomplete            BOOLEAN,
  last_purchase         TIMESTAMP WITH TIME ZONE,
  medium_purchase_value NUMERIC,
  last_pruchase_value   NUMERIC
);

CREATE OR REPLACE INDEX idx_clients_cpf
  ON clients (cpf);

CREATE OR REPLACE INDEX idx_clients_last_purchase_store
  ON clients (last_purchase_store);

CREATE OR REPLACE INDEX idx_clients_most_frequent_store
  ON clients (most_frequent_store);

CREATE OR REPLACE INDEX idx_clients_medium_purchase_value
  ON clients (medium_purchase_value);

CREATE OR REPLACE INDEX idx_clients_last_pruchase_value
  ON clients (last_pruchase_value);

