package stuff

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const numComics = 1200

func BenchmarkSyncFetch(b *testing.B) {
	for i:=0;i<b.N;i++ {
		_ = Fetch(numComics)
	}
}

func BenchmarkAsyncFetch(b *testing.B) {

	for i := 0; i < b.N; i++ {
		_ = AsyncFetch(numComics)
	}
}

func TestFetch(t *testing.T) {
	assert.ObjectsAreEqual(Fetch(numComics), AsyncFetch(numComics))
}
