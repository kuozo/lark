package pkg


// Tweet API
func Tweet(token string) error{
	
	rels, err := GitHub()
	if err != nil {
		return err
	}
	err = ding(token, rels)
	if err != nil {
		return err
	}
	return nil
}