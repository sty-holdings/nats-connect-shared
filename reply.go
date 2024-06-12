package nats_connect_shared

import (
	"time"
)

type GetSystemReply struct {
	Error struct {
		ErrorAdditionalInfo string `json:"error_additional_info"`
		ErrorFilename       string `json:"error_filename"`
		ErrorFunctionName   string `json:"error_function_name"`
		ErrorLineNumber     int    `json:"error_line_number"`
		ErrorMessage        string `json:"error_message"`
	} `json:"error"`
	Response struct {
		ConnectionType          string    `json:"connection_type"`
		Created                 time.Time `json:"created"`
		HasManagedOperator      bool      `json:"has_managed_operator"`
		HasManagedSystemAccount bool      `json:"has_managed_system_account"`
		HostSystemId            string    `json:"host_system_id"`
		Id                      string    `json:"id"`
		IsTenant                bool      `json:"is_tenant"`
		JetstreamDomain         string    `json:"jetstream_domain"`
		JetstreamEnabled        bool      `json:"jetstream_enabled"`
		JetstreamTiers          []string  `json:"jetstream_tiers"`
		Name                    string    `json:"name"`
		State                   string    `json:"state"`
		Team                    struct {
			Id   string `json:"id"`
			Name string `json:"name"`
		} `json:"team"`
		UserJwtExpiresInSecs int `json:"user_jwt_expires_in_secs"`
	} `json:"response"`
}

type GetSystemLimitsReply struct {
	Error struct {
		ErrorAdditionalInfo string `json:"error_additional_info"`
		ErrorFilename       string `json:"error_filename"`
		ErrorFunctionName   string `json:"error_function_name"`
		ErrorLineNumber     int    `json:"error_line_number"`
		ErrorMessage        string `json:"error_message"`
	} `json:"error"`
	Response struct {
		Allocated struct {
			NumAccounts    int `json:"num_accounts"`
			ResourceLimits struct {
				Conn               int  `json:"conn"`
				Consumer           int  `json:"consumer"`
				Data               int  `json:"data"`
				DisallowBearer     bool `json:"disallow_bearer"`
				DiskMaxStreamBytes int  `json:"disk_max_stream_bytes"`
				DiskStorage        int  `json:"disk_storage"`
				Exports            int  `json:"exports"`
				Imports            int  `json:"imports"`
				Leaf               int  `json:"leaf"`
				MaxAckPending      int  `json:"max_ack_pending"`
				MaxBytesRequired   bool `json:"max_bytes_required"`
				MemMaxStreamBytes  int  `json:"mem_max_stream_bytes"`
				MemStorage         int  `json:"mem_storage"`
				Payload            int  `json:"payload"`
				Streams            int  `json:"streams"`
				Subs               int  `json:"subs"`
				TieredLimits       struct {
					R1 struct {
						Consumer           int   `json:"consumer"`
						DiskMaxStreamBytes int64 `json:"disk_max_stream_bytes"`
						DiskStorage        int64 `json:"disk_storage"`
						MaxAckPending      int   `json:"max_ack_pending"`
						MaxBytesRequired   bool  `json:"max_bytes_required"`
						MemMaxStreamBytes  int   `json:"mem_max_stream_bytes"`
						MemStorage         int   `json:"mem_storage"`
						Streams            int   `json:"streams"`
					} `json:"R1"`
				} `json:"tiered_limits"`
				Wildcards bool `json:"wildcards"`
			} `json:"resource_limits"`
		} `json:"allocated"`
		Available struct {
			NumAccounts    int `json:"num_accounts"`
			ResourceLimits struct {
				Conn               int  `json:"conn"`
				Consumer           int  `json:"consumer"`
				Data               int  `json:"data"`
				DisallowBearer     bool `json:"disallow_bearer"`
				DiskMaxStreamBytes int  `json:"disk_max_stream_bytes"`
				DiskStorage        int  `json:"disk_storage"`
				Exports            int  `json:"exports"`
				Imports            int  `json:"imports"`
				Leaf               int  `json:"leaf"`
				MaxAckPending      int  `json:"max_ack_pending"`
				MaxBytesRequired   bool `json:"max_bytes_required"`
				MemMaxStreamBytes  int  `json:"mem_max_stream_bytes"`
				MemStorage         int  `json:"mem_storage"`
				Payload            int  `json:"payload"`
				Streams            int  `json:"streams"`
				Subs               int  `json:"subs"`
				TieredLimits       struct {
					R1 struct {
						Consumer           int   `json:"consumer"`
						DiskMaxStreamBytes int64 `json:"disk_max_stream_bytes"`
						DiskStorage        int   `json:"disk_storage"`
						MaxAckPending      int   `json:"max_ack_pending"`
						MaxBytesRequired   bool  `json:"max_bytes_required"`
						MemMaxStreamBytes  int   `json:"mem_max_stream_bytes"`
						MemStorage         int   `json:"mem_storage"`
						Streams            int   `json:"streams"`
					} `json:"R1"`
				} `json:"tiered_limits"`
				Wildcards bool `json:"wildcards"`
			} `json:"resource_limits"`
		} `json:"available"`
		Total struct {
			NumAccounts    int `json:"num_accounts"`
			ResourceLimits struct {
				Conn               int  `json:"conn"`
				Consumer           int  `json:"consumer"`
				Data               int  `json:"data"`
				DisallowBearer     bool `json:"disallow_bearer"`
				DiskMaxStreamBytes int  `json:"disk_max_stream_bytes"`
				DiskStorage        int  `json:"disk_storage"`
				Exports            int  `json:"exports"`
				Imports            int  `json:"imports"`
				Leaf               int  `json:"leaf"`
				MaxAckPending      int  `json:"max_ack_pending"`
				MaxBytesRequired   bool `json:"max_bytes_required"`
				MemMaxStreamBytes  int  `json:"mem_max_stream_bytes"`
				MemStorage         int  `json:"mem_storage"`
				Payload            int  `json:"payload"`
				Streams            int  `json:"streams"`
				Subs               int  `json:"subs"`
				TieredLimits       struct {
					R1 struct {
						Consumer           int   `json:"consumer"`
						DiskMaxStreamBytes int64 `json:"disk_max_stream_bytes"`
						DiskStorage        int64 `json:"disk_storage"`
						MaxAckPending      int   `json:"max_ack_pending"`
						MaxBytesRequired   bool  `json:"max_bytes_required"`
						MemMaxStreamBytes  int   `json:"mem_max_stream_bytes"`
						MemStorage         int   `json:"mem_storage"`
						Streams            int   `json:"streams"`
					} `json:"R1"`
				} `json:"tiered_limits"`
				Wildcards bool `json:"wildcards"`
			} `json:"resource_limits"`
		} `json:"total"`
	} `json:"response"`
}
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

