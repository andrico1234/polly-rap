package spotifyhelper

const dummyText = `
Yeah, passed the dawgs a celly
Sendin' texts, ain't sendin' kites, yeah
He said "keep that on lock"
I say "you know this shit, it's stife", yeah
It's absolute, yeah (yeah) I'm back, reboot (it's lit!)
LaFerrari to Jamba Juice, yeah (skrrt, skrrt)
We back on the road, they jumpin' off, no parachute, yeah
Shawty in the back
She said she workin' on her glutes, yeah (oh my God)
Ain't by the book, yeah, this how it look, yeah
'Bout a check, yeah, (check) just check the foots, yeah
Pass this to my daughter, I'ma show her what it took (yeah)
Baby mama cover Forbes, got these other bitches shook
Yeah
`

/*
SongData is a struct whose instances hold the data retrieved from spotify
*/
type SongData struct {
	Name   string
	Lyrics string
	Artist string
	Title  string
}

/*
GetSong calls the Spotify API and retrieves the relevant song data.
*/
func GetSong() *SongData {
	mockSong := SongData{Name: "hello", Lyrics: dummyText, Artist: "Travis Scott", Title: "Sicko Mode"}

	return &mockSong
}
