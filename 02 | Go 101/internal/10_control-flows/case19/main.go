package main

func main() {
	goto Label1 // error
	{
	Label1:
		goto Label2 // error
	}
	{
	Label2:
	}
}
