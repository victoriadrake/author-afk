# authorAFK

Post tweets with RSS links using AWS Lambda when you're afk.

# How it works

Written in Go, this AWS Lambda function will choose a random RSS item from one of any number of RSS feed links and post it to your Twitter timeline upon invocation.

Optionally, you may add text to the Tweet in front or after the link.

# Set up

For a full walkthrough with screenshots on creating a Lambda function and uploading the code, read [this blog post](https://vickylai.com/verbose/free-twitter-bot-aws-lambda/). Skip to setting environment variables at [this link](https://vickylai.com/verbose/free-twitter-bot-aws-lambda/#2-configure-your-function).

Lambda environment variables you can set are as follows, with example values:

```
// Required variables
TWITTER_CONSUMER_KEY        = TWITTER_API_VAR
TWITTER_CONSUMER_SECRET     = TWITTER_API_VAR
TWITTER_ACCESS_TOKEN        = TWITTER_API_VAR
TWITTER_ACCESS_TOKEN_SECRET = TWITTER_API_VAR
RSS_FEEDS                   = https://blog.com/index.xml;https://anotherblog.com/index.xml

// Optional variables
PREFIX                      = "Here's a post from my blog."
SUFFIX                      = "#awesome"
```

You will need to [create a new Twitter application and generate API keys](https://apps.twitter.com/) for each `TWITTER_API_VAR`.

The `RSS_FEEDS` variable must be any number of RSS URLs beginning with `http://` or `https://` and ending in `.xml` or `.rss`. It will also parse `.atom` For more on how feeds are parsed, see [gofeed](https://github.com/mmcdole/gofeed). The program expects any more than one URL to be separated by a `;`.

# update.sh

This handy bash script is included to help you upload your function code to Lambda. It requires [AWS Command Line Interface](https://aws.amazon.com/cli/). To set up, do `pip install awscli` and follow these instructions for [Quick Configuration](https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-getting-started.html).

# AWS SAM CLI

You can use [AWS Serverless Application Model CLI](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/what-is-sam.html) with this repository to test your function locally. __If you test this program locally, it will act on your real live timeline.__

Find installation and getting started guide here: [AWS SAM Reference](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-sam-reference.html)

To run locally, AWS SAM requires a `template.yaml` with your environment variables (among other things) [in it](https://github.com/awslabs/aws-sam-cli/issues/139#issuecomment-334977285). __For this reason, `template.yaml` should be included in `.gitignore` in this repository.__ You can generate a sample template with `sam init`.

To build the program:

```shell
GOOS=linux go build -o authorAFK
```

**NOTE**: If you're not building the function on a Linux machine, you will need to specify the `GOOS` and `GOARCH` environment variables. This allows Go to build your function for another system architecture and ensure compatibility.

## Local development

You can invoke this Lambda function locally by running:

```bash
sam local invoke --no-event
```

> **See [Serverless Application Model (SAM) HOWTO Guide](https://github.com/awslabs/serverless-application-model/blob/master/HOWTO.md) for more details on how to get started.**

# Contributing

Pull requests for bug fixes and improvements are always welcome.