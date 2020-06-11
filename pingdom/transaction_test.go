package pingdom

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransactionCheckServiceList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/tms/check", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{
			"checks": [
				{
					"active": true,
					"created_at": 1553070682,
					"id": 1,
					"interval": 10,
					"modified_at": 1553070968,
					"name": "My check 1",
					"region": "us-west",
					"status": "successful",
					"type": "script",
					"tags": [
						{
							"name": "apache",
							"type": "a",
							"count": 2
						}
					]
				},
				{
					"active": true,
					"created_at": 1553070682,
					"id": 2,
					"interval": 10,
					"modified_at": 1553070968,
					"name": "My check 2",
					"region": "us-west",
					"status": "successful",
					"type": "script",
					"tags": [
						{
							"name": "nginx",
							"type": "u",
							"count": 1
						}
					]
				},			
				{
					"active": true,
					"created_at": 1553070682,
					"id": 3,
					"interval": 10,
					"modified_at": 1553070968,
					"name": "My check 3",
					"region": "us-west",
					"status": "successful",
					"type": "script",
					"tags": [
						{
							"name": "apache",
							"type": "a",
							"count": 2
						}
					]
				}
			]
		}`)
	})

	var countA, countB float64 = 1, 2

	want := []TransactionCheckResponse{
		{
			ID:         1,
			Active:     true,
			CreatedAt:  1553070682,
			ModifiedAt: 1553070968,
			Interval:   10,
			Name:       "My check 1",
			Region:     "us-west",
			Status:     "successful",
			Type:       "script",
			Tags: []TransactionCheckResponseTag{
				{
					Name:  "apache",
					Type:  "a",
					Count: countB,
				},
			},
		},
		{
			ID:         2,
			Active:     true,
			CreatedAt:  1553070682,
			ModifiedAt: 1553070968,
			Interval:   10,
			Name:       "My check 2",
			Region:     "us-west",
			Status:     "successful",
			Type:       "script",
			Tags: []TransactionCheckResponseTag{
				{
					Name:  "nginx",
					Type:  "u",
					Count: countA,
				},
			},
		},
		{
			ID:         3,
			Active:     true,
			CreatedAt:  1553070682,
			ModifiedAt: 1553070968,
			Interval:   10,
			Name:       "My check 3",
			Region:     "us-west",
			Status:     "successful",
			Type:       "script",
			Tags: []TransactionCheckResponseTag{
				{
					Name:  "apache",
					Type:  "a",
					Count: countB,
				},
			},
		},
	}

	checks, err := client.TransactionChecks.List()
	assert.NoError(t, err)
	assert.Equal(t, want, checks)
}

func TestTransactionCheckServiceRead(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/tms/check/85975", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{
			"check": {
				"id": 85975,
				"active": true,
				"contact_ids": [
					12345678,
					19876654
				],
				"created_at": 1553070682,
				"modified_at": 1553070968,
				"custom_message": "My custom message",
				"interval": 10,
				"name": "My awesome check",
				"region": "us-west",
				"send_notification_when_down": 1,
				"severity_level": "low",
				"status": "successful",
				"steps": [{}],
				"team_ids": [
					12345678,
					135790
				],
				"integration_ids": [
					1234,
					1359
				],
				"metadata": {
					"width": 1950,
					"height": 1080,
					"disableWebSecurity": true,
					"authentications": {}
				},
				"tags": [
					{
						"name": "apache",
						"type": "a",
						"count": 2
					}
				],
				"type": "script"
			}
		}`)
	})

	var count float64 = 2
	want := &TransactionCheckResponse{
		ID:                       85975,
		Active:                   true,
		ContactIds:               []int{12345678, 19876654},
		CreatedAt:                1553070682,
		ModifiedAt:               1553070968,
		CustomMessage:            "My custom message",
		Interval:                 10,
		Name:                     "My awesome check",
		Region:                   "us-west",
		SendNotificationWhenDown: 1,
		SeverityLevel:            "low",
		Status:                   "successful",
		TeamIds:                  []int{12345678, 135790},
		IntegrationIds:           []int{1234, 1359},
		Tags: []TransactionCheckResponseTag{{
			Name:  "apache",
			Type:  "a",
			Count: count,
		}},
		Type:  "script",
		Steps: []interface{}{make(map[string]interface{})},
		Metadata: map[string]interface{}{
			"width":              1950.0,
			"height":             1080.0,
			"disableWebSecurity": true,
			"authentications":    make(map[string]interface{}),
		},
	}

	check, err := client.TransactionChecks.Read(85975)
	assert.NoError(t, err)
	assert.Equal(t, want, check)
}
