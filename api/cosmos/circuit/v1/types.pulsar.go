// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package circuitv1

import (
	fmt "fmt"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
	reflect "reflect"
	sync "sync"
)

var _ protoreflect.List = (*_Permissions_2_list)(nil)

type _Permissions_2_list struct {
	list *[]string
}

func (x *_Permissions_2_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_Permissions_2_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfString((*x.list)[i])
}

func (x *_Permissions_2_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.String()
	concreteValue := valueUnwrapped
	(*x.list)[i] = concreteValue
}

func (x *_Permissions_2_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.String()
	concreteValue := valueUnwrapped
	*x.list = append(*x.list, concreteValue)
}

func (x *_Permissions_2_list) AppendMutable() protoreflect.Value {
	panic(fmt.Errorf("AppendMutable can not be called on message Permissions at list field LimitTypeUrls as it is not of Message kind"))
}

func (x *_Permissions_2_list) Truncate(n int) {
	*x.list = (*x.list)[:n]
}

func (x *_Permissions_2_list) NewElement() protoreflect.Value {
	v := ""
	return protoreflect.ValueOfString(v)
}

func (x *_Permissions_2_list) IsValid() bool {
	return x.list != nil
}

var (
	md_Permissions                 protoreflect.MessageDescriptor
	fd_Permissions_level           protoreflect.FieldDescriptor
	fd_Permissions_limit_type_urls protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_circuit_v1_types_proto_init()
	md_Permissions = File_cosmos_circuit_v1_types_proto.Messages().ByName("Permissions")
	fd_Permissions_level = md_Permissions.Fields().ByName("level")
	fd_Permissions_limit_type_urls = md_Permissions.Fields().ByName("limit_type_urls")
}

var _ protoreflect.Message = (*fastReflection_Permissions)(nil)

type fastReflection_Permissions Permissions

func (x *Permissions) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Permissions)(x)
}

func (x *Permissions) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_circuit_v1_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Permissions_messageType fastReflection_Permissions_messageType
var _ protoreflect.MessageType = fastReflection_Permissions_messageType{}

type fastReflection_Permissions_messageType struct{}

