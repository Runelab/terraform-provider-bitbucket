---
layout: "bitbucket"
page_title: "Provider: Bitbucket"
sidebar_current: "docs-bitbucket-index"
description: |-
  The Bitbucket provider to interact with repositories, projects, etc..
---

This is a fork of [thamdub/terraform-provider-bitbucket](https://github.com/thamdub/terraform-provider-bitbucket) for security purposes only (code review).

# Bitbucket Provider

The Bitbucket provider allows you to manage resources including repositories,
webhooks, and default reviewers.

Use the navigation to the left to read about the available resources.

## Example Usage

```hcl
# Configure the Bitbucket Provider
provider "bitbucket" {
  username = "GobBluthe"
  password = "idoillusions" # you can also use app passwords
}

resource "bitbucket_repository" "illusions" {
  owner      = "theleagueofmagicians"
  name       = "illusions"
  scm        = "hg"
  is_private = true
}

resource "bitbucket_project" "project" {
  owner      = "theleagueofmagicians" # must be a team
  name       = "illusions-project"
  key        = "ILLUSIONSPROJ"
  is_private = true
}
```

## Argument Reference

The following arguments are supported in the `provider` block:

* `username` - (Required) Your username used to connect to bitbucket. You can
  also set this via the environment variable. `BITBUCKET_USERNAME`

* `password` - (Required) Your password used to connect to bitbucket. You can
  also set this via the environment variable. `BITBUCKET_PASSWORD`
