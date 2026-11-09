package tests

import (
	"dnd-character-dashboard/db/factories"
	"dnd-character-dashboard/db/models"
	"dnd-character-dashboard/services"
	"dnd-character-dashboard/tests/helpers"
	"dnd-character-dashboard/tests/mocks"
	"errors"
	"fmt"
	"net/http"
	"testing"
)

func TestRace_List(t *testing.T) {
	ts.ClearTable("race_traits")
	ts.ClearTable("races")
	ts.ClearTable("traits")
	ts.ClearTable("files")
	ts.SetupDefaultUsers()

	cardBackground := &models.File{}
	factories.NewFile(ts.S.DB, cardBackground)

	trait := &models.Trait{}
	factories.NewTrait(ts.S.DB, trait)
	trait2 := &models.Trait{}
	factories.NewTrait(ts.S.DB, trait2)
	trait3 := &models.Trait{}
	factories.NewTrait(ts.S.DB, trait3)

	race := &models.Race{
		Name:                 "AAA",
		CardBackgroundFileID: cardBackground.ID,
		Traits:               []*models.Trait{trait, trait2},
	}
	factories.NewRace(ts.S.DB, race)
	race2 := &models.Race{
		Name:   "BBB",
		Traits: []*models.Trait{trait3},
	}
	factories.NewRace(ts.S.DB, race2)
	namedRace := &models.Race{Name: "Test Race"}
	factories.NewRace(ts.S.DB, namedRace)

	getRequest := func(query string) helpers.Request {
		return helpers.Request{
			Method: http.MethodGet,
			Url:    fmt.Sprintf("/races%v", query),
		}
	}

	permissionRequest := getRequest("")
	RunNoAuthenticationTests(t, permissionRequest.Method, permissionRequest.Url)

	cases := []helpers.TestCase{
		{
			Name:    "Can get races",
			Request: getRequest(""),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"name":"%v"`, race.Name),
					fmt.Sprintf(`"name":"%v"`, race2.Name),
					fmt.Sprintf(`"name":"%v"`, namedRace.Name),

					fmt.Sprintf(`"location":"%v"`, cardBackground.Location),
					fmt.Sprintf(`"name":"%v"`, cardBackground.Name),

					fmt.Sprintf(`"name":"%v"`, trait.Name),
					fmt.Sprintf(`"name":"%v"`, trait2.Name),
					fmt.Sprintf(`"name":"%v"`, trait3.Name),

					`"total_count":3`,
				},
			},
		},
		{
			Name:    "Can get page 0 of races",
			Request: getRequest("?page=0&page_size=1"),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"name":"%v"`, race.Name),

					fmt.Sprintf(`"location":"%v"`, cardBackground.Location),
					fmt.Sprintf(`"name":"%v"`, cardBackground.Name),

					fmt.Sprintf(`"name":"%v"`, trait.Name),
					fmt.Sprintf(`"name":"%v"`, trait2.Name),

					`"total_count":3`,
				},
				BodyPartsMissing: []string{
					fmt.Sprintf(`"name":"%v"`, race2.Name),
					fmt.Sprintf(`"name":"%v"`, namedRace.Name),

					fmt.Sprintf(`"name":"%v"`, trait3.Name),
				},
			},
		},
		{
			Name:    "Can get page 1 of races",
			Request: getRequest("?page=1&page_size=1"),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"name":"%v"`, race2.Name),

					fmt.Sprintf(`"name":"%v"`, trait3.Name),

					`"total_count":3`,
				},
				BodyPartsMissing: []string{
					fmt.Sprintf(`"name":"%v"`, race.Name),
					fmt.Sprintf(`"name":"%v"`, namedRace.Name),

					fmt.Sprintf(`"location":"%v"`, cardBackground.Location),
					fmt.Sprintf(`"name":"%v"`, cardBackground.Name),

					fmt.Sprintf(`"name":"%v"`, trait.Name),
					fmt.Sprintf(`"name":"%v"`, trait2.Name),
				},
			},
		},
		{
			Name:    "Can search races by name",
			Request: getRequest("?search=Test"),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"name":"%v"`, namedRace.Name),

					`"total_count":1`,
				},
				BodyPartsMissing: []string{
					fmt.Sprintf(`"name":"%v"`, race.Name),
					fmt.Sprintf(`"name":"%v"`, race2.Name),

					fmt.Sprintf(`"location":"%v"`, cardBackground.Location),
					fmt.Sprintf(`"name":"%v"`, cardBackground.Name),

					fmt.Sprintf(`"name":"%v"`, trait.Name),
					fmt.Sprintf(`"name":"%v"`, trait2.Name),
					fmt.Sprintf(`"name":"%v"`, trait3.Name),
				},
			},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			RunAuthorisedTestCase(t, test)
		})
	}
}

