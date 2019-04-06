

package main




import  "fmt" 
import "math"
//import "strings"
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

		tempString += string(int(math.Mod(float64(int(myString[i]) - 97 + ceaserCount), 26)) +97)



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







