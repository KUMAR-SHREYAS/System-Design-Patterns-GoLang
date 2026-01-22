package main

type Playlist struct {
	songs []string
}

func NewPlaylist() *Playlist {
	return &Playlist{songs: []string{}}
}

func (p *Playlist) AddSong(song string) {
	p.songs = append(p.songs, song)
}

func (p *Playlist) getSongs(index int) string {
	return p.songs[index]
}

func (p *Playlist) getSize() int {
	return len(p.songs)
}

func (p *Playlist) CreateIterator() Iterator[string] {
	return &PlaylistIterator{playlist: p}
}
