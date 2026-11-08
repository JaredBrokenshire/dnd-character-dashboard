package tests

import (
	"dnd-character-dashboard/db/factories"
	"dnd-character-dashboard/db/models"
	"dnd-character-dashboard/tests/helpers"
	"fmt"
	"net/http"
	"testing"
)

func TestClass_List(t *testing.T) {
	ts.ClearTable("classes")
	ts.ClearTable("files")
	ts.SetupDefaultUsers()

	logo := &models.File{}
	factories.NewFile(ts.S.DB, logo)

	cardBackground := &models.File{}
	factories.NewFile(ts.S.DB, cardBackground)

	class := &models.Class{
		Name:                 "AAA",
		LogoFileID:           logo.ID,
		CardBackgroundFileID: cardBackground.ID,
	}
	factories.NewClass(ts.S.DB, class)
	class2 := &models.Class{Name: "BBB"}
	factories.NewClass(ts.S.DB, class2)
	namedClass := &models.Class{Name: "Test Class"}
	factories.NewClass(ts.S.DB, namedClass)

	getRequest := func(query string) helpers.Request {
		return helpers.Request{
			Method: http.MethodGet,
			Url:    fmt.Sprintf("/classes%v", query),
		}
	}

	permissionRequest := getRequest("")
	RunNoAuthenticationTests(t, permissionRequest.Method, permissionRequest.Url)

	cases := []helpers.TestCase{
		{
			Name:    "Can get classes",
			Request: getRequest(""),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"name":"%v"`, class.Name),
					fmt.Sprintf(`"name":"%v"`, class2.Name),
					fmt.Sprintf(`"name":"%v"`, namedClass.Name),

					fmt.Sprintf(`"location":"%v"`, logo.Location),
					fmt.Sprintf(`"name":"%v"`, logo.Name),

					fmt.Sprintf(`"location":"%v"`, cardBackground.Location),
					fmt.Sprintf(`"name":"%v"`, cardBackground.Name),

					fmt.Sprintf(`"tagline":"%v"`, class.Tagline),
					fmt.Sprintf(`"card_description":"%v"`, class.CardDescription),
					fmt.Sprintf(`"primary_ability":"%v"`, class.PrimaryAbility),
					fmt.Sprintf(`"hit_point_die_value":%v`, class.HitPointDieValue),
					fmt.Sprintf(`"save":"%v"`, class.Save),
					fmt.Sprintf(`"secondary_save":"%v"`, class.SecondarySave),

					`"total_count":3`,
				},
			},
		},
		{
			Name:    "Can get page 0 of classes",
			Request: getRequest("?page=0&page_size=1"),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"name":"%v"`, class.Name),

					fmt.Sprintf(`"location":"%v"`, logo.Location),
					fmt.Sprintf(`"name":"%v"`, logo.Name),

					fmt.Sprintf(`"location":"%v"`, cardBackground.Location),
					fmt.Sprintf(`"name":"%v"`, cardBackground.Name),

					`"total_count":3`,
				},
				BodyPartsMissing: []string{
					fmt.Sprintf(`"name":"%v"`, class2.Name),
					fmt.Sprintf(`"name":"%v"`, namedClass.Name),
				},
			},
		},
		{
			Name:    "Can get page 1 of classes",
			Request: getRequest("?page=1&page_size=1"),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"name":"%v"`, class2.Name),

					`"total_count":3`,
				},
				BodyPartsMissing: []string{
					fmt.Sprintf(`"name":"%v"`, class.Name),
					fmt.Sprintf(`"name":"%v"`, namedClass.Name),

					fmt.Sprintf(`"location":"%v"`, logo.Location),
					fmt.Sprintf(`"name":"%v"`, logo.Name),

					fmt.Sprintf(`"location":"%v"`, cardBackground.Location),
					fmt.Sprintf(`"name":"%v"`, cardBackground.Name),
				},
			},
		},
		{
			Name:    "Can search classes by name",
			Request: getRequest("?search=Test"),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"name":"%v"`, namedClass.Name),

					`"total_count":1`,
				},
				BodyPartsMissing: []string{
					fmt.Sprintf(`"name":"%v"`, class.Name),
					fmt.Sprintf(`"name":"%v"`, class2.Name),

					fmt.Sprintf(`"location":"%v"`, logo.Location),
					fmt.Sprintf(`"name":"%v"`, logo.Name),

					fmt.Sprintf(`"location":"%v"`, cardBackground.Location),
					fmt.Sprintf(`"name":"%v"`, cardBackground.Name),
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
