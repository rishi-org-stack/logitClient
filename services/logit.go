package services

import (
	"encoding/json"
	"logClient/lib/logit"
	ta "logClient/types/auth"
)

type LogitService struct{}

func (ls *LogitService) GetToken(auth *ta.AuthRequest, output chan *ta.AuthResponse, exit chan bool) error {
	res, err := logit.GetToken(auth)
	if err != nil {
		return err
	}
	// fmt.Println("res=>", res)
	ar := &ta.AuthResponse{}
	err = json.Unmarshal(res, ar)

	if err != nil {
		return err
	}
	for {
		select {
		case output <- ar:
		// output <- ar
		case b := <-exit:
			if b {
				return nil

			}

		}
	}

	// return nil
}
