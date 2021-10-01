package omdb_test

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/muhfaa/omdb-server/repository"
	omdbrepo "github.com/muhfaa/omdb-server/repository/omdb"
)

var _ = Describe("Impl", func() {

	baseURL := "http://localhost/"
	apiKey := "anystring"

	Describe("Search", func() {
		When("there's any result", func() {
			It("ok", func() {
				c := repository.NewMockHTTPClient(func(req *http.Request) (*http.Response, error) {

					jsonResponse := `
					{
						"Search": [
							{
								"Title": "Batman: The Killing Joke",
								"Year": "2016",
								"imdbID": "tt4853102",
								"Type": "movie",
								"Poster": "https://m.media-amazon.com/images/M/MV5BMTdjZTliODYtNWExMi00NjQ1LWIzN2MtN2Q5NTg5NTk3NzliL2ltYWdlXkEyXkFqcGdeQXVyNTAyODkwOQ@@._V1_SX300.jpg"
							}
						],
						"totalResults": "397",
						"Response": "True"
					}
					`

					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       ioutil.NopCloser(bytes.NewBufferString(jsonResponse)),
					}, nil
				})

				impl := omdbrepo.NewOMDBRepo(c, baseURL, apiKey)

				out, err := impl.Search(context.Background(), "Batman", 2)
				Expect(err).NotTo(HaveOccurred())
				Expect(len(out.Search)).To(Equal(1))
			})
		})

		When("no data", func() {
			It("error", func() {
				c := repository.NewMockHTTPClient(func(req *http.Request) (*http.Response, error) {

					jsonResponse := `
					{
						"Response": "False"
					}
					`

					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       ioutil.NopCloser(bytes.NewBufferString(jsonResponse)),
					}, nil
				})

				impl := omdbrepo.NewOMDBRepo(c, baseURL, apiKey)

				out, err := impl.Search(context.Background(), "Batman", 2)
				Expect(err).To(HaveOccurred())
				Expect(out).To(BeNil())
			})
		})
	})

	Describe("GetByID", func() {
		When("there's any result", func() {
			It("ok", func() {
				c := repository.NewMockHTTPClient(func(req *http.Request) (*http.Response, error) {

					jsonResponse := `
					{
						"Title": "Batman: The Killing Joke",
						"Year": "2016",
						"Rated": "R",
						"Released": "25 Jul 2016",
						"Runtime": "76 min",
						"Genre": "Animation, Action, Crime, Drama, Thriller",
						"Director": "Sam Liu",
						"Writer": "Brian Azzarello, Brian Bolland (based on the graphic novel illustrated by), Bob Kane (Batman created by), Bill Finger (Batman created by)",
						"Actors": "Kevin Conroy, Mark Hamill, Tara Strong, Ray Wise",
						"Plot": "As Batman hunts for the escaped Joker, the Clown Prince of Crime attacks the Gordon family to prove a diabolical point mirroring his own fall into madness.",
						"Language": "English",
						"Country": "USA",
						"Awards": "1 win & 2 nominations.",
						"Poster": "https://m.media-amazon.com/images/M/MV5BMTdjZTliODYtNWExMi00NjQ1LWIzN2MtN2Q5NTg5NTk3NzliL2ltYWdlXkEyXkFqcGdeQXVyNTAyODkwOQ@@._V1_SX300.jpg",
						"Ratings": [
							{
								"Source": "Internet Movie Database",
								"Value": "6.4/10"
							},
							{
								"Source": "Rotten Tomatoes",
								"Value": "40%"
							}
						],
						"Metascore": "N/A",
						"imdbRating": "6.4",
						"imdbVotes": "48,535",
						"imdbID": "tt4853102",
						"Type": "movie",
						"DVD": "N/A",
						"BoxOffice": "N/A",
						"Production": "Warner Bros.",
						"Website": "N/A",
						"Response": "True"
					}
					`

					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       ioutil.NopCloser(bytes.NewBufferString(jsonResponse)),
					}, nil
				})

				impl := omdbrepo.NewOMDBRepo(c, baseURL, apiKey)

				out, err := impl.GetByID(context.Background(), "tt4853102")
				Expect(err).NotTo(HaveOccurred())
				Expect(out.IMDBID).To(Equal("tt4853102"))
			})
		})

		When("no data", func() {
			It("error", func() {
				c := repository.NewMockHTTPClient(func(req *http.Request) (*http.Response, error) {

					jsonResponse := `
					{
						"Response": "False"
					}
					`

					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       ioutil.NopCloser(bytes.NewBufferString(jsonResponse)),
					}, nil
				})

				impl := omdbrepo.NewOMDBRepo(c, baseURL, apiKey)

				out, err := impl.GetByID(context.Background(), "tt4853102")
				Expect(err).To(HaveOccurred())
				Expect(out).To(BeNil())
			})
		})
	})

})
