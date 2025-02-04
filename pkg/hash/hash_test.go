package hash_test

import (
	"testing"

	"github.com/henriquepw/pobrin-api/pkg/hash"
	"github.com/henriquepw/pobrin-api/pkg/testutil/assert"
)

func TestHash(t *testing.T) {
	t.Run("big secret", func(t *testing.T) {
		s := "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum."
		s72 := "Lorem Ipsum is simply dummy text of the printing and typesetting industr"
		result := hash.MustGenerate(s)
		valid := hash.Validate(result, s72)
		assert.Equal(t, true, valid)
	})

	t.Run("small secret", func(t *testing.T) {
		s := "secureSecret"
		result := hash.MustGenerate(s)
		valid := hash.Validate(result, s)
		assert.Equal(t, true, valid)
	})

	t.Run("Generate", func(t *testing.T) {
		s := "secureSecret"
		result, err := hash.Generate(s)
		assert.Nil(t, err)
		valid := hash.Validate(result, s)
		assert.Equal(t, true, valid)
	})
}
