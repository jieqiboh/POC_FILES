namespace go multiply

/*
 C like comments are supported
*/
// This is also a valid comment

typedef i32 int // We can use typedef to get pretty names for the types we are using
service MultiplicationService
{
        int multiply(1:int n1, 2:int n2),
}