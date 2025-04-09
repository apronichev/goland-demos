package main

import (
	"testing"
	"time"
)

func TestVerifyPassword(t *testing.T) {
	tests := []struct {
		name           string
		password       string
		wantSixOrMore  bool
		wantNumber     bool
		wantUpper      bool
		wantSpecial    bool
		expectedResult bool
	}{
		{
			name:          "Valid password with all requirements",
			password:      "myZelda@123",
			wantSixOrMore: true,
			wantNumber:    true,
			wantUpper:     true,
			wantSpecial:   true,
		},
		{
			name:          "Password too short",
			password:      "Ab1",
			wantSixOrMore: false,
			wantNumber:    true,
			wantUpper:     true,
			wantSpecial:   true,
		},
		{
			name:          "Password without numbers",
			password:      "myZeldaPass@",
			wantSixOrMore: true,
			wantNumber:    false,
			wantUpper:     true,
			wantSpecial:   true,
		},
		{
			name:          "Password without uppercase",
			password:      "myzelda@123",
			wantSixOrMore: true,
			wantNumber:    true,
			wantUpper:     false,
			wantSpecial:   true,
		},
		{
			name:          "Password without special chars",
			password:      "myZelda123",
			wantSixOrMore: true,
			wantNumber:    true,
			wantUpper:     true,
			wantSpecial:   false,
		},
		{
			name:          "Empty password",
			password:      "",
			wantSixOrMore: false,
			wantNumber:    false,
			wantUpper:     false,
			wantSpecial:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSixOrMore, gotNumber, gotUpper, gotSpecial := verifyPassword(tt.password)

			if gotSixOrMore != tt.wantSixOrMore {
				t.Errorf("verifyPassword() sixOrMore = %v, want %v", gotSixOrMore, tt.wantSixOrMore)
			}
			if gotNumber != tt.wantNumber {
				t.Errorf("verifyPassword() number = %v, want %v", gotNumber, tt.wantNumber)
			}
			if gotUpper != tt.wantUpper {
				t.Errorf("verifyPassword() upper = %v, want %v", gotUpper, tt.wantUpper)
			}
			if gotSpecial != tt.wantSpecial {
				t.Errorf("verifyPassword() special = %v, want %v", gotSpecial, tt.wantSpecial)
			}
		})
	}
}

func TestTellUserPassword(t *testing.T) {
	tests := []struct {
		name             string
		userJSON         string
		registrationDate time.Time
		wantError        bool
	}{
		{
			name: "Valid JSON input",
			userJSON: `{
				"email": "test@example.com",
				"password": "Test@123",
				"first_name": "Test",
				"last_name": "User"
			}`,
			registrationDate: time.Date(2022, 11, 1, 0, 0, 0, 0, time.UTC),
			wantError:        false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Since tellUserPassword doesn't return anything, we can only test that it doesn't panic
			defer func() {
				if r := recover(); (r != nil) != tt.wantError {
					t.Errorf("tellUserPassword() panic = %v, wantError %v", r, tt.wantError)
				}
			}()

			tellUserPassword(tt.userJSON, tt.registrationDate)
		})
	}
}
