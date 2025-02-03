package id_test

import (
	"encoding/json"
	"testing"

	"github.com/henriquepw/pobrin-api/pkg/id"
	"github.com/henriquepw/pobrin-api/pkg/testutil/assert"
)

func TestIDMarshaling(t *testing.T) {
	type TestStruct struct {
		ID   id.ID  `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
	}

	id := id.New()
	x := map[string]any{"id": id, "name": "Jose"}
	xBytes, err := json.Marshal(x)
	assert.Nil(t, err)
	testStruct := TestStruct{}
	err = json.Unmarshal(xBytes, &testStruct)
	assert.Nil(t, err)
}

func TestIDScanValue(t *testing.T) {
	t.Run("null id", func(t *testing.T) {
		var id id.ID

		err := id.Scan(nil)
		assert.Nil(t, err)
		assert.Equal(t, "", id)
	})

	t.Run("valid id", func(t *testing.T) {
		var testID id.ID
		expectedID := id.New()
		err := testID.Scan(expectedID)
		assert.Nil(t, err)
		assert.Equal(t, expectedID, testID)
	})

	t.Run("valid string id", func(t *testing.T) {
		var testID id.ID
		expectedID := id.New()
		err := testID.Scan(string(expectedID))
		assert.Nil(t, err)
		assert.Equal(t, expectedID, testID)
	})

	t.Run("valid byte id", func(t *testing.T) {
		var testID id.ID
		expectedID := id.New()
		err := testID.Scan([]byte(expectedID))
		assert.Nil(t, err)
		assert.Equal(t, expectedID, testID)
	})

	t.Run("invalid type id", func(t *testing.T) {
		var testID id.ID

		err := testID.Scan(123)
		assert.NotNil(t, err)
	})
}

func TestParseAndValidate(t *testing.T) {
	t.Run("Valid ID", func(t *testing.T) {
		id, err := id.Parse("ABC")
		assert.Nil(t, err)
		assert.NotEmptyString(t, id)
	})

	t.Run("Invalid ID", func(t *testing.T) {
		id, err := id.Parse("1w")
		assert.Nil(t, err)
		assert.NotEmptyString(t, id)
	})

}
