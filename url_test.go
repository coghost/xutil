package xutil_test

import (
	"testing"

	"github.com/coghost/xutil"

	"github.com/google/go-querystring/query"
	"github.com/gookit/goutil/dump"
	"github.com/stretchr/testify/suite"
)

type UrlSuite struct {
	suite.Suite
}

func TestUrl(t *testing.T) {
	suite.Run(t, new(UrlSuite))
}

func (s *UrlSuite) SetupSuite() {
}

func (s *UrlSuite) TearDownSuite() {
}

func (s *UrlSuite) TestDecodeUrl() {
	baseurl := "https://www.vivian.com/browse-jobs?employmentType=Permanent&prevent-filtering=&page=1&configure%5BhitsPerPage%5D=25&configure%5BfacetingAfterDistinct%5D=true&configure%5Bfilters%5D=&configure%5BuserToken%5D=&configure%5BclickAnalytics%5D=true&configure%5BenablePersonalization%5D=false&configure%5BmaxValuesPerFacet%5D=1000&refinementList%5BemploymentType%5D=&refinementList%5Blocation%5D%5B0%5D=Texas"
	dat := xutil.DecodeUrl(baseurl)
	s.Equal("25", dat["configure[hitsPerPage]"][0])
	s.Equal("Permanent", dat["employmentType"][0])
}

func (s *UrlSuite) TestEncode() {
	baseurl := "https://www.vivian.com/browse-jobs?employmentType=Permanent&prevent-filtering=&page=1&refinementList%5Blocation%5D%5B0%5D=Texas"
	type args struct {
		baseurl string
		params  map[string]interface{}
		pth     string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "set page to 10",
			args: args{
				baseurl: baseurl,
				params: map[string]interface{}{
					"page": 10,
				},
				pth: "",
			},
			want: "https://www.vivian.com/browse-jobs?employmentType=Permanent&page=10&prevent-filtering=&refinementList%5Blocation%5D%5B0%5D=Texas",
		},
		{
			name: "set employmentType to Temporary",
			args: args{
				baseurl: baseurl,
				params: map[string]interface{}{
					"employmentType": "Temporary",
				},
				pth: "",
			},
			want: "https://www.vivian.com/browse-jobs?employmentType=Temporary&page=1&prevent-filtering=&refinementList%5Blocation%5D%5B0%5D=Texas",
		},
		{
			name: "set employmentType to Temporary",
			args: args{
				baseurl: baseurl,
				params: map[string]interface{}{
					"employmentType": "Temporary",
				},
				pth: "userid=b1291kasdk12",
			},
			want: "https://www.vivian.com/browse-jobs/userid=b1291kasdk12?employmentType=Temporary&page=1&prevent-filtering=&refinementList%5Blocation%5D%5B0%5D=Texas",
		},
	}
	for _, tt := range tests {
		got := xutil.EncodeUrl(tt.args.baseurl, tt.args.params, tt.args.pth)
		s.Equal(tt.want, got.String(), tt.name)
	}
}

func (s *UrlSuite) Test02_ExampleOfUrlWithStruct() {
	// ExampleOfUrlWithStruct
	type Options struct {
		Ordering string `url:"ordering"`
		Query    string `url:"query"`
		Page     int    `url:"page"`
	}
	// ordering=relevancy&query=Medical-Surgical&page=1
	baseUrl := "https://nomadhealth.com/api/jobposts/jobpost_search/?"
	opt := Options{
		Ordering: "relevancy",
		Query:    "Critical Care",
		Page:     2,
	}
	v, _ := query.Values(opt)
	dump.P(v.Encode())
	dump.P(baseUrl + v.Encode())
}
