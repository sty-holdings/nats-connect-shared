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

type GetTeamLimitsReply struct {
	Error struct {
		ErrorAdditionalInfo string `json:"error_additional_info"`
		ErrorFilename       string `json:"error_filename"`
		ErrorFunctionName   string `json:"error_function_name"`
		ErrorLineNumber     int    `json:"error_line_number"`
		ErrorMessage        string `json:"error_message"`
	} `json:"error"`
	Response struct {
		Allocated struct {
			DisablePrometheusFederation bool `json:"disable_prometheus_federation"`
			NumUsers                    int  `json:"num_users"`
		} `json:"allocated"`
		Available struct {
			DisablePrometheusFederation bool `json:"disable_prometheus_federation"`
			NumUsers                    int  `json:"num_users"`
		} `json:"available"`
		Total struct {
			DisablePrometheusFederation bool `json:"disable_prometheus_federation"`
			NumUsers                    int  `json:"num_users"`
		} `json:"total"`
	} `json:"response"`
}

type GetVersionReply struct {
	Error struct {
		ErrorAdditionalInfo string `json:"error_additional_info"`
		ErrorFilename       string `json:"error_filename"`
		ErrorFunctionName   string `json:"error_function_name"`
		ErrorLineNumber     int    `json:"error_line_number"`
		ErrorMessage        string `json:"error_message"`
	} `json:"error"`
	Response struct {
		Commit  string    `json:"commit"`
		Date    time.Time `json:"date"`
		Version string    `json:"version"`
	} `json:"response"`
}

type ListInfoAppUsersTeamReply struct {
	Error struct {
		ErrorAdditionalInfo string `json:"error_additional_info"`
		ErrorFilename       string `json:"error_filename"`
		ErrorFunctionName   string `json:"error_function_name"`
		ErrorLineNumber     int    `json:"error_line_number"`
		ErrorMessage        string `json:"error_message"`
	} `json:"error"`
	Response struct {
		Items []struct {
			AppUser struct {
				Id         string `json:"id"`
				Identifier string `json:"identifier"`
				Name       string `json:"name"`
			} `json:"app_user"`
			Created           time.Time `json:"created"`
			Id                string    `json:"id"`
			PendingInvitation bool      `json:"pending_invitation"`
			RoleId            string    `json:"role_id"`
			RoleName          string    `json:"role_name"`
			Team              struct {
				Id   string `json:"id"`
				Name string `json:"name"`
			} `json:"team"`
		} `json:"items"`
	} `json:"response"`
}

type ListPersonalAccessTokenReply struct {
	Error struct {
		ErrorAdditionalInfo string `json:"error_additional_info"`
		ErrorFilename       string `json:"error_filename"`
		ErrorFunctionName   string `json:"error_function_name"`
		ErrorLineNumber     int    `json:"error_line_number"`
		ErrorMessage        string `json:"error_message"`
	} `json:"error"`
	Response struct {
		Items []struct {
			Created    time.Time `json:"created"`
			Id         string    `json:"id"`
			Name       string    `json:"name"`
			Expires    time.Time `json:"expires,omitempty"`
			LastActive time.Time `json:"last_active,omitempty"`
		} `json:"items"`
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
