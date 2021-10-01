package presenters

import (
	"reflect"
	"testing"

	"github.com/muhfaa/omdb-server/core"
)

func TestSearchResultToProto(t *testing.T) {
	type args struct {
		inp *core.OMDBSearchResult
	}
	tests := []struct {
		name string
		args args
		want *core.SearchReply
	}{
		{
			name: "1",
			args: args{
				inp: &core.OMDBSearchResult{
					Search: []core.OMDBResultCompact{
						{
							IMDBID: "DavidBowie",
						},
					},
					Response:     "True",
					TotalResults: "3",
				},
			},
			want: &core.SearchReply{
				Search: []*core.SearchEntry{
					{
						ImdbID: "DavidBowie",
					},
				},
				Response:     "True",
				TotalResults: "3",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SearchResultToProto(tt.args.inp); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SearchResultToProto() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSingleToProto(t *testing.T) {
	type args struct {
		inp *core.OMDBResultSingle
	}
	tests := []struct {
		name string
		args args
		want *core.SingleReply
	}{
		{
			name: "1",
			args: args{
				inp: &core.OMDBResultSingle{
					Actors: "Sukab, David Bowie",
					Ratings: []core.OMDBRating{
						{
							Source: "Metacrit",
							Value:  "100",
						},
					},
				},
			},
			want: &core.SingleReply{
				Actors: "Sukab, David Bowie",
				Ratings: []*core.SingleRating{
					{
						Source: "Metacrit",
						Value:  "100",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SingleToProto(tt.args.inp); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SingleToProto() = %v, want %v", got, tt.want)
			}
		})
	}
}
