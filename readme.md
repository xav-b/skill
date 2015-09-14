# [Skill][releases]

[![Circle CI](https://circleci.com/gh/hackliff/skill.svg?style=svg)](https://circleci.com/gh/hackliff/skill)

> Teach your system a new trick

`skill` makes it easy to download and install remote binaries like ones hosted on [Bintray][1] or [Github releases][2].

## Getting started (with docker)

Installing relies currently on Docker but binaries and instructions for building from source are coming soon.

```sh
# installation:
#   - build the container
#   - cross compile the project in ./bin directory
#   - copy the relevant one in /usr/local/bin
make

# testing
make tests

# Usage
./bin/skill-darwin-amd64 -help
./bin/skill-darwin-amd64 -dest bin https://dl.bintray.com/mitchellh/consul/0.5.2_darwin_amd64.zip
```

[1]: https://bintray.com/
[2]: https://github.com/blog/1547-release-your-software
[releases]: https://github.com/hackliff/skill/releases
