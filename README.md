**The objective of this task is to write an API with one GET endpoint.**

**Requirements:**

- Endpoint should return the data taken from the mocked-up database using the provided helpers functions `GetUsers().` This functions returns an slice of `Users (ID - int, Name - string, Role - string)`
- A request GET `/users?name=Jhon` should return all users with given name and status code `200`.
- A request GET `/users?name=SomeNameNotExist` should return an empty slice and status code `200`.
