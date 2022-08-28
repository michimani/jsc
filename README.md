jsc
===

'jsc' is a tool that joins multiple slack channels to display posts in chronological order.

# Install

```bash
brew install michimani/jsc/jsc
```

# Usage

```bash
jsc -h
```

```text
     _
    (_)___  ___
    | / __|/ __|
    | \__ \ (__
   _/ |___/\___|
  |__/

Usage:
  jsc [flags] [values]
Flags:
  -c (string)
    ID of the channels you want to join.
    This option can be used multiple times to specify multiple channels.

  -l (integer)
    Limit of messages to be retrieved for each channel (default: 10)

  -o (string)
    Specify the output method as 'json' or 'timeline'. The default is 'timeline',
    which outputs in a format suitable for display.

  -t (string)
    Specify the OAuth Token for the Slack API. It must either be specified
    with this option or set in the environment variable 'SLACK_OAUTH_TOKEN'.
    If both are set, the one specified in this option takes precedence.

Example:
  jsc -c CHANNEL1 -c CHANNEL2 -l 5
```


# License

[MIT](https://github.com/michimani/jsc/blob/main/LICENSE)

# Author

[michimani210](https://twitter.com/michimani210)