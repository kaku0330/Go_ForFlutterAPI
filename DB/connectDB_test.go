package DB

import (
	"reflect"
	"testing"

	"gorm.io/gorm"
)

func TestConnectDB(t *testing.T) {
	tests := []struct {
		name string
		want *gorm.DB
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConnectDB(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConnectDB() = %v, want %v", got, tt.want)
			}
		})
	}
}
