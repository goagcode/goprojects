package main

import (
	"errors"
	"reflect"
	"testing"
)

type FakeReleaseInfoer struct {
	Tag string
	Err error
}

func (f FakeReleaseInfoer) GetLatestReleaseTag(repo string) (string, error) {
	if f.Err != nil {
		return "", f.Err
	}
	return f.Tag, nil
}

func TestGetReleaseTagMessage(t *testing.T) {
	cases := []struct {
		f               FakeReleaseInfoer
		repo            string
		expectedMessage string
		expectedErr     error
	}{
		{
			f: FakeReleaseInfoer{
				Tag: "v0.1.0",
				Err: nil,
			},
			repo:            "miguellgt/test",
			expectedMessage: "The latest release is v0.1.0",
			expectedErr:     nil,
		},
		{
			f: FakeReleaseInfoer{
				Tag: "v0.1.0",
				Err: errors.New("TCP timeout"),
			},
			repo:            "doesnt/foo",
			expectedMessage: "",
			expectedErr:     errors.New("Error querying Github API: TCP timeout"),
		},
	}

	for _, c := range cases {
		msg, err := getReleaseTagMessage(c.f, c.repo)
		if !reflect.DeepEqual(err, c.expectedErr) {
			t.Errorf("Expected err to be %q but it was %q", c.expectedErr, err)
		}

		if c.expectedMessage != msg {
			t.Fatalf("Expected %q but got %q", c.expectedMessage, msg)
		}
	}
}
