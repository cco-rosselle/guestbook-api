package services

import (
	"fmt"
	"home/zellie/Code/guestbook-api/internal/interfaces"
	"home/zellie/Code/guestbook-api/internal/models"
	"testing"
)

// mock repo
type noOpCommentsRepo struct{}

func (r noOpCommentsRepo) InsertComment(req *models.Comment) error {
	return fmt.Errorf("Not implemented")
}

func (r noOpCommentsRepo) GetAllComments() (*models.Comments, error) {
	return nil, fmt.Errorf("Not implemented")
}

func (r noOpCommentsRepo) DeleteComment(req string) error {
	return nil
}

func Test_NewCommentsService(t *testing.T) {
	type args struct {
		repo interfaces.CommentsRepo
	}

	testRepo := noOpCommentsRepo{}
	tests := []struct {
		name    string
		args    args
		want    interfaces.CommentsRepo
		wantErr bool
	}{
		{
			name: "should return a CommentsService",
			args: args{
				repo: testRepo,
			},
			want:    testRepo,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewCommentsService(tt.args.repo)

			if err != nil && !tt.wantErr {
				t.Errorf("NewCommesService had unexpected error: %s ", err)
				return
			}

			if got.repo != tt.want {
				t.Errorf("NewCommentsService did not set repo properly; want: '%v', got '%v'", tt.want, got.repo)
			}

		})
	}

}