func TestRace_Get(t *testing.T) {
	ts.ClearTable("race_traits")
	ts.ClearTable("race_variants")
	ts.ClearTable("races")
	ts.ClearTable("traits")
	ts.ClearTable("files")
	ts.SetupDefaultUsers()

	detailsImage := &models.File{}
	factories.NewFile(ts.S.DB, detailsImage)
	detailsImage2 := &models.File{}
	factories.NewFile(ts.S.DB, detailsImage2)

	trait := &models.Trait{}
	factories.NewTrait(ts.S.DB, trait)
	trait2 := &models.Trait{}
	factories.NewTrait(ts.S.DB, trait2)
	trait3 := &models.Trait{}
	factories.NewTrait(ts.S.DB, trait3)

	race := &models.Race{
		DetailsImageFileID: detailsImage.ID,
		Traits:             []*models.Trait{trait, trait2},
	}
	factories.NewRace(ts.S.DB, race)
	race2 := &models.Race{
		DetailsImageFileID: detailsImage2.ID,
		Traits:             []*models.Trait{trait3},
	}
	factories.NewRace(ts.S.DB, race2)

	variant := &models.RaceVariant{
		RaceID: race.ID,
	}
	factories.NewRaceVariant(ts.S.DB, variant)
	variant2 := &models.RaceVariant{
		RaceID: race.ID,
	}
	factories.NewRaceVariant(ts.S.DB, variant2)
	variant3 := &models.RaceVariant{
		RaceID: race2.ID,
	}
	factories.NewRaceVariant(ts.S.DB, variant3)

	getRequest := func(id interface{}) helpers.Request {
		return helpers.Request{
			Method: http.MethodGet,
			Url:    fmt.Sprintf("/races/%v", id),
		}
	}

	permissionRequest := getRequest(race.ID)
	RunNoAuthenticationTests(t, permissionRequest.Method, permissionRequest.Url)

	cases := []helpers.TestCase{
		{
			Name:    "Can't get race with invalid id",
			Request: getRequest("invalid-id"),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Race not found",
			},
		},
		{
			Name:    "Can't get race that doesn't exist",
			Request: getRequest(999),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Race not found",
			},
		},
		{
			Name: "Can't get race if there is an error from the race service",
			Setup: func(test *helpers.TestCase) {
				mockRaceService := mocks.NewMockRaceService()
				mockRaceService.GetError = errors.New("test error")

				ts.S.Dependencies.SetRaceService(mockRaceService)
			},
			Request: getRequest(race.ID),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusInternalServerError,
				BodyPart:   "There was an error getting the race",
			},
			Teardown: func() {
				ts.S.Dependencies.SetRaceService(services.NewRaceService(ts.S.Repos.Race))
			},
		},
		{
			Name:    "Can get race",
			Request: getRequest(race.ID),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"id":%v`, race.ID),
					fmt.Sprintf(`"name":"%v"`, race.Name),
					fmt.Sprintf(`"description":"%v"`, race.Description),
					fmt.Sprintf(`"size":"%v"`, race.Size),
					fmt.Sprintf(`"base_speed":%v`, race.BaseSpeed),
					fmt.Sprintf(`"creature_type":"%v"`, race.CreatureType),

					fmt.Sprintf(`"name":"%v"`, trait.Name),
					fmt.Sprintf(`"description":"%v"`, trait.Description),
					fmt.Sprintf(`"name":"%v"`, trait2.Name),
					fmt.Sprintf(`"description":"%v"`, trait2.Description),

					fmt.Sprintf(`"name":"%v"`, variant.Name),
					fmt.Sprintf(`"name":"%v"`, variant2.Name),

					fmt.Sprintf(`"name":"%v"`, detailsImage.Name),
					fmt.Sprintf(`"location":"%v"`, detailsImage.Location),
				},
				BodyPartsMissing: []string{
					fmt.Sprintf(`"name":"%v"`, trait3.Name),
					fmt.Sprintf(`"name":"%v"`, variant3.Name),
					fmt.Sprintf(`"name":"%v"`, detailsImage2.Name),
					fmt.Sprintf(`"location":"%v"`, detailsImage2.Location),
				},
			},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			RunAuthorisedTestCase(t, test)
		})
	}
}
