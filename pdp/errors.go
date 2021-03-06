package pdp

/* AUTOMATICALLY GENERATED FROM errors.yaml - DO NOT EDIT */

import (
	"fmt"
	"github.com/google/uuid"
	"strings"
)

// Numeric identifiers of errors.
const (
	externalErrorID                           = 0
	multiErrorID                              = 1
	missingAttributeErrorID                   = 2
	missingValueErrorID                       = 3
	unknownTypeStringCastErrorID              = 4
	invalidTypeStringCastErrorID              = 5
	notImplementedStringCastErrorID           = 6
	invalidBooleanStringCastErrorID           = 7
	invalidAddressStringCastErrorID           = 8
	invalidNetworkStringCastErrorID           = 9
	invalidAddressNetworkStringCastErrorID    = 10
	attributeValueTypeErrorID                 = 11
	duplicateAttributeValueErrorID            = 12
	unknownTypeSerializationErrorID           = 13
	invalidTypeSerializationErrorID           = 14
	assignmentTypeMismatchID                  = 15
	mapperArgumentTypeErrorID                 = 16
	UntaggedPolicyModificationErrorID         = 17
	MissingPolicyTagErrorID                   = 18
	PolicyTagsNotMatchErrorID                 = 19
	emptyPathModificationErrorID              = 20
	invalidRootPolicyItemTypeErrorID          = 21
	hiddenRootPolicyAppendErrorID             = 22
	invalidRootPolicyErrorID                  = 23
	hiddenPolicySetModificationErrorID        = 24
	invalidPolicySetItemTypeErrorID           = 25
	tooShortPathPolicySetModificationErrorID  = 26
	missingPolicySetChildErrorID              = 27
	hiddenPolicyAppendErrorID                 = 28
	policyTransactionTagsNotMatchErrorID      = 29
	failedPolicyTransactionErrorID            = 30
	unknownPolicyUpdateOperationErrorID       = 31
	hiddenPolicyModificationErrorID           = 32
	tooLongPathPolicyModificationErrorID      = 33
	tooShortPathPolicyModificationErrorID     = 34
	invalidPolicyItemTypeErrorID              = 35
	hiddenRuleAppendErrorID                   = 36
	missingPolicyChildErrorID                 = 37
	missingContentErrorID                     = 38
	invalidContentStorageItemID               = 39
	missingContentItemErrorID                 = 40
	invalidContentItemErrorID                 = 41
	invalidContentItemTypeErrorID             = 42
	invalidSelectorPathErrorID                = 43
	networkMapKeyValueTypeErrorID             = 44
	mapContentSubitemErrorID                  = 45
	invalidContentModificationErrorID         = 46
	missingPathContentModificationErrorID     = 47
	tooLongPathContentModificationErrorID     = 48
	invalidContentValueModificationErrorID    = 49
	UntaggedContentModificationErrorID        = 50
	MissingContentTagErrorID                  = 51
	ContentTagsNotMatchErrorID                = 52
	unknownContentUpdateOperationErrorID      = 53
	failedContentTransactionErrorID           = 54
	contentTransactionIDNotMatchErrorID       = 55
	contentTransactionTagsNotMatchErrorID     = 56
	tooShortRawPathContentModificationErrorID = 57
	tooLongRawPathContentModificationErrorID  = 58
	invalidContentUpdateDataErrorID           = 59
	invalidContentUpdateResultTypeErrorID     = 60
	invalidContentUpdateKeysErrorID           = 61
	unknownContentItemResultTypeErrorID       = 62
	invalidContentItemResultTypeErrorID       = 63
	invalidContentKeyTypeErrorID              = 64
	invalidContentStringMapErrorID            = 65
	invalidContentNetworkMapErrorID           = 66
	invalidContentDomainMapErrorID            = 67
	invalidContentValueErrorID                = 68
	invalidContentValueTypeErrorID            = 69
)

type externalError struct {
	errorLink
	err error
}

func newExternalError(err error) *externalError {
	return &externalError{
		errorLink: errorLink{id: externalErrorID},
		err:       err}
}

func (e *externalError) Error() string {
	return e.errorf("%s", e.err)
}

type multiError struct {
	errorLink
	errs []error
}

func newMultiError(errs []error) *multiError {
	return &multiError{
		errorLink: errorLink{id: multiErrorID},
		errs:      errs}
}

func (e *multiError) Error() string {
	msgs := make([]string, len(e.errs))
	for i, err := range e.errs {
		msgs[i] = err.Error()
	}
	msg := strings.Join(msgs, ", ")

	return e.errorf("multiple errors: %s", msg)
}

type missingAttributeError struct {
	errorLink
}

func newMissingAttributeError() *missingAttributeError {
	return &missingAttributeError{
		errorLink: errorLink{id: missingAttributeErrorID}}
}

