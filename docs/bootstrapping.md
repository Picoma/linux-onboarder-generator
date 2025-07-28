# Bootstrapping

## The problem
During the initial attempts to configure workstations in my company (which was entirely grassroots: we were allowed to use Linux in exchange for no IT support), we wanted to manage tool versions to help with troubleshooting. Between the deprecations of certain tools (such as OpenLens) and specific configurations, it was difficult to have a coherent image of one system compared to another.

So the CLI was responsible for installing **all possible tools** through Ansible. Which was impossible, of course: you can't have an op mode for all the tools in the world...

To solve these problems, we decided to use **[mise](https://mise.jdx.dev/)** to manage multiple versions of tools. Through a command, we had the list of all the tools used. But the policy being that the CLI is the unique entry point, the CLI drove mise.

But then, another problem encountered was that the evolutions of the CLI itself could cause regressions. The CLI not being managed by mise (and not being in Go, and not using versioned Ansible collections...), it was impossible to revert to a previous version in case of a problem.

So, how to integrate the CLI into mise if mise must be installed by the CLI?

## Solution: Bootstrapping the CLI with Mise at installation
To address this problem, we decided to integrate mise during the CLI installation. This allows us to maintain transparency for the end user, who ultimately installs a CLI, while integrating the version management of the CLI itself, and thus is able to revert to a previous version in case of regression.
