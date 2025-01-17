package resto

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/dduuddeekk/go-restaurant-app/internal/mocks"
	"github.com/dduuddeekk/go-restaurant-app/internal/model"
	"github.com/dduuddeekk/go-restaurant-app/internal/model/constant"
	"github.com/dduuddeekk/go-restaurant-app/internal/repository/menu"
	"github.com/golang/mock/gomock"
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
		{
			name: "success get menu list",
			r: &restoUsecase{
				menuRepo: func() menu.Repository {
					ctrl := gomock.NewController(t)
					mock := mocks.NewMockMenuRepository(ctrl)

					mock.EXPECT().GetMenuList(gomock.Any(), string(constant.MenuTypeFood)).
						Times(1).
						Return([]model.MenuItem{
							{
								OrderCode: "babiguling",
								Name:      "Babi Guling",
								Price:     30000,
								Type:      constant.MenuTypeFood,
							},
						}, nil)

					return mock
				}(),
			},
			args: args{
				ctx:      context.Background(),
				menuType: string(constant.MenuTypeFood),
			},
			want: []model.MenuItem{
				{
					OrderCode: "babiguling",
					Name:      "Babi Guling",
					Price:     30000,
					Type:      constant.MenuTypeFood,
				},
			},
			wantErr: false,
		},
		{
			name: "success get menu list empty",
			r: &restoUsecase{
				menuRepo: func() menu.Repository {
					ctrl := gomock.NewController(t)
					mock := mocks.NewMockMenuRepository(ctrl)

					mock.EXPECT().GetMenuList(gomock.Any(), string(constant.MenuTypeFood)).
						Times(1).
						Return(nil, errors.New("mock errors"))

					return mock
				}(),
			},
			args: args{
				ctx:      context.Background(),
				menuType: string(constant.MenuTypeFood),
			},
			want:    nil,
			wantErr: true,
		},
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