func (e *missingAttributeError) Error() string {
	return e.errorf("Missing attribute")
}

type missingValueError struct {
	errorLink
}

func newMissingValueError() *missingValueError {
	return &missingValueError{
		errorLink: errorLink{id: missingValueErrorID}}
}

func (e *missingValueError) Error() string {
	return e.errorf("Missing value")
}

type unknownTypeStringCastError struct {
	errorLink
	t int
}

func newUnknownTypeStringCastError(t int) *unknownTypeStringCastError {
	return &unknownTypeStringCastError{
		errorLink: errorLink{id: unknownTypeStringCastErrorID},
		t:         t}
}

func (e *unknownTypeStringCastError) Error() string {
	return e.errorf("Unknown type id %d", e.t)
}

type invalidTypeStringCastError struct {
	errorLink
	t int
}

func newInvalidTypeStringCastError(t int) *invalidTypeStringCastError {
	return &invalidTypeStringCastError{
		errorLink: errorLink{id: invalidTypeStringCastErrorID},
		t:         t}
}

func (e *invalidTypeStringCastError) Error() string {
	return e.errorf("Can't convert string to value of %q type", TypeNames[e.t])
}

type notImplementedStringCastError struct {
	errorLink
	t int
}

func newNotImplementedStringCastError(t int) *notImplementedStringCastError {
	return &notImplementedStringCastError{
		errorLink: errorLink{id: notImplementedStringCastErrorID},
		t:         t}
}

func (e *notImplementedStringCastError) Error() string {
	return e.errorf("Conversion from string to value of %q type hasn't been implemented", TypeNames[e.t])
}

type invalidBooleanStringCastError struct {
	errorLink
	s   string
	err error
}

func newInvalidBooleanStringCastError(s string, err error) *invalidBooleanStringCastError {
	return &invalidBooleanStringCastError{
		errorLink: errorLink{id: invalidBooleanStringCastErrorID},
		s:         s,
		err:       err}
}

func (e *invalidBooleanStringCastError) Error() string {
	return e.errorf("Can't treat %q as boolean (%s)", e.s, e.err)
}

type invalidAddressStringCastError struct {
	errorLink
	s string
}

func newInvalidAddressStringCastError(s string) *invalidAddressStringCastError {
	return &invalidAddressStringCastError{
		errorLink: errorLink{id: invalidAddressStringCastErrorID},
		s:         s}
}

func (e *invalidAddressStringCastError) Error() string {
	return e.errorf("Can't treat %q as IP address", e.s)
}

type invalidNetworkStringCastError struct {
	errorLink
	s   string
	err error
}

func newInvalidNetworkStringCastError(s string, err error) *invalidNetworkStringCastError {
	return &invalidNetworkStringCastError{
		errorLink: errorLink{id: invalidNetworkStringCastErrorID},
		s:         s,
		err:       err}
}

func (e *invalidNetworkStringCastError) Error() string {
	return e.errorf("Can't treat %q as network address (%s)", e.s, e.err)
}

type invalidAddressNetworkStringCastError struct {
	errorLink
	s   string
	err error
}

func newInvalidAddressNetworkStringCastError(s string, err error) *invalidAddressNetworkStringCastError {
	return &invalidAddressNetworkStringCastError{
		errorLink: errorLink{id: invalidAddressNetworkStringCastErrorID},
		s:         s,
		err:       err}
}

func (e *invalidAddressNetworkStringCastError) Error() string {
	return e.errorf("Can't treat %q as address or network (%s)", e.s, e.err)
}

type attributeValueTypeError struct {
	errorLink
	expected int
	actual   int
}

func newAttributeValueTypeError(expected, actual int) *attributeValueTypeError {
	return &attributeValueTypeError{
		errorLink: errorLink{id: attributeValueTypeErrorID},
		expected:  expected,
		actual:    actual}
}

func (e *attributeValueTypeError) Error() string {
	return e.errorf("Expected %s value but got %s", TypeNames[e.expected], TypeNames[e.actual])
}

type duplicateAttributeValueError struct {
	errorLink
	ID   string
	t    int
	curr AttributeValue
	prev AttributeValue
}

func newDuplicateAttributeValueError(ID string, t int, curr, prev AttributeValue) *duplicateAttributeValueError {
	return &duplicateAttributeValueError{
		errorLink: errorLink{id: duplicateAttributeValueErrorID},
		ID:        ID,
		t:         t,
		curr:      curr,
		prev:      prev}
}

func (e *duplicateAttributeValueError) Error() string {
	return e.errorf("Duplicate attribute %q of type %q in request %s - %s", e.ID, TypeNames[e.t], e.curr.describe(), e.prev.describe())
}

type unknownTypeSerializationError struct {
	errorLink
	t int
}

