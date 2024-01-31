package model

type DiscordUser struct {
    Id     string `json:"id"`
    Handle string `json:"handle"`
    Avatar string `json:"avatar" default:""`
}
