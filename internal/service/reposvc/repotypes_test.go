package reposvc

import (
	"testing"
)

func TestExpludeRepoUrl(t *testing.T) {
	tests := []struct {
		repo     string
		expected *RepoExplosionStruct
		hasError bool
	}{
		{
			repo: "https://github.com/user/repo.git",
			expected: &RepoExplosionStruct{
				Protocol: "https",
				Domain:   "github.com",
				Repopath: "github.com/user/repo",
			},
			hasError: false,
		},
		{
			repo: "https://github.com/user/repo",
			expected: &RepoExplosionStruct{
				Protocol: "https",
				Domain:   "github.com",
				Repopath: "github.com/user/repo",
			},
			hasError: false,
		},
		{
			repo: "git@github.com:user/repo",
			expected: &RepoExplosionStruct{
				Protocol: "ssh",
				Domain:   "github.com",
				Repopath: "github.com/user/repo",
			},
			hasError: false,
		},
		{
			repo: "git@github.com:user/repo.git",
			expected: &RepoExplosionStruct{
				Protocol: "ssh",
				Domain:   "github.com",
				Repopath: "github.com/user/repo",
			},
			hasError: false,
		},
		{
			repo:     "invalid-url",
			expected: nil,
			hasError: true,
		},
	}

	for _, test := range tests {
		result, err := ExpludeRepoURL(test.repo)
		if test.hasError {
			if err == nil {
				t.Errorf("Expected error for repo %s, but got none", test.repo)
			}
		} else {
			if err != nil {
				t.Errorf("Did not expect error for repo %s, but got %v", test.repo, err)
			}
			if result.Protocol != test.expected.Protocol || result.Domain != test.expected.Domain || result.Repopath != test.expected.Repopath {
				t.Errorf("For repo %s, expected %v, but got %v", test.repo, test.expected, result)
			}
		}
	}
}
