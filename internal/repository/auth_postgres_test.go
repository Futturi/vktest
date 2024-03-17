package repository

import (
	"testing"

	"github.com/Futturi/vktest/internal/models"
	"github.com/magiconair/properties/assert"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
)

func TestAuthPostgres_SignUp(t *testing.T) {

	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Log(err)
	}
	defer db.Close()

	r := NewAuthRepo(db)

	type mockBehaviour func(user models.User)

	testTable := []struct {
		name          string
		username      string
		password      string
		mockBehaviour mockBehaviour
		wantErr       bool
	}{
		{
			name:     "OK",
			username: "username",
			password: "password",
			mockBehaviour: func(user models.User) {
				mock.ExpectExec("INSERT INTO").WithArgs(user.Username, user.Password).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			wantErr: false,
		},
		{
			name:     "Error",
			username: "username",
			password: "",
			mockBehaviour: func(user models.User) {
				mock.ExpectExec("INSERT INTO").WithArgs(user.Username, user.Password).WillReturnError(err)
			},
			wantErr: true,
		},
	}

	for _, testCase := range testTable {

		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehaviour(models.User{Username: testCase.username, Password: testCase.password})
			got, err := r.SignUp(models.User{Username: testCase.username, Password: testCase.password})
			if testCase.wantErr {
				if err == nil {
					t.Errorf("expected error got nil")
				}
			} else {
				if err == nil {
					t.Errorf("expected nil got error")
				}
			}
			assert.Equal(t, got, 0)
		})
	}
}

func TestAuthPostgres_SignIn(t *testing.T) {

	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Log(err)
	}
	defer db.Close()

	r := NewAuthRepo(db)

	type mockBehaviour func(user models.User)

	testTable := []struct {
		name          string
		username      string
		password      string
		id            int
		mockBehaviour mockBehaviour
		wantErr       bool
	}{
		{
			name:     "OK",
			username: "username",
			password: "password",
			id:       1,
			mockBehaviour: func(user models.User) {
				mock.ExpectQuery("SELECT id FROM").WithArgs(user.Username, user.Password).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
			},
			wantErr: false,
		},
		{
			name:     "Error",
			username: "username",
			password: "",
			id:       0,
			mockBehaviour: func(user models.User) {
				mock.ExpectQuery("SELECT id FROM").WithArgs(user.Username, user.Password).WillReturnError(err)
			},
			wantErr: true,
		},
	}

	for _, testCase := range testTable {

		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehaviour(models.User{Username: testCase.username, Password: testCase.password})
			got, err := r.SignIn(models.User{Username: testCase.username, Password: testCase.password})
			if testCase.wantErr {
				if err == nil {
					t.Errorf("expected error got nil")
				}
			} else {
				if err != nil {
					t.Errorf("expected nil got error")
				}
			}
			assert.Equal(t, got, testCase.id)
		})
	}
}

func TestActorRepository_NewAuthRepo(t *testing.T) {
	db, _, err := sqlmock.Newx()
	if err != nil {
		t.Log(err)
	}
	defer db.Close()
	r := NewAuthRepo(db)
	assert.Equal(t, db, r.db)
}

func TestAuthPostgres_SignUpAdmin(t *testing.T) {

	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Log(err)
	}
	defer db.Close()

	r := NewAuthRepo(db)

	type mockBehaviour func(user models.Admin)

	testTable := []struct {
		name          string
		username      string
		id            int
		password      string
		mockBehaviour mockBehaviour
		wantErr       bool
	}{
		{
			name:     "OK",
			username: "username",
			id:       0,
			password: "password",
			mockBehaviour: func(user models.Admin) {
				mock.ExpectExec("INSERT INTO").WithArgs(user.Username, user.Password).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			wantErr: false,
		},
		{
			name:     "Error",
			username: "username",
			id:       0,
			password: "",
			mockBehaviour: func(user models.Admin) {
				mock.ExpectExec("INSERT INTO").WithArgs(user.Username, user.Password).WillReturnError(err)
			},
			wantErr: true,
		},
	}

	for _, testCase := range testTable {

		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehaviour(models.Admin{Username: testCase.username, Password: testCase.password})
			got, err := r.SignUpAdmin(models.Admin{Username: testCase.username, Password: testCase.password})
			if testCase.wantErr {
				if err == nil {
					t.Errorf("expected error got nil")
				}
			} else {
				if err == nil {
					t.Errorf("expected nil got error")
				}
			}
			assert.Equal(t, got, testCase.id)
		})
	}
}

func TestAuthPostgres_SignInAdmin(t *testing.T) {

	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Log(err)
	}
	defer db.Close()

	r := NewAuthRepo(db)

	type mockBehaviour func(user models.Admin)

	testTable := []struct {
		name          string
		username      string
		password      string
		id            int
		mockBehaviour mockBehaviour
		wantErr       bool
	}{
		{
			name:     "OK",
			username: "username",
			password: "password",
			id:       1,
			mockBehaviour: func(user models.Admin) {
				mock.ExpectQuery("SELECT id FROM").WithArgs(user.Username, user.Password).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
			},
			wantErr: false,
		},
		{
			name:     "Error",
			username: "username",
			password: "",
			id:       0,
			mockBehaviour: func(user models.Admin) {
				mock.ExpectQuery("SELECT id FROM").WithArgs(user.Username, user.Password).WillReturnError(err)
			},
			wantErr: true,
		},
	}

	for _, testCase := range testTable {

		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehaviour(models.Admin{Username: testCase.username, Password: testCase.password})
			got, err := r.SignInAdmin(models.Admin{Username: testCase.username, Password: testCase.password})
			if testCase.wantErr {
				if err == nil {
					t.Errorf("expected error got nil")
				}
			} else {
				if err != nil {
					t.Errorf("expected nil got error")
				}
			}
			assert.Equal(t, got, testCase.id)
		})
	}
}
