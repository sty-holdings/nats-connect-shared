package nats_connect_shared

import (
	"time"
)

type GetTeamReply struct {
	Error struct {
		ErrorAdditionalInfo string `json:"error_additional_info"`
		ErrorFilename       string `json:"error_filename"`
		ErrorFunctionName   string `json:"error_function_name"`
		ErrorLineNumber     int    `json:"error_line_number"`
		ErrorMessage        string `json:"error_message"`
	} `json:"error"`
	Response struct {
		Created time.Time `json:"created"`
		Id      string    `json:"id"`
		Limits  struct {
			DisablePrometheusFederation bool `json:"disable_prometheus_federation"`
			NumUsers                    int  `json:"num_users"`
		} `json:"limits"`
		Name string `json:"name"`
	} `json:"response"`
}

type ListTeamsReply struct {
	Error struct {
		ErrorAdditionalInfo string `json:"error_additional_info"`
		ErrorFilename       string `json:"error_filename"`
		ErrorFunctionName   string `json:"error_function_name"`
		ErrorLineNumber     int    `json:"error_line_number"`
		ErrorMessage        string `json:"error_message"`
	} `json:"error"`
	Response struct {
		Items []struct {
			Created time.Time `json:"created"`
			Id      string    `json:"id"`
			Limits  struct {
				DisablePrometheusFederation bool `json:"disable_prometheus_federation"`
				NumUsers                    int  `json:"num_users"`
			} `json:"limits"`
			Name string `json:"name"`
		} `json:"items"`
	} `json:"response"`
}
