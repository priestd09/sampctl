package rook

import (
	"testing"

	"github.com/Masterminds/semver"
)

func Test_generateVersionIncString(t *testing.T) {
	type args struct {
		repo    string
		version string
	}
	tests := []struct {
		name       string
		args       args
		wantResult string
	}{
		{"1", args{"samp-ini", "1.1.1"},
			`// This file was generated by "sampctl package release"
// DO NOT EDIT THIS FILE MANUALLY!
// To update the version number for a new release, run "sampctl package release"

#define SAMP_INI_VERSION_MAJOR (1)
#define SAMP_INI_VERSION_MINOR (1)
#define SAMP_INI_VERSION_PATCH (1)
`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ver, _ := semver.NewVersion(tt.args.version)
			if gotResult := generateVersionIncString(tt.args.repo, ver); gotResult != tt.wantResult {
				t.Errorf("generateVersionIncString() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}