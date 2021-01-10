package main

import "fmt"

func main() {
	helloWorld()
	dataTypeBool()
	dataTypeString()
	dataTypeNumber()
	assignVariable()
	constants()
	typeDeclaration()
}

func helloWorld() {
	fmt.Println("Hello World")
}

func dataTypeBool() {
	var trueBool bool
	var falseBool bool
	trueBool = true
	falseBool = false
	fmt.Println("True", trueBool)
	fmt.Println("False", falseBool)
}

func dataTypeString() {
	var name string
	name = "Monki"
	fmt.Println("Name", name)
	// Print String Length
	fmt.Println("Length", len(name))
	// Print Character Byte Value At Index 0
	fmt.Println("Character Byte Value Index 0:", name[0])
	// Print Character At Index 0
	nameChar := string(name[0])
	fmt.Println("Character Index 0:", nameChar)
}

func dataTypeNumber() {
	// * Integer
	// 8-bit signed integer
	var numberInt8 int8
	// 16-bit signed integer
	var numberInt16 int16
	// 32-bit signed integer
	var numberInt32 int32
	// 64-bit signed integer
	var numberInt64 int64
	// 8-bit unsigned integer
	var numberUInt8 uint8
	// 16-bit unsigned integer
	var numberUInt16 uint16
	// 32-bit unsigned integer
	var numberUInt32 uint32
	// 64-bit unsigned integer
	var numberUInt64 uint64

	// * Float
	// 32-bit float
	var numberFloat32 float32
	// 64-bit float
	var numberFloat64 float64
	// complex number with float32 real and imaginary parts
	var numberComplex64 complex64
	// complex number with float64 real and imaginary parts
	var numberComplex128 complex128

	// * Alias
	// byte = uint8
	var numberByte byte
	// rune = int32
	var numberRune rune
	// int = minimal int32 (based on the OS)
	var numberInt int
	// uint = minimal uint32 (based on the OS)
	var numberUInt uint

	// Assign Variable
	numberInt8 = 127
	numberInt16 = 32767
	numberInt32 = 2147483647
	numberInt64 = 9223372036854775807
	numberUInt8 = 255
	numberUInt16 = 65535
	numberUInt32 = 4294967295
	numberUInt64 = 18446744073709551615
	numberFloat32 = 3.4
	numberFloat64 = 1.8
	numberComplex64 = 3.4
	numberComplex128 = 1.8
	numberByte = 255
	numberRune = 2147483647
	numberInt = 9223372036854775807
	numberUInt = 18446744073709551615

	// Print Value
	fmt.Println("Int8", numberInt8)
	fmt.Println("Int16", numberInt16)
	fmt.Println("Int32", numberInt32)
	fmt.Println("Int64", numberInt64)
	fmt.Println("UInt8", numberUInt8)
	fmt.Println("UInt16", numberUInt16)
	fmt.Println("UInt32", numberUInt32)
	fmt.Println("UInt64", numberUInt64)
	fmt.Println("Float32", numberFloat32)
	fmt.Println("Float64", numberFloat64)
	fmt.Println("Complex64", numberComplex64)
	fmt.Println("Complex128", numberComplex128)
	fmt.Println("Byte", numberByte)
	fmt.Println("Rune", numberRune)
	fmt.Println("Int", numberInt)
	fmt.Println("UInt", numberUInt)
}

func assignVariable() {
	// Auto Assign As String
	var anotherName = "Darted"
	fmt.Println(anotherName)

	// Auto Assign As Int
	var anotherNumber = 22
	fmt.Println(anotherNumber)

	// Assign Value and DataType
	var anotherNumberInt8 int8 = 22
	fmt.Println(anotherNumberInt8)

	// Assign Variable Without "Var"
	// ! Only in first declaration
	withoutVar := 22
	fmt.Println(withoutVar)
	withoutVarInt8 := int8(22)
	fmt.Println(withoutVarInt8)

	// Assign Multiple Variable
	var (
		firstName = "Darted"
		lastName  = "Monki"
	)
	boolA, boolB, stringA := true, false, "no!"
	var i, j int = 1, 2

	fmt.Println(firstName, lastName)
	fmt.Println(boolA, boolB, stringA)
	fmt.Println(i, j)
}

func constants() {
	const constantNumber = 99
	const (
		myName           = "DartedMonki"
		yourName         = "Good"
		valueInt16 int16 = 22
		valueInt8        = int8(22)
	)
	fmt.Println(myName, yourName, valueInt16, valueInt8)
}

func typeDeclaration() {
	type kata string
	var name kata = "Darted"
	fmt.Println(name, kata("Monki"))
}
