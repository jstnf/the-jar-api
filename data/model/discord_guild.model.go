package model

type DiscordGuild struct {
    Id      string   `json:"id"`
    Name    string   `json:"name"`
    Members []string `json:"members"`
}
