package reposvc

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/rs/zerolog/log"
)

type RepoExplosionStruct struct {
	Protocol string
	Domain   string
	Repopath string
}

type CreateRepoReq struct {
	Name      string
	Org       int64
	Provider  string
	Repourl   string
	VcsRepoID int64
}

func ExpludeRepoURL(repo string) (*RepoExplosionStruct, error) {
	httpPattern := `^(https?)://([^/]+)/([^/]+)/([^/]+)(?:\.git)?$`
	sshPattern := `^git@([^:]+):([^/]+)/([^/]+)(?:\.git)?$`

	p, d, r, err := func() (string, string, string, error) {
		httpRegex, err := regexp.Compile(httpPattern)
		if err != nil {
			log.Info().Str("repo", repo).Msg("cannot compile regex")
			return "", "", "", err
		}
		if httpRegex.MatchString(repo) {
			matches := httpRegex.FindStringSubmatch(repo)
			protocol := matches[1]
			domain := matches[2]
			organization := matches[3]
			repo := matches[4]
			return protocol, domain, fmt.Sprintf("%s/%s/%s",
				domain, organization, strings.TrimSuffix(repo, ".git")), nil
		}
		sshRegex, err := regexp.Compile(sshPattern)
		if err != nil {
			log.Info().Str("repo", repo).Msg("cannot compile regex")
			return "", "", "", err
		}
		if sshRegex.MatchString(repo) {
			matches := sshRegex.FindStringSubmatch(repo)
			protocol := "ssh"
			domain := matches[1]
			organization := matches[2]
			repo = matches[3]
			return protocol, domain, fmt.Sprintf("%s/%s/%s",
				domain, organization, strings.TrimSuffix(repo, ".git")), nil
		}
		return "", "", "", fmt.Errorf("invalid repo URL: %s", repo)
	}()
	if err != nil {
		log.Info().Err(err).Msg("cannot explode repo url")
		return nil, err
	}
	return &RepoExplosionStruct{
		Protocol: p,
		Repopath: r,
		Domain:   d,
	}, nil
}
