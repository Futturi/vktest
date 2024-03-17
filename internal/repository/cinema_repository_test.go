package repository

import (
	"testing"

	"github.com/Futturi/vktest/internal/models"
	"github.com/magiconair/properties/assert"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
)

func TestCinemaRepository_InsertCinema(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Log(err)
	}

	defer db.Close()
	r := NewCinemaRepo(db)
	type mockBehaviour func(args models.Cinema, id int)

	testTable := []struct {
		name          string
		args          models.Cinema
		id            int
		mockBehaviour mockBehaviour
		wantErr       bool
	}{
		{
			name: "OK",
			args: models.Cinema{
				Name:        "name",
				Description: "desc",
				Data:        "data",
				Rating:      1,
			},
			id: 1,
			mockBehaviour: func(args models.Cinema, id int) {
				mock.ExpectQuery("INSERT INTO").WithArgs(args.Name, args.Description, args.Data, args.Rating).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(id))
			},
			wantErr: false,
		},
		{
			name: "Error",
			args: models.Cinema{
				Name:        "name",
				Description: "desc",
				Data:        "data",
				Rating:      1,
			},
			id: 1,
			mockBehaviour: func(args models.Cinema, id int) {
				mock.ExpectQuery("INSERT INTO").WithArgs(args.Name, args.Description, args.Data, args.Rating).WillReturnError(err)
			},
			wantErr: true,
		},
	}

	for _, testCase := range testTable {

		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehaviour(testCase.args, testCase.id)
			got, err := r.InsertCinema(testCase.args)
			if testCase.wantErr {
				if err == nil {
					t.Errorf("unexpected err: %v", err)
				}
				return
			}
			if err != nil {
				t.Errorf("error %v", err)
			}

			assert.Equal(t, got, testCase.id)
		})
	}

}

func TestCinemaRepository_GetIdActorByName(t *testing.T) {

	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Log(err)
	}

	defer db.Close()

	r := NewCinemaRepo(db)

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
				mock.ExpectQuery("SELECT id FROM").WithArgs(name).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(id))
			},
			wantErr: false,
		},
		{
			name:       "Error",
			nameCinema: "",
			id:         0,
			mockBehaviour: func(name string, id int) {
				mock.ExpectQuery("SELECT id FROM").WithArgs(name).WillReturnError(err)
			},
			wantErr: true,
		},
	}

	for _, testCase := range testTable {

		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehaviour(testCase.nameCinema, testCase.id)
			got, err := r.GetIdActorByName(testCase.nameCinema)
			if testCase.wantErr {
				if err == nil {
					t.Errorf("unexpected err: %v", err)
				}
				return
			}
			if err != nil {
				t.Errorf("error %v", err)
			}
			assert.Equal(t, got, testCase.id)
		})
	}
}

func TestActorRepository_NewCinemaRepo(t *testing.T) {
	db, _, err := sqlmock.Newx()
	if err != nil {
		t.Log(err)
	}
	defer db.Close()
	r := NewCinemaRepo(db)
	assert.Equal(t, db, r.db)
}

func TestCinemaRepository_DeleteFilm(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Log(err)
	}

	defer db.Close()

	r := NewCinemaRepo(db)

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
			err := r.DeleteFilm(testCase.id)
			if testCase.wantErr {
				if err == nil {
					t.Errorf("unexpected err: %v", err)
				}
				return
			}
			if err != nil {
				t.Errorf("error %v", err)
			}
		})
	}
}

func TestCinemaRepository_GetCinemas(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil {
		t.Log(err)
	}
	defer db.Close()

	r := NewCinemaRepo(db)

	type mockBehaviour func([]models.Cinema)

	testTable := []struct {
		name          string
		result        []models.Cinema
		mockBehaviour mockBehaviour
		sort          string
	}{
		{
			name: "OK",
			mockBehaviour: func([]models.Cinema) {
				mock.ExpectQuery("SELECT id, name, description, data, rating").WillReturnRows(sqlmock.NewRows([]string{"name", "description", "data", "rating"}).AddRow("name", "description", "data", 1))
			},
			sort:   "",
			result: []models.Cinema{},
		},
	}

	for _, testCase := range testTable {

		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehaviour(testCase.result)
			got, err := r.GetCinemas(testCase.sort)
			if err != nil {
				t.Errorf("error %v", err)
			}
			assert.Equal(t, got, testCase.result)
		})
	}

}
