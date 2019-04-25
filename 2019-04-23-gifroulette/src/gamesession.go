package main

type GameSession struct {
	SessionID          string     `json:"sessionID"`
	SessionCookieToken string     `json:"sessionCookieToken"`
	SessionAdminToken  string     `json:"sessionAdminToken"`
	PlayerCount        int        `json:"playerCount"`
	Questions          []Question `json:"questions"`
}

type Question struct {
	Prompt  string `json:"prompt"`
	GifLink string `json:"gifLink"`
}