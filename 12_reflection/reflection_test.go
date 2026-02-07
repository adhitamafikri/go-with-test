package reflection

import (
	"reflection/test_utils"
	"testing"
)

func TestWalk(t *testing.T) {
	t.Run("Should be able to handle these", func(t *testing.T) {
		cases := []struct {
			Title         string
			Input         interface{}
			ExpectedCalls []string
		}{
			{
				Title: "struct with one string field",
				Input: struct {
					Name string
				}{Name: "Quavo"},
				ExpectedCalls: []string{"Quavo"},
			},
			{
				Title: "struct with two string field",
				Input: struct {
					Name string
					Game string
				}{Name: "Quavo", Game: "Apex Legends"},
				ExpectedCalls: []string{"Quavo", "Apex Legends"},
			},
			{
				"struct with non string field",
				struct {
					Name string
					Age  int
				}{"Chris", 33},
				[]string{"Chris"},
			},
			{
				Title:         "struct with nested fields",
				Input:         Person{Name: "Quavo", Profile: Profile{Age: 29, City: "Depok"}, Game: Game{Title: "Apex Legends", Genre: "FPS"}},
				ExpectedCalls: []string{"Quavo", "Depok", "Apex Legends", "FPS"},
			},
			{
				Title:         "struct with a pointer field",
				Input:         &Person{Name: "Quavo", Profile: Profile{Age: 29, City: "Depok"}, Game: Game{Title: "Apex Legends", Genre: "FPS"}},
				ExpectedCalls: []string{"Quavo", "Depok", "Apex Legends", "FPS"},
			},
			{
				Title: "struct with a slices",
				Input: [3]Game{
					{Title: "Valorant", Genre: "FPS"},
					{Title: "NBA 2k25", Genre: "Sport"},
					{Title: "Balenciaga Challenge", Genre: "Fashion"},
				},
				ExpectedCalls: []string{"Valorant", "FPS", "NBA 2k25", "Sport", "Balenciaga Challenge", "Fashion"},
			},
		}

		for _, tc := range cases {
			t.Run(tc.Title, func(t *testing.T) {
				var got []string

				walk(tc.Input, func(input string) {
					got = append(got, input)
				})

				test_utils.AssertValueDeepEqual(t, got, tc.ExpectedCalls)
			})
		}
	})

	t.Run("Test with Map (bcus maps does not guarantee order)", func(t *testing.T) {
		value := map[string]string{
			"Book":   "The book #1",
			"Author": "Playboi Carti",
		}

		var got []string
		walk(value, func(input string) {
			got = append(got, input)
		})

		test_utils.AssertContains(t, got, "The book #1")
		test_utils.AssertContains(t, got, "Playboi Carti")
	})

	t.Run("Test with Channel", func(t *testing.T) {
		pChannel := make(chan Person)

		go func() {
			pChannel <- Person{Name: "Adhitama Fikri", Profile: Profile{Age: 29, City: "Jakarta"}, Game: Game{Title: "PAN", Genre: "Action"}}
			pChannel <- Person{Name: "Adhitama Kuro", Profile: Profile{Age: 29, City: "Depok"}, Game: Game{Title: "GOW", Genre: "Action"}}
			close(pChannel)
		}()

		var got []string
		expected := []string{"Adhitama Fikri", "Jakarta", "PAN", "Action", "Adhitama Kuro", "Depok", "GOW", "Action"}

		walk(pChannel, func(input string) {
			got = append(got, input)
		})

		test_utils.AssertValueDeepEqual(t, got, expected)
	})

	t.Run("Test with function", func(t *testing.T) {
		aFunc := func() (Game, Game) {
			return Game{Title: "Valorant", Genre: "FPS"}, Game{Title: "GuitarHero", Genre: "Music"}
		}

		var got []string
		expected := []string{"Valorant", "FPS", "GuitarHero", "Music"}

		walk(aFunc, func(input string) {
			got = append(got, input)
		})

		test_utils.AssertValueDeepEqual(t, got, expected)
	})
}
