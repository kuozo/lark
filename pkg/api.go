package pkg

import (
	"errors"
)


// Tweet API
func Tweet(token string) error{
	
	rels, err := GitHub()
	if err != nil {
		return err
	}
	if len(rels) < 1 {
		return errors.New("It's empty releases")
	}
	err = ding(token, rels)
	if err != nil {
		return err
	}
	return nil
}