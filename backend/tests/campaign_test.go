package tests

import (
	"dnd-character-dashboard/db/factories"
	"dnd-character-dashboard/db/models"
	"dnd-character-dashboard/tests/helpers"
	"fmt"
	"net/http"
	"testing"
)

func TestCampaign_List(t *testing.T) {
	ts.ClearTable("campaigns")
	ts.ClearTable("characters")
	ts.ClearTable("files")
	ts.SetupDefaultUsers()

	campaign := &models.Campaign{Name: "AAA"}
	factories.NewCampaign(ts.S.Db, campaign)
	campaign2 := &models.Campaign{Name: "BBB"}
	factories.NewCampaign(ts.S.Db, campaign2)
	dmCampaign := &models.Campaign{Name: "DM Campaign", UserID: ts.AdminUser.ID}
	factories.NewCampaign(ts.S.Db, dmCampaign)
	namedCampaign := &models.Campaign{Name: "Test Campaign"}
	factories.NewCampaign(ts.S.Db, namedCampaign)
	otherUserCampaign := &models.Campaign{}
	factories.NewCampaign(ts.S.Db, otherUserCampaign)
	neverReturnedCampaign := &models.Campaign{UserID: 999}
	factories.NewCampaign(ts.S.Db, neverReturnedCampaign)

	profilePicture := &models.File{}
	factories.NewFile(ts.S.Db, profilePicture)

	character := &models.Character{CampaignID: campaign.ID, ProfilePictureFileID: profilePicture.ID}
	factories.NewCharacter(ts.S.Db, character)
	sameCampaignCharacter := &models.Character{CampaignID: campaign.ID}
	factories.NewCharacter(ts.S.Db, sameCampaignCharacter)
	character2 := &models.Character{CampaignID: campaign2.ID}
	factories.NewCharacter(ts.S.Db, character2)
	character3 := &models.Character{CampaignID: namedCampaign.ID}
	factories.NewCharacter(ts.S.Db, character3)
	otherUserCharacter := &models.Character{CampaignID: otherUserCampaign.ID, UserID: 999}
	factories.NewCharacter(ts.S.Db, otherUserCharacter)

	getRequest := func(query string) helpers.Request {
		return helpers.Request{
			Method: http.MethodGet,
			Url:    fmt.Sprintf("/campaigns%v", query),
		}
	}

	permissionRequest := getRequest("")
	RunNoAuthenticationTests(t, permissionRequest.Method, permissionRequest.Url)

	cases := []helpers.TestCase{
		{
			Name:    "Can get campaigns",
			Request: getRequest(""),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"name":"%v"`, campaign.Name),
					fmt.Sprintf(`"name":"%v"`, campaign2.Name),
					fmt.Sprintf(`"name":"%v"`, dmCampaign.Name),
					fmt.Sprintf(`"name":"%v"`, namedCampaign.Name),
					fmt.Sprintf(`"name":"%v"`, character.Name),
					fmt.Sprintf(`"name":"%v"`, profilePicture.Name),
					fmt.Sprintf(`"location":"%v"`, profilePicture.Location),
					fmt.Sprintf(`"name":"%v"`, sameCampaignCharacter.Name),
					fmt.Sprintf(`"name":"%v"`, character2.Name),
					fmt.Sprintf(`"name":"%v"`, character3.Name),
					`"total_count":4`,
				},
				BodyPartsMissing: []string{
					fmt.Sprintf(`"name":"%v"`, otherUserCampaign.Name),
					fmt.Sprintf(`"name":"%v"`, otherUserCharacter.Name),
					fmt.Sprintf(`"name":"%v"`, neverReturnedCampaign.Name),
				},
			},
		},
		{
			Name:    "Can get page 0 of campaigns",
			Request: getRequest("?page=0&page_size=1"),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"name":"%v"`, campaign.Name),
					fmt.Sprintf(`"name":"%v"`, character.Name),
					fmt.Sprintf(`"name":"%v"`, profilePicture.Name),
					fmt.Sprintf(`"location":"%v"`, profilePicture.Location),
					fmt.Sprintf(`"name":"%v"`, sameCampaignCharacter.Name),
					`"total_count":4`,
				},
				BodyPartsMissing: []string{
					fmt.Sprintf(`"name":"%v"`, campaign2.Name),
					fmt.Sprintf(`"name":"%v"`, dmCampaign.Name),
					fmt.Sprintf(`"name":"%v"`, namedCampaign.Name),
					fmt.Sprintf(`"name":"%v"`, character2.Name),
					fmt.Sprintf(`"name":"%v"`, character3.Name),
					fmt.Sprintf(`"name":"%v"`, otherUserCampaign.Name),
					fmt.Sprintf(`"name":"%v"`, otherUserCharacter.Name),
					fmt.Sprintf(`"name":"%v"`, neverReturnedCampaign.Name),
				},
			},
		},
		{
			Name:    "Can get page 1 of campaigns",
			Request: getRequest("?page=1&page_size=1"),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"name":"%v"`, campaign2.Name),
					fmt.Sprintf(`"name":"%v"`, character2.Name),
					`"total_count":4`,
				},
				BodyPartsMissing: []string{
					fmt.Sprintf(`"name":"%v"`, campaign.Name),
					fmt.Sprintf(`"name":"%v"`, dmCampaign.Name),
					fmt.Sprintf(`"name":"%v"`, namedCampaign.Name),
					fmt.Sprintf(`"name":"%v"`, character.Name),
					fmt.Sprintf(`"name":"%v"`, profilePicture.Name),
					fmt.Sprintf(`"location":"%v"`, profilePicture.Location),
					fmt.Sprintf(`"name":"%v"`, sameCampaignCharacter.Name),
					fmt.Sprintf(`"name":"%v"`, character3.Name),
					fmt.Sprintf(`"name":"%v"`, otherUserCampaign.Name),
					fmt.Sprintf(`"name":"%v"`, otherUserCharacter.Name),
					fmt.Sprintf(`"name":"%v"`, neverReturnedCampaign.Name),
				},
			},
		},
		{
			Name:    "Can search campaigns by name",
			Request: getRequest("?search=Test"),
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					fmt.Sprintf(`"name":"%v"`, namedCampaign.Name),
					fmt.Sprintf(`"name":"%v"`, character3.Name),
					`"total_count":1`,
				},
				BodyPartsMissing: []string{
					fmt.Sprintf(`"name":"%v"`, campaign.Name),
					fmt.Sprintf(`"name":"%v"`, campaign2.Name),
					fmt.Sprintf(`"name":"%v"`, dmCampaign.Name),
					fmt.Sprintf(`"name":"%v"`, otherUserCampaign.Name),
					fmt.Sprintf(`"name":"%v"`, character.Name),
					fmt.Sprintf(`"name":"%v"`, profilePicture.Name),
					fmt.Sprintf(`"location":"%v"`, profilePicture.Location),
					fmt.Sprintf(`"name":"%v"`, sameCampaignCharacter.Name),
					fmt.Sprintf(`"name":"%v"`, character2.Name),
					fmt.Sprintf(`"name":"%v"`, otherUserCharacter.Name),
					fmt.Sprintf(`"name":"%v"`, neverReturnedCampaign.Name),
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