func newUnknownTypeSerializationError(t int) *unknownTypeSerializationError {
	return &unknownTypeSerializationError{
		errorLink: errorLink{id: unknownTypeSerializationErrorID},
		t:         t}
}

func (e *unknownTypeSerializationError) Error() string {
	return e.errorf("Unknown type id %d", e.t)
}

type invalidTypeSerializationError struct {
	errorLink
	t int
}

func newInvalidTypeSerializationError(t int) *invalidTypeSerializationError {
	return &invalidTypeSerializationError{
		errorLink: errorLink{id: invalidTypeSerializationErrorID},
		t:         t}
}

func (e *invalidTypeSerializationError) Error() string {
	return e.errorf("Can't serialize value of %q type", TypeNames[e.t])
}

type assignmentTypeMismatch struct {
	errorLink
	a Attribute
	t int
}

func newAssignmentTypeMismatch(a Attribute, t int) *assignmentTypeMismatch {
	return &assignmentTypeMismatch{
		errorLink: errorLink{id: assignmentTypeMismatchID},
		a:         a,
		t:         t}
}

func (e *assignmentTypeMismatch) Error() string {
	return e.errorf("Can't assign %q value to attribute %q of type %q", TypeNames[e.t], e.a.id, TypeNames[e.a.t])
}

type mapperArgumentTypeError struct {
	errorLink
	actual int
}

func newMapperArgumentTypeError(actual int) *mapperArgumentTypeError {
	return &mapperArgumentTypeError{
		errorLink: errorLink{id: mapperArgumentTypeErrorID},
		actual:    actual}
}

func (e *mapperArgumentTypeError) Error() string {
	return e.errorf("Expected %s, %s or %s as argument but got %s", TypeNames[TypeString], TypeNames[TypeSetOfStrings], TypeNames[TypeListOfStrings], TypeNames[e.actual])
}

// UntaggedPolicyModificationError indicates attempt to modify incrementally a policy which has no tag.
type UntaggedPolicyModificationError struct {
	errorLink
}

func newUntaggedPolicyModificationError() *UntaggedPolicyModificationError {
	return &UntaggedPolicyModificationError{
		errorLink: errorLink{id: UntaggedPolicyModificationErrorID}}
}

// Error implements error interface.
func (e *UntaggedPolicyModificationError) Error() string {
	return e.errorf("Can't modify policies with no tag")
}

// MissingPolicyTagError indicates that update has no tag to match policy before modification.
type MissingPolicyTagError struct {
	errorLink
}

func newMissingPolicyTagError() *MissingPolicyTagError {
	return &MissingPolicyTagError{
		errorLink: errorLink{id: MissingPolicyTagErrorID}}
}

// Error implements error interface.
func (e *MissingPolicyTagError) Error() string {
	return e.errorf("Update has no previous policy tag")
}

// PolicyTagsNotMatchError indicates that update tag doesn't match policy before modification.
type PolicyTagsNotMatchError struct {
	errorLink
	cntTag *uuid.UUID
	updTag *uuid.UUID
}

func newPolicyTagsNotMatchError(cntTag, updTag *uuid.UUID) *PolicyTagsNotMatchError {
	return &PolicyTagsNotMatchError{
		errorLink: errorLink{id: PolicyTagsNotMatchErrorID},
		cntTag:    cntTag,
		updTag:    updTag}
}

// Error implements error interface.
func (e *PolicyTagsNotMatchError) Error() string {
	return e.errorf("Update tag %s doesn't match policies tag %s", e.updTag.String(), e.cntTag.String())
}

type emptyPathModificationError struct {
	errorLink
}

func newEmptyPathModificationError() *emptyPathModificationError {
	return &emptyPathModificationError{
		errorLink: errorLink{id: emptyPathModificationErrorID}}
}

func (e *emptyPathModificationError) Error() string {
	return e.errorf("Can't modify items by empty path")
}

type invalidRootPolicyItemTypeError struct {
	errorLink
	item interface{}
}

func newInvalidRootPolicyItemTypeError(item interface{}) *invalidRootPolicyItemTypeError {
	return &invalidRootPolicyItemTypeError{
		errorLink: errorLink{id: invalidRootPolicyItemTypeErrorID},
		item:      item}
}

func (e *invalidRootPolicyItemTypeError) Error() string {
	return e.errorf("Expected policy or policy set as new root policy but got %T", e.item)
}

type hiddenRootPolicyAppendError struct {
	errorLink
}

func newHiddenRootPolicyAppendError() *hiddenRootPolicyAppendError {
	return &hiddenRootPolicyAppendError{
		errorLink: errorLink{id: hiddenRootPolicyAppendErrorID}}
}

func (e *hiddenRootPolicyAppendError) Error() string {
	return e.errorf("Can't append hidden policy to as root policy")
}