func (x fastReflection_Permissions_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Permissions)(nil)
}
func (x fastReflection_Permissions_messageType) New() protoreflect.Message {
	return new(fastReflection_Permissions)
}
func (x fastReflection_Permissions_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Permissions
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Permissions) Descriptor() protoreflect.MessageDescriptor {
	return md_Permissions
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Permissions) Type() protoreflect.MessageType {
	return _fastReflection_Permissions_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Permissions) New() protoreflect.Message {
	return new(fastReflection_Permissions)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Permissions) Interface() protoreflect.ProtoMessage {
	return (*Permissions)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Permissions) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Level != 0 {
		value := protoreflect.ValueOfEnum((protoreflect.EnumNumber)(x.Level))
		if !f(fd_Permissions_level, value) {
			return
		}
	}
	if len(x.LimitTypeUrls) != 0 {
		value := protoreflect.ValueOfList(&_Permissions_2_list{list: &x.LimitTypeUrls})
		if !f(fd_Permissions_limit_type_urls, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_Permissions) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.circuit.v1.Permissions.level":
		return x.Level != 0
	case "cosmos.circuit.v1.Permissions.limit_type_urls":
		return len(x.LimitTypeUrls) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.circuit.v1.Permissions"))
		}
		panic(fmt.Errorf("message cosmos.circuit.v1.Permissions does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Permissions) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.circuit.v1.Permissions.level":
		x.Level = 0
	case "cosmos.circuit.v1.Permissions.limit_type_urls":
		x.LimitTypeUrls = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.circuit.v1.Permissions"))
		}
		panic(fmt.Errorf("message cosmos.circuit.v1.Permissions does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Permissions) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.circuit.v1.Permissions.level":
		value := x.Level
		return protoreflect.ValueOfEnum((protoreflect.EnumNumber)(value))
	case "cosmos.circuit.v1.Permissions.limit_type_urls":
		if len(x.LimitTypeUrls) == 0 {
			return protoreflect.ValueOfList(&_Permissions_2_list{})
		}
		listValue := &_Permissions_2_list{list: &x.LimitTypeUrls}
		return protoreflect.ValueOfList(listValue)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.circuit.v1.Permissions"))
		}
		panic(fmt.Errorf("message cosmos.circuit.v1.Permissions does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Permissions) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.circuit.v1.Permissions.level":
		x.Level = (Permissions_Level)(value.Enum())
	case "cosmos.circuit.v1.Permissions.limit_type_urls":
		lv := value.List()
		clv := lv.(*_Permissions_2_list)
		x.LimitTypeUrls = *clv.list
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.circuit.v1.Permissions"))
		}
		panic(fmt.Errorf("message cosmos.circuit.v1.Permissions does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Permissions) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.circuit.v1.Permissions.limit_type_urls":
		if x.LimitTypeUrls == nil {
			x.LimitTypeUrls = []string{}
		}
		value := &_Permissions_2_list{list: &x.LimitTypeUrls}
		return protoreflect.ValueOfList(value)
	case "cosmos.circuit.v1.Permissions.level":
		panic(fmt.Errorf("field level of message cosmos.circuit.v1.Permissions is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.circuit.v1.Permissions"))
		}
		panic(fmt.Errorf("message cosmos.circuit.v1.Permissions does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Permissions) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.circuit.v1.Permissions.level":
		return protoreflect.ValueOfEnum(0)
	case "cosmos.circuit.v1.Permissions.limit_type_urls":
		list := []string{}
		return protoreflect.ValueOfList(&_Permissions_2_list{list: &list})
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.circuit.v1.Permissions"))
		}
		panic(fmt.Errorf("message cosmos.circuit.v1.Permissions does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Permissions) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.circuit.v1.Permissions", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Permissions) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Permissions) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_Permissions) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Permissions) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Permissions)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.Level != 0 {
			n += 1 + runtime.Sov(uint64(x.Level))
		}
		if len(x.LimitTypeUrls) > 0 {
			for _, s := range x.LimitTypeUrls {
				l = len(s)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*Permissions)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.LimitTypeUrls) > 0 {
			for iNdEx := len(x.LimitTypeUrls) - 1; iNdEx >= 0; iNdEx-- {
				i -= len(x.LimitTypeUrls[iNdEx])
				copy(dAtA[i:], x.LimitTypeUrls[iNdEx])
				i = runtime.EncodeVarint(dAtA, i, uint64(len(x.LimitTypeUrls[iNdEx])))
				i--
				dAtA[i] = 0x12
			}
		}
		if x.Level != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Level))
			i--
			dAtA[i] = 0x8
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*Permissions)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Permissions: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Permissions: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Level", wireType)
				}
				x.Level = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Level |= Permissions_Level(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field LimitTypeUrls", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.LimitTypeUrls = append(x.LimitTypeUrls, string(dAtA[iNdEx:postIndex]))
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: cosmos/circuit/v1/types.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Level is the permission level.
type Permissions_Level int32

const (
	// LEVEL_NONE_UNSPECIFIED indicates that the account will have no circuit
	// breaker permissions.
	Permissions_LEVEL_NONE_UNSPECIFIED Permissions_Level = 0
	// LEVEL_SOME_MSGS indicates that the account will have permission to
	// trip or reset the circuit breaker for some Msg type URLs. If this level
	// is chosen, a non-empty list of Msg type URLs must be provided in
	// limit_type_urls.
	Permissions_LEVEL_SOME_MSGS Permissions_Level = 1
	// LEVEL_ALL_MSGS indicates that the account can trip or reset the circuit
	// breaker for Msg's of all type URLs.
	Permissions_LEVEL_ALL_MSGS Permissions_Level = 2
	// LEVEL_SUPER_ADMIN indicates that the account can take all circuit breaker
	// actions and can grant permissions to other accounts.
	Permissions_LEVEL_SUPER_ADMIN Permissions_Level = 3
)

// Enum value maps for Permissions_Level.
var (
	Permissions_Level_name = map[int32]string{
		0: "LEVEL_NONE_UNSPECIFIED",
		1: "LEVEL_SOME_MSGS",
		2: "LEVEL_ALL_MSGS",
		3: "LEVEL_SUPER_ADMIN",
	}
	Permissions_Level_value = map[string]int32{
		"LEVEL_NONE_UNSPECIFIED": 0,
		"LEVEL_SOME_MSGS":        1,
		"LEVEL_ALL_MSGS":         2,
		"LEVEL_SUPER_ADMIN":      3,
	}
)

func (x Permissions_Level) Enum() *Permissions_Level {
	p := new(Permissions_Level)
	*p = x
	return p
}

func (x Permissions_Level) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Permissions_Level) Descriptor() protoreflect.EnumDescriptor {
	return file_cosmos_circuit_v1_types_proto_enumTypes[0].Descriptor()
}

func (Permissions_Level) Type() protoreflect.EnumType {
	return &file_cosmos_circuit_v1_types_proto_enumTypes[0]
}

func (x Permissions_Level) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Permissions_Level.Descriptor instead.
func (Permissions_Level) EnumDescriptor() ([]byte, []int) {
	return file_cosmos_circuit_v1_types_proto_rawDescGZIP(), []int{0, 0}
}

// Permissions are the permissions that an account has to trip
// or reset the circuit breaker.
type Permissions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// level is the level of permissions granted to this account.
	Level Permissions_Level `protobuf:"varint,1,opt,name=level,proto3,enum=cosmos.circuit.v1.Permissions_Level" json:"level,omitempty"`
	// limit_type_urls is used with LEVEL_SOME_MSGS to limit the lists of Msg type
	// URLs that the account can pause. It is an error to use limit_type_urls with
	// a level other than LEVEL_SOME_MSGS.
	LimitTypeUrls []string `protobuf:"bytes,2,rep,name=limit_type_urls,json=limitTypeUrls,proto3" json:"limit_type_urls,omitempty"`
}

func (x *Permissions) Reset() {
	*x = Permissions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_circuit_v1_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Permissions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Permissions) ProtoMessage() {}

