

package main




import  "fmt" 
import "math"
import "github.com/ehteshamz/greetings"




func ConcatSlice(sliceToConcat []byte) string {

	myString := "cancatinating "

	for i := 0 ; i < len(sliceToConcat) ; i++ {

	

		myString += string(sliceToConcat[i])
		myString += "-"
	}

	myString += string(sliceToConcat[len(sliceToConcat)-1])


	return myString

}


func Encrypt(sliceToEncrypt []byte, ceaserCount int) {

	myString := ""

	for i := 0 ; i < len(sliceToEncrypt) ; i++ {

	

		myString += string(sliceToEncrypt[i])
	}


	tempString := ""

	for i := 0 ; i < len(myString) ; i++ {

		ascii_value := int(myString[i])
		ascii_value = ascii_value - 97
		simplified_Number := float64( ascii_value + ceaserCount)
		tempString += string(int(math.Mod(simplified_Number, 26)) +97)



	}

	for i := 0 ; i < len(sliceToEncrypt) ; i++ {

	

		sliceToEncrypt[i] += tempString[i]
	}


	fmt.Println("Encrypted String is :   " , sliceToEncrypt )
	//fmt.Println(sliceToEncrypt)

	//return tempString



}



func EZGreetings(name string) string{
	return (greetings.PrintGreetings(name))

}






func main(){

	byteArray := []byte{'J', 'O', 'H', 'N'}

	Encrypt(byteArray, 10)

}







