# gowipetweet : Twitter history eraser

## Requirements

This application uses OAuth 1 for user authentication. To obtain a consumer key and a consumer key secret, the application has to be registered at Twitter Developer Platform https://developer.twitter.com.

## Usage

### Help

```
gowipetweet h
gowipetweet -h
gowipetweet help
gowipetweet --help
```

__Common parameters__

- -c / --config - path to configuration file. See [config.example.yaml](config.example.yaml). Default value is `config.yaml`, so configuration file will be read from current directory.

### Conversion of Javascript tweet dump file into JSON Lines

Twitter provides dumps as JavaScript files which are inapproptiate for analysis and filtering of records. This command converts the original JavaScript file into [JSON Lines format](https://jsonlines.org/).

```bash
gowipetweettweets:dump:to_jsonl -c $PWD/.local/config.yaml -i /home/john/somefolder/twitter_dump/data/tweet.js -o /home/john/somefolder/twitter_dump_processed/tweets.js
```

### Deletion of tweets using CSV file

__Parameters__

- -i / --input-file - path to CSV file where IDs of tweets to delete are listed. Each line contains a single column which is tweet ID. No CSV header. _Note: it was probably a bad idea to mention the CSV format in this command, because records are not actually comma-separated :D Probably this name will be changed later._

```bash
gowipetweet tweets:delete:using_csv -c $PWD/.local/config.yaml -i /home/john/somefolder/tweets_to_delete.csv
```

## TODO

- Config file should NOT be mandatory for all commands. Only those commands which interact with Twitter API need it
- Filtering of dump records, preferably using exression evaluation library (`Govaluate`?). The goal is to filter tweets using some customizable logic e.g. number of likes/retweets