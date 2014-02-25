/*
HTTPmock provides tools for mocking HTTP responses.

Simple Example:
	func TestFetchArticles(t *testing.T) {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder("GET", "https://api.mybiz.com/articles.json",
			httpmock.NewStringResponder(`[{"id": 1, "name": "My Great Article"}]`, 200))

		// do stuff that makes a request to articles.json
	}

Advanced Example:
	func TestFetchArticles(t *testing.T) {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		// our database of articles
		articles := make([]map[string]interface{}, 0)

		// mock to list out the articles
		httpmock.RegisterResponder("GET", "https://api.mybiz.com/articles.json",
			func(req *http.Request) (*http.Response, error) {
				if resp, err := httpmock.NewJsonResponse(articles, 200); err != nil {
					return httpmock.NewStringResponse("", 500), nil
				}
				return resp
			},
		)

		// mock to add a new article
		httpmock.RegisterResponder("POST", "https://api.mybiz.com/articles.json",
			func(req *http.Request) (*http.Response, error) {
				article := make(map[string]interface{})
				if err := json.NewDecoder(req.Body).Decode(&article); err != nil {
					return httpmock.NewStringResponse("", 400)
				}

				articles = append(articles, article)

				if resp, err := httpmock.NewJsonResponse(article, 200); err != nil {
					return httpmock.NewStringResponse("", 500), nil
				}
				return resp
			},
		)

		// do stuff that adds and checks articles
	}

*/
package httpmock