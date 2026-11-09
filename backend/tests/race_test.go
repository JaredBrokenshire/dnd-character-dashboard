package tests

import (
	"dnd-character-dashboard/db/factories"
	"dnd-character-dashboard/db/models"
	"dnd-character-dashboard/tests/helpers"
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
