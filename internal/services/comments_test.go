package services

import (
	"home/zellie/Code/guestbook-api/internal/interfaces"
	"home/zellie/Code/guestbook-api/internal/models"
	"reflect"
	"testing"

	"github.com/rs/zerolog/log"
)

type mockCommentsRepo struct {
}

func (r mockCommentsRepo) InsertComment(req *models.Comment) error {
	return nil
}

func (r mockCommentsRepo) GetAllComments() (*models.Comments, error) {
	return &models.Comments{}, nil
}

func (r mockCommentsRepo) DeleteComment(req string) error {
	return nil
}

func Test_NewCommentsService(t *testing.T) {
	type args struct {
		repo interfaces.CommentsRepo
	}

	testRepo := mockCommentsRepo{}
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
				t.Errorf("NewCommentsService did not set repo properly; want: '%v', got: '%v'", tt.want, got.repo)
			}

		})
	}

}

func Test_InsertComment(t *testing.T) {
	type args struct {
		c *models.Comment
	}

	testRepo := mockCommentsRepo{}
	testService, _ := NewCommentsService(testRepo)

	mockComment := models.Comment{
		Description: "test description",
	}

	mockEmptyDescriptionComment := models.Comment{
		Description: "",
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "should post comment successfully",
			args: args{
				c: &mockComment,
			},
			wantErr: false,
		},
		{
			name: "should return an error message: 'comment description required but was empty'",
			args: args{
				c: &mockEmptyDescriptionComment,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := testService.InsertComment(tt.args.c)

			if err != nil {
				if !tt.wantErr {
					log.Err(err).Str("package", "services").Msgf("%v", err)

					t.Errorf("InsertComment had an unexpected error: %s ", err)
					return
				}
			}

			if err == nil && tt.args.c.CommentID == "" {
				t.Errorf("InsertComment did not insert comment populate commentid")
				return
			}

		})
	}
}

func Test_GetAllComments(t *testing.T) {
	testRepo := mockCommentsRepo{}
	testService, _ := NewCommentsService(testRepo)

	mockComments := &models.Comments{}

	tests := []struct {
		name    string
		want    *models.Comments
		wantErr bool
	}{
		{
			name:    "should return comments struct successfully",
			want:    mockComments,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := testService.GetAllComments()

			if err != nil {
				if !tt.wantErr {
					log.Err(err).Str("package", "services").Msgf("%v", err)

					t.Errorf("GetAllComments had an unexpected error: %s ", err)
					return
				}
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllComments got: %v, but expected: %v", got, tt.want)
				return
			}
		})
	}
}
