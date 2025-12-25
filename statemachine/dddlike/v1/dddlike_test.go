package dddlike

import (
	"testing"
)

func TestBook_Arrived(t *testing.T) {
	tests := []struct {
		testName string // description of this test case
		// Named input parameters for receiver constructor.
		name    string
		wantErr bool
	}{
		{
			testName: "在庫なしの状態で到着したら在庫有りに変わる",
			name : "在庫ある本",
			wantErr : false,
		},
		// {
		// 	testName: "在庫ありの状態で到着したらエラー",
		// 	name: "在庫ある本",
		// 	wantErr: true,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			b := NewBook(tt.name)
			gotErr := b.Arrived()
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("Arrived() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("Arrived() succeeded unexpectedly")
			}
		})
	}
}

func TestBook_Sold(t *testing.T) {
	tests := []struct {
		testName string // description of this test case
		// Named input parameters for receiver constructor.
		name    string
		wantErr bool
	}{
		{
			testName: "在庫ありで売れたらOK",
			name: "在庫あり本",
			wantErr: false,
		},
		// {
		// 	testName: "在庫なしで売れたらエラー",
		// 	name: "在庫なし本",
		// 	wantErr: true,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			b := NewBook(tt.name)
			// 状態変化のために実行
			b.Arrived()
			gotErr := b.Sold()
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("Sold() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("Sold() succeeded unexpectedly")
			}
		})
	}
}