type invalidRootPolicyError struct {
	errorLink
	actual   string
	expected string
}

func newInvalidRootPolicyError(actual, expected string) *invalidRootPolicyError {
	return &invalidRootPolicyError{
		errorLink: errorLink{id: invalidRootPolicyErrorID},
		actual:    actual,
		expected:  expected}
}

func (e *invalidRootPolicyError) Error() string {
	return e.errorf("Root policy is %q but got %q as first path element", e.expected, e.actual)
}

type hiddenPolicySetModificationError struct {
	errorLink
}

func newHiddenPolicySetModificationError() *hiddenPolicySetModificationError {
	return &hiddenPolicySetModificationError{
		errorLink: errorLink{id: hiddenPolicySetModificationErrorID}}
}

func (e *hiddenPolicySetModificationError) Error() string {
	return e.errorf("Can't modify hidden policy set")
}

type invalidPolicySetItemTypeError struct {
	errorLink
	item interface{}
}

func newInvalidPolicySetItemTypeError(item interface{}) *invalidPolicySetItemTypeError {
	return &invalidPolicySetItemTypeError{
		errorLink: errorLink{id: invalidPolicySetItemTypeErrorID},
		item:      item}
}

func (e *invalidPolicySetItemTypeError) Error() string {
	return e.errorf("Expected policy or policy set to append but got %T", e.item)
}

type tooShortPathPolicySetModificationError struct {
	errorLink
}

func newTooShortPathPolicySetModificationError() *tooShortPathPolicySetModificationError {
	return &tooShortPathPolicySetModificationError{
		errorLink: errorLink{id: tooShortPathPolicySetModificationErrorID}}
}

func (e *tooShortPathPolicySetModificationError) Error() string {
	return e.errorf("Path to item to delete is too short")
}

type missingPolicySetChildError struct {
	errorLink
	ID string
}

func newMissingPolicySetChildError(ID string) *missingPolicySetChildError {
	return &missingPolicySetChildError{
		errorLink: errorLink{id: missingPolicySetChildErrorID},
		ID:        ID}
}

func (e *missingPolicySetChildError) Error() string {
	return e.errorf("Policy set has no child policy or policy set with id %q", e.ID)
}

type hiddenPolicyAppendError struct {
	errorLink
}

func newHiddenPolicyAppendError() *hiddenPolicyAppendError {
	return &hiddenPolicyAppendError{
		errorLink: errorLink{id: hiddenPolicyAppendErrorID}}
}

func (e *hiddenPolicyAppendError) Error() string {
	return e.errorf("Can't append hidden policy or policy set")
}

type policyTransactionTagsNotMatchError struct {
	errorLink
	tTag uuid.UUID
	uTag uuid.UUID
}

func newPolicyTransactionTagsNotMatchError(tTag, uTag uuid.UUID) *policyTransactionTagsNotMatchError {
	return &policyTransactionTagsNotMatchError{
		errorLink: errorLink{id: policyTransactionTagsNotMatchErrorID},
		tTag:      tTag,
		uTag:      uTag}
}

func (e *policyTransactionTagsNotMatchError) Error() string {
	return e.errorf("Update tag %s doesn't match policies transaction tag %s", e.uTag.String(), e.tTag.String())
}

type failedPolicyTransactionError struct {
	errorLink
	t   uuid.UUID
	err error
}

func newFailedPolicyTransactionError(t uuid.UUID, err error) *failedPolicyTransactionError {
	return &failedPolicyTransactionError{
		errorLink: errorLink{id: failedPolicyTransactionErrorID},
		t:         t,
		err:       err}
}

func (e *failedPolicyTransactionError) Error() string {
	return e.errorf("Can't operate with failed transaction on policies %s. Previous failure %s", e.t, e.err)
}

type unknownPolicyUpdateOperationError struct {
	errorLink
	op int
}

func newUnknownPolicyUpdateOperationError(op int) *unknownPolicyUpdateOperationError {
	return &unknownPolicyUpdateOperationError{
		errorLink: errorLink{id: unknownPolicyUpdateOperationErrorID},
		op:        op}
}

func (e *unknownPolicyUpdateOperationError) Error() string {
	return e.errorf("Unknown operation %d", e.op)
}

type hiddenPolicyModificationError struct {
	errorLink
}

func newHiddenPolicyModificationError() *hiddenPolicyModificationError {
	return &hiddenPolicyModificationError{
		errorLink: errorLink{id: hiddenPolicyModificationErrorID}}
}

func (e *hiddenPolicyModificationError) Error() string {
	return e.errorf("Can't modify hidden policy")
}

type tooLongPathPolicyModificationError struct {
	errorLink
	path []string
}

