package models

import "github.com/supabase-community/supabase-go"

type SupabaseRepo struct {
	supabaseClient    *supabase.Client // anon client
	serviceRoleClient *supabase.Client // service role client
	url               string
	anonKey           string
	serviceKey        string
}

func SupabaseNewRepo(supabaseClient *supabase.Client, url, anonKey, serviceKey string) *SupabaseRepo {
	var serviceClient *supabase.Client
	if serviceKey != "" {
		serviceClient, _ = supabase.NewClient(url, serviceKey, nil)
	}

	return &SupabaseRepo{
		supabaseClient:    supabaseClient,
		serviceRoleClient: serviceClient,
		url:               url,
		anonKey:           anonKey,
		serviceKey:        serviceKey,
	}
}
