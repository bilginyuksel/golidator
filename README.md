# gorify

Gorify is a validation library. Gorify allows you to define constraint tag's to your structs. After you write your constraints then you can use `Validate` method and easily verify your struct. There are a lot already defined tags out there for any kind of field but if you have a custom scenario in mind you can define your custom constraint very easy without breaking or overriding already defined constraints.  

### Example Usage

```go
import (
    "fmt"

    "github.com/bilginyuksel/gorify"
)


type TestStruct struct {
	Email      string    `blank:"false" pattern:"^[\\w-\\.]+@([\\w-]+\\.)+[\\w-]{2,4}$" between:"10-40"`
	Username   string    `size:"12"`
	Firstname  string    `default:"yuksel"`
	Age        int       `between:"20-50"`
	ExpireTime time.Time `default:"now,add-9h,local,utc,round-1000ms"`
}

func main() {
    test := &TestStruct{
        Email: "test.test@gmail.com",
        Username: "bilginyuksel",
        Age: 25,
    }

    err := gorify.Validate(test)

    if err != nil {
        panic(err)
    }

    fmt.Println(test)
}
```

Like I have mentioned earlier you can easily define custom tags for your fields. Let's see a custom constraint example below.

__Explanation:__ In the example below we create a lengthEqual method as you can see in the method we have two parameters first one is value and the second one is the tags. You can check the tag you wish to control via tags then you can make your controls over value. The same method signature (just the value type is changing) are defined for known types (int, string, time,Time etc.).

> If you have very special case you can use the globalValidator method. If you wish you can use the interface signatured method but the safer way would be to use global validator which is `func(field reflect.StructField, value reflect.Value)`.


> __Special Note__: Also you can create a new kind. We covered the types that we frequently use like (time.Time, float32, float64, int, int64, string, interface{}, struct{} etc.). But if you want to write special configurations to specific types that you have created in your project. Also you can do it very easily too.

```go
import (
    "fmt"

    "github.com/bilginyuksel/gorify"
)

var lengthEqual = func(value string, tags reflect.StructTag) error {
    if strSize, exists := tags.Lookup("size"); exists {
        size, err := strconv.Atoi(strSize)

        if err != nil {
            return errors.New("value of the size should be int")
        }

        if len(value) != size {
            return errors.New("value length is not equal to size")
        }
    }
    return nil
}


func main() {
    NewStringValidator(lengthEqual)
}

```