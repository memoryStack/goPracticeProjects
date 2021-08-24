# goPracticeProjects
after learning some basics about GO i will be solving exercises from https://courses.calhoun.io/courses/cor_gophercises course
one by one until i get a good enough grasp on the GO

Learnings from Problem 1
    PartA
        1. how to open files from filesystem
        2. how to use "csv" package to read entries
        3. what's the use of "defer" keyword
        4. how to use "strconv" package
    PartB
        1. how to use channels
        2. wrote an interesting fact about os.Open("file.txt") in partB.go file {about where this file should be present}
        3. how to use "select" statement while listening from 2 or more go routines
    Bonus
        1. learned how to use flags
        2. learned how to randomize the array
            command to run the program would be "go run Problem_1/cmd/main.go -shuffle=true A"

Learning from Choose Your Own Adventure Exercise
    1. this awesome tool will convert json object to go type struct automatically https://mholt.github.io/json-to-go/
    2. if you run this command "go run cmd/main.go --help" in the "goPracticeProjects/ChooseYourOwnAdventure" folder then
        we can see some "flags" to pass or maybe other info (not sure what) about running the program. it's pretty awesome to me.
    3. was facing an issue with the auto-import saw how that will be resolved.
    4. wrote a TODO for error-handling error using Panic. some research will help me grow there.
    5. learned how to parse json to go structs.
    6. At this point i understand how to make a web server.
    7. TODO: need to understand html template in GO.

Learnings from HTML Link Parser Exercise
    1. Added 2 todos which i need too investigate.
    2. Learned how to play with HTML API a little better.
    3. Learned how nodes work in the HTML when it gets parsed. specially with the nested nodes.
    4. The solution doesn't work properly with the nested links, like "<a>Text1 <a>Text3</a> Text2 </a>".
        this test case is also mentioned in the "html1.html" file and in this case "Text2" doesn't get recoroded
        in the textValue of the first <a> tag, which was expected to be recorded ideally.