type ListPersonalAccessTokensReply struct {
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

type ListSystemsReply struct {
	Error struct {
		ErrorAdditionalInfo string `json:"error_additional_info"`
		ErrorFilename       string `json:"error_filename"`
		ErrorFunctionName   string `json:"error_function_name"`
		ErrorLineNumber     int    `json:"error_line_number"`
		ErrorMessage        string `json:"error_message"`
	} `json:"error"`
	Response struct {
		Items []struct {
			ConnectionType          string    `json:"connection_type"`
			Created                 time.Time `json:"created"`
			HasManagedOperator      bool      `json:"has_managed_operator"`
			HasManagedSystemAccount bool      `json:"has_managed_system_account"`
			HostSystemId            string    `json:"host_system_id"`
			Id                      string    `json:"id"`
			IsTenant                bool      `json:"is_tenant"`
			JetstreamDomain         string    `json:"jetstream_domain"`
			JetstreamEnabled        bool      `json:"jetstream_enabled"`
			JetstreamTiers          []string  `json:"jetstream_tiers"`
			Name                    string    `json:"name"`
			State                   string    `json:"state"`
			Team                    struct {
				Id   string `json:"id"`
				Name string `json:"name"`
			} `json:"team"`
			UserJwtExpiresInSecs int `json:"user_jwt_expires_in_secs"`
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

type ListTeamServerAccountsReply struct {
	Error struct {
		ErrorAdditionalInfo string `json:"error_additional_info"`
		ErrorFilename       string `json:"error_filename"`
		ErrorFunctionName   string `json:"error_function_name"`
		ErrorLineNumber     int    `json:"error_line_number"`
		ErrorMessage        string `json:"error_message"`
	} `json:"error"`
}
