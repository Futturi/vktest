package repository

import (
	"testing"

	"github.com/Futturi/vktest/internal/models"
	"github.com/magiconair/properties/assert"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
)

func TestActorRepository_GetActors(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Log(err)
	}
	defer db.Close()

	r := NewActorRepo(db)

	type mockBehavior func([]models.ActorSelect)

	testTable := []struct {
		name         string
		result       []models.ActorSelect
		wantErr      bool
		mockBehavior mockBehavior
	}{
		{
			name: "OK",
			result: []models.ActorSelect{
				{},
			},
			wantErr: false,
			mockBehavior: func(result []models.ActorSelect) {
				mock.ExpectQuery("SELECT name, genre, data FROM").WillReturnRows(sqlmock.NewRows([]string{"name", "genre", "data"}).AddRow("name", "genre", "data"))
			},
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior(testCase.result)
			_, err := r.GetActors()
			if testCase.wantErr {
				if err == nil {
					t.Errorf("unexpected err: %v", err)
				}
				return
			}
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestActorRepository_InsertActor(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Log(err)
	}
	defer db.Close()
	r := NewActorRepo(db)
	type mockBehaviour func(args models.Actor, id int)

	testTable := []struct {
		name          string
		args          models.Actor
		id            int
		mockBehaviour mockBehaviour
		wantErr       bool
	}{
		{
			name: "OK",
			args: models.Actor{
				Name: "name",
			},
			id: 1,
			mockBehaviour: func(args models.Actor, id int) {
				mock.ExpectQuery("INSERT INTO").WithArgs(args.Name, args.Genre, args.Data).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(id))
			},
			wantErr: false,
		},
		{
			name: "Error",
			args: models.Actor{
				Name: "",
			},
			id: 1,
			mockBehaviour: func(args models.Actor, id int) {
				mock.ExpectQuery("INSERT INTO").WithArgs(args.Name, args.Genre, args.Data).WillReturnError(err)
			},
			wantErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehaviour(testCase.args, testCase.id)
			got, err := r.InsertActor(testCase.args)
			if testCase.wantErr {
				if err == nil {
					t.Errorf("unexpected err: %v", err)
				}
				return
			}
			assert.Equal(t, got, testCase.id)
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestActorRepository_NewActorRepo(t *testing.T) {
	db, _, err := sqlmock.Newx()
	if err != nil {
		t.Log(err)
	}
	defer db.Close()
	r := NewActorRepo(db)
	assert.Equal(t, db, r.db)
}

func TestActorRepository_FindIdCinemaByName(t *testing.T) {

	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Log(err)
	}
	defer db.Close()
	r := NewActorRepo(db)

	type mockBehaviour func(name string, id int)

	testTable := []struct {
		name          string
		nameCinema    string
		id            int
		mockBehaviour mockBehaviour
		wantErr       bool
	}{
		{
			name:       "OK",
			nameCinema: "cinema",
			id:         1,
			mockBehaviour: func(name string, id int) {
				mock.ExpectQuery("SELECT id FROm").WithArgs(name).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(id))
			},
			wantErr: false,
		},
		{
			name:       "Error",
			nameCinema: "",
			id:         1,
			mockBehaviour: func(name string, id int) {
				mock.ExpectQuery("SELECT id FROm").WithArgs(name).WillReturnError(err)
			},
			wantErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehaviour(testCase.nameCinema, testCase.id)
			got, err := r.FindIdCinemaByName(testCase.nameCinema)
			if testCase.wantErr {
				if err == nil {
					t.Errorf("unexpected err: %v", err)
				}
				return
			}
			assert.Equal(t, got, testCase.id)
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestActorRepository_DeleteActor(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Log(err)
	}
	defer db.Close()
	r := NewActorRepo(db)

	type mockBehaviour func(id string)

	testTable := []struct {
		name          string
		id            string
		mockBehaviour mockBehaviour
		wantErr       bool
	}{
		{
			name: "OK",
			id:   "1",
			mockBehaviour: func(id string) {
				mock.ExpectExec("DELETE FROM").WithArgs(id).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			wantErr: false,
		},
		{
			name: "Error",
			id:   "1",
			mockBehaviour: func(id string) {
				mock.ExpectExec("DELETE FROM").WithArgs(id).WillReturnError(err)
			},
			wantErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehaviour(testCase.id)
			err := r.DeleteActor(testCase.id)
			if testCase.wantErr {
				if err == nil {
					t.Errorf("unexpected err: %v", err)
				}
				return
			}
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}