func newTooLongPathPolicyModificationError(path []string) *tooLongPathPolicyModificationError {
	return &tooLongPathPolicyModificationError{
		errorLink: errorLink{id: tooLongPathPolicyModificationErrorID},
		path:      path}
}

func (e *tooLongPathPolicyModificationError) Error() string {
	return e.errorf("Trailing path \"%s\"", strings.Join(e.path, "/"))
}

type tooShortPathPolicyModificationError struct {
	errorLink
}

func newTooShortPathPolicyModificationError() *tooShortPathPolicyModificationError {
	return &tooShortPathPolicyModificationError{
		errorLink: errorLink{id: tooShortPathPolicyModificationErrorID}}
}

func (e *tooShortPathPolicyModificationError) Error() string {
	return e.errorf("Path to item to delete is too short")
}

type invalidPolicyItemTypeError struct {
	errorLink
	item interface{}
}

func newInvalidPolicyItemTypeError(item interface{}) *invalidPolicyItemTypeError {
	return &invalidPolicyItemTypeError{
		errorLink: errorLink{id: invalidPolicyItemTypeErrorID},
		item:      item}
}

func (e *invalidPolicyItemTypeError) Error() string {
	return e.errorf("Expected rule to append but got %T", e.item)
}

type hiddenRuleAppendError struct {
	errorLink
}

func newHiddenRuleAppendError() *hiddenRuleAppendError {
	return &hiddenRuleAppendError{
		errorLink: errorLink{id: hiddenRuleAppendErrorID}}
}

func (e *hiddenRuleAppendError) Error() string {
	return e.errorf("Can't append hidden rule")
}

type missingPolicyChildError struct {
	errorLink
	ID string
}

func newMissingPolicyChildError(ID string) *missingPolicyChildError {
	return &missingPolicyChildError{
		errorLink: errorLink{id: missingPolicyChildErrorID},
		ID:        ID}
}

func (e *missingPolicyChildError) Error() string {
	return e.errorf("Policy has no rule with id %q", e.ID)
}

type missingContentError struct {
	errorLink
	ID string
}

func newMissingContentError(ID string) *missingContentError {
	return &missingContentError{
		errorLink: errorLink{id: missingContentErrorID},
		ID:        ID}
}

func (e *missingContentError) Error() string {
	return e.errorf("Missing content %s", e.ID)
}

type invalidContentStorageItem struct {
	errorLink
	ID string
	v  interface{}
}

func newInvalidContentStorageItem(ID string, v interface{}) *invalidContentStorageItem {
	return &invalidContentStorageItem{
		errorLink: errorLink{id: invalidContentStorageItemID},
		ID:        ID,
		v:         v}
}

func (e *invalidContentStorageItem) Error() string {
	return e.errorf("Invalid value at %s (expected *LocalContent but got %T)", e.ID, e.v)
}

type missingContentItemError struct {
	errorLink
	ID string
}

func newMissingContentItemError(ID string) *missingContentItemError {
	return &missingContentItemError{
		errorLink: errorLink{id: missingContentItemErrorID},
		ID:        ID}
}

func (e *missingContentItemError) Error() string {
	return e.errorf("Missing content item %q", e.ID)
}

type invalidContentItemError struct {
	errorLink
	v interface{}
}

func newInvalidContentItemError(v interface{}) *invalidContentItemError {
	return &invalidContentItemError{
		errorLink: errorLink{id: invalidContentItemErrorID},
		v:         v}
}

func (e *invalidContentItemError) Error() string {
	return e.errorf("Invalid value (expected *ContentItem but got %T)", e.v)
}

type invalidContentItemTypeError struct {
	errorLink
	expected int
	actual   int
}

func newInvalidContentItemTypeError(expected, actual int) *invalidContentItemTypeError {
	return &invalidContentItemTypeError{
		errorLink: errorLink{id: invalidContentItemTypeErrorID},
		expected:  expected,
		actual:    actual}
}

func (e *invalidContentItemTypeError) Error() string {
	return e.errorf("Invalid conent item type. Expected %q but got %q", TypeNames[e.expected], TypeNames[e.actual])
}

type invalidSelectorPathError struct {
	errorLink
	expected []int
	actual   []Expression
}

func newInvalidSelectorPathError(expected []int, actual []Expression) *invalidSelectorPathError {
	return &invalidSelectorPathError{
		errorLink: errorLink{id: invalidSelectorPathErrorID},
		expected:  expected,
		actual:    actual}
}

func (e *invalidSelectorPathError) Error() string {
	expStrs := make([]string, len(e.expected))
	for i, t := range e.expected {
		expStrs[i] = TypeNames[t]
	}
	expected := strings.Join(expStrs, "/")

	actual := "nothing"
	if len(e.actual) > 0 {
		strs := make([]string, len(e.actual))
		for i, e := range e.actual {
			strs[i] = TypeNames[e.GetResultType()]
		}
		actual = strings.Join(strs, "/")
	}

	return e.errorf("Invalid selector path. Expected %s but got %s", expected, actual)
}

