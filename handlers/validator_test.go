package handlers

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

type mockReadCloser int

var testErr = errors.New("test error")

func (mrc mockReadCloser) Read(p []byte) (n int, err error) {
	return 0, testErr
}

func (mrc mockReadCloser) Close() error {
	return testErr
}

func TestValidate(t *testing.T) {
	gin.SetMode(gin.TestMode)

	testCases := map[string]struct {
		model      interface{}
		file       string
		json       string
		wantErr    bool
		errors     []error
		readCloser io.ReadCloser
	}{
		"validates a course successfully": {
			file:  "./../handlers/schemas/course_schema.json",
			model: &Course{},
			json: `{
				"name": "Nursing 101",
				"institution": "Tampa Bay Nurses United University",
				"creditHours": 3,
				"length": 16,
				"goal": "8-10 hours"
			}`,
			wantErr: false,
		},
		"validates a module successfully": {
			file:  "./../handlers/schemas/module_schema.json",
			model: &Module{},
			json: `{
				"name": "Module 9",
				"number": 9,
				"moduleActivities":[
					{
						"input": 30,
						"notes": "A note",
						"activityId": 1
					},
					{
						"input": 180,
						"activityId": 11
					}
				]
			}`,
			wantErr: false,
		},
		"returns an error if the request is empty": {
			file:    "./../handlers/schemas/module_schema.json",
			model:   &Course{},
			wantErr: true,
			errors: []error{
				errors.New("invalid request"),
			},
		},
		"returns errors for an invalid course": {
			file:  "./../handlers/schemas/course_schema.json",
			model: &Course{},
			json: `{
				"creditHours": 3,
				"length": 16,
				"goal": "8-10 hours"
			}`,
			wantErr: true,
			errors: []error{
				errors.New("name is required"),
				errors.New("institution is required"),
			},
		},
		"returns an error for an invalid schema": {
			file:  "schema.json",
			model: &Course{},
			json: `{
				"name": "Nursing 101",
				"institution": "Tampa Bay Nurses United University",
				"creditHours": 3,
				"length": 16,
				"goal": "8-10 hours"
			}`,
			wantErr: true,
			errors: []error{
				errors.New("Reference file://schema.json must be canonical"),
			},
		},
		"returns json reading errors": {
			file:  "./../handlers/schemas/course_schema.json",
			model: &Course{},
			json: `{
				"name": "Nursing 101",
				"institution": "Tampa Bay Nurses United University",
				"creditHours": 3,
				"length": 16,
				"goal": "8-10 hours"
			}`,
			wantErr:    true,
			readCloser: mockReadCloser(0),
			errors: []error{
				testErr,
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			if tc.json != "" {
				reader := strings.NewReader(tc.json)
				ctx.Request = &http.Request{
					Header: make(http.Header),
					Body: func() io.ReadCloser {
						if tc.readCloser != nil {
							return tc.readCloser
						}
						return io.NopCloser(reader)
					}(),
				}
			}

			errs := Validate(ctx, tc.model, tc.file)

			if tc.wantErr {
				for i := 0; i < len(errs); i++ {
					if errs[i].Error() != tc.errors[i].Error() {
						t.Errorf("got %v want %v for errors[%d]", errs[i], tc.errors[i], i)
					}
				}
			} else {
				if errs != nil {
					for _, e := range errs {
						t.Errorf("unexpected error: %v", e)
					}
				}
			}
		})
	}
}
