package postgres

import "context"

func (a *API) Migrate() error {
	query := `
	CREATE TABLE IF NOT EXISTS Subscriptions(
  		ID UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  		Service_name TEXT NOT NULL,
  		Price INT NOT NULL,
  		User_ID UUID NOT NULL,
  		Start_date TIMESTAMPTZ,
  		Exp_date TIMESTAMPTZ
	); 
	CREATE UNIQUE INDEX IF NOT EXISTS idx_subscriptions_user_service
    ON subscriptions(user_id, service_name);
	`

	_, err := a.DB.Exec(context.Background(), query)
	if err != nil {
		return err
	}
	return nil
}