type networkMapKeyValueTypeError struct {
	errorLink
	t int
}

func newNetworkMapKeyValueTypeError(t int) *networkMapKeyValueTypeError {
	return &networkMapKeyValueTypeError{
		errorLink: errorLink{id: networkMapKeyValueTypeErrorID},
		t:         t}
}

func (e *networkMapKeyValueTypeError) Error() string {
	return e.errorf("Expected %s or %s as network map key but got %s", TypeNames[TypeAddress], TypeNames[TypeNetwork], TypeNames[e.t])
}

type mapContentSubitemError struct {
	errorLink
}

func newMapContentSubitemError() *mapContentSubitemError {
	return &mapContentSubitemError{
		errorLink: errorLink{id: mapContentSubitemErrorID}}
}

func (e *mapContentSubitemError) Error() string {
	return e.errorf("Not a map of the content")
}

type invalidContentModificationError struct {
	errorLink
}

func newInvalidContentModificationError() *invalidContentModificationError {
	return &invalidContentModificationError{
		errorLink: errorLink{id: invalidContentModificationErrorID}}
}

func (e *invalidContentModificationError) Error() string {
	return e.errorf("Can't modify non-mapping content item")
}

type missingPathContentModificationError struct {
	errorLink
}

func newMissingPathContentModificationError() *missingPathContentModificationError {
	return &missingPathContentModificationError{
		errorLink: errorLink{id: missingPathContentModificationErrorID}}
}

func (e *missingPathContentModificationError) Error() string {
	return e.errorf("Missing path for content item change")
}

type tooLongPathContentModificationError struct {
	errorLink
	expected []int
	actual   []AttributeValue
}

func newTooLongPathContentModificationError(expected []int, actual []AttributeValue) *tooLongPathContentModificationError {
	return &tooLongPathContentModificationError{
		errorLink: errorLink{id: tooLongPathContentModificationErrorID},
		expected:  expected,
		actual:    actual}
}

func (e *tooLongPathContentModificationError) Error() string {
	expected := "no"
	if len(e.expected) > 0 {
		expStrs := make([]string, len(e.expected))
		for i, t := range e.expected {
			expStrs[i] = fmt.Sprintf("%q", TypeNames[t])
		}
		expected = strings.Join(expStrs, "/")
	}

	actStrs := make([]string, len(e.actual))
	for i, e := range e.actual {
		actStrs[i] = fmt.Sprintf("%q", TypeNames[e.GetResultType()])
	}
	actual := strings.Join(actStrs, "/")

	return e.errorf("Too long modification path. Expected %s path but got %s", expected, actual)
}

type invalidContentValueModificationError struct {
	errorLink
}

func newInvalidContentValueModificationError() *invalidContentValueModificationError {
	return &invalidContentValueModificationError{
		errorLink: errorLink{id: invalidContentValueModificationErrorID}}
}

func (e *invalidContentValueModificationError) Error() string {
	return e.errorf("Can't modify final content value")
}

// UntaggedContentModificationError indicates attempt to modify incrementally a content which has no tag.
type UntaggedContentModificationError struct {
	errorLink
	ID string
}

func newUntaggedContentModificationError(ID string) *UntaggedContentModificationError {
	return &UntaggedContentModificationError{
		errorLink: errorLink{id: UntaggedContentModificationErrorID},
		ID:        ID}
}

// Error implements error interface.
func (e *UntaggedContentModificationError) Error() string {
	return e.errorf("Can't modify content %q with no tag", e.ID)
}

// MissingContentTagError indicates that update has no tag to match content before modification.
type MissingContentTagError struct {
	errorLink
}

func newMissingContentTagError() *MissingContentTagError {
	return &MissingContentTagError{
		errorLink: errorLink{id: MissingContentTagErrorID}}
}

// Error implements error interface.
func (e *MissingContentTagError) Error() string {
	return e.errorf("Update has no previous content tag")
}

// ContentTagsNotMatchError indicates that update tag doesn't match content before modification.
type ContentTagsNotMatchError struct {
	errorLink
	ID     string
	cntTag *uuid.UUID
	updTag *uuid.UUID
}

func newContentTagsNotMatchError(ID string, cntTag, updTag *uuid.UUID) *ContentTagsNotMatchError {
	return &ContentTagsNotMatchError{
		errorLink: errorLink{id: ContentTagsNotMatchErrorID},
		ID:        ID,
		cntTag:    cntTag,
		updTag:    updTag}
}

// Error implements error interface.
func (e *ContentTagsNotMatchError) Error() string {
	return e.errorf("Update tag %s doesn't match content %q tag %s", e.cntTag.String(), e.ID, e.updTag.String())
}

