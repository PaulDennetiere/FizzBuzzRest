# Fizz Buzz exercise

Warning : this version is not production ready.


This version is one of the simplest ways to implement the Fizz Buzz server.
However there are several things that should be improve : 

- API is not Restful
- Data validation is ugly 
- We use the default http server (http.ListenAndServe)
- Nothing is tested
- The way code is organized doesn't allow the server to grow easily (eg: adding more handlers properly)

In conclusion this version is the first experiment, that make sure the problem is understood.