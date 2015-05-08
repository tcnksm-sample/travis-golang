TravisCI for golang
====

[![Build Status](https://travis-ci.org/tcnksm-sample/travis-golang.svg?branch=master)](https://travis-ci.org/tcnksm-sample/travis-golang)

This is sample of releasing golang project from [TravisCI](https://travis-ci.org)

1. Get source code
1. Run test
1. Send test coverage to [coveralls.io](https://coveralls.io/) by [mattn/goveralls](https://github.com/mattn/goveralls)
1. Cross-compile by [mitchellh/gox](https://github.com/mitchellh/gox)
1. Release binary to Github Release by [tcnksm/ghr](https://github.com/tcnksm/ghr)

See `.travis.yml`. To use this, you need to set `GITHUB_TOKEN` and `COVERALLS_TOKEN` in settings tab on travis-ci.org


