package tests

import (
	"dnd/db/factories"
	"dnd/db/models"
	"dnd/tests/helpers"
	"dnd/tests/mocks"
	"errors"
	"fmt"
	"net/http"
	"testing"
)

func TestClass_List(t *testing.T) {
	ts.ClearTable("classes")
	ts.ClearTable("class_save_proficiencies")
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

	classSaveProficiency := &models.ClassSaveProficiency{ClassID: class.ID}
	factories.NewClassSaveProficiency(ts.S.DB, classSaveProficiency)

	classSaveProficiency2 := &models.ClassSaveProficiency{ClassID: class.ID}
	factories.NewClassSaveProficiency(ts.S.DB, classSaveProficiency2)

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

					fmt.Sprintf(`"save":"%v"`, classSaveProficiency.Ability),
					fmt.Sprintf(`"save":"%v"`, classSaveProficiency2.Ability),

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

					fmt.Sprintf(`"tagline":"%v"`, class.Tagline),
					fmt.Sprintf(`"card_description":"%v"`, class.CardDescription),
					fmt.Sprintf(`"primary_ability":"%v"`, class.PrimaryAbility),
					fmt.Sprintf(`"hit_point_die_value":%v`, class.HitPointDieValue),

					fmt.Sprintf(`"save":"%v"`, classSaveProficiency.Ability),
					fmt.Sprintf(`"save":"%v"`, classSaveProficiency2.Ability),

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

					fmt.Sprintf(`"tagline":"%v"`, class2.Tagline),
					fmt.Sprintf(`"card_description":"%v"`, class2.CardDescription),
					fmt.Sprintf(`"primary_ability":"%v"`, class2.PrimaryAbility),
					fmt.Sprintf(`"hit_point_die_value":%v`, class2.HitPointDieValue),

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

					fmt.Sprintf(`"tagline":"%v"`, namedClass.Tagline),
					fmt.Sprintf(`"card_description":"%v"`, namedClass.CardDescription),
					fmt.Sprintf(`"primary_ability":"%v"`, namedClass.PrimaryAbility),
					fmt.Sprintf(`"hit_point_die_value":%v`, namedClass.HitPointDieValue),

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

func TestClass_Get(t *testing.T) {
	ts.ClearTable("classes")
	ts.ClearTable("class_save_proficiencies")
	ts.ClearTable("class_weapon_proficiencies")
	ts.ClearTable("class_skill_proficiencies")
	ts.ClearTable("class_armour_proficiencies")
	ts.ClearTable("subclasses")
	ts.ClearTable("files")
	ts.SetupDefaultUsers()

	logo := &models.File{}
	factories.NewFile(ts.S.DB, logo)

	cardBackground := &models.File{}
	factories.NewFile(ts.S.DB, cardBackground)

	class := &models.Class{
		LogoFileID:           logo.ID,
		CardBackgroundFileID: cardBackground.ID,
	}
	factories.NewClass(ts.S.DB, class)
	class2 := &models.Class{}
	factories.NewClass(ts.S.DB, class2)

	classSaveProficiency := &models.ClassSaveProficiency{ClassID: class.ID}
	factories.NewClassSaveProficiency(ts.S.DB, classSaveProficiency)
	classSaveProficiency2 := &models.ClassSaveProficiency{ClassID: class.ID}
	factories.NewClassSaveProficiency(ts.S.DB, classSaveProficiency2)

	classSkillProficiency := &models.ClassSkillProficiency{ClassID: class.ID}
	factories.NewClassSkillProficiency(ts.S.DB, classSkillProficiency)
	classSkillProficiency2 := &models.ClassSkillProficiency{ClassID: class.ID}
	factories.NewClassSkillProficiency(ts.S.DB, classSkillProficiency2)

	classWeaponProficiency := &models.ClassWeaponProficiency{ClassID: class.ID}
	factories.NewClassWeaponProficiency(ts.S.DB, classWeaponProficiency)
	classWeaponProficiency2 := &models.ClassWeaponProficiency{ClassID: class.ID}
	factories.NewClassWeaponProficiency(ts.S.DB, classWeaponProficiency2)

	classArmourProficiency := &models.ClassArmourProficiency{ClassID: class.ID}
	factories.NewClassArmourProficiency(ts.S.DB, classArmourProficiency)
	classArmourProficiency2 := &models.ClassArmourProficiency{ClassID: class.ID}
	factories.NewClassArmourProficiency(ts.S.DB, classArmourProficiency2)

	subclass := &models.Subclass{ClassID: class.ID}
	factories.NewSubclass(ts.S.DB, subclass)
	subclass2 := &models.Subclass{ClassID: class.ID}
	factories.NewSubclass(ts.S.DB, subclass2)

	getRequest := func(id interface{}) helpers.Request {
		return helpers.Request{
			Method: http.MethodGet,
			Url:    fmt.Sprintf("/classes/%v", id),
		}
	}

	permissionRequest := getRequest(class.ID)
	RunNoAuthenticationTests(t, permissionRequest.Method, permissionRequest.Url)

	cases := []helpers.TestCase{
		{
			Name:    "Can't get class that doesn't exist",
			Request: getRequest(999),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Class not found",
			},
		},
		{
			Name:    "Can't get class with invalid id",
			Request: getRequest("invalid-id"),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Class not found",
			},
		},
		{
			Name: "Can't get class if there is an error from the class service",
			Setup: func(test *helpers.TestCase) {
				mockClassService := mocks.NewMockClassService()
				mockClassService.GetError = errors.New("test error")

				ts.S.Dependencies.SetClassService(mockClassService)
			},
			Request: getRequest(class.ID),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusInternalServerError,
				BodyPart:   "There was an error getting the class",
			},
			Teardown: func() {
				ts.S.Dependencies.SetClassService(nil)
			},
		},
		{
			Name:    "Can get class",
			Request: getRequest(class.ID),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"name":"%v"`, class.Name),

					fmt.Sprintf(`"location":"%v"`, logo.Location),
					fmt.Sprintf(`"name":"%v"`, logo.Name),

					fmt.Sprintf(`"location":"%v"`, cardBackground.Location),
					fmt.Sprintf(`"name":"%v"`, cardBackground.Name),

					fmt.Sprintf(`"tagline":"%v"`, class.Tagline),
					fmt.Sprintf(`"card_description":"%v"`, class.CardDescription),
					fmt.Sprintf(`"primary_ability":"%v"`, class.PrimaryAbility),
					fmt.Sprintf(`"hit_point_die_value":%v`, class.HitPointDieValue),

					fmt.Sprintf(`"save":"%v"`, classSaveProficiency.Ability),
					fmt.Sprintf(`"save":"%v"`, classSaveProficiency2.Ability),

					fmt.Sprintf(`"skill":"%v"`, classSkillProficiency.Skill),
					fmt.Sprintf(`"skill":"%v"`, classSkillProficiency2.Skill),
					fmt.Sprintf(`"skill_proficiency_max_options":%v`, class.SkillProficiencyMaxOptions),

					fmt.Sprintf(`"weapon_type":"%v"`, classWeaponProficiency.WeaponType),
					fmt.Sprintf(`"weapon_type":"%v"`, classWeaponProficiency2.WeaponType),
					fmt.Sprintf(`"condition":"%v"`, classWeaponProficiency.Condition),
					fmt.Sprintf(`"condition":"%v"`, classWeaponProficiency2.Condition),

					fmt.Sprintf(`"armour_type":"%v"`, classArmourProficiency.Armour),
					fmt.Sprintf(`"armour_type":"%v"`, classArmourProficiency2.Armour),

					fmt.Sprintf(`"name":"%v"`, subclass.Name),
					fmt.Sprintf(`"name":"%v"`, subclass2.Name),
					fmt.Sprintf(`"description":"%v"`, subclass.Description),
					fmt.Sprintf(`"description":"%v"`, subclass2.Description),
				},
				BodyPartsMissing: []string{
					fmt.Sprintf(`"name":"%v"`, class2.Name),
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
