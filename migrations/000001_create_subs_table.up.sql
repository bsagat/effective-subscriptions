CREATE TABLE IF NOT EXISTS Subscriptions(
    ID UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  	Service_name TEXT NOT NULL,
  	Price INT NOT NULL CHECK(Price >= 0),
  	User_ID UUID NOT NULL,
  	Start_date TIMESTAMPTZ DEFAULT NOW(),
  	Exp_date TIMESTAMPTZ DEFAULT NOW(),
	CHECK (Exp_date IS NULL OR Start_date IS NULL OR Exp_date >= Start_date)
);
 
CREATE UNIQUE INDEX IF NOT EXISTS idx_subscriptions_user_service
    ON subscriptions(user_id, service_name);