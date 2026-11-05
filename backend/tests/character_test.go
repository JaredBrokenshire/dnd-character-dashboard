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

func TestCharacter_List(t *testing.T) {
	ts.ClearTable("characters")
	ts.ClearTable("character_classes")
	ts.ClearTable("character_subclasses")
	ts.ClearTable("classes")
	ts.ClearTable("subclasses")
	ts.ClearTable("races")
	ts.ClearTable("campaigns")
	ts.ClearTable("files")
	ts.SetupDefaultUsers()

	classLogo := &models.File{}
	factories.NewFile(ts.S.Db, classLogo)
	classLogo2 := &models.File{}
	factories.NewFile(ts.S.Db, classLogo2)

	class := &models.Class{LogoFileID: classLogo.ID}
	factories.NewClass(ts.S.Db, class)
	class2 := &models.Class{LogoFileID: classLogo2.ID}
	factories.NewClass(ts.S.Db, class2)

	subclass := &models.Subclass{ClassID: class.ID}
	factories.NewSubclass(ts.S.Db, subclass)
	subclass2 := &models.Subclass{ClassID: class2.ID}
	factories.NewSubclass(ts.S.Db, subclass2)

	race := &models.Race{}
	factories.NewRace(ts.S.Db, race)
	race2 := &models.Race{}
	factories.NewRace(ts.S.Db, race2)

	campaign := &models.Campaign{}
	factories.NewCampaign(ts.S.Db, campaign)
	campaign2 := &models.Campaign{}
	factories.NewCampaign(ts.S.Db, campaign2)

	profilePicture := &models.File{}
	factories.NewFile(ts.S.Db, profilePicture)
	profilePicture2 := &models.File{}
	factories.NewFile(ts.S.Db, profilePicture2)
	otherUserProfilePicture := &models.File{}
	factories.NewFile(ts.S.Db, otherUserProfilePicture)

	character := &models.Character{
		Name:                 "AAA",
		CampaignID:           campaign.ID,
		RaceID:               race.ID,
		Classes:              []*models.Class{class},
		Subclasses:           []*models.Subclass{subclass},
		ProfilePictureFileID: profilePicture.ID,
	}
	factories.NewCharacter(ts.S.Db, character)
	character2 := &models.Character{
		Name:                 "BBB",
		CampaignID:           campaign2.ID,
		RaceID:               race2.ID,
		Classes:              []*models.Class{class2},
		Subclasses:           []*models.Subclass{subclass2},
		ProfilePictureFileID: profilePicture2.ID,
	}
	factories.NewCharacter(ts.S.Db, character2)
	namedCharacter := &models.Character{
		Name: "Test Character",
	}
	factories.NewCharacter(ts.S.Db, namedCharacter)
	otherUserCharacter := &models.Character{
		UserID:               999,
		ProfilePictureFileID: otherUserProfilePicture.ID,
	}
	factories.NewCharacter(ts.S.Db, otherUserCharacter)

	getRequest := func(query string) helpers.Request {
		return helpers.Request{
			Method: http.MethodGet,
			Url:    fmt.Sprintf("/characters%v", query),
		}
	}

	permissionRequest := getRequest("")
	RunNoAuthenticationTests(t, permissionRequest.Method, permissionRequest.Url)

	cases := []helpers.TestCase{
		{
			Name:    "Can get characters",
			Request: getRequest(""),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"name":"%v"`, character.Name),
					fmt.Sprintf(`"name":"%v"`, character2.Name),
					fmt.Sprintf(`"name":"%v"`, namedCharacter.Name),
					fmt.Sprintf(`"name":"%v"`, class.Name),
					fmt.Sprintf(`"name":"%v"`, class2.Name),
					fmt.Sprintf(`"name":"%v"`, subclass.Name),
					fmt.Sprintf(`"name":"%v"`, subclass2.Name),
					fmt.Sprintf(`"name":"%v"`, race.Name),
					fmt.Sprintf(`"name":"%v"`, race2.Name),
					fmt.Sprintf(`"name":"%v"`, campaign.Name),
					fmt.Sprintf(`"name":"%v"`, campaign2.Name),
					fmt.Sprintf(`"name":"%v"`, profilePicture.Name),
					fmt.Sprintf(`"location":"%v"`, profilePicture.Location),
					fmt.Sprintf(`"name":"%v"`, profilePicture2.Name),
					fmt.Sprintf(`"location":"%v"`, profilePicture2.Location),
					fmt.Sprintf(`"name":"%v"`, classLogo.Name),
					fmt.Sprintf(`"location":"%v"`, classLogo.Location),
					fmt.Sprintf(`"name":"%v"`, classLogo2.Name),
					fmt.Sprintf(`"location":"%v"`, classLogo2.Location),
					`"total_count":3`,
				},
				BodyPartsMissing: []string{
					fmt.Sprintf(`"name":"%v"`, otherUserCharacter.Name),
					fmt.Sprintf(`"name":"%v"`, otherUserProfilePicture.Name),
					fmt.Sprintf(`"location":"%v"`, otherUserProfilePicture.Location),
				},
			},
		},
		{
			Name:    "Can get page 0 of characters",
			Request: getRequest("?page=0&page_size=1"),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"name":"%v"`, character.Name),
					fmt.Sprintf(`"name":"%v"`, class.Name),
					fmt.Sprintf(`"name":"%v"`, subclass.Name),
					fmt.Sprintf(`"name":"%v"`, race.Name),
					fmt.Sprintf(`"name":"%v"`, campaign.Name),
					fmt.Sprintf(`"name":"%v"`, profilePicture.Name),
					fmt.Sprintf(`"location":"%v"`, profilePicture.Location),
					fmt.Sprintf(`"name":"%v"`, classLogo.Name),
					fmt.Sprintf(`"location":"%v"`, classLogo.Location),
					`"total_count":3`,
				},
				BodyPartsMissing: []string{
					fmt.Sprintf(`"name":"%v"`, character2.Name),
					fmt.Sprintf(`"name":"%v"`, namedCharacter.Name),
					fmt.Sprintf(`"name":"%v"`, otherUserCharacter.Name),
					fmt.Sprintf(`"name":"%v"`, class2.Name),
					fmt.Sprintf(`"name":"%v"`, subclass2.Name),
					fmt.Sprintf(`"name":"%v"`, race2.Name),
					fmt.Sprintf(`"name":"%v"`, campaign2.Name),
					fmt.Sprintf(`"name":"%v"`, profilePicture2.Name),
					fmt.Sprintf(`"location":"%v"`, profilePicture2.Location),
					fmt.Sprintf(`"name":"%v"`, otherUserProfilePicture.Name),
					fmt.Sprintf(`"location":"%v"`, otherUserProfilePicture.Location),
					fmt.Sprintf(`"name":"%v"`, classLogo2.Name),
					fmt.Sprintf(`"location":"%v"`, classLogo2.Location),
				},
			},
		},
		{
			Name:    "Can get page 1 of characters",
			Request: getRequest("?page=1&page_size=1"),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"name":"%v"`, character2.Name),
					fmt.Sprintf(`"name":"%v"`, class2.Name),
					fmt.Sprintf(`"name":"%v"`, subclass2.Name),
					fmt.Sprintf(`"name":"%v"`, race2.Name),
					fmt.Sprintf(`"name":"%v"`, campaign2.Name),
					fmt.Sprintf(`"name":"%v"`, profilePicture2.Name),
					fmt.Sprintf(`"location":"%v"`, profilePicture2.Location),
					fmt.Sprintf(`"name":"%v"`, classLogo2.Name),
					fmt.Sprintf(`"location":"%v"`, classLogo2.Location),
					`"total_count":3`,
				},
				BodyPartsMissing: []string{
					fmt.Sprintf(`"name":"%v"`, character.Name),
					fmt.Sprintf(`"name":"%v"`, namedCharacter.Name),
					fmt.Sprintf(`"name":"%v"`, otherUserCharacter.Name),
					fmt.Sprintf(`"name":"%v"`, class.Name),
					fmt.Sprintf(`"name":"%v"`, subclass.Name),
					fmt.Sprintf(`"name":"%v"`, race.Name),
					fmt.Sprintf(`"name":"%v"`, campaign.Name),
					fmt.Sprintf(`"name":"%v"`, profilePicture.Name),
					fmt.Sprintf(`"location":"%v"`, profilePicture.Location),
					fmt.Sprintf(`"name":"%v"`, otherUserProfilePicture.Name),
					fmt.Sprintf(`"location":"%v"`, otherUserProfilePicture.Location),
					fmt.Sprintf(`"name":"%v"`, classLogo.Name),
					fmt.Sprintf(`"location":"%v"`, classLogo.Location),
				},
			},
		},
		{
			Name:    "Can search characters by name",
			Request: getRequest("?search=Test"),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"name":"%v"`, namedCharacter.Name),
					`"total_count":1`,
				},
				BodyPartsMissing: []string{
					fmt.Sprintf(`"name":"%v"`, character.Name),
					fmt.Sprintf(`"name":"%v"`, character2.Name),
					fmt.Sprintf(`"name":"%v"`, otherUserCharacter.Name),
					fmt.Sprintf(`"name":"%v"`, class.Name),
					fmt.Sprintf(`"name":"%v"`, subclass.Name),
					fmt.Sprintf(`"name":"%v"`, race.Name),
					fmt.Sprintf(`"name":"%v"`, campaign.Name),
					fmt.Sprintf(`"name":"%v"`, class2.Name),
					fmt.Sprintf(`"name":"%v"`, subclass2.Name),
					fmt.Sprintf(`"name":"%v"`, race2.Name),
					fmt.Sprintf(`"name":"%v"`, campaign2.Name),
					fmt.Sprintf(`"name":"%v"`, profilePicture.Name),
					fmt.Sprintf(`"location":"%v"`, profilePicture.Location),
					fmt.Sprintf(`"name":"%v"`, profilePicture2.Name),
					fmt.Sprintf(`"location":"%v"`, profilePicture2.Location),
					fmt.Sprintf(`"name":"%v"`, otherUserProfilePicture.Name),
					fmt.Sprintf(`"location":"%v"`, otherUserProfilePicture.Location),
					fmt.Sprintf(`"name":"%v"`, classLogo.Name),
					fmt.Sprintf(`"location":"%v"`, classLogo.Location),
					fmt.Sprintf(`"name":"%v"`, classLogo2.Name),
					fmt.Sprintf(`"location":"%v"`, classLogo2.Location),
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

func TestCharacter_Get(t *testing.T) {
	ts.ClearTable("characters")
	ts.ClearTable("character_classes")
	ts.ClearTable("character_subclasses")
	ts.ClearTable("classes")
	ts.ClearTable("subclasses")
	ts.ClearTable("races")
	ts.ClearTable("campaigns")
	ts.ClearTable("files")
	ts.SetupDefaultUsers()

	class := &models.Class{}
	factories.NewClass(ts.S.Db, class)
	subclass := &models.Subclass{ClassID: class.ID}
	factories.NewSubclass(ts.S.Db, subclass)
	race := &models.Race{}
	factories.NewRace(ts.S.Db, race)
	campaign := &models.Campaign{}
	factories.NewCampaign(ts.S.Db, campaign)
	profilePicture := &models.File{}
	factories.NewFile(ts.S.Db, profilePicture)
	otherUserProfilePicture := &models.File{}
	factories.NewFile(ts.S.Db, otherUserProfilePicture)

	character := &models.Character{
		CampaignID:           campaign.ID,
		RaceID:               race.ID,
		Classes:              []*models.Class{class},
		Subclasses:           []*models.Subclass{subclass},
		ProfilePictureFileID: profilePicture.ID,
	}
	factories.NewCharacter(ts.S.Db, character)

	otherUserCharacter := &models.Character{
		UserID:               999,
		ProfilePictureFileID: otherUserProfilePicture.ID,
	}
	factories.NewCharacter(ts.S.Db, otherUserCharacter)

	getRequest := func(id interface{}) helpers.Request {
		return helpers.Request{
			Method: http.MethodGet,
			Url:    fmt.Sprintf("/characters/%v", id),
		}
	}

	permissionRequest := getRequest(character.ID)
	RunNoAuthenticationTests(t, permissionRequest.Method, permissionRequest.Url)

	cases := []helpers.TestCase{
		{
			Name:    "Can't get character with an invalid id",
			Request: getRequest("invalid-id"),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Character not found",
			},
		},
		{
			Name:    "Can't get character that doesn't exist",
			Request: getRequest(999),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Character not found",
			},
		},
		{
			Name:    "Can't get character that belongs to a different user",
			Request: getRequest(otherUserCharacter.ID),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusNotFound,
				BodyPart:   "Character not found",
			},
		},
		{
			Name: "Can't get character if there is an error from the character service",
			Setup: func(test *helpers.TestCase) {
				mockCharacterService := mocks.NewMockCharacterService()
				mockCharacterService.GetError = errors.New("test error")

				ts.S.Dependencies.SetCharacterService(mockCharacterService)
			},
			Request: getRequest(character.ID),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusInternalServerError,
				BodyPart:   "There was an error getting the character",
			},
			Teardown: func() {
				ts.S.Dependencies.SetCharacterService(services.NewCharacterService(ts.S.Repos.Character))
			},
		},
		{
			Name:    "Can get character",
			Request: getRequest(character.ID),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"name":"%v"`, character.Name),
					fmt.Sprintf(`"name":"%v"`, class.Name),
					fmt.Sprintf(`"name":"%v"`, subclass.Name),
					fmt.Sprintf(`"name":"%v"`, race.Name),
					fmt.Sprintf(`"name":"%v"`, campaign.Name),
					fmt.Sprintf(`"name":"%v"`, profilePicture.Name),
					fmt.Sprintf(`"location":"%v"`, profilePicture.Location),
				},
				BodyPartsMissing: []string{
					fmt.Sprintf(`"name":"%v"`, otherUserCharacter.Name),
					fmt.Sprintf(`"name":"%v"`, otherUserProfilePicture.Name),
					fmt.Sprintf(`"location":"%v"`, otherUserProfilePicture.Location),
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
