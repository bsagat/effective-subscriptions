CREATE TABLE IF NOT EXISTS Subscriptions(
  ID UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  Service_name TEXT NOT NULL,
  Price INT NOT NULL,
  User_ID UUID DEFAULT gen_random_uuid(),
  Start_date TIMESTAMPTZ,
  Exp_date TIMESTAMPTZ
); 

CREATE INDEX IF NOT EXISTS idx_subscriptions_user_id ON subscriptions(user_id);

CREATE INDEX IF NOT EXISTS idx_subscriptions_service_name ON subscriptions(Service_name);