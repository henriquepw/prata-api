package id_test

import (
	"encoding/json"
	"testing"

	"github.com/henriquepw/pobrin-api/pkg/id"
	"github.com/henriquepw/pobrin-api/pkg/testutil"
)

func TestIDMarshaling(t *testing.T) {
	type TestStruct struct {
		ID   id.ID  `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
	}

	id := id.New()
	x := map[string]any{"id": id, "name": "Jose"}
	xBytes, err := json.Marshal(x)
	testutil.NilError(t, err)
	testStruct := TestStruct{}
	err = json.Unmarshal(xBytes, &testStruct)
	testutil.NilError(t, err)
}

func TestIDScanValue(t *testing.T) {
	t.Run("null id", func(t *testing.T) {
		var id id.ID

		err := id.Scan(nil)
		testutil.NilError(t, err)
		testutil.Equal(t, "", id)
	})

	t.Run("valid id", func(t *testing.T) {
		var testID id.ID
		expectedID := id.New()
		err := testID.Scan(expectedID)
		testutil.NilError(t, err)
		testutil.Equal(t, expectedID, testID)
	})

	t.Run("valid string id", func(t *testing.T) {
		var testID id.ID
		expectedID := id.New()
		err := testID.Scan(string(expectedID))
		testutil.NilError(t, err)
		testutil.Equal(t, expectedID, testID)
	})

	t.Run("valid byte id", func(t *testing.T) {
		var testID id.ID
		expectedID := id.New()
		err := testID.Scan([]byte(expectedID))
		testutil.NilError(t, err)
		testutil.Equal(t, expectedID, testID)
	})

	t.Run("invalid type id", func(t *testing.T) {
		var testID id.ID

		err := testID.Scan(123)
		testutil.NotNilError(t, err)
	})
}
