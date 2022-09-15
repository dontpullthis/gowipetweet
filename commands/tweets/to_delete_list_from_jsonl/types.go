package to_delete_list_from_jsonl

type Tweet struct {
	Id            string `json:"id"`
	FavoriteCount string `json:"favorite_count"`
	FullText      string `json:"full_text"`
	RetweetCount  string `json:"retweet_count"`
}

type Record struct {
	Tweet Tweet `json:"tweet"`
}