// Deprecated: Use Permissions.ProtoReflect.Descriptor instead.
func (*Permissions) Descriptor() ([]byte, []int) {
	return file_cosmos_circuit_v1_types_proto_rawDescGZIP(), []int{0}
}

func (x *Permissions) GetLevel() Permissions_Level {
	if x != nil {
		return x.Level
	}
	return Permissions_LEVEL_NONE_UNSPECIFIED
}

func (x *Permissions) GetLimitTypeUrls() []string {
	if x != nil {
		return x.LimitTypeUrls
	}
	return nil
}

var File_cosmos_circuit_v1_types_proto protoreflect.FileDescriptor

var file_cosmos_circuit_v1_types_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x69, 0x72, 0x63, 0x75, 0x69, 0x74,
	0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x11, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x63, 0x69, 0x72, 0x63, 0x75, 0x69, 0x74, 0x2e,
	0x76, 0x31, 0x22, 0xd6, 0x01, 0x0a, 0x0b, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x3a, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x24, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x63, 0x69, 0x72, 0x63, 0x75,
	0x69, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x26,
	0x0a, 0x0f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x75, 0x72, 0x6c,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x55, 0x72, 0x6c, 0x73, 0x22, 0x63, 0x0a, 0x05, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12,
	0x1a, 0x0a, 0x16, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x4c,
	0x45, 0x56, 0x45, 0x4c, 0x5f, 0x53, 0x4f, 0x4d, 0x45, 0x5f, 0x4d, 0x53, 0x47, 0x53, 0x10, 0x01,
	0x12, 0x12, 0x0a, 0x0e, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x41, 0x4c, 0x4c, 0x5f, 0x4d, 0x53,
	0x47, 0x53, 0x10, 0x02, 0x12, 0x15, 0x0a, 0x11, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x53, 0x55,
	0x50, 0x45, 0x52, 0x5f, 0x41, 0x44, 0x4d, 0x49, 0x4e, 0x10, 0x03, 0x42, 0xb7, 0x01, 0x0a, 0x15,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x63, 0x69, 0x72, 0x63, 0x75,
	0x69, 0x74, 0x2e, 0x76, 0x31, 0x42, 0x0a, 0x54, 0x79, 0x70, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x2c, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x73, 0x64, 0x6b, 0x2e, 0x69,
	0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x69, 0x72,
	0x63, 0x75, 0x69, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x69, 0x72, 0x63, 0x75, 0x69, 0x74, 0x76,
	0x31, 0xa2, 0x02, 0x03, 0x43, 0x43, 0x58, 0xaa, 0x02, 0x11, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2e, 0x43, 0x69, 0x72, 0x63, 0x75, 0x69, 0x74, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x11, 0x43, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x5c, 0x43, 0x69, 0x72, 0x63, 0x75, 0x69, 0x74, 0x5c, 0x56, 0x31, 0xe2,
	0x02, 0x1d, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5c, 0x43, 0x69, 0x72, 0x63, 0x75, 0x69, 0x74,
	0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x13, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x3a, 0x3a, 0x43, 0x69, 0x72, 0x63, 0x75, 0x69,
	0x74, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cosmos_circuit_v1_types_proto_rawDescOnce sync.Once
	file_cosmos_circuit_v1_types_proto_rawDescData = file_cosmos_circuit_v1_types_proto_rawDesc
)

func file_cosmos_circuit_v1_types_proto_rawDescGZIP() []byte {
	file_cosmos_circuit_v1_types_proto_rawDescOnce.Do(func() {
		file_cosmos_circuit_v1_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_cosmos_circuit_v1_types_proto_rawDescData)
	})
	return file_cosmos_circuit_v1_types_proto_rawDescData
}

var file_cosmos_circuit_v1_types_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cosmos_circuit_v1_types_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_cosmos_circuit_v1_types_proto_goTypes = []interface{}{
	(Permissions_Level)(0), // 0: cosmos.circuit.v1.Permissions.Level
	(*Permissions)(nil),    // 1: cosmos.circuit.v1.Permissions
}
var file_cosmos_circuit_v1_types_proto_depIdxs = []int32{
	0, // 0: cosmos.circuit.v1.Permissions.level:type_name -> cosmos.circuit.v1.Permissions.Level
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_cosmos_circuit_v1_types_proto_init() }
func file_cosmos_circuit_v1_types_proto_init() {
	if File_cosmos_circuit_v1_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cosmos_circuit_v1_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Permissions); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cosmos_circuit_v1_types_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cosmos_circuit_v1_types_proto_goTypes,
		DependencyIndexes: file_cosmos_circuit_v1_types_proto_depIdxs,
		EnumInfos:         file_cosmos_circuit_v1_types_proto_enumTypes,
		MessageInfos:      file_cosmos_circuit_v1_types_proto_msgTypes,
	}.Build()
	File_cosmos_circuit_v1_types_proto = out.File
	file_cosmos_circuit_v1_types_proto_rawDesc = nil
	file_cosmos_circuit_v1_types_proto_goTypes = nil
	file_cosmos_circuit_v1_types_proto_depIdxs = nil
}