type unknownContentUpdateOperationError struct {
	errorLink
	op int
}

func newUnknownContentUpdateOperationError(op int) *unknownContentUpdateOperationError {
	return &unknownContentUpdateOperationError{
		errorLink: errorLink{id: unknownContentUpdateOperationErrorID},
		op:        op}
}

func (e *unknownContentUpdateOperationError) Error() string {
	return e.errorf("Unknown operation %d", e.op)
}

type failedContentTransactionError struct {
	errorLink
	id  string
	t   uuid.UUID
	err error
}

func newFailedContentTransactionError(id string, t uuid.UUID, err error) *failedContentTransactionError {
	return &failedContentTransactionError{
		errorLink: errorLink{id: failedContentTransactionErrorID},
		id:        id,
		t:         t,
		err:       err}
}

func (e *failedContentTransactionError) Error() string {
	return e.errorf("Can't operate with failed transaction on content %q tagged %s. Previous failure %s", e.id, e.t, e.err)
}

type contentTransactionIDNotMatchError struct {
	errorLink
	tID string
	uID string
}

func newContentTransactionIDNotMatchError(tID, uID string) *contentTransactionIDNotMatchError {
	return &contentTransactionIDNotMatchError{
		errorLink: errorLink{id: contentTransactionIDNotMatchErrorID},
		tID:       tID,
		uID:       uID}
}

func (e *contentTransactionIDNotMatchError) Error() string {
	return e.errorf("Update content ID %q doesn't match transaction content ID %q", e.uID, e.tID)
}

type contentTransactionTagsNotMatchError struct {
	errorLink
	id   string
	tTag uuid.UUID
	uTag uuid.UUID
}

func newContentTransactionTagsNotMatchError(id string, tTag, uTag uuid.UUID) *contentTransactionTagsNotMatchError {
	return &contentTransactionTagsNotMatchError{
		errorLink: errorLink{id: contentTransactionTagsNotMatchErrorID},
		id:        id,
		tTag:      tTag,
		uTag:      uTag}
}

func (e *contentTransactionTagsNotMatchError) Error() string {
	return e.errorf("Update tag %s doesn't match content %q transaction tag %s", e.uTag.String(), e.id, e.tTag.String())
}

type tooShortRawPathContentModificationError struct {
	errorLink
}

func newTooShortRawPathContentModificationError() *tooShortRawPathContentModificationError {
	return &tooShortRawPathContentModificationError{
		errorLink: errorLink{id: tooShortRawPathContentModificationErrorID}}
}

func (e *tooShortRawPathContentModificationError) Error() string {
	return e.errorf("Expected at least content item ID in path but got nothing")
}

type tooLongRawPathContentModificationError struct {
	errorLink
	expected []int
	actual   []string
}

func newTooLongRawPathContentModificationError(expected []int, actual []string) *tooLongRawPathContentModificationError {
	return &tooLongRawPathContentModificationError{
		errorLink: errorLink{id: tooLongRawPathContentModificationErrorID},
		expected:  expected,
		actual:    actual}
}

func (e *tooLongRawPathContentModificationError) Error() string {
	expected := "no"
	if len(e.expected) > 0 {
		expStrs := make([]string, len(e.expected))
		for i, t := range e.expected {
			expStrs[i] = fmt.Sprintf("%q", TypeNames[t])
		}
		expected = strings.Join(expStrs, "/")
	}

	actStrs := make([]string, len(e.actual))
	for i, s := range e.actual {
		actStrs[i] = fmt.Sprintf("%q", s)
	}
	actual := strings.Join(actStrs, "/")

	return e.errorf("Too long modification path. Expected %s path but got %s", expected, actual)
}

type invalidContentUpdateDataError struct {
	errorLink
	v interface{}
}

func newInvalidContentUpdateDataError(v interface{}) *invalidContentUpdateDataError {
	return &invalidContentUpdateDataError{
		errorLink: errorLink{id: invalidContentUpdateDataErrorID},
		v:         v}
}

func (e *invalidContentUpdateDataError) Error() string {
	return e.errorf("Expected content update data but got %T", e.v)
}

type invalidContentUpdateResultTypeError struct {
	errorLink
	actual   int
	expected int
}

func newInvalidContentUpdateResultTypeError(actual, expected int) *invalidContentUpdateResultTypeError {
	return &invalidContentUpdateResultTypeError{
		errorLink: errorLink{id: invalidContentUpdateResultTypeErrorID},
		actual:    actual,
		expected:  expected}
}

func (e *invalidContentUpdateResultTypeError) Error() string {
	return e.errorf("Expected %q as a result type but got %q", TypeNames[e.expected], TypeNames[e.actual])
}

