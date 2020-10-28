package types

// trustchain module event types
const (
	// TODO: Create your event types
	// EventType<Action>    		= "action"
	EventTypeCreatePromise = "CreatePromise"

	// TODO: Create keys fo your events, the values will be derivided from the msg
	// AttributeKeyAddress  		= "address"
	AttributePromiseDescription = "promiseDescription"
	AttributePromiseKeeper      = "promiseKeeper"
	AttributeReward             = "reward"

	// TODO: Some events may not have values for that reason you want to emit that something happened.
	// AttributeValueDoubleSign = "double_sign"

	AttributeValueCategory = ModuleName
)
