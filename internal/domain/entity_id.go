package domain

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
)

type EntityID struct {
	value string
}

func NewEntityID() EntityID {
	return EntityID{value: uuid.New().String()}
}

func (eid EntityID) String() string {
	return eid.value
}

func (eid EntityID) IsZero() bool {
	return eid.value == ""
}

func ParseEntityID(s string) (EntityID, error) {
	if _, err := uuid.Parse(s); err != nil {
		return EntityID{}, fmt.Errorf("invalid EntityID: %v", err)
	}
	return EntityID{value: s}, nil
}

func (eid EntityID) MarshalJSON() ([]byte, error) {
	return json.Marshal(eid.value)
}

func (eid *EntityID) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	parsed, err := ParseEntityID(s)
	if err != nil {
		return err
	}
	*eid = parsed
	return nil
}

func (eid EntityID) Value() (driver.Value, error) {
	return eid.value, nil
}

func (eid *EntityID) Scan(value interface{}) error {
	if value == nil {
		return fmt.Errorf("cannot scan nil into EntityID")
	}

	switch v := value.(type) {
	case []byte:
		return eid.UnmarshalJSON(v)
	case string:
		parsed, err := ParseEntityID(v)
		if err != nil {
			return err
		}
		*eid = parsed
		return nil
	default:
		return fmt.Errorf("cannot scan %T into EntityID", value)
	}
}
