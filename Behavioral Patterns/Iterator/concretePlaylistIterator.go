package main

type PlaylistIterator struct {
	playlist *Playlist
	index    int
}

func NewPlaylistIterator(playlist Playlist) *PlaylistIterator {
	return &PlaylistIterator{
		playlist: &playlist,
		index:    0,
	}
}

func (pi *PlaylistIterator) HasNext() bool {
	return pi.index < pi.playlist.getSize()
}

func (pi *PlaylistIterator) Next() string {
	song := pi.playlist.getSongs(pi.index)
	pi.index++
	return song
}
