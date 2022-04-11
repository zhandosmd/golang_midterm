# golang_midterm

A little bit documentation:
1.	http://localhost:8080/store - by this get request we can get all our key-value pair. example below:
    ![image](https://user-images.githubusercontent.com/64863365/162787882-53720b09-a87b-455e-85c5-4811e1df6fe8.png)
    
2.	http://localhost:8080/store/:key –by this get request we can get specific key-value pair with key
    ![image](https://user-images.githubusercontent.com/64863365/162787940-96b4fd8e-4b47-47cc-904e-4910aa961122.png)

3.	http://localhost:8080/store - by this POST request we can create new key-value pair. we should send json data(both required) in the body of request like this:
    ![image](https://user-images.githubusercontent.com/64863365/162787962-d78cf116-78d3-475e-ac26-ddc033aa4f13.png)

4.	http://localhost:8080/store/:key - by this PUT request we can edit existing key-value pair. example below:
    ![image](https://user-images.githubusercontent.com/64863365/162788032-37cb618f-0cc3-4537-9274-1b293e0770a6.png)
    ![image](https://user-images.githubusercontent.com/64863365/162788043-44d09428-03c6-4761-a98f-f2e2b5eb73e9.png)
 
5.	http://localhost:8080/store/:key – by this DELETE request we can delete existing key-value pair. example below:
    ![image](https://user-images.githubusercontent.com/64863365/162788069-a95789ab-15ba-4021-a68a-9649dd3ce0d5.png)
