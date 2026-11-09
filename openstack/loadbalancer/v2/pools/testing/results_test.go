package testing

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/gophercloud/gophercloud/openstack/loadbalancer/v2/pools"
)

func TestMemberUnmarshalJSONTimeFormats(t *testing.T) {
	cases := []struct {
		name        string
		payload     string
		wantCreated time.Time
		wantUpdated time.Time
		wantErr     bool
	}{
		{
			name: "older neutron format without timezone",
			payload: `{"id":"member-1","address":"192.168.153.10","protocol_port":80,
				"created_at":"2026-11-09T02:19:20","updated_at":"2026-11-09T02:19:21"}`,
			wantCreated: time.Date(2026, 11, 9, 2, 19, 20, 0, time.UTC),
			wantUpdated: time.Date(2026, 11, 9, 2, 19, 21, 0, time.UTC),
		},
		{
			name: "newer neutron format with UTC designator",
			payload: `{"id":"member-1","address":"192.168.153.10","protocol_port":80,
				"created_at":"2026-11-09T02:19:20Z","updated_at":"2026-11-09T02:19:21Z"}`,
			wantCreated: time.Date(2026, 11, 9, 2, 19, 20, 0, time.UTC),
			wantUpdated: time.Date(2026, 11, 9, 2, 19, 21, 0, time.UTC),
		},
		{
			name: "newer neutron format with numeric offset",
			payload: `{"id":"member-1","address":"192.168.153.10","protocol_port":80,
				"created_at":"2026-11-09T10:19:20+08:00","updated_at":"2026-11-09T10:19:21+08:00"}`,
			wantCreated: time.Date(2026, 11, 9, 2, 19, 20, 0, time.UTC),
			wantUpdated: time.Date(2026, 11, 9, 2, 19, 21, 0, time.UTC),
		},
		{
			name: "fractional seconds are preserved",
			payload: `{"id":"member-1","address":"192.168.153.10","protocol_port":80,
				"created_at":"2026-11-09T02:19:20.123456Z","updated_at":"2026-11-09T02:19:21.654321Z"}`,
			wantCreated: time.Date(2026, 11, 9, 2, 19, 20, 123456000, time.UTC),
			wantUpdated: time.Date(2026, 11, 9, 2, 19, 21, 654321000, time.UTC),
		},
		{
			name:    "missing timestamps",
			payload: `{"id":"member-1","address":"192.168.153.10","protocol_port":80}`,
		},
		{
			name: "null timestamps",
			payload: `{"id":"member-1","address":"192.168.153.10","protocol_port":80,
				"created_at":null,"updated_at":null}`,
		},
		{
			name: "empty timestamps",
			payload: `{"id":"member-1","address":"192.168.153.10","protocol_port":80,
				"created_at":"","updated_at":""}`,
		},
		{
			name: "invalid timestamp",
			payload: `{"id":"member-1","address":"192.168.153.10","protocol_port":80,
				"created_at":"not-a-time","updated_at":"2026-11-09T02:19:21Z"}`,
			wantErr: true,
		},
		{
			name: "unrelated field type error is not swallowed",
			payload: `{"id":"member-1","address":"192.168.153.10","protocol_port":"eighty",
				"created_at":"2026-11-09T02:19:20Z","updated_at":"2026-11-09T02:19:21Z"}`,
			wantErr: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var member pools.Member
			err := json.Unmarshal([]byte(tc.payload), &member)
			if tc.wantErr {
				if err == nil {
					t.Fatal("expected an error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unmarshal Member: %v", err)
			}

			if member.ID != "member-1" {
				t.Errorf("ID = %q, want %q", member.ID, "member-1")
			}
			if member.Address != "192.168.153.10" {
				t.Errorf("Address = %q, want %q", member.Address, "192.168.153.10")
			}
			if member.ProtocolPort != 80 {
				t.Errorf("ProtocolPort = %d, want 80", member.ProtocolPort)
			}
			if !member.CreatedAt.Equal(tc.wantCreated) {
				t.Errorf("CreatedAt = %v, want %v", member.CreatedAt, tc.wantCreated)
			}
			if !member.UpdatedAt.Equal(tc.wantUpdated) {
				t.Errorf("UpdatedAt = %v, want %v", member.UpdatedAt, tc.wantUpdated)
			}
		})
	}
}
