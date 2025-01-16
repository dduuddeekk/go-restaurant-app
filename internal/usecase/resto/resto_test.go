package resto

import (
	"context"
	"reflect"
	"testing"

	"github.com/dduuddeekk/go-restaurant-app/internal/model"
)

func Test_restoUsecase_GetMenuList(t *testing.T) {
	type args struct {
		ctx      context.Context
		menuType string
	}
	tests := []struct {
		name    string
		r       *restoUsecase
		args    args
		want    []model.MenuItem
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.GetMenuList(tt.args.ctx, tt.args.menuType)
			if (err != nil) != tt.wantErr {
				t.Errorf("restoUsecase.GetMenuList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("restoUsecase.GetMenuList() = %v, want %v", got, tt.want)
			}
		})
	}
}
