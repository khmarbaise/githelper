# TODO

Need to improve build pipeline
```
- name: unit-test
  image: golang:1.24.1
  commands:
  - make unit-test-coverage
  settings:
    group: test
  when:
    event:
    - push
    - pull_request

- name: release-test
  image: golang:1.24.1
  commands:
  - make test
  settings:
    group: test
  when:
    branch:
    - "release/*"
    event:
    - push
    - pull_request

- name: tag-test
  pull: always
  image: golang:1.24.1
  commands:
  - make test
  settings:
    group: test
  when:
    event:
    - tag

- name: static
  image: golang:1.24.1
  commands:
  - make release
  when:
    event:
    - tag

- name: github
  pull: always
  image: plugins/github-release:1
  settings:
    files:
      - "dist/release/*"
  environment:
    GITHUB_TOKEN:
      from_secret: github_token
  when:
    event:
      - tag

```