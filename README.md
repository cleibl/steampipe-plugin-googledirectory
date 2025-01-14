![image](https://hub.steampipe.io/images/plugins/turbot/googledirectory-social-graphic.png)

# Google Directory Plugin for Steampipe

Use SQL to query users, groups, org units and more from your Google Workspace directory.

- **[Get started →](https://hub.steampipe.io/plugins/turbot/googledirectory)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/googledirectory/tables)
- Community: [Slack Channel](https://join.slack.com/t/steampipe/shared_invite/zt-oij778tv-lYyRTWOTMQYBVAbtPSWs3g)
- Get involved: [Issues](https://github.com/turbot/steampipe-plugin-googledirectory/issues)

## Quick start

Install the plugin with [Steampipe](https://steampipe.io):

```shell
steampipe plugin install googledirectory
```

Run a query:

```sql
select
  id,
  primary_email,
  full_name
from
  googledirectory_user;
```
## Credentials

This plugin uses the Admin SDK Directory API to gather information from Google Workspace.  The recommended method for Directory API authentication is to use a GCP Service Account with Domain Wide Delegation.

To setup Domain Wide Delegation follow these [instructions](https://developers.google.com/admin-sdk/directory/v1/guides/delegation).

When asked for scopes, input the following

```
https://www.googleapis.com/auth/admin.directory.group.member.readonly, https://www.googleapis.com/auth/admin.directory.group.readonly,https://www.googleapis.com/auth/admin.directory.orgunit.readonly,https://www.googleapis.com/auth/admin.directory.user.readonly,https://www.googleapis.com/auth/admin.directory.user.alias.readonly,https://www.googleapis.com/auth/admin.directory.user.security,https://www.googleapis.com/auth/admin.directory.rolemanagement.readonly,https://www.googleapis.com/auth/admin.directory.customer.readonly,https://www.googleapis.com/auth/admin.directory.domain.readonly
```

Once the above steps have been completed, follow the instructions in [Developing](#developing).

## Developing

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone git@github.com:turbot/steampipe-plugin-googledirectory
cd steampipe-plugin-googledirectory
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```
make
```

Configure the plugin:

```
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/googledirectory.spc
```

Try it!

```
steampipe query
> .inspect googledirectory
```

Further reading:

- [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
- [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). All contributions are subject to the [Apache 2.0 open source license](https://github.com/turbot/steampipe-plugin-googledirectory/blob/main/LICENSE).

`help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [Google Directory Plugin](https://github.com/turbot/steampipe-plugin-googledirectory/labels/help%20wanted)