type invalidContentUpdateKeysError struct {
	errorLink
	start    int
	actual   []int
	expected []int
}

func newInvalidContentUpdateKeysError(start int, actual, expected []int) *invalidContentUpdateKeysError {
	return &invalidContentUpdateKeysError{
		errorLink: errorLink{id: invalidContentUpdateKeysErrorID},
		start:     start,
		actual:    actual,
		expected:  expected}
}

func (e *invalidContentUpdateKeysError) Error() string {
	enames := make([]string, len(e.expected)-e.start)
	for i, t := range e.expected[e.start:] {
		enames[i] = fmt.Sprintf("%q", TypeNames[t])
	}
	expected := strings.Join(enames, "/")

	actual := "nothing"
	if len(e.actual) > 0 {
		anames := make([]string, len(e.actual))
		for i, t := range e.actual {
			anames[i] = fmt.Sprintf("%q", TypeNames[t])
		}
		actual = strings.Join(anames, "/")
	}

	return e.errorf("Expected %s path after position %d but got %s", expected, e.start, actual)
}

type unknownContentItemResultTypeError struct {
	errorLink
	t int
}

func newUnknownContentItemResultTypeError(t int) *unknownContentItemResultTypeError {
	return &unknownContentItemResultTypeError{
		errorLink: errorLink{id: unknownContentItemResultTypeErrorID},
		t:         t}
}

func (e *unknownContentItemResultTypeError) Error() string {
	return e.errorf("Unknown result type for content item: %d", e.t)
}

type invalidContentItemResultTypeError struct {
	errorLink
	t int
}

func newInvalidContentItemResultTypeError(t int) *invalidContentItemResultTypeError {
	return &invalidContentItemResultTypeError{
		errorLink: errorLink{id: invalidContentItemResultTypeErrorID},
		t:         t}
}

func (e *invalidContentItemResultTypeError) Error() string {
	return e.errorf("Invalid result type for content item: %s", TypeNames[e.t])
}

type invalidContentKeyTypeError struct {
	errorLink
	t        int
	expected map[int]bool
}

func newInvalidContentKeyTypeError(t int, expected map[int]bool) *invalidContentKeyTypeError {
	return &invalidContentKeyTypeError{
		errorLink: errorLink{id: invalidContentKeyTypeErrorID},
		t:         t,
		expected:  expected}
}

func (e *invalidContentKeyTypeError) Error() string {
	names := make([]string, len(e.expected))
	i := 0
	for t := range e.expected {
		names[i] = TypeNames[t]
		i++
	}
	s := strings.Join(names, ", ")

	return e.errorf("Invalid key type for content item: %s (expected %s)", TypeNames[e.t], s)
}

type invalidContentStringMapError struct {
	errorLink
	value interface{}
}

func newInvalidContentStringMapError(value interface{}) *invalidContentStringMapError {
	return &invalidContentStringMapError{
		errorLink: errorLink{id: invalidContentStringMapErrorID},
		value:     value}
}

func (e *invalidContentStringMapError) Error() string {
	return e.errorf("Expected string map but got %T", e.value)
}

type invalidContentNetworkMapError struct {
	errorLink
	value interface{}
}

func newInvalidContentNetworkMapError(value interface{}) *invalidContentNetworkMapError {
	return &invalidContentNetworkMapError{
		errorLink: errorLink{id: invalidContentNetworkMapErrorID},
		value:     value}
}

func (e *invalidContentNetworkMapError) Error() string {
	return e.errorf("Expected network map but got %T", e.value)
}

type invalidContentDomainMapError struct {
	errorLink
	value interface{}
}

func newInvalidContentDomainMapError(value interface{}) *invalidContentDomainMapError {
	return &invalidContentDomainMapError{
		errorLink: errorLink{id: invalidContentDomainMapErrorID},
		value:     value}
}

func (e *invalidContentDomainMapError) Error() string {
	return e.errorf("Expected domain map but got %T", e.value)
}

type invalidContentValueError struct {
	errorLink
	value interface{}
}

func newInvalidContentValueError(value interface{}) *invalidContentValueError {
	return &invalidContentValueError{
		errorLink: errorLink{id: invalidContentValueErrorID},
		value:     value}
}

func (e *invalidContentValueError) Error() string {
	return e.errorf("Expected value but got %T", e.value)
}

type invalidContentValueTypeError struct {
	errorLink
	value    interface{}
	expected int
}

func newInvalidContentValueTypeError(value interface{}, expected int) *invalidContentValueTypeError {
	return &invalidContentValueTypeError{
		errorLink: errorLink{id: invalidContentValueTypeErrorID},
		value:     value,
		expected:  expected}
}

func (e *invalidContentValueTypeError) Error() string {
	return e.errorf("Expected value of type %s but got %T", TypeNames[e.expected], e.value)
}
